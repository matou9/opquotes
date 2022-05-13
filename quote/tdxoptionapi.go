package quote

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	base "opquotes/basedata"
	"opquotes/log"
	"opquotes/pub"
	"opquotes/utils"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	op        *PoolOptions
	code_chan                = make(chan string, 1)
	Tdx_tick_int int64       =0
)
type Tdx struct{
	StartFlag  int32
	StdPool      IPool
	ExtPool   IPool
	OptionMgr  *Options
	Ctx  *pub.Context
}
func NewTdx(ctx *pub.Context)*Tdx{
	tdx:= &Tdx{StartFlag:0,Ctx:ctx}
	cext := NewPoolConfig(ctx,tdx,"EXT")
	if pext, err := NewChannelPool(cext); err == nil {
		tdx.ExtPool=pext
	}
	cstd := NewPoolConfig(ctx,tdx,"STD")
	if pstd, err := NewChannelPool(cstd); err == nil {
		tdx.StdPool=pstd
	}
	return tdx
}
func (tdx *Tdx)init() {
	tdx.Start()
}
func (tdx *Tdx)Start(){
	if tdx.StartFlag==0{
		log.Logger.Info("开始启动行情服务器连接")
	}
}
func (tdx *Tdx)Stop(){
	//if tdx.Pool==nil{
	//	log.Logger.Panic("连接池未初始化，程序退出")
	//	return
	//}
	//tdx.Pool.mu.Lock()
	//defer tdx.Pool.mu.Unlock()
	//log.Logger.Info("开始关闭行情服务器连接")
	//tdx.Pool.Stop()

}
func (tdx *Tdx)Min(){
	var (
		codelist  []string
		wg sync.WaitGroup
		nmarket uint8 = 8
	)
	if tdx.ExtPool==nil{
		log.Logger.Error("pool为空，程序出错")
		return
	}
	codelist = OptionMgr.GetOptionCodes()

	wg.Add(len(codelist))
	var i int =0
	for _, code := range codelist {
		if v, err := tdx.ExtPool.Get(); err == nil {
			i+=1
			if exclient,ok:=v.(*SyncExternClient);ok{
				go exclient.GetMinuteTimeData(nmarket, code,&wg)
			}
		}
	}
	wg.Wait()
	fmt.Println("分时mindata执行结束")
}
func (tdx *Tdx)UnderlyingTick(){
	if v, err := tdx.StdPool.Get(); err == nil {
		if client,ok:=v.(*SyncQuoteClient);ok{
			 client.UnderlyingTick()
		}
	}
}
func (tdx *Tdx)GetMinData(codes ...string) {
	var (
		codelist  []string
		wg sync.WaitGroup
		nmarket uint8 = 8
		i int
	)
	begin:=time.Now()
	if tdx.ExtPool==nil{
		log.Logger.Panic("连接池未初始化，程序退出")
		return
	}
	if len(codes) == 0 {
		codelist = OptionMgr.GetOptionCodes()
	} else if len(codes) == 1 {
		codelist = strings.Split(codes[0], ",")
	}

	wg.Add(len(codelist))
	for _, code := range codelist {
		if v, err := tdx.ExtPool.Get(); err == nil {
			i+=1
			fmt.Println("获取分时数据第",i,"条记录")
			if exclient,ok:=v.(*SyncExternClient);ok{
				go exclient.GetMinuteTimeData(nmarket, code,&wg)
			}
		}else{
			wg.Done()
		}
	}
	wg.Wait()
	fmt.Println("分时mindata执行结束,用时=",time.Now().Sub(begin).Seconds())
}
func (tdx *Tdx)SavetMindataToRds(min *base.Mindata){
	mindata,_:=json.Marshal(min)
	score:= utils.Int64ToString(min.Datetime.Unix())
	pub.Rds.ZRemRangeByScore(pub.Ctx,"min"+min.Code,score,score)
	pub.Rds.ZAdd(pub.Ctx,"min"+min.Code,&redis.Z{
		Score:  float64(min.Datetime.Unix()),
		Member: mindata,
	})
	//pub.Rds.Set(pub.Ctx,"min:"+min.Code+":"+min.Datetime.Format("2006-01-02 1504"),mindata,time.Hour*5)
	//t.Uninx().Format("2006-01-02 15:04:05")
}

