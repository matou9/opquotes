package genk
import(
	"math"
	"opquotes/basedata"
	"opquotes/utils"
	"sync"
	"time"
)
type BarKey struct{
	Code string
	Exchange string
	Datetime time.Time
}
type BarGenMgr struct{
	Bars sync.Map
	BarGen *BarGen
}
func NewBarGenMgr()*BarGenMgr{
	return &BarGenMgr{BarGen:NewBarGen()}
}
func (mgr *BarGenMgr)UpdateTick(t interface{}){
	if tick,ok:=t.(*basedata.Optiontick);ok{
		key:=BarKey{
			Code:     tick.Code,
			Exchange: tick.Exchange,
			Datetime: tick.Datetime,
		}
		if _,ok:=mgr.Bars.Load(key);!ok{
			mgr.BarGen.UpdateTick(t)
		}
	}


}
type BarGen struct{

}

func NewBarGen()*BarGen{
	return &BarGen{}
}
func (b *BarGen)UpdateTick(t interface{}){
	var new_minute bool=false
	var bar *basedata.Bar

	if utils.IsTradeTime(){
		if tick,ok:=t.(*basedata.Optiontick);ok{
			if tick.Lastprice<=0.0{
				return
			}
			if bar.Datetime.Hour()!= tick.Datetime.Hour() || (bar.Datetime.Minute()!= tick.Datetime.Minute())||bar==nil{
				new_minute=true
			}
			if new_minute{
				bar=basedata.NewBar()
				bar.Code=tick.Code
				bar.Exchange=tick.Exchange
				bar.PreClose=tick.Preclose
				bar.LowLimit=tick.Lowlimit
				bar.HighLimit=tick.Highlimit
				bar.Hold=tick.Hold
			}else{
				bar.Close=tick.Lastprice
				bar.High=math.Max(tick.Lastprice,bar.High)
				bar.Low=math.Min(tick.Lastprice,bar.Low)
				bar.HighLimit=tick.Highlimit
				bar.LowLimit=tick.Lowlimit
				bar.Vol+=tick.Vol
				bar.Amount+=tick.Amount
			}
		}

	}

}