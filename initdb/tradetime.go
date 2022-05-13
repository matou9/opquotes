package initdb

import (
	"opquotes/pub"
	"opquotes/log"
	"time"
)

type Tradetime struct {
	Time   time.Time `json:"time" xorm:"not null pk unique(pk_tbl_primary_tradetime) DATETIME"`
	Market string    `json:"market" xorm:"not null pk unique(pk_tbl_primary_tradetime) VARCHAR(10)"`
	Period string    `json:"period" xorm:"not null pk unique(pk_tbl_primary_tradetime) VARCHAR(10)"`
	Memo   string    `json:"memo" xorm:"VARCHAR(50)"`
}
func (s *Tradetime)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Tradetime{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Tradetime{})
		if err != nil {
			log.Logger.Debug("初始化Tradetime表失败!")
			return
		}
	}
}