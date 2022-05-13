package basedata

import (
	"bytes"
	"fmt"
	"opquotes/log"
	"opquotes/pub"
	"opquotes/utils"
	"strings"
	"time"
)

type Bar struct {
	Code      string    `json:"code" xorm:"not null pk unique() VARCHAR(20)"`
	Datetime  time.Time `json:"datetime" xorm:"not null pk unique() DATETIME"`
	Exchange  string    `json:"exchange" xorm:"default '''::character varying' VARCHAR(10)"`
	Type      int       `json:"type" xorm:"not null pk unique() INTEGER"`
	Open      float64   `json:"open" xorm:"not null default 0 DOUBLE"`
	High      float64   `json:"high" xorm:"not null default 0 DOUBLE"`
	Low       float64   `json:"low" xorm:"not null default 0 DOUBLE"`
	Settle     float64   `json:"settle" xorm:"not null default 0 DOUBLE"`
	PreSettle     float64   `json:"pre_settle" xorm:"not null default 0 DOUBLE"`
	Close     float64   `json:"close" xorm:"not null default 0 DOUBLE"`
	Vol       int64     `json:"vol" xorm:"not null default 0 BIGINT"`
	Hold       int64     `json:"hold" xorm:"not null default 0 BIGINT"`
	Amount    float64   `json:"amount" xorm:"not null default 0 DOUBLE"`
	AvgPrice  float64   `json:"avg_price" xorm:"not null default 0 DOUBLE"`
	HighLimit float64   `json:"high_limit" xorm:"not null default 0 DOUBLE"`
	LowLimit  float64   `json:"low_limit" xorm:"not null default 0 DOUBLE"`
	PreClose  float64   `json:"pre_close" xorm:"not null default 0 DOUBLE"`
	Paused    int       `json:"paused" xorm:"not null default 0 INTEGER"`
}

func NewBar()*Bar{
	return &Bar{}
}
func (bar *Bar)Save(tablename string)error{
	if tablename==""{
		log.Logger.Error("bar数据对应的表为空，不能插入")
		return nil
	}
	sql:=fmt.Sprintf(`insert into %s(code,datetime,exchange,type,open,high,low,settle,pre_settle,close,vol,hold,amount,avg_price,high_limit,low_limit,pre_close,paused) values ('%s','%s','%s',%d,%f,%f,%f,%f,%f,%f,%d,%d,%f,%f,%f,%f,%f,%d) on conflict(Code,Datetime,Type) do update set code=EXCLUDED.code,datetime=EXCLUDED.datetime,exchange=EXCLUDED.exchange,type=EXCLUDED.type,open=EXCLUDED.open,high=EXCLUDED.high,low=EXCLUDED.low,settle=EXCLUDED.settle,pre_settle=EXCLUDED.pre_settle,close=EXCLUDED.close,vol=EXCLUDED.vol,hold=EXCLUDED.hold,amount=EXCLUDED.amount,avg_price=EXCLUDED.avg_price,high_limit=EXCLUDED.high_limit,low_limit=EXCLUDED.low_limit,pre_close=EXCLUDED.pre_close,paused=EXCLUDED.paused`,tablename,bar.Code,utils.Time2Str(bar.Datetime),bar.Exchange,bar.Type,bar.Open,bar.High,bar.Low,bar.Settle,bar.PreSettle,bar.Close,bar.Vol,bar.Hold,bar.Amount,bar.AvgPrice,bar.HighLimit,bar.LowLimit,bar.PreClose,bar.Paused)
	_,err:=pub.DB.Exec(sql)
	if err!=nil{
		log.Logger.Error("tushare接口K线数据插入失败，错误代码="+err.Error())
		return err
	}
	return nil
}
func (bar *Bar)KLine(code string,period string)*Bar{
	t :=NewBar()
	var tablename string
	switch strings.ToUpper(period){
	case "15MIN":
		tablename="min15"
	case "1MIN":
		tablename="min1"
	case "30MIN":
		tablename="min30"
	case "60MIN":
		tablename="min60"
	case "DAY":
		tablename="DAY"
	default:
		tablename="min1"
	}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE CODE='%s' ORDER BY DATETIME DESC LIMIT 1",tablename,code)
	pub.DB.SQL(sql).Get(t)
	return t
}

func (bar *Bar)KLines(codes string,period string,count int64)[]*Bar{
	var pjcodes string
	var buffer bytes.Buffer
	for _,code :=range strings.Split(codes,","){
		buffer.WriteString("'")
		buffer.WriteString(code)
		buffer.WriteString("'")
		buffer.WriteString(",")
	}
	if strings.HasSuffix(buffer.String(),","){
		pjcodes=strings.TrimSuffix(buffer.String(),",")
	}
	bars:=make([]*Bar,0)
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

	sql:=fmt.Sprintf("SELECT * FROM %s WHERE CODE IN(%s)  ORDER BY DATETIME LIMIT %d",tablename,pjcodes,count)
	if err:=pub.DB.SQL(sql).Find(&bars);err!=nil{
		return nil
	}
	return bars
}
