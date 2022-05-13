package market

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"opquotes/basedata"
	"opquotes/quote"
	"opquotes/log"
	"opquotes/utils"
	"sort"
	"strings"
	"time"
)
type OptionTTick struct{
	CallCode string //认购代码
	PutCode string //认沽代码
	StrikePrice float64 //行权价

	CallBidPrice1 float64
	PutBidPrice1 float64

	CallBidVolume1 float64
	PutBidVolume1 float64

	CallAskPrice1 float64
	PutAskPrice1 float64

	CallAskVolume1 float64
	PutAskVolume1 float64
	CallChg   float64 //认购涨幅
	PutChg    float64 //认沽涨幅
	CallPrice  float64 //认购最新价
	PutPrice   float64 //认沽最新价
	CallVol int64
	PutVol  int64
	CallHold int64
	PutHold int64
}
type OptionTTicks []*OptionTTick
func (s OptionTTicks) Len() int {
	return len(s)
}

//Less():成绩将有低到高排序
func (s OptionTTicks) Less(i, j int) bool {
	return s[i].StrikePrice<s[j].StrikePrice
}

//Swap()
func (s OptionTTicks) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Mindata(c *gin.Context){
	var lasttradedate string
	code :=c.PostForm("code")
	log.Logger.Info("获取期权分时行情,code="+code)
	if len(quote.OptionMgr.TradeDatemgr.Dates)>1 {
		lasttradedate = utils.Time2StrDate(quote.OptionMgr.TradeDatemgr.Dates[0].Date)
	}
	sql:=fmt.Sprintf(`select  row_to_json(t1)  from
(select to_char(datetime, 'YYYYMMDD') date,pre_close yclose,option.code symbol,"option".name, (
select array_to_json(array_agg(row_to_json(t))) from
(select price,vol,amount,cast(to_char(datetime, 'HH24MI') as integer) "time" ,avg_price,case pre_close when 0 then 0 else (price-pre_close)/pre_close end risefall from mindata where code='%s' and date(mindata.datetime)='%s' order by datetime)t) "minute"  from mindata,option  where mindata.code=option.code and date(mindata.datetime)='%s' and option.code='%s' )t1 limit 1
 `,code,lasttradedate,lasttradedate,code)
	if data,err:=utils.GetFromDBToJson(sql);err!=nil{
		c.JSON(http.StatusOK, gin.H{"status": 204, "data": gin.H{}, "message": "获取mindata数据失败，失败的错误="+err.Error()})
	}else if data==nil {
		c.JSON(http.StatusOK, gin.H{"status": 204, "data": gin.H{}, "message": "没有查到该代码的mindata数据"})
	}else{
		c.JSON(http.StatusOK,gin.H{"status":200,"message":"获取分时数据成功","data":gin.H{"stock":data}})
	}
}
func Tick(c *gin.Context) {
	code :=c.PostForm("code")
	sql:=fmt.Sprintf(`select row_to_json(optiontick) from optiontick where code='%s' `,code)
	if data,err:=utils.GetFromDBToJson(sql);err!=nil{
		c.JSON(http.StatusOK, gin.H{"status": 204, "data": gin.H{}, "message": "获取tick数据失败，失败的错误="+err.Error()})
	}else if data==nil||len(data)==0{
		c.JSON(http.StatusOK, gin.H{"status": 204, "data": gin.H{}, "message": "没有查到该代码的tick数据,代码="+code})
	}else if len(data)==1{
		c.JSON(http.StatusOK,gin.H{"status":200,"message":"返回成功","data":gin.H{"tick":data[0]}})
	}
}
func Ticks(c *gin.Context) {
	code :=c.PostForm("code")
	log.Logger.Info("获取期权多品种tick行情,code="+code)
	symbols:=utils.SplitStrQuote(code)
	sql:=fmt.Sprintf(`select row_to_json(optiontick) from optiontick where code in (%s)`,symbols)
	if data,err:=utils.GetFromDBToJson(sql);err!=nil{
		c.JSON(http.StatusOK, gin.H{"status": 204, "data": gin.H{}, "message": "获取tick数据失败，失败的错误="+err.Error()})
	}else if data==nil {
		c.JSON(http.StatusOK, gin.H{"status": 204, "data": gin.H{}, "message": "没有查到该代码的tick数据"})
	}else{
		c.JSON(http.StatusOK,gin.H{"status":200,"message":"返回成功","data":gin.H{"tick":data}})
	}

}
func Bars(c *gin.Context){
	var(
		sql string
		start,end time.Time
	)
	code :=c.PostForm("code")
	period :=c.PostForm("period")
	count:=utils.ParseInt64(c.PostForm("count"))
	log.Logger.Info("获取期权bar行情,code="+code)
	if len(quote.OptionMgr.TradeDatemgr.Dates)>4{
		start=quote.OptionMgr.TradeDatemgr.Dates[4].Date
		end=quote.OptionMgr.TradeDatemgr.Dates[0].Date
	}else{
		d, _ := time.ParseDuration("-240h")
		start = time.Now().Add(d)
		end=time.Now()
	}

	var tablename string
	switch strings.ToUpper(period){
	case "15MIN":
		tablename="min15"
	case "1MIN":
		tablename="min1"
	case "30MIN":
		tablename="min30"
	case "5MIN":
		tablename="min5"
	case "60MIN":
		tablename="min60"
	case "DAY":
		tablename="DAY"
	default:
		tablename="min1"
	}
	if tablename=="DAY"{
		sql=fmt.Sprintf(`select array_to_json(array_agg(row_to_json(t))) from (select to_char(datetime, 'YYYYMMDD')::int as date,pre_close as preclose,open,high,low,close,vol as volume,amount,EXTRACT(epoch FROM datetime)::int as timestamp,hold from day where code='%s' order by datetime limit %d)t`,code,count)
	}else{
		sql=fmt.Sprintf(`select array_to_json(array_agg(array[to_char(datetime, 'YYYYMMDD')::float, pre_close::float,open::float,high::float,low::float,close::float,vol::float,
amount::float,to_char(datetime, 'HH24MI')::float,hold::float]))from(select datetime,pre_close,open,high,low,close,vol,amount,hold from  %s where code='%s' and datetime>='%s' and datetime<='%s' order by datetime)t`,tablename,code,utils.Time2StrDate(start),utils.Time2StrLast(end))

	}


	if data,err:=utils.GetFromDBToJson(sql);err!=nil{
		c.JSON(http.StatusOK, gin.H{"status": 204, "data": gin.H{}, "message": "获取K线数据失败，失败的错误="+err.Error()})
	}else if data==nil {
		c.JSON(http.StatusOK, gin.H{"status": 204, "data": gin.H{}, "message": "没有查到该代码的K线数据"})
	}else if len(data)==1{
		c.JSON(http.StatusOK,gin.H{"status":200,"message":"获取K线数据成功","data":data[0]})
	}else if len(data)>1{
		c.JSON(http.StatusOK,gin.H{"status":200,"message":"获取K线数据成功","data":data})
	}

}
func TickUnderlying(c *gin.Context){
	code:=c.PostForm("code")
	sql:=fmt.Sprintf(`select row_to_json(optiontick) from optiontick  WHERE code='%s'`,code)
	if data,err:=utils.GetFromDBToJson(sql);err!=nil{
		c.JSON(http.StatusOK, gin.H{"status": 204, "data": gin.H{}, "message": "获取tick数据失败，失败的错误="+err.Error()})
	}else if data==nil || len(data)==0{
		c.JSON(http.StatusOK, gin.H{"status": 204, "data": gin.H{}, "message": "没有查到该代码的tick数据，请检查代码是否正确"})
	}else if len(data)>0{
		c.JSON(http.StatusOK,gin.H{"status":200,"message":"获取标的的tick数据成功","data":gin.H{"tick":data[0]}})
     }
}
func TickByMonth(c *gin.Context){
	var underlying string
	month:=c.PostForm("month")
	underlying=c.PostForm("underlying")
	if strings.Contains(underlying,"300"){
		underlying="510300"
	}else if strings.Contains(underlying,"50"){
		underlying="510050"
	}
	optick:=basedata.Optiontick{}
	ticks:=optick.TicksByMonth(month,underlying)
	strikemap:=make(map[float64]*OptionTTick)
	tticks:=make([]*OptionTTick,0)
	for _,tick :=range ticks{
		if t,ok:=strikemap[tick.Strikeprice];ok{
			if tick.Flag=="认购"||tick.Flag=="C"{
				t.StrikePrice=tick.Strikeprice
				t.CallCode=tick.Code
				t.CallChg=tick.Chg
				t.CallAskPrice1=tick.Askprice1
				t.CallAskVolume1=tick.Askvolume1
				t.CallBidPrice1=tick.Bidprice1
				t.CallBidVolume1=tick.Bidvolume1
				t.CallHold=tick.Hold
				t.CallPrice=tick.Lastprice
				t.CallVol=tick.Vol
			}else if tick.Flag=="认沽"||tick.Flag=="P"{
				t.PutCode=tick.Code
				t.PutChg=tick.Chg
				t.PutHold=tick.Hold
				t.PutPrice=tick.Lastprice
				t.PutVol=tick.Vol
				t.StrikePrice=tick.Strikeprice
				t.PutAskPrice1=tick.Askprice1
				t.PutAskVolume1=tick.Askvolume1
				t.PutBidPrice1=tick.Bidprice1
				t.PutBidVolume1=tick.Bidvolume1
			}
		}else{
			strikemap[tick.Strikeprice]=&OptionTTick{}
			if tick.Flag=="认购"||tick.Flag=="C"{
				strikemap[tick.Strikeprice].CallCode=tick.Code
				strikemap[tick.Strikeprice].CallChg=tick.Chg
				strikemap[tick.Strikeprice].CallHold=tick.Hold
				strikemap[tick.Strikeprice].CallPrice=tick.Lastprice
				strikemap[tick.Strikeprice].CallVol=tick.Vol
				strikemap[tick.Strikeprice].StrikePrice=tick.Strikeprice
				strikemap[tick.Strikeprice].CallAskPrice1=tick.Askprice1
				strikemap[tick.Strikeprice].CallAskVolume1=tick.Askvolume1
				strikemap[tick.Strikeprice].CallBidPrice1=tick.Bidprice1
				strikemap[tick.Strikeprice].CallBidVolume1=tick.Bidvolume1
			}else if tick.Flag=="认沽"||tick.Flag=="P"{
				strikemap[tick.Strikeprice].PutCode=tick.Code
				strikemap[tick.Strikeprice].PutChg=tick.Chg
				strikemap[tick.Strikeprice].PutHold=tick.Hold
				strikemap[tick.Strikeprice].PutPrice=tick.Lastprice
				strikemap[tick.Strikeprice].PutVol=tick.Vol
				strikemap[tick.Strikeprice].StrikePrice=tick.Strikeprice
				strikemap[tick.Strikeprice].PutAskPrice1=tick.Askprice1
				strikemap[tick.Strikeprice].PutAskVolume1=tick.Askvolume1
				strikemap[tick.Strikeprice].PutBidPrice1=tick.Bidprice1
				strikemap[tick.Strikeprice].PutBidVolume1=tick.Bidvolume1
			}
		}
	}
	for _,v:=range strikemap{
		tticks=append(tticks,v)
	}
	sort.Sort(OptionTTicks(tticks))
	if tticks==nil{
		c.JSON(http.StatusOK,gin.H{"status":204,"data":gin.H{},"message":"没有查到对应月份的合约tick数据"})

	}else{
		c.JSON(http.StatusOK,gin.H{"status":200,"message":"获取tick数据成功","data":gin.H{"tick":tticks}})
	}

}
/*

select  array_to_json(array_agg(row_to_json(option)))  from option
select  array_to_json(array_agg(row_to_json(t1)))  from
(select to_char(datetime, 'YYYYMMDD') date,pre_close yclose,option.code symbol,"option".name, (
select array_to_json(array_agg(row_to_json(t))) from
(select price,vol,amount,to_char(datetime, 'HH12MI')"time" ,case pre_close when 0 then 0 else (price-pre_close)/pre_close end risefall from mindata where code='10003010')t) "minute"  from mindata,option  where mindata.code=option.code and option.code='10003010')t1 limit 1
 */