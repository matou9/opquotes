package initdb

import "opquotes/pub"
import "opquotes/log"
type Subaccount struct {
	Uid          int64   `json:"uid" xorm:"not null BIGINT"`
	Accountid    int64   `json:"accountid" xorm:"not null pk unique BIGINT"`
	Balance      float64 `json:"balance" xorm:"not null default 0 DOUBLE"`
	Cash         float64 `json:"cash" xorm:"not null default 0 DOUBLE"`
	Available    float64 `json:"available" xorm:"not null default 0 DOUBLE"`
	Margin       float64 `json:"margin" xorm:"not null default 0 DOUBLE"`
	FrozenMargin float64 `json:"frozen_margin" xorm:"not null default 0 DOUBLE"`
	PositionCost float64 `json:"position_cost" xorm:"not null default 0 DOUBLE"`
	MarketValue  float64 `json:"market_value" xorm:"not null default 0 DOUBLE"`
	Maxopencount float64 `json:"maxopencount" xorm:"not null default 1000 DOUBLE"`
	Maxholdcount float64 `json:"maxholdcount" xorm:"not null default 1000 DOUBLE"`
	Prior        int64   `json:"prior" xorm:"not null default 0 BIGINT"`
	Memo         string  `json:"memo" xorm:"VARCHAR(50)"`
	Initcash     float64 `json:"initcash" xorm:"not null default 0 DOUBLE"`
	Floatprofit  float64 `json:"floatprofit" xorm:"not null DOUBLE"`
	Closeprofit  float64 `json:"closeprofit" xorm:"not null DOUBLE"`
	Status       int64   `json:"status" xorm:"not null default 0 BIGINT"`
	Maxamount    float64 `json:"maxamount" xorm:"not null default 1000 DOUBLE"`
	Tradetype    int64   `json:"tradetype" xorm:"not null default 0 BIGINT"`
}

func (s *Subaccount)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Subaccount{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Subaccount{})
		if err != nil {
			log.Logger.Debug("初始化Subaccount表失败!")
			return
		}
	}
}