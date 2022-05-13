package initdb

import "opquotes/pub"
import "opquotes/log"

type Position struct {
	Accountid       int64   `json:"accountid" xorm:"not null pk unique(pk_tbl_primary_position) BIGINT"`
	Parentaccountid int64   `json:"parentaccountid" xorm:"not null pk unique(pk_tbl_primary_position) BIGINT"`
	Symbol          string  `json:"symbol" xorm:"not null pk unique(pk_tbl_primary_position) VARCHAR(20)"`
	Exchange        string  `json:"exchange" xorm:"not null VARCHAR(10)"`
	Direction       int     `json:"direction" xorm:"not null pk unique(pk_tbl_primary_position) INTEGER"`
	Volume          float64 `json:"volume" xorm:"not null default 0 DOUBLE"`
	FrozenVolume    float64 `json:"frozen_volume" xorm:"not null default 0 DOUBLE"`
	CloseableVolume float64 `json:"closeable_volume" xorm:"not null default 0 DOUBLE"`
	PositionCost    float64 `json:"position_cost" xorm:"not null default 0 DOUBLE"`
	AvgPrice        float64 `json:"avg_price" xorm:"not null default 0 DOUBLE"`
	LastPrice       float64 `json:"last_price" xorm:"not null default 0 DOUBLE"`
	CloseProfit     float64 `json:"close_profit" xorm:"not null default 0 DOUBLE"`
	Memo            string  `json:"memo" xorm:"VARCHAR(50)"`
	Margin          float64 `json:"margin" xorm:"not null default 0 DOUBLE"`
}
func (p *Position)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Position{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Position{})
		if err != nil {
			log.Logger.Debug("初始化Position表失败!")
			return
		}
	}
}