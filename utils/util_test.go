package utils

import (
	"testing"
	"time"
)
type Mindata struct {
	Code     string    `json:"code" xorm:"not null pk VARCHAR(20)"`
	Datetime time.Time `json:"datetime" xorm:"not null pk DATETIME"`
	Price    float64   `json:"price" xorm:"not null default 0 DOUBLE"`
	Avgprice float64   `json:"avgprice" xorm:"not null default 0 DOUBLE"`
	Vol      int64     `json:"vol" xorm:"not null default 0 BIGINT"`
	Amount   float64   `json:"amount" xorm:"not null default 0 DOUBLE"`
	Hold     int64     `json:"hold" xorm:"not null default 0 BIGINT"`
	PreClose float64   `json:"pre_close" xorm:"not null default 0 DOUBLE"`
}
type Option struct {
	Code        string    `json:"code" xorm:"not null pk unique VARCHAR(255)"`
	Syscode     string    `json:"syscode" xorm:"VARCHAR(255)"`
	Name        string    `json:"name" xorm:"VARCHAR(255)"`
	Underlying  string    `json:"underlying" xorm:"VARCHAR(255)"`
	Strikeprice float64   `json:"strikeprice" xorm:"DOUBLE"`
	Type        string    `json:"type" xorm:"VARCHAR(255)"`
	Unit        int64     `json:"unit" xorm:"BIGINT"`
	Expire      time.Time `json:"expire" xorm:"DATETIME"`
	Remain      int64     `json:"remain" xorm:"BIGINT"`
	Month       string    `json:"month" xorm:"VARCHAR(10)"`
	Exchange    string    `json:"exchange" xorm:"VARCHAR(255)"`
	Preclose float64  `json:"-" xorm:"-"`
}
type Bar struct {
	Code      string    `json:"code" xorm:"not null pk unique() VARCHAR(20)"`
	Datetime  time.Time `json:"datetime" xorm:"not null pk unique() DATETIME"`
	Exchange  string    `json:"exchange" xorm:"default '''::character varying' VARCHAR(10)"`
	Type      int       `json:"type" xorm:"not null pk unique() INTEGER"`
	Open      float64   `json:"open" xorm:"not null default 0 DOUBLE"`
	High      float64   `json:"high" xorm:"not null default 0 DOUBLE"`
	Low       float64   `json:"low" xorm:"not null default 0 DOUBLE"`
	Settle     float64   `json:"settle" xorm:"not null default 0 DOUBLE"`
	PreSettle     float64   `json:"pre_settle" xorm:"not null default 0 DOUBLE"`
	Close     float64   `json:"close" xorm:"not null default 0 DOUBLE"`
	Vol       int64     `json:"vol" xorm:"not null default 0 BIGINT"`
	Hold       int64     `json:"hold" xorm:"not null default 0 BIGINT"`
	Amount    float64   `json:"amount" xorm:"not null default 0 DOUBLE"`
	AvgPrice  float64   `json:"avg_price" xorm:"not null default 0 DOUBLE"`
	HighLimit float64   `json:"high_limit" xorm:"not null default 0 DOUBLE"`
	LowLimit  float64   `json:"low_limit" xorm:"not null default 0 DOUBLE"`
	PreClose  float64   `json:"pre_close" xorm:"not null default 0 DOUBLE"`
	Paused    int       `json:"paused" xorm:"not null default 0 INTEGER"`
}

func Test_Ticks(t *testing.T){
	GenInsertSQL(Bar{},"bar")
}
