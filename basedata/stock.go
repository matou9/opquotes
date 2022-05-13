package  basedata
import (
	"errors"
	"opquotes/log"
	"opquotes/pub"
	"time"
)
type ContractKey struct{
	code string
	market string
}
type Contract struct {
	Code             string    `json:"code" xorm:"not null pk unique(pk_tbl_primary_contract) VARCHAR(20)"`
	Name             string    `json:"name" xorm:"not null VARCHAR(20)"`
	Exchange         string    `json:"exchange" xorm:"not null pk unique(pk_tbl_primary_contract) VARCHAR(10)"`
	MarketTplus      int       `json:"market_tplus" xorm:"not null default 1 INTEGER"`
	ContractMultiple float64   `json:"contract_multiple" xorm:"not null default 1 DOUBLE"`
	MarginRate       float64   `json:"margin_rate" xorm:"not null default 1 DOUBLE"`
	StartDate        time.Time `json:"start_date" xorm:"not null DATE"`
	EndDate          time.Time `json:"end_date" xorm:"not null DATE"`
	Type             string    `json:"type" xorm:"not null VARCHAR(10)"`
	Size             int       `json:"size" xorm:"not null default 1 INTEGER"`
	Pricetick        float64   `json:"pricetick" xorm:"not null default 0.0001 DOUBLE"`
	OptionStrike     float64   `json:"option_strike" xorm:"not null default 0 DOUBLE"`
	OptionUnderlying string    `json:"option_underlying" xorm:"VARCHAR(20)"`
	OptionType       string    `json:"option_type" xorm:"default ''E'::character varying' VARCHAR(2)"`
}
type ContractMgr struct{
	Contracts map[ContractKey]Contract
}
func NewStockMgr()*ContractMgr{
	mgr:= &ContractMgr{Contracts:make(map[ContractKey]Contract)}
	mgr.Load()
	return mgr
}

func (cm *ContractMgr)Load()error{
	var contracts []Contract
	defer func(){
		if p:=recover();p!=nil{
			log.Logger.Info("马上即将关闭程序")
		}
	}()
	if cm.Contracts==nil{
		log.Logger.Error("Contracts未初始化，请检查，重新载入")
		panic("即将关闭程序")
	}else{
		contracts=make([]Contract,0) //获取多组值
		pub.DB.Where("type='股票'").OrderBy("CODE").Find(&contracts)
		if len(contracts)>0{
			for _,v :=range contracts{
				key:=&ContractKey{code:v.Code,market:v.Exchange}
				cm.Contracts[*key]=v
			}
		}
	}
	return nil
}
func (cm *ContractMgr)Query(code string,market string)(contract Contract,err error){

	key:=ContractKey{code:code,market:market}
	if symbol,ok:=cm.Contracts[key];ok{
		err=nil
		contract=symbol
	}else{
		err=errors.New("未找到该合约")
	}
	return contract,err
}

