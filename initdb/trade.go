package initdb

import (
	"opquotes/pub"
	"opquotes/log"
	"time"
)

type Trade struct {
	Accountid       int64     `json:"accountid" xorm:"not null pk unique(pk_tbl_primary_trade) BIGINT"`
	Parentaccountid int64     `json:"parentaccountid" xorm:"not null default 0 BIGINT"`
	Orderid         string    `json:"orderid" xorm:"not null VARCHAR(100)"`
	Tradeid         string    `json:"tradeid" xorm:"not null pk unique(pk_tbl_primary_trade) VARCHAR(100)"`
	Datetime        time.Time `json:"datetime" xorm:"not null DATETIME"`
	Symbol          string    `json:"symbol" xorm:"not null VARCHAR(20)"`
	Exchange        string    `json:"exchange" xorm:"not null VARCHAR(10)"`
	Direction       int       `json:"direction" xorm:"not null default 0 INTEGER"`
	Offsettype      int       `json:"offsettype" xorm:"not null default 0 INTEGER"`
	Price           float64   `json:"price" xorm:"not null default 0 DOUBLE"`
	Volume          float64   `json:"volume" xorm:"not null default 0 DOUBLE"`
	Commission      float64   `json:"commission" xorm:"not null default 0 DOUBLE"`
	Tax             float64   `json:"tax" xorm:"not null default 0 DOUBLE"`
	Flag            int64     `json:"flag" xorm:"not null default 0 BIGINT"`
	Memo            string    `json:"memo" xorm:"VARCHAR(50)"`
}
func (s *Trade)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Trade{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Trade{})
		if err != nil {
			log.Logger.Debug("初始化trade表失败!")
			return
		}
	}
}