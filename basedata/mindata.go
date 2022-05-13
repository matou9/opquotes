package basedata

import (
	"fmt"
	"opquotes/log"
	"opquotes/pub"
	"opquotes/utils"
	"time"
)

type Mindata struct {
	Code     string    `json:"code" xorm:"not null pk VARCHAR(20)"`
	Datetime time.Time `json:"datetime" xorm:"not null pk DATETIME"`
	Price    float64   `json:"price" xorm:"not null default 0 DOUBLE"`
	AvgPrice float64   `json:"avg_price" xorm:"not null default 0 DOUBLE"`
	Vol      int64     `json:"vol" xorm:"not null default 0 BIGINT"`
	Amount   float64   `json:"amount" xorm:"not null default 0 DOUBLE"`
	Hold     int64     `json:"hold" xorm:"not null default 0 BIGINT"`
	PreClose float64   `json:"pre_close" xorm:"not null default 0 DOUBLE"`
}

func NewMin()*Mindata{
	return &Mindata{}
}
func (m *Mindata)Min(code string)[]*Mindata{
	var data []*Mindata
	start:=time.Now().Format("2006-01-02")+" 09:25"
	end:=time.Now().Format("2006-01-02")+" 15:00"
	sql := fmt.Sprintf("SELECT * FROM MINDATA WHERE CODE='%s' AND DATETIME>='%s' AND DATETIME<='%s' ORDER BY DATETIME ASC",code,start,end)
	pub.DB.SQL(sql).Find(&data)
	return data
}
func (m *Mindata)Save()error{
	sql:=fmt.Sprintf(`insert into Mindata(Code,Datetime,Price,Avg_price,Vol,Amount,Hold,Pre_Close) values ('%s','%s',%f,%f,%d,%f,%d,%f) on conflict(Code,Datetime) do update set Code=EXCLUDED.Code,Datetime=EXCLUDED.Datetime,Price=EXCLUDED.Price,Avg_price=EXCLUDED.Avg_price,Vol=EXCLUDED.Vol,Amount=EXCLUDED.Amount,Hold=EXCLUDED.Hold,Pre_Close=EXCLUDED.Pre_Close`,m.Code,utils.Time2Str(m.Datetime),m.Price,m.AvgPrice,m.Vol,m.Amount,m.Hold,m.PreClose)
	_,err:=pub.DB.Exec(sql)
	if err!=nil{
		log.Logger.Error("mindata分时数据插入失败，错误代码="+err.Error())
		return err
	}
	return nil
}