package basedata
import (
	"fmt"
	"opquotes/log"
	"opquotes/pub"
	"opquotes/utils"
	"time"
)
type Tradedate struct {
	Date time.Time `json:"date" xorm:"DATE"`
}
func NewTradeDate(date time.Time)*Tradedate{
	return &Tradedate{Date:date}
}
func (t *Tradedate)Save()error{
		sql:=fmt.Sprintf(`insert into tradedate(date)values('%s')  on conflict(date) do nothing`,utils.Time2StrDate(t.Date))
		_,err:=pub.DB.Exec(sql)
		if err!=nil{
			log.Logger.Error("插入tradedate数据失败，错误代码="+err.Error())
			return err
		}
		return nil
}
type TradeDateMgr struct{
	Dates []*Tradedate
}
func NewTradeDateMgr()*TradeDateMgr{
	return &TradeDateMgr{
		Dates: make([]*Tradedate,0),
	}
}
func (mgr *TradeDateMgr)Load()error{
	nowday := time.Now().Format("2006-01-02")
	sql:=fmt.Sprintf(`select * from  tradedate where date<='%s' order by date desc limit 100`,nowday)
	err:=pub.DB.SQL(sql).Find(&mgr.Dates)
	if err!=nil{
		log.Logger.Error("获取交易日历失败，错误代码="+err.Error())
		return err
	}
	return nil
}