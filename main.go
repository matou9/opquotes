package main

import (
	"opquotes/config"
	"opquotes/cronjob"
	"opquotes/quote"
	"opquotes/router"
)

func main(){
	//tushare:=&quote.Tushare{}
	job:=cronjob.NewCronJob()
	quote.OptionMgr=quote.NewOptions(job.Ctx)
	//quote.OptionMgr.QuoteFunc=tushare.GetDayData
	job.OptionMgr=quote.OptionMgr
	job.Init()
	go job.Start()

	r:=router.NewRouter()
	addr:=config.Getstring("server.addr")
	if addr==""{
		addr=":9999"
	}
	r.Run(addr)
	////AddFunc
	//spec := "*/1000 * * * * ?"
	//c.AddFunc(spec, func() {
	//	if _,rdserror:=rds.Publish(ctx.Ctx,"bar","20s").Result();rdserror!=nil{
	//		log.Logger.Panic("redis publish错误,err="+rdserror.Error())
	//	}
	//	if _,rdserror:=rds.Publish(ctx.Ctx,"pickstock","duma|600133.sh,002563.sz").Result();rdserror!=nil{
	//		log.Logger.Panic("redis publish错误,err="+rdserror.Error())
	//	}
	//	log.Logger.Info("cron running:"+"20秒行情")
	//})
	//c.AddFunc("5min", func() {
	//	log.Logger.Info("cron running:"+"5min行情")
	//})
	//
	//c.AddFunc("1min", func() {
	//	log.Logger.Info("cron running:"+"1min行情")
	//})
	//
	//c.AddFunc("15min", func() {
	//	log.Logger.Info("cron running:"+"15min行情")
	//})
	//c.AddFunc("30min", func() {
	//	log.Logger.Info("cron running:"+"30min行情")
	//})
	//c.AddFunc("60min", func() {
	//	log.Logger.Info("cron running:"+"60min行情")
	//})
	//c.AddFunc("120min", func() {
	//	log.Logger.Info("cron running:"+"120min行情")
	//})


}