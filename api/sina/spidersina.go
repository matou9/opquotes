package sina

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/gocolly/colly"
	"github.com/tidwall/gjson"
	base "opquotes/basedata"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)
const(
    month_url string =`http://stock.finance.sina.com.cn/futures/api/openapi.php/StockOptionService.getStockName?exchange=null&cate=`
    calllist_url string =`http://hq.sinajs.cn/list=OP_UP_`
	putlist_url string =`http://hq.sinajs.cn/list=OP_DOWN_`
	market_url string=`http://hq.sinajs.cn/list=`
)


var (
	Underlying = map[string][]string {"SH":{"300ETF","50ETF"}}
	Spider *colly.Collector=colly.NewCollector(colly.Async(true))
	Option_month map[string]*hashset.Set
	Option_Codes *hashset.Set =hashset.New()
	reqContext *colly.Context
	lock sync.RWMutex =sync.RWMutex{}
	month_req int = 0
	codes_req int = 0
	pre  bool = false
	tick_chan chan bool=make(chan bool,0)
)

func Get(url string,key string,value interface{}){
	reqContext := colly.NewContext()
	reqContext.Put(key,value)
	Spider.Request("GET",url,nil,reqContext,nil)
}

func init(){
	Option_month=make(map[string]*hashset.Set)
	//Spider.SetRequestTimeout(time.Second*1)
	Spider.OnResponse(ResponseCallback)
	Spider.OnRequest(RequestCallback)



}
func  RequestCallback(e *colly.Request) {
	//e.Ctx.Put("month","getmonth")
	e.Headers.Set("Referer","https://stock.finance.sina.com.cn/option/quotes.html")
	e.Headers.Set("Host","hq.sinajs.cn")
	e.Headers.Set("Connection","keep-alive")
}
func  ResponseCallback(e *colly.Response){
	if e.Ctx.Get("type")=="month" {
		dealmonth(e)
	}else if e.Ctx.Get("type")=="code"{
		dealcode(e)
	} else if e.Ctx.Get("type")=="tick"{
		dealtick(e)
	}

	//fmt.Println("月份=",Option_month["510050"],Option_month)
}
func dealmonth(e *colly.Response){
	lock.Lock()
	defer lock.Unlock()
	month_req-=1
	month:=gjson.Get(string(e.Body),"result.data.contractMonth")
	stockid:=gjson.Get(string(e.Body),"result.data.stockId").String()
	for _,item :=range month.Array(){
		if _,ok:=Option_month[stockid];!ok{
			Option_month[stockid] =hashset.New()
		}
		Option_month[stockid].Add(item.String())
	}
	if month_req==0 {
		fmt.Println("最后记录=",Option_month)
		GetOptionCodes()

	}

}
func dealcode(e *colly.Response){
	lock.Lock()
	defer lock.Unlock()
	codes_req-=1
	content:=string(e.Body)
	reg:=`\"([^\"]*)\"`
	r,_:=regexp.Compile(reg)
	codes:=r.FindAllString(content,-1)
	for _,code :=range strings.Split(codes[0],","){
		if strings.HasPrefix(code,`"`){
			Option_Codes.Add(strings.TrimPrefix(code,`"`))
		}else if strings.HasSuffix(code,`"`){
			Option_Codes.Add(strings.TrimSuffix(code,`"`))
		}else {
			Option_Codes.Add(code)
		}
	}
	Option_Codes.Remove("")
	if codes_req==0{
		tick_chan<-true
		fmt.Println("codes最后一条记录")
		fmt.Println(Option_Codes.Size())
	}




}
func float_64(lastprice string)(price float64){
	if price,err:= strconv.ParseFloat(lastprice, 64);err==nil{
		return price
	}else{
		return 0
	}
}
func int_64(lastprice string)(price int64){
	if price,err:= strconv.ParseInt(lastprice, 10,64);err==nil{
		return price
	}else{
		return 0
	}
}
func  GetOptiontick(){
		Getmonth()
		select{
		case <-tick_chan:
			for{
				codes :=make([]string,0)
				for _,code:=range Option_Codes.Values(){
					codes=append(codes,code.(string))
				}
				url:=market_url+strings.Join(codes,",")
				Get(url,"type","tick")
				time.Sleep(time.Second*1)
			}
		}


}

