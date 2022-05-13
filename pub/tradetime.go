package pub

import "time"

type Tradetime struct {
	Time   string `json:"time" xorm:"not null pk unique(pk_tbl_primary_tradetime) varchar(20)"`
	Market string    `json:"market" xorm:"not null pk unique(pk_tbl_primary_tradetime) VARCHAR(20)"`
	Freq string    `json:"freq" xorm:"not null pk unique(pk_tbl_primary_tradetime) VARCHAR(20)"`
	Memo   string    `json:"memo" xorm:"VARCHAR(50)"`
}

type Tradetime2 struct {
	Time   time.Time `json:"time" xorm:"not null pk unique(pk_tbl_primary_tradetime) DATETIME"`
	Market string    `json:"market" xorm:"not null pk unique(pk_tbl_primary_tradetime) VARCHAR(20)"`
	Freq string    `json:"freq" xorm:"not null pk unique(pk_tbl_primary_tradetime) VARCHAR(20)"`
	Memo   string    `json:"memo" xorm:"VARCHAR(50)"`
}

