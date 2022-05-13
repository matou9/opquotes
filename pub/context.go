package pub

import (
	"context"
	"sync"
	"time"
)
type FuncWrap func()
func (f FuncWrap) Run() { f() }
type TransKey struct{
	Code string
	TimeUnix int64
}
type Context struct {
	// Keys is a key/value pair exclusively for the context of each request.
	Lock sync.RWMutex
	Keys map[string]interface{}
	Transkey sync.Map//线程并发安全
	TransData sync.Map
	Cron *Cron
	Ctx context.Context
	Stop context.CancelFunc
	PrivateCancel map[string]context.CancelFunc
	SyncOrderData chan interface{}
	TdxErrorChan chan  interface{}
}
func NewContext()*Context{
	ctx :=&Context{TdxErrorChan:make(chan  interface{},1)}
	ctx.PrivateCancel = make(map[string]context.CancelFunc)
	ctx.Ctx, ctx.Stop= context.WithCancel(context.Background())
	ctx.Cron= NewCron()
	ctx.Cron.Start()
	return ctx
}

func (c *Context)AddCronJob(spec string,job func()){
	c.Cron.AddJob(spec,FuncWrap(job))
}
func(c *Context)StopCron(){
	c.Cron.Stop()
}




func (c *Context) Set(key string, value interface{}) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[key] = value
}
func (c *Context) Get(key string) (value interface{}, exists bool) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	value, exists = c.Keys[key]
	if exists{
		return value,exists
	}
	return nil,false
}
func (c *Context) GetString(key string) (s string) {
	if val, ok := c.Get(key); ok && val != nil {
		s, _ = val.(string)
	}
	return
}

// GetBool returns the value associated with the key as a boolean.
func (c *Context) GetBool(key string) (b bool) {
	if val, ok := c.Get(key); ok && val != nil {
		b, _ = val.(bool)
	}
	return
}

// GetInt returns the value associated with the key as an integer.
func (c *Context) GetInt(key string) (i int) {
	if val, ok := c.Get(key); ok && val != nil {
		i, _ = val.(int)
	}
	return
}

// GetInt64 returns the value associated with the key as an integer.
func (c *Context) GetInt64(key string) (i64 int64) {
	if val, ok := c.Get(key); ok && val != nil {
		i64, _ = val.(int64)
	}
	return
}

// GetFloat64 returns the value associated with the key as a float64.
func (c *Context) GetFloat64(key string) (f64 float64) {
	var fint int
	var i64 int64
	if val, ok := c.Get(key); ok && val != nil {
		f64, ok = val.(float64)
		if ok{
			return f64
		}
		fint,ok = val.(int)
		if ok{
			return float64(fint)
		}
		i64,ok = val.(int64)
		if ok{
			return float64(i64)
		}
	}
	return
}

// GetTime returns the value associated with the key as time.
func (c *Context) GetTime(key string) (t time.Time) {

	if val, ok := c.Get(key); ok && val != nil {
		t, _ = val.(time.Time)
	}
	return
}

// GetDuration returns the value associated with the key as a duration.
func (c *Context) GetDuration(key string) (d time.Duration) {

	if val, ok := c.Get(key); ok && val != nil {
		d, _ = val.(time.Duration)
	}
	return
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func (c *Context) GetStringSlice(key string) (ss []string) {

	if val, ok := c.Get(key); ok && val != nil {
		ss, _ = val.([]string)
	}
	return
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func (c *Context) GetStringMap(key string) (sm map[string]interface{}) {

	if val, ok := c.Get(key); ok && val != nil {
		sm, _ = val.(map[string]interface{})
	}
	return
}

// GetStringMapString returns the value associated with the key as a map of strings.
func (c *Context) GetStringMapString(key string) (sms map[string]string) {

	if val, ok := c.Get(key); ok && val != nil {
		sms, _ = val.(map[string]string)
	}
	return
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func (c *Context) GetStringMapStringSlice(key string) (smss map[string][]string) {

	if val, ok := c.Get(key); ok && val != nil {
		smss, _ = val.(map[string][]string)
	}
	return
}
/*
定时器样本代码
func crontest(){
	fmt.Println("定时器测试",time.Now())
	fmt.Println("现金=",xm.SlaveMgr.Cash(),"可用资金=",xm.SlaveMgr.Available(),"冻结资金=",xm.SlaveMgr.Frozenmargin(),"总金额=",xm.SlaveMgr.Balance(),"持仓=",xm.SlaveMgr.Positioncost())
}
xm.Ctx.AddCronJob(spec,crontest)//定时器测试
	xm.Ctx.AddCronJob(spec2,crontest2)
https://www.pppet.net/解析定时格式
*/