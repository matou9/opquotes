package initdb

import (
	"opquotes/pub"
	"opquotes/log"
	"time"
)

type Tradedate struct {
	Date time.Time `json:"date" xorm:"DATE"`
}
func (s *Tradedate)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Tradedate{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Tradedate{})
		if err != nil {
			log.Logger.Debug("初始化Tradedate表失败!")
			return
		}
	}
}
