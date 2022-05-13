package initdb

import (
	"opquotes/pub"
	"opquotes/log"
	"time"
)

type Agent struct {
	Nodeid         int64     `json:"nodeid" xorm:"not null pk unique BIGINT"`
	Parentid       int64     `json:"parentid" xorm:"not null BIGINT"`
	Username       string    `json:"username" xorm:"not null VARCHAR(50)"`
	Password       string    `json:"password" xorm:"not null VARCHAR(100)"`
	Name           string    `json:"name" xorm:"not null VARCHAR(100)"`
	Avatar         string    `json:"avatar" xorm:"VARCHAR(100)"`
	RememberToken  string    `json:"remember_token" xorm:"VARCHAR(60)"`
	Createdatetime time.Time `json:"createdatetime" xorm:"not null DATETIME"`
	Updatedatetime time.Time `json:"updatedatetime" xorm:"not null DATETIME"`
	Email          string    `json:"email" xorm:"VARCHAR(30)"`
	Mobile         string    `json:"mobile" xorm:"VARCHAR(20)"`
	Sex            int       `json:"sex" xorm:"SMALLINT"`
	Status         int       `json:"status" xorm:"SMALLINT"`
	Buyfee         float64   `json:"buyfee" xorm:"not null default 5 DOUBLE"`
	Closefee       float64   `json:"closefee" xorm:"not null default 5 DOUBLE"`
	Bankname       string    `json:"bankname" xorm:"not null VARCHAR(20)"`
	Bankcard       string    `json:"bankcard" xorm:"not null VARCHAR(30)"`
	Bank           string    `json:"bank" xorm:"not null VARCHAR(100)"`
}
func (ag *Agent)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Agent{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Agent{})
		if err != nil {
			log.Logger.Debug("初始化Agent表失败!")
			return
		}
	}
}