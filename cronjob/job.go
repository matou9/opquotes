package cronjob

import (
	"fmt"
	"opquotes/log"
	"opquotes/pub"
	"opquotes/quote"
	"sync/atomic"
	"time"
)
func (f FuncWrap) Run() { f() }
type FuncWrap func()
type CronJob struct{
	Status int64
	Tradetime map[string]string
	Cron *pub.Cron
	TickRunStatus int64
	Ctx *pub.Context
	MarketQuotes *quote.MarketQuotes
	OptionMgr *quote.Options
}
func NewCronJob()*CronJob{
	cj:= &CronJob{Status:0,Tradetime:make(map[string]string),TickRunStatus:0,Ctx:pub.NewContext()}
	cj.MarketQuotes= quote.NewMarketQuotes(cj.Ctx)
	return cj
}
func (c *CronJob)AddCronJob(spec string,job func()){
	c.Cron.AddJob(spec,FuncWrap(job))
}
func (c *CronJob)Start(){
	if atomic.LoadInt64(&c.Status)==0 {
		atomic.AddInt64(&c.Status, 1)
	}
	for _,v:=range c.MarketQuotes.Quote{
		v.Start()
	}
	c.Cron.Start()
	c.RunTick()
}
func (c *CronJob)Stop(){
	if atomic.LoadInt64(&c.Status)==1{
		atomic.AddInt64(&c.Status,-1)
	}
	for _,v:=range c.MarketQuotes.Quote{
		v.Stop()
	}
	c.Cron.Stop()
}
func (c *CronJob)Init(){
	c.Tradetime["EVERY_DAY"]="0 0 9 * * ?"
	c.Tradetime["AM_START"]="0 15 9 * * ?"
	c.Tradetime["PM_START"]="0 0 13 * * ?"
	c.Tradetime["AM_END"]="0 30 11 * * ?"
	c.Tradetime["PM_END"]="0 1 15 * * ?"
	c.Tradetime["SETTLE"]="0 31 16 * * ?"
	if c.OptionMgr==nil{
		log.Logger.Panic("未初始化optionmgr，程序自动关闭!")
	}
	c.Cron=pub.NewCron()
	c.MarketQuotes.CreateTickObj("tdx")
	c.MarketQuotes.CreateTickObj("sina")
	c.MarketQuotes.CreateTickObj("tushare")
	c.OptionMgr.QuoteFunc=c.MarketQuotes.Quote["SINA"].GetTickData
	c.OptionMgr.Init()
	c.FlushRedisDB()
	c.AddCronJob("1MIN",c.RunMin1)
	c.AddCronJob(c.Tradetime["AM_START"],c.RunTick)
	c.AddCronJob(c.Tradetime["SETTLE"],c.Settle)
}
func (c *CronJob)Settle(){
	if Q,ok:=c.MarketQuotes.Quote["TUSHARE"];ok{
		Q.GetBarData("DAY",100)
		c.OptionMgr.Init()
	}
}
func  (c *CronJob)FlushRedisDB(){
	pub.Rds.FlushDB(pub.Ctx)
}
func (c *CronJob)RunTick(){
	timer:=time.NewTicker(time.Second)
	timer2:=time.NewTicker(time.Second*5)
	defer timer.Stop()
	defer timer2.Stop()
	for {
		select {
		case <-timer.C:
			if atomic.LoadInt64(&c.Status)==1{
				if pub.IsTradeDay(time.Now())==true&&pub.IsTradeTime() == true {
					start:=time.Now()
					if Q,ok:=c.MarketQuotes.Quote["SINA"];ok{
						begin:=time.Now()
						Q.GetTickData()
						Q.UnderlyingTick()
						fmt.Println("获取tick数据用时=",time.Now().Sub(begin).Seconds(),"秒")
					}

					fmt.Println("job用时=", time.Now().Sub(start).Seconds())
				}
			}else{
				time.Sleep(time.Second*30)
			}
			case <-timer2.C:
				if atomic.LoadInt64(&c.Status)==1{
					if pub.IsTradeDay(time.Now())==true&&pub.IsTradeTime() == true {
						if Q,ok:=c.MarketQuotes.Quote["TDX"];ok{
							go Q.GetMinData()
							//go Q.GetBarData("1min",5)
							//Q.GetTransDetailData(100)
						}
					}
				}else{
					time.Sleep(time.Second*30)
				}

		case <-c.Ctx.Ctx.Done():
			fmt.Println("tick結束")
			return
		}

	}
}
func (c *CronJob)RunMin1(){
	log.Logger.Info("开始1分钟事件触发")
	if Q,ok:=c.MarketQuotes.Quote["TDX"];ok{
		Q.GetBarData("1min",2)
	}


}

