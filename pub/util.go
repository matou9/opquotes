package pub

import (
	"fmt"
	"strings"
	"sync"
	"time"
)
var(
TradePeriod sync.Map//map[time.Time]bool=make(map[time.Time]bool)
TradeDay sync.Map//map[time.Time]bool=make(map[time.Time]bool)
)
type TransTime struct{
	Tradetime time.Time
}

func IsTradePeriod()bool{
	return true
	if IsTradeDay(time.Now())&&IsTrade(){
		return true
	}
	return false
}
func GetPreTradeDate(day time.Time)time.Time{
	type sDate struct{
		Date time.Time
	}
	retdate:=new(sDate)
	sql:=fmt.Sprintf(`SELECT date FROM tradedate where date<'%s' order by date desc limit 1`,Time2StrDate(day))
	if has,err:=DB.SQL(sql).Get(retdate);err==nil&&has{
		return retdate.Date

	}
	return time.Time{}
}
func IsTradeDay(day time.Time)bool{

	if istradeday,ok:=TradeDay.Load(day);ok{
		return istradeday.(bool)
	}else{
		sql:=fmt.Sprintf("select date from tradedate where date='%s'",Time2StrDate(day))
		has,err:=DB.SQL(sql).Exist()
		if has && err==nil{
			TradeDay.Store(day,true)
			return true
		}else if err!=nil{
			TradeDay.Store(day,false)
			return false
		}
	}
	return false

}
func IsSettleTime()bool{
	settle_start:=Strtotime(time.Now().Format("2006-01-02")+" 15:30:00")
	now:=time.Now()
	if now.After(settle_start)&&IsTradePeriod()==false{
		return true
	}
	return false
}
func IsTrade()bool{
	morning_trade_start:=Strtotime(time.Now().Format("2006-01-02")+" 09:15:00")
	after_trade_end:=Strtotime(time.Now().Format("2006-01-02")+" 15:00:00")
	now:=time.Now()
	if now.After(morning_trade_start)&&now.Before(after_trade_end){
		return true
	}
	return false
}
func IsTradeTime()bool{
	morning_trade_start:=Strtotime(time.Now().Format("2006-01-02")+" 09:15:00")
	morning_trade_end:=Strtotime(time.Now().Format("2006-01-02")+" 11:30:00")
	after_trade_start:=Strtotime(time.Now().Format("2006-01-02")+" 13:00:00")
	after_trade_end:=Strtotime(time.Now().Format("2006-01-02")+" 15:00:00")
	now:=time.Now()
	if (now.After(morning_trade_start)&&now.Before(morning_trade_end))||(now.After(after_trade_start)&&now.Before(after_trade_end)){
		return true
	}
	return false
}
func InTradeTime(now time.Time)bool{
	morning_trade_start:=Strtotime(time.Now().Format("2006-01-02")+" 09:15:00")
	morning_trade_end:=Strtotime(time.Now().Format("2006-01-02")+" 11:30:00")
	after_trade_start:=Strtotime(time.Now().Format("2006-01-02")+" 13:00:00")
	after_trade_end:=Strtotime(time.Now().Format("2006-01-02")+" 15:00:00")
	if (now.After(morning_trade_start)&&now.Before(morning_trade_end))||(now.After(after_trade_start)&&now.Before(after_trade_end)){
		return true
	}
	return false
}
func Time2StrDate(t time.Time)string{
	const shortForm = "2006-01-02"
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	return str
}
func GenTimeSerial(day string,period string)([]*TransTime){
	var(
		sql string
		err error
		trade_morning []*TransTime
		trade_after []*TransTime
		tradetime []*TransTime
		morning_time_start string
		morning_time_start_end string
		after_time_start string
		after_time_start_end string

	)
	if !IsTradeDay(StrtoDate(day)){
		return nil
	}
	if strings.Contains(period,"min")||strings.Contains(period,"minute")||strings.Contains(period,"hours"){
		trade_morning=make([]*TransTime,0)
		trade_after =make([]*TransTime,0)
		morning_time_start=day+" 9:30:00"
		morning_time_start_end =day+" 11:30:00"
		sql=fmt.Sprintf("select tradetime from (SELECT generate_series as tradetime FROM generate_series('%s'::timestamp,'%s', '%s'))  as trade where tradetime>'%s'",morning_time_start,morning_time_start_end,period,morning_time_start)
		err=DB.SQL(sql).Find(&trade_morning)
		if err!=nil{
			fmt.Println(err)
			return nil
		}

		after_time_start=day+" 13:00:00"
		after_time_start_end=day+" 15:00:00"
		sql=fmt.Sprintf("select tradetime from (SELECT generate_series as tradetime FROM generate_series('%s'::timestamp,'%s', '%s'))  as trade where tradetime>'%s'",after_time_start,after_time_start_end,period,after_time_start)
		err=DB.SQL(sql).Find(&trade_after)
		tradetime=append(trade_morning,trade_after...)

	}

	return tradetime
}

func StrtoDate(timestr string)time.Time{
	todayZero, err := time.ParseInLocation("2006-01-02",timestr ,time.Local)
	if err==nil{
		return todayZero
	}
	return time.Time{}


}

func Strtotime(timestr string)time.Time{
	todayZero, err := time.ParseInLocation("2006-01-02 15:04:05",timestr ,time.Local)
	if err==nil{
		return todayZero
	}
	return time.Time{}


}
func Time2Str(t time.Time) string {
	const shortForm = "2006-01-02 15:04:05"
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	return str
}
func GenTradeTime(market string,day string,period string)bool{
	if market=="" ||day==""||period==""{
		return false
	}else{
		timeserial:=GenTimeSerial(day,period)
		trans:=BeginTrans()
		for _,ts :=range timeserial{
			tradetime:=&Tradetime2{
				Time:ts.Tradetime ,
				Market: market,
				Freq: period,
				Memo:   "",
			}
			_,err:=DB.Table("tradetime").Insert(tradetime)
			if err!=nil{
				trans.Rollback()
				return false
			}

		}

		return true

	}

}