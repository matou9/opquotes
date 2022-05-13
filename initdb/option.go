package initdb

import (
	"opquotes/pub"
	"opquotes/log"
	"time"
)

type Option struct {
	Code        string    `json:"code" xorm:"VARCHAR(255)"`
	Syscode     string    `json:"syscode" xorm:"VARCHAR(255)"`
	Name        string    `json:"name" xorm:"VARCHAR(255)"`
	Underlying  string    `json:"underlying" xorm:"VARCHAR(255)"`
	Strikeprice float64   `json:"strikeprice" xorm:"DOUBLE"`
	Type        string    `json:"type" xorm:"VARCHAR(255)"`
	Unit        int64     `json:"unit" xorm:"BIGINT"`
	Expire      time.Time `json:"expire" xorm:"DATETIME"`
	Remain      int64     `json:"remain" xorm:"BIGINT"`
	Month       string    `json:"month" xorm:"VARCHAR(10)"`
}
func (o *Option)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Option{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Option{})
		if err != nil {
			log.Logger.Debug("初始化Option表失败!")
			return
		}
	}
}