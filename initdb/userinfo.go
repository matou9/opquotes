package initdb

import (
	"opquotes/pub"
	"opquotes/log"
	"time"
)

type Userinfo struct {
	Uid         int64     `json:"uid" xorm:"not null pk unique BIGINT"`
	Code        string    `json:"code" xorm:"not null VARCHAR(20)"`
	Password    string    `json:"password" xorm:"not null VARCHAR(20)"`
	Name        string    `json:"name" xorm:"not null VARCHAR(20)"`
	Idcard      string    `json:"idcard" xorm:"VARCHAR(18)"`
	Mobile      string    `json:"mobile" xorm:"not null VARCHAR(11)"`
	Nodeid      int       `json:"nodeid" xorm:"not null INTEGER"`
	Bank        string    `json:"bank" xorm:"VARCHAR(50)"`
	Bankname    string    `json:"bankname" xorm:"VARCHAR(50)"`
	Bankaccount string    `json:"bankaccount" xorm:"VARCHAR(50)"`
	Email       string    `json:"email" xorm:"VARCHAR(20)"`
	Imei        string    `json:"imei" xorm:"VARCHAR(50)"`
	Msystem     string    `json:"msystem" xorm:"VARCHAR(50)"`
	Mversion    string    `json:"mversion" xorm:"VARCHAR(50)"`
	Lastip      string    `json:"lastip" xorm:"VARCHAR(50)"`
	Lasttime    time.Time `json:"lasttime" xorm:"DATETIME"`
	Memo        string    `json:"memo" xorm:"VARCHAR(50)"`
	Buylimit    int64     `json:"buylimit" xorm:"not null default 0 BIGINT"`
	Selllimit   int64     `json:"selllimit" xorm:"not null default 0 BIGINT"`
	Status      int64     `json:"status" xorm:"not null default 0 BIGINT"`
}
func (u *Userinfo)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Userinfo{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Userinfo{})
		if err != nil {
			log.Logger.Debug("初始化Userinfo表失败!")
			return
		}
	}
}