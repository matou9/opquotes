package initdb

import (
	"opquotes/pub"
	"opquotes/log"
	"time"
)

type Min1 struct {
	Code      string    `json:"code" xorm:"not null pk unique(pk_tbl_primary_min1) VARCHAR(20)"`
	Datetime  time.Time `json:"datetime" xorm:"not null pk unique(pk_tbl_primary_min1) DATETIME"`
	Open      float64   `json:"open" xorm:"not null default 0 DOUBLE"`
	High      float64   `json:"high" xorm:"not null default 0 DOUBLE"`
	Low       float64   `json:"low" xorm:"not null default 0 DOUBLE"`
	Close     float64   `json:"close" xorm:"not null default 0 DOUBLE"`
	Vol       int64     `json:"vol" xorm:"not null default 0 BIGINT"`
	Amount    float64   `json:"amount" xorm:"not null default 0 DOUBLE"`
	Hold      int64     `json:"hold" xorm:"not null default 0 BIGINT"`
	AvgPrice  float64   `json:"avg_price" xorm:"not null default 0 DOUBLE"`
	HighLimit float64   `json:"high_limit" xorm:"not null default 0 DOUBLE"`
	LowLimit  float64   `json:"low_limit" xorm:"not null default 0 DOUBLE"`
	PreClose  float64   `json:"pre_close" xorm:"not null default 0 DOUBLE"`
	Exchange  string    `json:"exchange" xorm:"default '''::character varying' VARCHAR(10)"`
	Paused    int       `json:"paused" xorm:"not null default 0 INTEGER"`
	Type      int       `json:"type" xorm:"not null pk unique(pk_tbl_primary_min1) INTEGER"`
}
func (k *Min1)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Min1{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Min1{})
		if err != nil {
			log.Logger.Debug("初始化Min1表失败!")
			return
		}
	}
}