package initdb

import (
	"opquotes/pub"
	"opquotes/log"
	"time"
)

type Suborders struct {
	Accountid       int64     `json:"accountid" xorm:"not null pk unique(pk_tbl_primary_suborder) BIGINT"`
	Orderid         string    `json:"orderid" xorm:"not null pk unique(pk_tbl_primary_suborder) VARCHAR(100)"`
	Suborderid      string    `json:"suborderid" xorm:"not null VARCHAR(500)"`
	Symbol          string    `json:"symbol" xorm:"not null VARCHAR(20)"`
	Exchange        string    `json:"exchange" xorm:"not null VARCHAR(10)"`
	Datetime        time.Time `json:"datetime" xorm:"not null pk unique(pk_tbl_primary_suborder) DATETIME"`
	Direction       int       `json:"direction" xorm:"not null default 0 INTEGER"`
	Offsettype      int       `json:"offsettype" xorm:"not null default 0 INTEGER"`
	Price           float64   `json:"price" xorm:"not null default 0 DOUBLE"`
	Volume          float64   `json:"volume" xorm:"not null default 0 DOUBLE"`
	Traded          float64   `json:"traded" xorm:"not null default 0 DOUBLE"`
	Status          int       `json:"status" xorm:"not null default 0 INTEGER"`
	Ordertype       int       `json:"ordertype" xorm:"not null default 0 INTEGER"`
	Margin          float64   `json:"margin" xorm:"not null default 0 DOUBLE"`
	Memo            string    `json:"memo" xorm:"VARCHAR(50)"`
	Parentaccountid int64     `json:"parentaccountid" xorm:"not null default 0 BIGINT"`
}
func (o *Suborders)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Suborders{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Suborders{})
		if err != nil {
			log.Logger.Debug("初始化subOrders表失败!")
			return
		}
	}
}