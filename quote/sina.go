package quote

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
	"opquotes/basedata"
	"opquotes/log"
	"opquotes/pub"
	"opquotes/utils"
	"strconv"
	"strings"
	"time"
)


type SinaCrawl struct{
	Ctx *pub.Context
	TickMgr map[string]*basedata.Optiontick
}
func NewSina(ctx *pub.Context)*SinaCrawl{
	return &SinaCrawl{
		Ctx: ctx,
	}
}
func (sc *SinaCrawl)Start(){

}
func (sc *SinaCrawl)Stop(){

}
func (sc *SinaCrawl)GetBarData(ktype string, count int, codes ...string) {

}
func(sc *SinaCrawl)GetTransDetailData(count uint16,codes ...string){

}
func (sc *SinaCrawl)GetMinData(codes ...string){
	var code string
	if len(codes)==0{
		for code,_=range OptionMgr.Items{
			sc.DealMindata(code)
		}
	}else{
		for _,code=range codes{}
		sc.DealMindata(code)
	}
}
func (sc *SinaCrawl)DealMindata(code string){//,wg *sync.WaitGroup
	url:=fmt.Sprintf(`https://stock.finance.sina.com.cn/futures/api/openapi.php/StockOptionDaylineService.getOptionMinline?symbol=CON_OP_%s`,code)
	resp, err := grequests.Get(url,nil)
	if err != nil {
		log.Logger.Error("Unable to make request: "+err.Error())
	}
	content:=utils.StringFromGBK(resp.String())
	record:=gjson.Get(content,"result.data").Array()
	for i:=0;i<len(record)-4;i++{
		m:=basedata.NewMin()
		m.Code=code
		m.Datetime=utils.Strtotime(utils.Time2StrDate(time.Now())+" "+gjson.Get(record[i+4].String(),"i").String())
		if pub.InTradeTime(m.Datetime)==false{
			log.Logger.Error("该分时数据不属于交易时间内的，不保存")
			return
		}
		if op,ok:= OptionMgr.Items[code];ok{
			m.PreClose=op.Preclose
		}
		m.Price=float_64(gjson.Get(record[i+4].String(),"p").String())
		m.Hold=int_64(gjson.Get(record[i+4].String(),"t").String())
		m.Vol=int_64(gjson.Get(record[i+4].String(),"v").String())
		m.Amount=float64(m.Vol)*m.Price*float64(OptionMgr.Items[code].Unit)
		m.AvgPrice=float_64(gjson.Get(record[i+4].String(),"a").String())
		m.Save()
	}

}
func (sc *SinaCrawl)UnderlyingTick(){
	Underlying:=[]string{"sh510050","sh510300"}
	url:="http://hq.sinajs.cn/list="+ strings.Join(Underlying,",")
	resp, err := grequests.Get(url, nil)
	// You can modify the request by passing an optional RequestOptions struct

	if err != nil {
		log.Logger.Error("Unable to make request: "+err.Error())
	}
	content:=utils.StringFromGBK(resp.String())
	sc.DealUnderlyingTick(content)
}
func (sc *SinaCrawl)GetTickData(codes ...string){
	var url string
	if len(codes)==0{
		url="http://hq.sinajs.cn/list="+ OptionMgr.CreateSinaTickCodes()
	}else{
		keys:=make([]string,0)
		for _,k:=range codes{
			key:="CON_OP_"+k
			keys=append(keys,key)
		}
		url = strings.Join(keys,",")
	}
	resp, err := grequests.Get(url, nil)
	// You can modify the request by passing an optional RequestOptions struct

	if err != nil {
		log.Logger.Error("Unable to make request: "+err.Error())
	}
	content:=utils.StringFromGBK(resp.String())
	sc.DealTick(content)

}
func (sc *SinaCrawl)DealUnderlyingTick(content string){
	var ceshicode string
	defer func(){
		if p:=recover();p!=nil{
			log.Logger.Info("出错的代码="+ceshicode)
		}
	}()
	lines :=strings.Split(content,`;`)
	if len(lines)>2{
		for _,line:=range lines{
			cols :=strings.Split(line,`=`)
			if len(cols[0])<8{
				continue
			}
			code :=cols[0][(len(cols[0])-8):len(cols[0])]
			ceshicode=code
			tick :=strings.Split(cols[1],`,`)
			optick:=basedata.Optiontick{}
			optick.Code=code[2:]
			if strings.Contains(optick.Code,"300"){
				optick.Name="300ETF期权"
			}else if strings.Contains(optick.Code,"510050"){
				optick.Name="50ETF期权"
			}else{
				optick.Name=optick.Code
			}

			optick.Open=float_64(tick[1])
			optick.Preclose=float_64(tick[2])
			optick.Lastprice=float_64(tick[3])
			if optick.Preclose!=0{
				optick.Chg=(optick.Lastprice-optick.Preclose)*100/optick.Preclose
			}

			optick.High=float_64(tick[4])
			optick.Low=float_64(tick[5])
			optick.Bidprice1=float_64(tick[6])
			optick.Askprice1=float_64(tick[7])
			optick.Vol = int_64(tick[8])
			optick.Amount=float_64(tick[9])
			optick.Bidvolume1=float_64(tick[10])
			optick.Bidprice1=float_64(tick[11])
			optick.Bidvolume2=float_64(tick[12])
			optick.Bidprice2 =float_64(tick[13])
			optick.Bidvolume3 =float_64(tick[14])
			optick.Bidprice3 =float_64(tick[15])
			optick.Bidvolume4 =float_64(tick[16])
			optick.Bidprice4 =float_64(tick[17])
			optick.Bidvolume5 =float_64(tick[18])
			optick.Bidprice5 =float_64(tick[19])
			optick.Askvolume1 =float_64(tick[20])
			optick.Askprice1 =float_64(tick[21])
			optick.Askvolume2 =float_64(tick[22])
			optick.Askprice2 =float_64(tick[23])
			optick.Askvolume3 =float_64(tick[24])
			optick.Askprice3 =float_64(tick[25])
			optick.Askvolume4 =float_64(tick[26])
			optick.Askprice4 =float_64(tick[27])
			optick.Askvolume5 =float_64(tick[28])
			optick.Askprice5 =float_64(tick[29])
			datetime,_:=time.ParseInLocation("2006-01-02 15:04:05",tick[30]+" "+tick[31],time.Local)
			optick.Datetime=datetime
			optick.Exchange="SSE"
			optick.Save()
		}
	}else{
		log.Logger.Info("该代码获取行情出错,代码="+ceshicode)
	}

}
func (sc *SinaCrawl)DealTick(content string){
	var ceshicode string
	defer func(){
		if p:=recover();p!=nil{
			log.Logger.Info("出错的代码="+ceshicode)
		}
	}()
	if OptionMgr==nil{
		log.Logger.Error("OptionMgr为空，不能保存tick数据")
		return
	}
	lines :=strings.Split(content,`;`)
	if len(lines)>2{
		for _,line:=range lines{
			cols :=strings.Split(line,`=`)
			if len(cols[0])<8{
				continue
			}
			code :=cols[0][(len(cols[0])-8):len(cols[0])]
			ceshicode=code
			tick :=strings.Split(cols[1],`,`)
			optick:=basedata.Optiontick{}
			optick.Code=code
			optick.Bidvolume1=float_64(tick[0])
			optick.Bidprice1=float_64(tick[1])
			optick.Lastprice=float_64(tick[2])
			optick.Askprice1=float_64(tick[3])
			optick.Askvolume1=float_64(tick[4])
			optick.Hold=int_64(tick[5])
			optick.Chg=float_64(tick[6])
			optick.Strikeprice = float_64(tick[7])
			optick.Preclose=float_64(tick[8])
			optick.Open=float_64(tick[9])
			optick.Highlimit=float_64(tick[10])
			optick.Lowlimit=float_64(tick[11])
			optick.Askprice5 =float_64(tick[12])
			optick.Askvolume5 =float_64(tick[13])
			optick.Askprice4 =float_64(tick[14])
			optick.Askvolume4 =float_64(tick[15])
			optick.Askprice3 =float_64(tick[16])
			optick.Askvolume3 =float_64(tick[17])
			optick.Askprice2 =float_64(tick[18])
			optick.Askvolume2 =float_64(tick[19])
			optick.Askprice1 =float_64(tick[20])
			optick.Askvolume1 =float_64(tick[21])
			optick.Bidprice1 =float_64(tick[22])
			optick.Bidvolume1 =float_64(tick[23])
			optick.Bidprice2 =float_64(tick[24])
			optick.Bidvolume2 =float_64(tick[25])
			optick.Bidprice3 =float_64(tick[26])
			optick.Bidvolume3 =float_64(tick[27])
			optick.Bidprice4 =float_64(tick[28])
			optick.Bidvolume4 =float_64(tick[29])
			optick.Bidprice5 =float_64(tick[30])
			optick.Bidvolume5 =float_64(tick[31])
			datetime,_:=time.ParseInLocation("2006-01-02 15:04:05",tick[32],time.Local)
			optick.Datetime=datetime

			optick.Underlying=OptionMgr.Items[optick.Code].Underlying
			optick.Name=tick[37]
			optick.Exchange="SSE"
			optick.Amp=float_64(tick[38])
			optick.High=float_64(tick[39])
			optick.Low=float_64(tick[40])
			optick.Vol=int_64(tick[41])
			optick.Amount=float_64(tick[42])
			optick.Flag=tick[45]
			expire,_:=time.ParseInLocation("2006-01-02",tick[46],time.Local)
			optick.Expire=expire
			optick.Remain=int_64(tick[47])
			optick.Unit= OptionMgr.Items[optick.Code].Unit
			optick.Save()
			OptionMgr.Ticks.Store(optick.Code,optick)
			//pub.DB.Insert(&optick

		}
	}else{
		log.Logger.Info("该代码获取行情出错,代码="+ceshicode)
	}

}
func float_64(lastprice string)(price float64){
	if price,err:= strconv.ParseFloat(lastprice, 64);err==nil{
		return price
	}else{
		return 0
	}
}
func int_64(lastprice string)(price int64){
	if price,err:= strconv.ParseInt(lastprice, 10,64);err==nil{
		return price
	}else{
		return 0
	}
}