func (tdx *Tdx)SaveMindata(min *RspGetMinuteTimeData)error{


		m := &base.Mindata{}
		m.Code=min.Code
		hour := (min.Time / 60)
		minute := (min.Time % 60)
		hourstr:= strconv.Itoa(int(hour))
		if len(hourstr)==1{
			hourstr="0"+hourstr
		}
		minutestr:=strconv.Itoa(int(minute))
		if len(minutestr)==1{
			minutestr="0"+minutestr
		}
		datestr:=time.Now().Format("2006-01-02")
		datetimestr:=datestr+" "+hourstr+":"+minutestr+":00"

		m.Price = float64(min.Price)
		m.AvgPrice = float64(min.AveragePrice)
		m.Vol = int64(min.Volume)
		m.Amount = float64(min.Amount)
		if op,ok:=OptionMgr.Items[min.Code];ok{
			m.PreClose=op.Preclose
		}else{
			log.Logger.Error("没有发现改代码的昨日收盘价,代码="+op.Code)
			return errors.New("没有发现改代码的昨日收盘价,代码="+op.Code)
		}


		m.Datetime=utils.Str2Time(datetimestr)

		sql:=fmt.Sprintf(`INSERT INTO MINDATA(CODE,DATETIME,PRICE,AVG_PRICE,VOL,AMOUNT,HOLD,PRE_CLOSE) VALUES(%s,'%s',%f,%f,%d,%f,%d,%f) ON CONFLICT(CODE,DATETIME)
						DO UPDATE SET PRICE=EXCLUDED.PRICE,AVG_PRICE=EXCLUDED.AVG_PRICE,VOL=EXCLUDED.VOL,AMOUNT=EXCLUDED.AMOUNT,HOLD=EXCLUDED.HOLD,PRE_CLOSE=EXCLUDED.PRE_CLOSE `,min.Code,
						datetimestr,float64(min.Price),float64(min.AveragePrice),int64(min.Volume),float64(min.Amount),0,m.PreClose)

		_,err:=pub.DB.Exec(sql)
		if err!=nil{
			log.Logger.Error("通达信接口分时数据插入失败，错误代码="+err.Error())
			return err
		}
		tdx.SavetMindataToRds(m)
		return nil
}
func (tdx *Tdx)SavetTickToRds(tick *base.Optiontick){
	Tickdata,_:=json.Marshal(tick)
	pub.Rds.Set(pub.Ctx,tick.Code,Tickdata,time.Hour*12)
}
func (tdx *Tdx)SaveTick(tick *RspGetInstrumentQuote)error{
	optick := &base.Optiontick{}
	optick.Code = tick.Code
	OptionMgr.Convert(optick)
	optick.Low = float64(tick.Low)
	//optick.Amount=int64(tick.F3)
	optick.Hold = int64(tick.Position)
	optick.Vol = int64(tick.Volume)
	optick.High = float64(tick.High)
	optick.Low = float64(tick.Low)
	optick.Open = float64(tick.Open)
	optick.Lastprice = float64(tick.Price)
	optick.Askprice5 = float64(tick.Ask5)
	optick.Askprice4 = float64(tick.Ask4)
	optick.Askprice3 = float64(tick.Ask3)
	optick.Askprice2 = float64(tick.Ask2)
	optick.Askprice1 = float64(tick.Ask1)
	optick.Bidprice1 = float64(tick.Bid1)
	optick.Bidprice2 = float64(tick.Bid2)
	optick.Bidprice3 = float64(tick.Bid3)
	optick.Bidprice4 = float64(tick.Bid4)
	optick.Bidprice5 = float64(tick.Bid5)
	optick.Askvolume5 = float64(tick.AskVolume5)
	optick.Askvolume4 = float64(tick.AskVolume4)
	optick.Askvolume3 = float64(tick.AskVolume3)
	optick.Askvolume2 = float64(tick.AskVolume2)
	optick.Askvolume1 = float64(tick.AskVolume1)
	optick.Bidvolume1 = float64(tick.BidVolume1)
	optick.Bidvolume2 = float64(tick.BidVolume2)
	optick.Bidvolume3 = float64(tick.BidVolume3)
	optick.Bidvolume4 = float64(tick.BidVolume4)
	optick.Bidvolume5 = float64(tick.BidVolume5)
	if optick.Preclose!=0{
		optick.Chg = float64((optick.Lastprice - optick.Preclose) * 100 / optick.Preclose)
	}else{
		optick.Chg=0
	}

	if strings.HasPrefix(tick.Code,"6"){
		optick.Exchange = "SH"
	}else{
		optick.Exchange = "SZ"
	}

	optick.Datetime = time.Now()
	OptionMgr.Ticks.Store(optick.Code,optick)
	tdx.SavetTickToRds(optick)
	err:=optick.Save()
	//_,err:=pub.DB.Insert(optick)
	if err!=nil{
		fmt.Println("保存tick数据到数据库错误=",err)
	}
	return err



}
func (tdx *Tdx)SaveBar(k *TdxKline,tablename string)error{
	var datetime string
	bar := &base.Bar{}
	bar.Code = k.Code
	bar.Close = float64(k.Close)
	if k.Category==1{
		datetime= time.Unix(k.Time, 0).Format("2006-01-02")
	}else if k.Category==0{
		datetime= time.Unix(k.Time, 0).Format("2006-01-02 15:04:05")
	}
	Exchange:="SSE"
	preclose:=OptionMgr.Items[k.Code].Preclose
	sql:=fmt.Sprintf(`INSERT INTO %s(CODE,DATETIME,OPEN,HIGH,LOW,CLOSE,VOL,AMOUNT,HOLD,AVG_PRICE,HIGH_LIMIT,LOW_LIMIT,PRE_CLOSE,EXCHANGE,PAUSED,TYPE)VALUES(%s,'%s',%f,%f,%f,%f,%d,%f,%d,%f,%f,%f,%f,'%s',%d,%d) ON CONFLICT(CODE,DATETIME,TYPE) DO UPDATE SET OPEN=EXCLUDED.OPEN, HOLD=EXCLUDED.HOLD, HIGH=EXCLUDED.HIGH, CLOSE=EXCLUDED.CLOSE, LOW=EXCLUDED.LOW, VOL=EXCLUDED.VOL`,tablename,k.Code,
		datetime,k.Open,k.High,k.Low,k.Close,k.Volume,0.0,k.Position,0.0,0.0,0.0,preclose,Exchange,0,1)
	_,err:=pub.DB.Exec(sql)

	if err!=nil{
		log.Logger.Error("通达信接口bar数据1分钟周期k线插入失败，错误代码="+err.Error())
		return err
	}
	return nil
}
func (tdx *Tdx)GetBarData(ktype string, count int, codes ...string) {

	klinemap := map[string]TdxKlineType{"1MIN": TdxKlineType_EXHQ_1MIN, "5MIN": TdxKlineType_5MIN, "15MIN": TdxKlineType_15MIN,
		"30MIN": TdxKlineType_30MIN, "60MIN": TdxKlineType_1HOUR, "DAY": TdxKlineType_DAILY, "WEEK": TdxKlineType_WEEKLY,
		"MONTH": TdxKlineType_MONTHLY, "YEAR": TdxKlineType_YEARLY, "JIDU": TdxKlineType_3MONTH}
	var (
		codelist  []string
	 	wg sync.WaitGroup
		)
	begin:=time.Now()
	if tdx.ExtPool==nil{
		log.Logger.Panic("连接池未初始化，程序退出")
		return
	}
	if len(codes) == 0 {
		codelist = OptionMgr.GetOptionCodes()
	} else if len(codes) == 1 {
		codelist = strings.Split(codes[0], ",")
	}
	wg.Add(len(codelist))
	for _, code := range codelist {
		req := &ReqGetInstrumentBars{}
		req.Code = code
		req.Market = 8
		if _, ok := klinemap[strings.ToUpper(ktype)]; ok {
			req.Category = klinemap[strings.ToUpper(ktype)]
		}
		req.Start = uint32(0)
		req.Count = uint16(count)
		if v, err := tdx.ExtPool.Get(); err == nil {
			if exclient,ok:=v.(*SyncExternClient);ok{
				go exclient.GetInstrumentBars(req, &wg)
			}
		}else{
			wg.Done()
		}
		//select {
		//// 错误快返回,适用于get接口
		//	case err,ok := <-tdx.Ctx.TdxErrorChan:
		//		if ok{
		//
		//		}
		//	case <-done:
		//	case <-time.After(500 * time.Millisecond):
		//}

		/*
		if exclient, err := tdx.Pool.GetExternClient(0); err == nil {
			i+=1
			go  exclient.GetInstrumentBars(req, &wg)
		}*/
	}
	wg.Wait()
	fmt.Println("bar执行结束,周期=",time.Now(),"用时=",time.Now().Sub(begin).Seconds())

}

