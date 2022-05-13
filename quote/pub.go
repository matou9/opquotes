package quote
import(
	"opquotes/pub"
	"strings"
)
type IQuote interface {
	GetTickData(codes ...string)
	GetMinData(codes ...string)
	GetBarData(ktype string, count int, codes ...string)
	GetTransDetailData(count uint16,codes ...string)
	UnderlyingTick()
	Start()
	Stop()
}

type MarketQuotes struct{
	Ctx *pub.Context
	Quote map[string]IQuote
}
func NewMarketQuotes(ctx *pub.Context)*MarketQuotes{
	return &MarketQuotes{Ctx:ctx,Quote:make(map[string]IQuote)}
}
func (q *MarketQuotes)CreateTickObj(Type string){
	if strings.ToUpper(Type)=="TDX"{
		q.Quote["TDX"]= NewTdx(q.Ctx)
	}else if strings.ToUpper(Type)=="SINA"{
		q.Quote["SINA"]=NewSina(q.Ctx)
	}else if strings.ToUpper(Type)=="TUSHARE"{
		q.Quote["TUSHARE"]=NewTuShare(q.Ctx)
	}
}
