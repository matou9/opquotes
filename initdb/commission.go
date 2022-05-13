package initdb


import "opquotes/pub"
import "opquotes/log"

type Commission struct {
	Accountid int64   `json:"accountid" xorm:"not null pk unique BIGINT"`
	Openfee   float64 `json:"openfee" xorm:"not null default 0 DOUBLE"`
	Closefee  float64 `json:"closefee" xorm:"not null default 0 DOUBLE"`
	Memo      string  `json:"memo" xorm:"VARCHAR(50)"`
	Marginfee float64 `json:"marginfee" xorm:"not null default 0 DOUBLE"`
}
func (comm *Commission)CreateTable(){
	if ok,err:=pub.DB.IsTableExist(&Commission{});!ok &&err==nil{
		err := pub.DB.CreateTables(&Commission{})
		if err != nil {
			log.Logger.Debug("初始化Commission表失败!")
			return
		}
	}
}