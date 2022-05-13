package initdb

import (
	"opquotes/pub"
	"opquotes/log"
	"time"
)

type Dailyaccount struct {
	Datetime     time.Time `json:"datetime" xorm:"not null pk unique(pk_tbl_primary_dailyaccount) DATETIME"`
	Updatetime   time.Time `json:"updatetime" xorm:"DATETIME"`
	Accountid    int64     `json:"accountid" xorm:"not null pk unique(pk_tbl_primary_dailyaccount) BIGINT"`
	Initbalance  float64   `json:"initbalance" xorm:"not null default 0 DOUBLE"`
	Balance      float64   `json:"balance" xorm:"not null default 0 DOUBLE"`
	Initcash     float64   `json:"initcash" xorm:"not null default 0 DOUBLE"`
	Cash         float64   `json:"cash" xorm:"not null default 0 DOUBLE"`
	Available    float64   `json:"available" xorm:"not null default 0 DOUBLE"`
	Margin       float64   `json:"margin" xorm:"not null default 0 DOUBLE"`
	FrozenMargin float64   `json:"frozen_margin" xorm:"not null default 0 DOUBLE"`
	PositionCost float64   `json:"position_cost" xorm:"not null default 0 DOUBLE"`
	MarketValue  float64   `json:"market_value" xorm:"not null default 0 DOUBLE"`
	Memo         string    `json:"memo" xorm:"VARCHAR(50)"`
}
func (Dacc *Dailyaccount)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Dailyaccount{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Dailyaccount{})
		if err != nil {
			log.Logger.Debug("初始化Dailyaccount表失败!")
			return
		}
	}
}