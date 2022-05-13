package initdb

import (
	"opquotes/pub"
	"opquotes/log"
	"time"
)

type Optiontick struct {
	Code        string    `json:"code" xorm:"not null pk unique VARCHAR(20)"`
	Name        string    `json:"name" xorm:"VARCHAR(50)"`
	Exchange    string    `json:"exchange" xorm:"VARCHAR(10)"`
	Unit        int64     `json:"unit" xorm:"not null default 10000 BIGINT"`
	Underlying  string    `json:"underlying" xorm:"not null VARCHAR(50)"`
	Strikeprice float64   `json:"strikeprice" xorm:"not null default 0 DOUBLE"`
	Flag        string    `json:"flag" xorm:"default 'NULL::character varying' VARCHAR(5)"`
	Expire      time.Time `json:"expire" xorm:"DATE"`
	Remain      int64     `json:"remain" xorm:"default 0 BIGINT"`
	Hold        int64     `json:"hold" xorm:"default 0 BIGINT"`
	Datetime    time.Time `json:"datetime" xorm:"not null DATETIME"`
	Lastprice   float64   `json:"lastprice" xorm:"not null default 0 DOUBLE"`
	Chg         float64   `json:"chg" xorm:"not null default 0 DOUBLE"`
	Amp         float64   `json:"amp" xorm:"not null default 0 DOUBLE"`
	Vol         int64     `json:"vol" xorm:"not null default 0 BIGINT"`
	Amount      float64   `json:"amount" xorm:"not null default 0 DOUBLE"`
	Open        float64   `json:"open" xorm:"not null default 0 DOUBLE"`
	High        float64   `json:"high" xorm:"not null default 0 DOUBLE"`
	Low         float64   `json:"low" xorm:"not null default 0 DOUBLE"`
	Preclose    float64   `json:"preclose" xorm:"not null default 0 DOUBLE"`
	Bidprice1   float64   `json:"bidprice1" xorm:"not null default 0 DOUBLE"`
	Bidvolume1  float64   `json:"bidvolume1" xorm:"not null default 0 DOUBLE"`
	Askprice1   float64   `json:"askprice1" xorm:"not null default 0 DOUBLE"`
	Askvolume1  float64   `json:"askvolume1" xorm:"not null default 0 DOUBLE"`
	Bidprice2   float64   `json:"bidprice2" xorm:"not null default 0 DOUBLE"`
	Bidvolume2  float64   `json:"bidvolume2" xorm:"not null default 0 DOUBLE"`
	Askprice2   float64   `json:"askprice2" xorm:"not null default 0 DOUBLE"`
	Askvolume2  float64   `json:"askvolume2" xorm:"not null default 0 DOUBLE"`
	Bidprice3   float64   `json:"bidprice3" xorm:"not null default 0 DOUBLE"`
	Bidvolume3  float64   `json:"bidvolume3" xorm:"not null default 0 DOUBLE"`
	Askprice3   float64   `json:"askprice3" xorm:"not null default 0 DOUBLE"`
	Askvolume3  float64   `json:"askvolume3" xorm:"not null default 0 DOUBLE"`
	Bidprice4   float64   `json:"bidprice4" xorm:"not null default 0 DOUBLE"`
	Bidvolume4  float64   `json:"bidvolume4" xorm:"not null default 0 DOUBLE"`
	Askprice4   float64   `json:"askprice4" xorm:"not null default 0 DOUBLE"`
	Askvolume4  float64   `json:"askvolume4" xorm:"not null default 0 DOUBLE"`
	Bidprice5   float64   `json:"bidprice5" xorm:"not null default 0 DOUBLE"`
	Bidvolume5  float64   `json:"bidvolume5" xorm:"not null default 0 DOUBLE"`
	Askprice5   float64   `json:"askprice5" xorm:"not null default 0 DOUBLE"`
	Askvolume5  float64   `json:"askvolume5" xorm:"not null default 0 DOUBLE"`
}
func (t *Optiontick)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Optiontick{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Optiontick{})
		if err != nil {
			log.Logger.Debug("初始化Optiontick表失败!")
			return
		}
	}
}