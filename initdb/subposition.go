package initdb

import "opquotes/pub"
import "opquotes/log"

type Subposition struct {
	Accountid       int64   `json:"accountid" xorm:"not null pk unique(pk_tbl_primary_subposition) BIGINT"`
	Parentaccountid int64   `json:"parentaccountid" xorm:"not null pk unique(pk_tbl_primary_subposition) BIGINT"`
	Symbol          string  `json:"symbol" xorm:"not null pk unique(pk_tbl_primary_subposition) VARCHAR(20)"`
	Exchange        string  `json:"exchange" xorm:"not null VARCHAR(10)"`
	Direction       int     `json:"direction" xorm:"not null pk unique(pk_tbl_primary_subposition) INTEGER"`
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
func (o *Subposition)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Subposition{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Subposition{})
		if err != nil {
			log.Logger.Debug("初始化Subposition表失败!")
			return
		}
	}
}