package quote

import (
	"fmt"
	"opquotes/basedata"
	"opquotes/log"
	"opquotes/pub"
	"opquotes/utils"
	"strings"
	"sync"
	"time"
)
var(
	OptionMgr *Options
)


type Options struct{
	Ctx *pub.Context
	Items map[string]*basedata.Option
	Ticks sync.Map
	MinDatas sync.Map
	UpdateTime time.Time
	TradeDatemgr *basedata.TradeDateMgr
	QuoteFunc func(codes ...string)
}
func NewOptions(ctx *pub.Context)*Options{
	ops:= &Options{Items:make(map[string]*basedata.Option),Ctx:ctx,TradeDatemgr:basedata.NewTradeDateMgr()}
	//ops.Init()
	ops.TradeDatemgr.Load()
	return ops
}
func (ops *Options)GetTransDataDetail()error{

	tdx:=&Tdx{}
	tdx.GetTransDetailData(100)
	return nil

}
func (ops *Options)Init(){
	if ops.Items!=nil&&len(ops.Items)>0{
		ops.Items=make(map[string]*basedata.Option)
	}
	ops.Createtable()
	options:=ops.SpiderAllOption()
	if len(options)>0{
		ops.Del()
		ops.updatemonth()
		ops.Save()
	}else{
		ops.Load()
	}
	if ops.QuoteFunc==nil{
		log.Logger.Panic("行情接口函数没有定义")
	}else{
		ops.QuoteFunc()
	}
	ops.GetPreClose()

}
func (ops *Options)Load()error{
	options:=make([]*basedata.Option,0)
	sql:=fmt.Sprintf("SELECT * FROM  OPTION")
	if err:=pub.DB.SQL(sql).Find(&options);err==nil{
		if len(options)==0{
			log.Logger.Info("加载期权历史数据基本资料失败！")
		}
		for _,op:=range options{
			ops.Items[op.Code]=op
		}
		ops.UpdateTime=time.Now()
	}
	return nil
}
func (ops *Options)GenMin1Data()error{
	begin:=time.Now()
	log.Logger.Info("根据transaction生成分时1分钟K线数据")
	sql:=`with _min1  as(select distinct  o.code,o.exchange, mindatetime::timestamp + '1 min' datetime ,pre_close,first_value(price) over(partition by t.code||t.exchange||t.mindatetime ORDER BY datetime RANGE BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING) as  open,sum(volume*price*o.unit) over(w) as amount,
last_value(price)  over(partition by t.code||t.exchange||t.mindatetime ORDER BY datetime RANGE BETWEEN  UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING) as close,max(price) over(w) as high,min(price) over(w) low,sum(volume*price) over(w)/sum(volume) over(w)  as avg_price,3 as type,
sum(volume) over(w) as vol from transactiondata t,option o where t.code=o.code 
window w as (partition by t.code||t.exchange||t.mindatetime))
insert into min1(code,datetime,open,high,low,close,vol,amount,avg_price,pre_close,exchange,type)
select code,datetime,open,high,low,close,vol,amount,avg_price,pre_close,exchange,type from _min1 on conflict(code,datetime,type) do nothing`
	_,err:=pub.DB.Exec(sql)
	if err!=nil{
		log.Logger.Error("生成1分钟K线数据失败，请检查，错误代码="+err.Error())
		return err
	}
	fmt.Println("生成1分钟K线数据用时=",time.Now().Sub(begin).Seconds())
	return nil
}
func (ops *Options)GenMinData()error{
	begin:=time.Now()
	log.Logger.Info("根据transaction生成分时mindata数据")
	sql:=`with everymin as(
select distinct o.code,to_char(datetime, 'YYYY-MM-DD HH24:MI')datetime,last_value(price) over(partition by t.code||t.exchange||t.mindatetime ORDER BY datetime RANGE BETWEEN  UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING) as price,t.pre_close,sum(volume) over(w) as vol,sum(volume*price*o.unit) over(w) as amount,sum(volume*price) over(w)/sum(volume) over(w)  as avg_price  from transactiondata t,option o  where o.code=t.code  window w as (partition by t.code||t.exchange||t.mindatetime))
insert into mindata(code,datetime,price,avg_price,vol,amount,pre_close)
select code,datetime::TIMESTAMP,price,avg_price,vol,amount,pre_close from everymin on conflict(code,datetime) do nothing;
`
	_,err:=pub.DB.Exec(sql)
	if err!=nil{
		log.Logger.Error("生成分时数据失败，请检查，错误代码="+err.Error())
		return err
	}
	fmt.Println("生成分时数据用时=",time.Now().Sub(begin).Seconds())
	return nil
}
func (ops *Options)Del()error{
	if _,err:=pub.DB.Exec("delete from option");err!=nil{
		return err
	}
	sql:=fmt.Sprintf("delete from optiontick where datetime<='%s'",utils.Time2StrDate(time.Now()))
	if _,err:=pub.DB.Exec(sql);err!=nil{
		return err
	}
	return nil

}
func (ops *Options)Save()error{
	for _,v :=range ops.Items{
		if _,err:=pub.DB.Table("option").Insert(v);err!=nil{
			return err
		}
	}
	return nil
}
func (ops *Options)Createtable(){
	if err:=pub.DB.Sync2(&basedata.Option{});err!=nil{
		log.Logger.Panic("生成option表失败，直接退出"+err.Error())
	}
}
func (ops *Options)GetOptionsMap()(error,map[string]*basedata.Option){
	optionsmap:=make(map[string]*basedata.Option)
	options:=make([]*basedata.Option,0)
	sql:=fmt.Sprintf("SELECT * FROM  OPTION")
	if err:=pub.DB.SQL(sql).Find(&options);err!=nil{
		return err,nil
	}else{
		for _,op:=range options{
			optionsmap[op.Code]=op
		}
	}
	return nil,optionsmap
}
func (ops *Options)GetOptions()(error,[]*basedata.Option){
	options:=make([]*basedata.Option,0)
	sql:=fmt.Sprintf("SELECT * FROM  OPTION")
	if err:=pub.DB.SQL(sql).Find(&options);err!=nil{
		return err,nil
	}else{
		return nil,options
	}
}
func (ops *Options)SpiderAllOption()map[string]*basedata.Option{
	if len(ops.Items)>0 && time.Now().Day()==ops.UpdateTime.Day(){
		return ops.Items
	}else{
		sse:= NewSSE(ops)
		sse.GetOptionList()
		ops.UpdateTime=time.Now()
		ops.Items=sse.Ops.Items
		return ops.Items
	}
}
func (ops *Options)Contain(code string)bool{
	if _,ok:=ops.Items[code];ok{
		return true
	}else{
		return false
	}
}
func (ops *Options)GetOptionCodes()[]string{
	var values []string=make([]string,0)
	if ops.Items==nil{
		log.Logger.Error("options中的items為空，請檢查")
		return nil
	}
	for k,_ :=range ops.Items{
		values=append(values,k)
	}
	return values
}
func (ops *Options)Name(code string)string{
	if option,found:=ops.Items[code];found{
		return option.Name
	}
	return ""
}
func (ops *Options)Convert(op *basedata.Optiontick){
	if op !=nil{
		if option,found:=ops.Items[op.Code];found{
			op.Name=option.Name
			op.Unit=option.Unit
			op.Strikeprice=option.Strikeprice
			op.Underlying=option.Underlying
			op.Flag=option.Type
			op.Remain=option.Remain
			op.Expire=option.Expire
			op.Preclose=option.Preclose
		}
	}
	return
}