type TickStock struct {
	Code       string    `json:"code" xorm:"not null pk unique(pk_tbl_primary_tick) VARCHAR(20)"`
	Exchange   string    `json:"exchange" xorm:"not null pk unique(pk_tbl_primary_tick) VARCHAR(10)"`
	Datetime   time.Time `json:"datetime" xorm:"not null pk unique(pk_tbl_primary_tick) DATETIME"`
	Price      float64   `json:"price" xorm:"not null default 0 DOUBLE"`
	Hold       float64   `json:"hold" xorm:"not null default 0 DOUBLE"`
	Vol        int64     `json:"vol" xorm:"not null default 0 BIGINT"`
	Amount     float64   `json:"amount" xorm:"not null default 0 DOUBLE"`
	Open       float64   `json:"open" xorm:"not null default 0 DOUBLE"`
	High       float64   `json:"high" xorm:"not null default 0 DOUBLE"`
	Low        float64   `json:"low" xorm:"not null default 0 DOUBLE"`
	Preclose   float64   `json:"preclose" xorm:"not null default 0 DOUBLE"`
	Chg        float64   `json:"chg" xorm:"not null default 0 DOUBLE"`
	Amp        float64   `json:"amp" xorm:"not null default 0 DOUBLE"`
	Limitup    float64   `json:"limitup" xorm:"not null default 0 DOUBLE"`
	Limitdown  float64   `json:"limitdown" xorm:"not null default 0 DOUBLE"`
	Bidprice1  float64   `json:"bidprice1" xorm:"not null default 0 DOUBLE"`
	Bidvolume1 float64   `json:"bidvolume1" xorm:"not null default 0 DOUBLE"`
	Askprice1  float64   `json:"askprice1" xorm:"not null default 0 DOUBLE"`
	Askvolume1 float64   `json:"askvolume1" xorm:"not null default 0 DOUBLE"`
	Bidprice2  float64   `json:"bidprice2" xorm:"not null default 0 DOUBLE"`
	Bidvolume2 float64   `json:"bidvolume2" xorm:"not null default 0 DOUBLE"`
	Askprice2  float64   `json:"askprice2" xorm:"not null default 0 DOUBLE"`
	Askvolume2 float64   `json:"askvolume2" xorm:"not null default 0 DOUBLE"`
	Bidprice3  float64   `json:"bidprice3" xorm:"not null default 0 DOUBLE"`
	Bidvolume3 float64   `json:"bidvolume3" xorm:"not null default 0 DOUBLE"`
	Askprice3  float64   `json:"askprice3" xorm:"not null default 0 DOUBLE"`
	Askvolume3 float64   `json:"askvolume3" xorm:"not null default 0 DOUBLE"`
	Bidprice4  float64   `json:"bidprice4" xorm:"not null default 0 DOUBLE"`
	Bidvolume4 float64   `json:"bidvolume4" xorm:"not null default 0 DOUBLE"`
	Askprice4  float64   `json:"askprice4" xorm:"not null default 0 DOUBLE"`
	Askvolume4 float64   `json:"askvolume4" xorm:"not null default 0 DOUBLE"`
	Bidprice5  float64   `json:"bidprice5" xorm:"not null default 0 DOUBLE"`
	Bidvolume5 float64   `json:"bidvolume5" xorm:"not null default 0 DOUBLE"`
	Askprice5  float64   `json:"askprice5" xorm:"not null default 0 DOUBLE"`
	Askvolume5 float64   `json:"askvolume5" xorm:"not null default 0 DOUBLE"`
}

func (tick *TickStock)Save()(err error){
	if tick!=nil{
		_,err=pub.DB.Insert(tick)
	}
	return err
}
type TickstockMgr struct{
	Ticks map[ContractKey]*TickStock
}
func NewTickstockMgr()*TickstockMgr{
	mgr:= &TickstockMgr{Ticks:make(map[ContractKey]*TickStock)}
	return mgr
}
func (tmg *TickstockMgr)Save(){
	for _,tick:=range tmg.Ticks{
		tick.Save()
	}
}

type Barstock struct {
	Code      string    `json:"code" xorm:"not null pk unique() VARCHAR(20)"`
	Datetime  time.Time `json:"datetime" xorm:"not null pk unique() DATETIME"`
	Open      float64   `json:"open" xorm:"not null default 0 DOUBLE"`
	High      float64   `json:"high" xorm:"not null default 0 DOUBLE"`
	Low       float64   `json:"low" xorm:"not null default 0 DOUBLE"`
	Close     float64   `json:"close" xorm:"not null default 0 DOUBLE"`
	Vol       int64     `json:"vol" xorm:"not null default 0 BIGINT"`
	Amount    float64   `json:"amount" xorm:"not null default 0 DOUBLE"`
	AvgPrice  float64   `json:"avg_price" xorm:"not null default 0 DOUBLE"`
	HighLimit float64   `json:"high_limit" xorm:"not null default 0 DOUBLE"`
	LowLimit  float64   `json:"low_limit" xorm:"not null default 0 DOUBLE"`
	PreClose  float64   `json:"pre_close" xorm:"not null default 0 DOUBLE"`
	Exchange  string    `json:"exchange" xorm:"default '''::character varying' VARCHAR(10)"`
	Paused    int       `json:"paused" xorm:"not null default 0 INTEGER"`
	Type      int       `json:"type" xorm:"not null pk unique() INTEGER"`
}
func (bar *Barstock)Save(tablename string)(err error){
	pub.DB.Table(tablename).Insert(bar)
	return err
}
