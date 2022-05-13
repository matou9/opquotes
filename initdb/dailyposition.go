package initdb

import (
	"opquotes/pub"
	"opquotes/log"
	"time"
)

type Dailyposition struct {
	Datetime        time.Time `json:"datetime" xorm:"not null pk unique(pk_tbl_primary_dailyposition) DATETIME"`
	Updatetime      time.Time `json:"updatetime" xorm:"DATETIME"`
	Accountid       int64     `json:"accountid" xorm:"not null pk unique(pk_tbl_primary_dailyposition) BIGINT"`
	Symbol          string    `json:"symbol" xorm:"not null pk unique(pk_tbl_primary_dailyposition) VARCHAR(20)"`
	Exchange        string    `json:"exchange" xorm:"not null VARCHAR(10)"`
	Direction       int       `json:"direction" xorm:"not null pk unique(pk_tbl_primary_dailyposition) INTEGER"`
	Volume          float64   `json:"volume" xorm:"not null default 0 DOUBLE"`
	FrozenVolume    float64   `json:"frozen_volume" xorm:"not null default 0 DOUBLE"`
	CloseableVolume float64   `json:"closeable_volume" xorm:"not null default 0 DOUBLE"`
	PositionCost    float64   `json:"position_cost" xorm:"not null default 0 DOUBLE"`
	AvgPrice        float64   `json:"avg_price" xorm:"not null default 0 DOUBLE"`
	LastPrice       float64   `json:"last_price" xorm:"not null default 0 DOUBLE"`
	Memo            string    `json:"memo" xorm:"VARCHAR(50)"`
}
func (dp *Dailyposition)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Dailyposition{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Dailyposition{})
		if err != nil {
			log.Logger.Debug("初始化Dailyposition表失败!")
			return
		}
	}
}