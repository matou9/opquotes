package basedata

import (
	"fmt"
	"opquotes/log"
	"opquotes/pub"
	"opquotes/utils"
	"time"
)

type Option struct {
	Code        string    `json:"code" xorm:"not null pk unique VARCHAR(255)"`
	Syscode     string    `json:"syscode" xorm:"VARCHAR(255)"`
	Name        string    `json:"name" xorm:"VARCHAR(255)"`
	Underlying  string    `json:"underlying" xorm:"VARCHAR(255)"`
	Strikeprice float64   `json:"strikeprice" xorm:"DOUBLE"`
	Type        string    `json:"type" xorm:"VARCHAR(255)"`
	Unit        int64     `json:"unit" xorm:"BIGINT"`
	Expire      time.Time `json:"expire" xorm:"DATETIME"`
	Remain      int64     `json:"remain" xorm:"BIGINT"`
	Month       string    `json:"month" xorm:"VARCHAR(10)"`
	Exchange    string    `json:"exchange" xorm:"VARCHAR(255)"`
	Preclose float64  `json:"-" xorm:"-"`
}
func NewOpiton()*Option{
	return &Option{
		Code:        "",
		Syscode:     "",
		Name:        "",
		Underlying:  "",
		Strikeprice: 0,
		Type:        "",
		Unit:        0,
		Expire:      time.Time{},
		Remain:      0,
		Month:       "",
		Exchange:    "",
		Preclose:    0,
	}
}
func (op *Option)Save()error{
	sql:=fmt.Sprintf(`insert into Option(Code,Syscode,Name,Underlying,Strikeprice,Type,Unit,Expire,Remain,Month,Exchange) values ('%s','%s','%s','%s',%f,'%s',%d,'%s',%d,'%s','%s') on conflict(Code) do update set Syscode=EXCLUDED.Syscode,Name=EXCLUDED.Name,Underlying=EXCLUDED.Underlying,Strikeprice=EXCLUDED.Strikeprice,Type=EXCLUDED.Type,Unit=EXCLUDED.Unit,Expire=EXCLUDED.Expire,Remain=EXCLUDED.Remain,Month=EXCLUDED.Month,Exchange=EXCLUDED.Exchange`,op.Code,op.Syscode,op.Name,op.Underlying,op.Strikeprice,op.Type,op.Unit,utils.Time2Str(op.Expire),op.Remain,op.Month,op.Exchange)
	_,err:=pub.DB.Exec(sql)
	if err!=nil{
		log.Logger.Error("通达信接口分时数据插入失败，错误代码="+err.Error())
		return err
	}
	return nil
}

