package basedata

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"opquotes/log"
	"opquotes/pub"
	"opquotes/utils"
	"strings"
	"time"
)



type Optiontick struct {
	Code        string    `json:"code" xorm:"not null pk unique(pk_tbl_primary_optiontick) VARCHAR(20)"`
	Name        string    `json:"name" xorm:"VARCHAR(20)"`
	Exchange    string    `json:"exchange" xorm:"VARCHAR(10)"`
	Underlying  string    `json:"underlying" xorm:"not null VARCHAR(10)"`
	Strikeprice float64   `json:"strikeprice" xorm:"not null default 0 DOUBLE"`
	Flag        string    `json:"flag" xorm:"default 'NULL::character varying' VARCHAR(5)"`
	Expire      time.Time `json:"expire" xorm:"DATE"`
	Remain      int64     `json:"remain" xorm:"default 0 BIGINT"`
	Hold      int64     `json:"hold" xorm:"default 0 BIGINT"`
	Unit      int64     `json:"unit" xorm:"default 0 BIGINT"`
	Datetime    time.Time `json:"datetime" xorm:"not null  DATETIME"`
	Lastprice   float64   `json:"lastprice" xorm:"not null default 0 DOUBLE"`
	Chg         float64   `json:"chg" xorm:"not null default 0 DOUBLE"`
	Amp         float64   `json:"amp" xorm:"not null default 0 DOUBLE"`
	Vol         int64     `json:"vol" xorm:"not null default 0 BIGINT"`
	Amount      float64   `json:"amount" xorm:"not null default 0 DOUBLE"`
	Open        float64   `json:"open" xorm:"not null default 0 DOUBLE"`
	High        float64   `json:"high" xorm:"not null default 0 DOUBLE"`
	Low         float64   `json:"low" xorm:"not null default 0 DOUBLE"`
	Preclose    float64   `json:"preclose" xorm:"not null default 0 DOUBLE"`
	Highlimit   float64   `json:"highlimit" xorm:"not null default 0 DOUBLE"`
	Lowlimit   float64   `json:"lowlimit" xorm:"not null default 0 DOUBLE"`
	Bidprice1   float64   `json:"bidprice1" xorm:"not null default 0 DOUBLE"`
	Bidvolume1  float64   `json:"bidvolume1" xorm:"not null default 0 DOUBLE"`
	Askprice1   float64   `json:"askprice1" xorm:"not null default 0 DOUBLE"`
	Askvolume1  float64   `json:"askvolume1" xorm:"not null default 0 DOUBLE"`
	Bidprice2   float64   `json:"bidprice2" xorm:"not null default 0 DOUBLE"`
	Bidvolume2  float64   `json:"bidvolume2" xorm:"not null default 0 DOUBLE"`
	Askprice2   float64   `json:"askprice2" xorm:"not null default 0 DOUBLE"`
	Askvolume2  float64   `json:"askvolume2" xorm:"not null default 0 DOUBLE"`
	Bidprice3   float64   `json:"bidprice3" xorm:"not null default 0 DOUBLE"`
	Bidvolume3  float64   `json:"bidvolume3" xorm:"not null default 0 DOUBLE"`
	Askprice3   float64   `json:"askprice3" xorm:"not null default 0 DOUBLE"`
	Askvolume3  float64   `json:"askvolume3" xorm:"not null default 0 DOUBLE"`
	Bidprice4   float64   `json:"bidprice4" xorm:"not null default 0 DOUBLE"`
	Bidvolume4  float64   `json:"bidvolume4" xorm:"not null default 0 DOUBLE"`
	Askprice4   float64   `json:"askprice4" xorm:"not null default 0 DOUBLE"`
	Askvolume4  float64   `json:"askvolume4" xorm:"not null default 0 DOUBLE"`
	Bidprice5   float64   `json:"bidprice5" xorm:"not null default 0 DOUBLE"`
	Bidvolume5  float64   `json:"bidvolume5" xorm:"not null default 0 DOUBLE"`
	Askprice5   float64   `json:"askprice5" xorm:"not null default 0 DOUBLE"`
	Askvolume5  float64   `json:"askvolume5" xorm:"not null default 0 DOUBLE"`
	aaaa        float64   `xorm:"-"`
}
func NewOptionTick()*Optiontick{
	return &Optiontick{}
}
func (Op *Optiontick)Save()error{
	if strings.TrimSpace(Op.Code)==""||Op.Lastprice==0{
		return errors.New("最新价为空")
	}
	sql:=fmt.Sprintf(`insert into Optiontick(Code,Name,Exchange,Underlying,Strikeprice,Flag,Expire,Remain,Hold,Unit,Datetime,Lastprice,Chg,Amp,Vol,Amount,Open,High,Low,Preclose,Highlimit,Lowlimit,Bidprice1,Bidvolume1,Askprice1,Askvolume1,Bidprice2,Bidvolume2,Askprice2,Askvolume2,Bidprice3,Bidvolume3,Askprice3,Askvolume3,Bidprice4,Bidvolume4,Askprice4,Askvolume4,Bidprice5,Bidvolume5,Askprice5,Askvolume5) values ('%s','%s','%s','%s',%f,'%s','%s',%d,%d,%d,'%s',%f,%f,%f,%d,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f) on conflict(Code) do update set Name=EXCLUDED.Name,Exchange=EXCLUDED.Exchange,Underlying=EXCLUDED.Underlying,Strikeprice=EXCLUDED.Strikeprice,Flag=EXCLUDED.Flag,Expire=EXCLUDED.Expire,Remain=EXCLUDED.Remain,Hold=EXCLUDED.Hold,Unit=EXCLUDED.Unit,Datetime=EXCLUDED.Datetime,Lastprice=EXCLUDED.Lastprice,Chg=EXCLUDED.Chg,Amp=EXCLUDED.Amp,Vol=EXCLUDED.Vol,Amount=EXCLUDED.Amount,Open=EXCLUDED.Open,High=EXCLUDED.High,Low=EXCLUDED.Low,Preclose=EXCLUDED.Preclose,Highlimit=EXCLUDED.Highlimit,Lowlimit=EXCLUDED.Lowlimit,Bidprice1=EXCLUDED.Bidprice1,Bidvolume1=EXCLUDED.Bidvolume1,Askprice1=EXCLUDED.Askprice1,Askvolume1=EXCLUDED.Askvolume1,Bidprice2=EXCLUDED.Bidprice2,Bidvolume2=EXCLUDED.Bidvolume2,Askprice2=EXCLUDED.Askprice2,Askvolume2=EXCLUDED.Askvolume2,Bidprice3=EXCLUDED.Bidprice3,Bidvolume3=EXCLUDED.Bidvolume3,Askprice3=EXCLUDED.Askprice3,Askvolume3=EXCLUDED.Askvolume3,Bidprice4=EXCLUDED.Bidprice4,Bidvolume4=EXCLUDED.Bidvolume4,Askprice4=EXCLUDED.Askprice4,Askvolume4=EXCLUDED.Askvolume4,Bidprice5=EXCLUDED.Bidprice5,Bidvolume5=EXCLUDED.Bidvolume5,Askprice5=EXCLUDED.Askprice5,Askvolume5=EXCLUDED.Askvolume5`,Op.Code,Op.Name,Op.Exchange,Op.Underlying,Op.Strikeprice,Op.Flag,utils.Time2Str(Op.Expire),Op.Remain,Op.Hold,Op.Unit,utils.Time2Str(Op.Datetime),Op.Lastprice,Op.Chg,Op.Amp,Op.Vol,Op.Amount,Op.Open,Op.High,Op.Low,Op.Preclose,Op.Highlimit,Op.Lowlimit,Op.Bidprice1,Op.Bidvolume1,Op.Askprice1,Op.Askvolume1,Op.Bidprice2,Op.Bidvolume2,Op.Askprice2,Op.Askvolume2,Op.Bidprice3,Op.Bidvolume3,Op.Askprice3,Op.Askvolume3,Op.Bidprice4,Op.Bidvolume4,Op.Askprice4,Op.Askvolume4,Op.Bidprice5,Op.Bidvolume5,Op.Askprice5,Op.Askvolume5)
	_,err:=pub.DB.Exec(sql)
	if err!=nil{
		log.Logger.Error("通达信接口分时数据插入失败，错误代码="+err.Error())
		return err
	}
	return nil
}
func (Op *Optiontick)Tick(code string)*Optiontick{
	t :=NewOptionTick()
	pub.DB.SQL("SELECT * FROM OPTIONTICK WHERE CODE=? ORDER BY DATETIME DESC LIMIT 1",code).Get(t)
	if strings.TrimSpace(t.Code)==strings.TrimSpace(code){
		return t
	}else {
		return nil
	}

}
func (Op *Optiontick)Ticks(codes string)[]*Optiontick{
	var pjcodes string
	var sql string
	var buffer bytes.Buffer
	for _,code :=range strings.Split(codes,","){
		buffer.WriteString("'")
		buffer.WriteString(code)
		buffer.WriteString("'")
		buffer.WriteString(",")
	}
	if strings.HasSuffix(buffer.String(),","){
		pjcodes=strings.TrimSuffix(buffer.String(),",")
	}
	ops:=make([]*Optiontick,0)
	if strings.ToUpper(strings.TrimSpace(codes))=="ALL"{
		sql=fmt.Sprintf("SELECT * FROM optiontick")
	}else{
		sql=fmt.Sprintf("SELECT * FROM  optiontick WHERE  code in(%s)",pjcodes)
	}
	if err:=pub.DB.SQL(sql).Find(&ops);err!=nil{
		return nil
	}
	return ops
}