func (ops *Options)GetLastTick(){

}
func (ops *Options)CheckTickData()bool{
	datetime:=time.Now().Format("2006-01-02")
	sql:=fmt.Sprintf("SELECT count(*) FROM OPTIONTICK where cast(datetime as date) IN(SELECT date FROM TRADEDATE WHERE date<='%s' order by date desc limit 1)",datetime)
	has,err:=pub.DB.SQL(sql).Exist()
	if has && err==nil{
		return true
	}else if err!=nil{
		return false
	}
	return false
}
func (ops *Options)CheckMinData()bool{
	datetime:=time.Now().Format("2006-01-02")
	sql:=fmt.Sprintf("select * from mindata where cast(datetime as date)='%s' and pre_close=0 limit 1",datetime)
	has,err:=pub.DB.SQL(sql).Exist()
	if has && err==nil{
		return true
	}else if err!=nil||has==false{
		return false
	}
	return false
}
func (ops *Options)UpdatePreClose(code string,pre_close float64){
	datetime:=time.Now().Format("2006-01-02")
	sql:=fmt.Sprintf("update mindata set pre_close=%f where cast(datetime as date)=(SELECT date FROM TRADEDATE WHERE date<='%s' order by date desc limit 1) and code='%s' and pre_close=0",pre_close,datetime,code)
	_,err:=pub.DB.Exec(sql)
	if err!=nil{
		log.Logger.Panic("修改分时数据的前收盘价错误"+err.Error())
	}
}
func (ops *Options)Settle(){
	ts:=&Tushare{}
	ts.GetDayData()
}
func (ops *Options)GetPreClose()error{
	var upminflag bool=false
	type opdb struct{
		Datetime time.Time
		Code string
		Preclose float64
		Close float64
	}

	var opdbs = make([]*opdb,0)
	datetime:=time.Now().Format("2006-01-02")+" 00:00:00"
	sql:=fmt.Sprintf("select datetime::date ,code,pre_settle as preclose,close from day where datetime in(SELECT date FROM TRADEDATE WHERE date<='%s' order by date desc limit 1)",datetime)
	err:=pub.DB.SQL(sql).Find(&opdbs)
	if err==nil{
		if len(opdbs)==0{
			sql:=fmt.Sprintf("select datetime::date,code,preclose,lastprice as close from optiontick")
			err:=pub.DB.SQL(sql).Find(&opdbs)
			if err==nil&&len(opdbs)>0{
				goto Loop
			}else{
				log.Logger.Error("获取期权的昨日收盘价数据失败，请检查数据库期权的昨日收盘价数据是否齐全")
				return err
			}
		}
	Loop:
		if ops.CheckMinData()==false{
			upminflag=true
			log.Logger.Debug("分时数据前收盘价为空")

		}
		for _,op:=range opdbs{
			if o,ok:=ops.Items[op.Code];ok{
				if pub.GetPreTradeDate(time.Now())==op.Datetime{
					o.Preclose=op.Close
				}else{
					o.Preclose=op.Preclose
				}
				if upminflag==true{
					ops.UpdatePreClose(op.Code,op.Close)
				}
			}
		}

	}else{
		return err
	}
	return nil
}
func (ops *Options)Add(op *basedata.Option){
	if _,found:=ops.Items[op.Code];!found{
		ops.Items[op.Code]=op
	}

}
func (ops *Options)updatemonth(){
	for _,o :=range ops.Items{
		o.Month=o.Expire.Format("01")
	}
}
func (ops *Options)CreateSinaTickCodes()string{
	keys:=make([]string,0)
	for k,_:=range ops.Items{
		key:="CON_OP_"+k
		keys=append(keys,key)
	}
	return strings.Join(keys,",")
}