func (tdx *Tdx)GetTickData(codes ...string) {
	//defer atomic.AddInt64(&Tdx_tick_int, -1)
	var (
		allcode []string
		wg sync.WaitGroup
		//i int
	)
	begin:=time.Now()
	if tdx.ExtPool==nil{
		log.Logger.Panic("连接池未初始化，程序退出")
		return
	}
	if len(codes) == 0 {
		allcode = OptionMgr.GetOptionCodes()
	} else if len(codes) == 1 {
		allcode = strings.Split(codes[0], ",")
	}
	wg.Add(len(allcode))
	for _, code := range allcode {
		if v, err := tdx.ExtPool.Get(); err == nil {
			//i+=1
			//fmt.Println("第",i,"条记录")
			if exclient,ok:=v.(*SyncExternClient);ok{
				go exclient.GetLastTick("QS", code, &wg)
			}else{
				log.Logger.Error("获取连接池客户端出现错误,错误="+err.Error())
			}
		}else{
			wg.Done()
		}
		/*

		if exclient, err := tdx.Pool.GetExternClient(1); err == nil {
			go exclient.GetLastTick("QS", code, &wg)
		}else{
			log.Logger.Error("获取连接池客户端出现错误,错误="+err.Error())
		}*/
	}
	wg.Wait()
	fmt.Println("tick行情获取结束，用时=",time.Now().Sub(begin).Seconds())
}
func (tdx *Tdx)GetTransDetailData(count uint16,codes ...string){
	var (
		codelist  []string
		wg sync.WaitGroup
		nmarket uint8 = 8
	)
	begin:=time.Now()
	if len(codes) == 0 {
		codelist = OptionMgr.GetOptionCodes()
	} else if len(codes) == 1 {
		codelist = strings.Split(codes[0], ",")
	}

	wg.Add(len(codelist))
	for _, code := range codelist {
		if v, err := tdx.ExtPool.Get(); err == nil {
			if exclient,ok:=v.(*SyncExternClient);ok{
				go exclient.GetTransactionData(nmarket, code,0,count,&wg)
			}
		}else{
			wg.Done()
		}
	}
	wg.Wait()
	fmt.Println("transaction数据获取执行结束",time.Now(),"用时=",time.Now().Sub(begin).Seconds())
}