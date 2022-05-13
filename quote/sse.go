package quote

import (
	"github.com/gocolly/colly"
	"github.com/tidwall/gjson"
	"math"
	"opquotes/basedata"
	"strings"
	"time"
)
var(

)
const optionlisturl = "http://query.sse.com.cn/commonQuery.do?jsonCallBack=jsonpCallback77327&isPagination=true&expireDate=&securityId=&sqlId=SSE_ZQPZ_YSP_GGQQZSXT_XXPL_DRHY_SEARCH_L&pageHelp.pageSize=10000&pageHelp.pageNo=1&pageHelp.beginPage=1&pageHelp.cacheSize=1&pageHelp.endPage=5&_=1531102881526"
type SSE struct{
	colly.Collector
	Ops *Options
}
func NewSSE(ops *Options,options ...func(*colly.Collector))*SSE{
	sse:= &SSE{Collector:*colly.NewCollector(options...)}
	sse.Ops=ops
	return sse
}
func (sse *SSE)Check(){
	//判断是否已经爬虫过，根据date时间判断本地文件和map，两者同步，否则自动爬虫一次
}
func (sse *SSE)GetOptionList(){
	sse.Collector.OnRequest(func(e *colly.Request){
		e.Headers.Set("Accept","*/*")
		e.Headers.Set("Accept-Encoding", "gzip, deflate")
		e.Headers.Set("Referer","http://www.sse.com.cn/assortment/options/disclo/preinfo/")

	})
	sse.Collector.OnResponse(func(e *colly.Response){
		if e.StatusCode==200{
			content :=string(e.Body)
			start:=strings.Index(content,"(")+1
			end :=len(content)-1
			jsonbody:=content[start:end]


			//names:=gjson.Get(jsonbody,"pageHelp.data.#.CONTRACT_SYMBOL")
			//codes:=gjson.Get(jsonbody,"pageHelp.data.#.SECURITY_ID")
			names:="pageHelp.data.#.CONTRACT_SYMBOL"
			codes:="pageHelp.data.#.SECURITY_ID"
			syscodes:="pageHelp.data.#.CONTRACT_ID"
			types:="pageHelp.data.#.CALL_OR_PUT"
			units:="pageHelp.data.#.CONTRACT_UNIT"
			underlyings :="pageHelp.data.#.SECURITYNAMEBYID"
			strikeprices:="pageHelp.data.#.EXERCISE_PRICE"
			expires:="pageHelp.data.#.EXPIRE_DATE"

			nodes:=gjson.GetMany(jsonbody,names,codes,syscodes,types,units,underlyings,strikeprices,expires)
			if len(nodes)<=7{
				return
			}
			for i:=0;i<=len(nodes[0].Array())-1;i++{
				op :=new(basedata.Option)
				op.Code=nodes[1].Array()[i].String()
				op.Name=nodes[0].Array()[i].String()
				op.Syscode=nodes[2].Array()[i].String()
				op.Type=nodes[3].Array()[i].String()
				op.Unit=nodes[4].Array()[i].Int()
				if strings.Contains(nodes[5].Array()[i].String(),"510300"){
					op.Underlying="510300"
				}else if strings.Contains(nodes[5].Array()[i].String(),"510050"){
					op.Underlying="510050"
				}
				//op.Underlying =nodes[5].Array()[i].String()
				op.Strikeprice=nodes[6].Array()[i].Float()
				expire,_:=time.ParseInLocation("20060102",nodes[7].Array()[i].String(),time.Local)
				op.Expire=expire
				op.Remain=int64(math.Ceil(op.Expire.Sub(time.Now()).Hours()/24))
				sse.Ops.Add(op)
			}

		}
	})
	sse.Visit(optionlisturl)
}
//func main(){
//	sse:=NewSSE()
//	sse.OptionList()
//	fmt.Println(sse.Ops.GetOptionCodes())
//
//}