package initdb

import "opquotes/pub"
import "opquotes/log"

type Tradeapi struct {
	Accountid     int64  `json:"accountid" xorm:"not null pk unique BIGINT"`
	Brokerid      int    `json:"brokerid" xorm:"not null default 0 INTEGER"`
	Brokerind     int    `json:"brokerind" xorm:"not null default 0 INTEGER"`
	Broker        string `json:"broker" xorm:"not null VARCHAR(20)"`
	Branchid      string `json:"branchid" xorm:"VARCHAR(10)"`
	Clientversion string `json:"clientversion" xorm:"VARCHAR(10)"`
	Accounttype   string `json:"accounttype" xorm:"VARCHAR(20)"`
	Accountno     string `json:"accountno" xorm:"VARCHAR(20)"`
	Tradeaccount  string `json:"tradeaccount" xorm:"VARCHAR(30)"`
	Password      string `json:"password" xorm:"VARCHAR(20)"`
	Txpassword    string `json:"txpassword" xorm:"VARCHAR(20)"`
	Cancelprefix  string `json:"cancelprefix" xorm:"VARCHAR(40)"`
}
func (s *Tradeapi)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Tradeapi{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Tradeapi{})
		if err != nil {
			log.Logger.Debug("初始化Tradeapi表失败!")
			return
		}
	}
}