package initdb

import "opquotes/pub"
import "opquotes/log"
type Userbank struct {
	Uid         int64  `json:"uid" xorm:"not null pk unique(pk_tbl_primary_userbank) BIGINT"`
	Bank        string `json:"bank" xorm:"not null VARCHAR(50)"`
	Bankname    string `json:"bankname" xorm:"VARCHAR(50)"`
	Bankaccount string `json:"bankaccount" xorm:"not null pk unique(pk_tbl_primary_userbank) VARCHAR(50)"`
	Province    string `json:"province" xorm:"VARCHAR(10)"`
	City        string `json:"city" xorm:"VARCHAR(10)"`
	Memo        string `json:"memo" xorm:"VARCHAR(100)"`
}
func (u *Userbank)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Userbank{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Userbank{})
		if err != nil {
			log.Logger.Debug("初始化Userbank表失败!")
			return
		}
	}
}