func Getmonth(){
	for k,v :=range Underlying{
		if k=="SH"{
			for _,etf :=range v{
				visit_url:=strings.Join([]string{month_url,etf},"")
				month_req+=1
				Get(visit_url,"type","month")


			}

		}
	}
}

func GetOptionCodes(){
	for key,values:=range Option_month{
		for _,val :=range values.Values(){

			value:=val.(string)
			call_visit_url :=strings.Join([]string{calllist_url,key,value[2:4],value[5:7]},"")
			put_visit_url :=strings.Join([]string{putlist_url,key,value[2:4],value[5:7]},"")
			codes_req+=2
			Get(call_visit_url,"type","code")
			Get(put_visit_url,"type","code")
			fmt.Println(call_visit_url,put_visit_url)
		}

	}
}

func dealtick(e *colly.Response){
	lock.Lock()
	defer lock.Unlock()
	content:=string(e.Body)
	lines :=strings.Split(content,`;`)
	for _,line:=range lines{
		cols :=strings.Split(line,`=`)
		if len(cols[0])<8{
			continue
		}
		code :=cols[0][(len(cols[0])-8):len(cols[0])]
		tick :=strings.Split(cols[1],`,`)
		optick:=base.Optiontick{}
		optick.Code=code
		optick.Bidvolume1=float_64(tick[0])
		optick.Bidprice1=float_64(tick[1])
		optick.Lastprice=float_64(tick[2])
		optick.Askprice1=float_64(tick[3])
		optick.Askvolume1=float_64(tick[4])
		optick.Hold=int_64(tick[5])
		optick.Chg=float_64(tick[6])
		optick.Strikeprice = float_64(tick[7])
		optick.Preclose=float_64(tick[8])
		optick.Open=float_64(tick[9])
		optick.Highlimit=float_64(tick[10])
		optick.Lowlimit=float_64(tick[11])
		optick.Askprice5 =float_64(tick[12])
		optick.Askvolume5 =float_64(tick[13])
		optick.Askprice4 =float_64(tick[14])
		optick.Askvolume4 =float_64(tick[15])
		optick.Askprice3 =float_64(tick[16])
		optick.Askvolume3 =float_64(tick[17])
		optick.Askprice2 =float_64(tick[18])
		optick.Askvolume2 =float_64(tick[19])
		optick.Askprice1 =float_64(tick[20])
		optick.Askvolume1 =float_64(tick[21])
		optick.Bidprice1 =float_64(tick[22])
		optick.Bidvolume1 =float_64(tick[23])
		optick.Bidprice2 =float_64(tick[24])
		optick.Bidvolume2 =float_64(tick[25])
		optick.Bidprice3 =float_64(tick[26])
		optick.Bidvolume3 =float_64(tick[27])
		optick.Bidprice4 =float_64(tick[28])
		optick.Bidvolume4 =float_64(tick[29])
		optick.Bidprice5 =float_64(tick[30])
		optick.Bidvolume5 =float_64(tick[31])
		datetime,_:=time.ParseInLocation("2006-01-02 15:04:05",tick[32],time.Local)
		optick.Datetime=datetime
		optick.Underlying=tick[36]
		optick.Name=tick[37]
		optick.Amp=float_64(tick[38])
		optick.High=float_64(tick[39])
		optick.Low=float_64(tick[40])
		optick.Vol=int_64(tick[41])
		optick.Amount=float_64(tick[42])
		optick.Flag=tick[45]
		expire,_:=time.ParseInLocation("2006-01-02",tick[46],time.Local)
		optick.Expire=expire
		optick.Remain=int_64(tick[47])
		op:=base.Option{}
		op.Code=optick.Code
		op.Name=optick.Name
		op.Expire=optick.Expire
		op.Remain=optick.Remain
		op.Type=optick.Flag
		op.Underlying=optick.Underlying
		op.Strikeprice=optick.Strikeprice
		optick.Save()
		//pub.DB.Insert(&optick)


	}

}
