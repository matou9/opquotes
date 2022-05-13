package initdb

import "opquotes/pub"
import "opquotes/log"

type Payaccount struct {
	Payid       int64  `json:"payid" xorm:"not null pk unique BIGINT"`
	Type        int64  `json:"type" xorm:"not null BIGINT"`
	Account     string `json:"account" xorm:"not null VARCHAR(30)"`
	Accountname string `json:"accountname" xorm:"not null VARCHAR(30)"`
	Accountinfo string `json:"accountinfo" xorm:"VARCHAR(50)"`
	Status      int64  `json:"status" xorm:"not null default 0 BIGINT"`
	Memo        string `json:"memo" xorm:"VARCHAR(50)"`
}
func (p *Payaccount)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Payaccount{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Payaccount{})
		if err != nil {
			log.Logger.Debug("初始化Payaccount表失败!")
			return
		}
	}
}