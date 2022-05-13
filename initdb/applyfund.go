package initdb

import (
	"opquotes/pub"
	"opquotes/log"
	"time"
)

type Applyfund struct {
	Opuid          int64     `json:"opuid" xorm:"not null BIGINT"`
	Accountid      int64     `json:"accountid" xorm:"not null pk unique(pk_tbl_primary_applyfund) BIGINT"`
	Createdatetime time.Time `json:"createdatetime" xorm:"not null pk unique(pk_tbl_primary_applyfund) DATETIME"`
	Updatetime     time.Time `json:"updatetime" xorm:"not null DATETIME"`
	Payid          int64     `json:"payid" xorm:"not null default 1 BIGINT"`
	Account        string    `json:"account" xorm:"not null VARCHAR(30)"`
	Accountname    string    `json:"accountname" xorm:"not null VARCHAR(30)"`
	Optype         int64     `json:"optype" xorm:"not null default 0 BIGINT"`
	Status         int64     `json:"status" xorm:"not null default 0 BIGINT"`
	Channel        string    `json:"channel" xorm:"not null VARCHAR(10)"`
	Total          float64   `json:"total" xorm:"not null default 0 DOUBLE"`
	Memo           string    `json:"memo" xorm:"VARCHAR(50)"`
}
func (af *Applyfund)Createtable(){
	if ok,err:=pub.DB.IsTableExist(&Applyfund{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Applyfund{})
		if err != nil {
			log.Logger.Debug("初始化Applyfund表失败!")
			return
		}
	}
}