func (Op *Optiontick)TicksByMonth(month string,underlying string)[]*Optiontick{
	var sql string
	if strings.HasPrefix(month,"0"){
		month=strings.TrimPrefix(month,"0")
	}
	ops:=make([]*Optiontick,0)
	sql=fmt.Sprintf("SELECT * FROM  OPTIONTICK WHERE extract(month from expire)=%s and underlying='%s'",month,underlying)
	if err:=pub.DB.SQL(sql).Find(&ops);err!=nil{
		return nil
	}
	return ops
}

func (Op *Optiontick)TickUnderlying(code string)*Optiontick{
	var sql string
	tick:=new(Optiontick)
	sql=fmt.Sprintf("SELECT * FROM OPTIONTICK WHERE CODE='%s' ORDER BY DATETIME DESC LIMIT 1",code)
	if ok,err:=pub.DB.SQL(sql).Get(tick);err!=nil&&ok{
		return nil
	}
	return tick
}



//
//func main(){
//	tdxclient.Tick("10001994,10001995,10001996")
//	tdxclient.Bar("60min")
//	xx:=Getoptiontick("10001994,10001995,10001996")
//	yy,_:=json.Marshal(xx)
//	fmt.Println(string(yy))
//	x:=NewOptionTick()
//	ctx:=context.Background()
//	ctx,cancel:=context.WithCancel(ctx)
//	defer cancel()
//	go func(){
//		for{
//			y:=x.Tick("10001994")
//			jsondata,_:=json.Marshal(y)
//			pub.Publish("tick",string(jsondata))
//			time.Sleep(time.Second*1)
//		}
//	}()
//
//	go pub.Subscribe(ctx,func(msg  *redis.Message){
//		fmt.Println(msg.Payload,msg.Channel,"成功测试")
//	},"tick")
//	select{
//
//	}
//
//}