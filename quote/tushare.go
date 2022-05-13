package quote

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
	"opquotes/basedata"
	"opquotes/pub"
	"opquotes/utils"
	"strings"
	"time"
	"opquotes/log"
)
const apiurl="http://api.waditu.com"
const token="c400d142b09955b1929edc213f57a013186eece51c38de8dcd75eb4f"
type Tushare struct{
	Ctx *pub.Context
}
func NewTuShare(Ctx *pub.Context)*Tushare{
	return &Tushare{Ctx: Ctx}
}
func GetTradeDate(date time.Time,count int)[]*basedata.Tradedate{
	var dates =make([]*basedata.Tradedate,0)
	sql:=fmt.Sprintf(`select date from tradedate where date<='%s' order by date desc limit %d`,utils.Time2StrDate(date),count)
	if err:=pub.DB.SQL(sql).Find(&dates);err!=nil{
		log.Logger.Error("获取交易日历错误"+err.Error())
	}else{
		return dates
	}
	return nil
}
func (t *Tushare)UnderlyingTick(){

}
func (t *Tushare)GetTransDetailData(count uint16,codes ...string){

}
func  (t *Tushare)GetDayData(codes ...string){
 	t.GetBarData("1day",2)
}
func (t *Tushare)GetBarData(ktype string, count int, codes ...string){
	if strings.Contains(strings.ToUpper(ktype),"DAY"){
		dates:=GetTradeDate(time.Now(),count)
		for _,day:=range dates{
			t.GetDayK(day.Date)
		}
	}
}
func (t *Tushare)GetMinData(codes ...string){

}
func (t *Tushare)Stop(){

}
func (t *Tushare)Start(){

}
func (t *Tushare)GetTickData(codes ...string){

}

func (t *Tushare)GetDayK(day time.Time)error{
	parmas:=&grequests.RequestOptions{
		JSON: map[string]interface{}{
			"token": token,
			"api_name": "opt_daily",
			"params":map[string]string{"exchange":"sse","trade_date":day.Format("20060102")},
		},
	}
	resp,err:=grequests.Post(apiurl,parmas)
	if err != nil {
		log.Logger.Error("Unable to make request: "+err.Error())
	}else{
		content:=utils.StringFromGBK(resp.String())
		record:=gjson.Get(content,"data.items").Array()
		for i:=0;i<len(record)-1;i++{
			r:=record[i].Array()
			bar:=basedata.NewBar()
			bar.Code=strings.Split(r[0].String(),".")[0]
			bar.Exchange=strings.Split(r[0].String(),".")[1]
			bar.Datetime=utils.Strtotime2(r[1].String())
			bar.Type=3
			bar.PreClose=r[4].Float()
			bar.PreSettle=r[3].Float()
			bar.Settle=r[9].Float()
			bar.Open=r[5].Float()
			bar.High=r[6].Float()
			bar.Low=r[7].Float()
			bar.Close=r[8].Float()
			bar.Vol=r[10].Int()
			bar.Amount=r[11].Float()
			bar.Hold=r[12].Int()
			fmt.Println(bar)
			bar.Save("day")
		}
	}
	return nil
}
func (t *Tushare)GetTradeDate()error{
	parmas:=&grequests.RequestOptions{
		JSON: map[string]interface{}{
			"token": token,
			"api_name": "trade_cal",
			"params":map[string]string{"exchange":"sse","is_open":"1"},
		},
	}
	resp,err:=grequests.Post(apiurl,parmas)
	if err != nil {
		log.Logger.Error("Unable to make request: "+err.Error())
	}else{
		content:=utils.StringFromGBK(resp.String())
		record:=gjson.Get(content,"data.items").Array()
		for i:=0;i<len(record)-1;i++{
			r:=record[i].Array()
			day:=utils.Strtotime2(r[1].String())
			td:=basedata.NewTradeDate(day)
			td.Save()
		}
	}
	return nil
}
func (t *Tushare)GetOptions(){
	parmas:=&grequests.RequestOptions{
		JSON: map[string]interface{}{
			"token": token,
			"api_name": "opt_basic",
			"fields":"ts_code,opt_code,name,exercise_type,per_unit,opt_type,call_put,exercise_price,s_month,maturity_date,list_date,delist_date",
			"params":map[string]string{"exchange":"sse"},
		},
	}
	resp,err:=grequests.Post(apiurl,parmas)
	if err != nil {
		log.Logger.Error("Unable to make request: "+err.Error())
	}else{
		content:=utils.StringFromGBK(resp.String())
		record:=gjson.Get(content,"data.items").Array()
		fmt.Println()
		for i:=0;i<len(record)-1;i++{
			r:=record[i].Array()
			if utils.Strtotime2(r[9].String()).After(time.Now()){
				log.Logger.Error("该合约已经过期"+utils.ToString(time.Now())+","+r[9].String())
				continue
			}
			option:=basedata.NewOpiton()
			option.Code=strings.Split(r[0].String(),".")[0]
			option.Exchange=strings.Split(r[0].String(),".")[1]
			option.Name=r[1].String()
			option.Unit=r[2].Int()
			if strings.Contains(r[3].String(),"510050"){
				option.Underlying="50ETF(510050)"
			}else if strings.Contains(r[3].String(),"510300"){
				option.Underlying="300ETF(510300)"
			}
			if strings.ToUpper(r[5].String())=="C"{
				option.Type="认购"
			}else{
				option.Type="认沽"
			}
			option.Strikeprice=r[7].Float()
			option.Expire=utils.Strtotime2(r[9].String())
			option.Remain=utils.RemainDay(option.Expire)
			option.Month=r[8].String()[4:6]
			option.Save()
		}
	}
}