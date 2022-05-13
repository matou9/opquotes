package quote

import (
	"fmt"
	"opquotes/basedata"
	"opquotes/pub"
	util "opquotes/utils"
	"strings"
	"time"
)

var (
	StockMgr = basedata.NewStockMgr()
	Klinemap = map[string]TdxKlineType{"1MIN": TdxKlineType_1MIN, "5MIN": TdxKlineType_5MIN, "15MIN": TdxKlineType_15MIN,
		"30MIN": TdxKlineType_30MIN, "60MIN": TdxKlineType_1HOUR, "DAY": TdxKlineType_DAILY, "WEEK": TdxKlineType_WEEKLY,
		"MONTH": TdxKlineType_MONTHLY, "YEAR": TdxKlineType_YEARLY, "JIDU": TdxKlineType_3MONTH}

)

/*
[]*ReqGetInstrumentQuote{&ReqGetInstrumentQuote{Market: 0, Code: "000001"},
                &ReqGetInstrumentQuote{Market: 0, Code: "000002"}}
 */
type TdxStock struct{
	Pool *TdxPool
}

func (tdx *TdxStock)BarStock(code string,count uint16,ktype string) {
	var (
		tablename string
		symbol string
		category uint16
		_market uint16
		start uint16
		maxcount uint16 =800

	)
	switch ktype {
	case "1min":
		tablename = "min1"
		category = 7
	case "5min":
		tablename = "min5"
		category = 0
	case "15min":
		tablename = "min15"
		category = 1
	case "30min":
		tablename = "min30"
		category = 2
	case "60min":
		tablename="min60"
		category = 3
	case "day":
		tablename="day"
		category = 4
	case "week":
		tablename="week"
		category = 5
	}
	Kline_out_chan := make(chan *[]*SecurityBar, 0)
	if count>maxcount{
		count=maxcount
	}
	if strings.Contains(code,"."){
		symbol = strings.Split(code,".")[0]
		exchange :=strings.Split(code,".")[1]
		if strings.ToUpper(exchange)=="SH"{
			_market=1
		}else if strings.ToUpper(exchange)=="SZ"{
			_market=0
		}
	}else{
		symbol=code
		if strings.HasPrefix(code,"6"){
			_market=1
		}else{
			_market=0
		}
	}


	if client, err := tdx.Pool.GetQuoteClient(); err == nil {
			go client.ReqGetSecurityBars(category, _market, symbol, start, count,nil)
	}
	for {
		select {
		case kline := <-Kline_out_chan:
			{
				for _, k := range *kline {
					bar := &basedata.Barstock{}
					bar.Code = k.Code
					if k.Market==0{
						bar.Exchange="SZ"
					}else{
						bar.Exchange="SH"
					}
					bar.Close = k.Close
					v, err := time.Parse("20060102 15:04", fmt.Sprintf("%04d%02d%02d %02d:%02d", k.Year, k.Mon, k.Day, k.Hour, k.Minute))
					// log.Println("扩展行情查询K线", y, m, d, h, min, kline, num, i, err, v)
					if err == nil {
						bar.Datetime = util.Stamp2Time(v.Unix() - 3600*8)
					}
					bar.High = float64(k.High)
					bar.Open = float64(k.Open)
					bar.Low = float64(k.Low)
					bar.Vol = int64(k.Vol)
					bar.Amount=float64(k.DBVol)
					bar.Type = 3
					info,err:=pub.DB.Table(tablename).Insert(bar)
					fmt.Println(info,err)
				}

			}
		}
	}

}

func (tdx *TdxStock)TickStock() {
	Tick_out_chan := make(chan []*SecurityQuote, 1)
	var req []*ReqGetInstrumentQuote
	var num int64
	var _market uint8

	go func(){
		num=0
		for{
			if len(req)>0{
				req = req[0:0]
			}
			for _, contract := range StockMgr.Contracts {
				if num%80 == 0 && num != 0 {
					if client, err := tdx.Pool.GetQuoteClient(); err == nil {
						var _req =make([]*ReqGetInstrumentQuote,len(req))
						copy(_req,req)
						client.ReqGetSecurityQuotes(_req)
					}
					req = req[0:0] //清空数据，最大只能一次80条
				}
				if strings.ToUpper(contract.Exchange) == "SH" {
					_market = 1
				} else if strings.ToUpper(contract.Exchange) == "SZ" {
					_market = 0
				}
				req = append(req, &ReqGetInstrumentQuote{Market: _market, Code: contract.Code})
				num += 1
			}
		}

	}()


	go func() {
		var i int =0
		fmt.Println("开始时间=",time.Now())
		for {
			select {

			case tick := <-Tick_out_chan:
				{

					for _, t := range tick {
						Stocktick, _ := toTick(t)
						i+=1
						_,err:=pub.DB.Table("tick").Insert(&Stocktick)
						if err!=nil{
							fmt.Println("发现冲突异常")
						}
						fmt.Println(time.Now())
					}

				}
			}
		}
	}()


}


