package quote

import (
	"errors"
	"fmt"
	"opquotes/config"
	log "opquotes/log"
	"opquotes/pub"
	"strings"
	"sync"
	"time"
	//"reflect"
)

var (
	//ErrMaxActiveConnReached 连接池超限
	ErrMaxActiveConnReached = errors.New("MaxActiveConnReached")
)
var (
	//ErrClosed 连接池已经关闭Error
	ErrClosed = errors.New("pool is closed")
)

// Pool 基本方法
type IPool interface {
	Get() (interface{}, error)

	Put(interface{}) error

	PutError(interface{}) error

	Close(interface{}) error

	Release()

	Len() int
}
// Config 连接池相关配置
type Config struct {
	//连接池中拥有的最小连接数
	InitialCap int
	//最大并发存活连接数
	MaxCap int
	//最大空闲连接
	MaxIdle int
	//生成连接的方法
	Factory func(interface{}) (interface{}, error)
	//关闭连接的方法
	Close func(interface{}) error
	//检查连接是否有效的方法
	Ping func(interface{}) error
	//连接最大空闲时间，超过该事件则将失效
	IdleTimeout time.Duration
	Ctx *pub.Context
	tdx *Tdx
	CfgType string
}
func NewPoolConfig(ctx *pub.Context,tdx *Tdx,Type string)*Config{
	conf := &Config{Ctx:ctx,tdx:tdx,IdleTimeout:time.Minute,CfgType:Type}//std,ext
	conf.InitialCap = 100
	conf.IdleTimeout = time.Second*60
	conf.Factory=conf.CreateConn
	conf.Close=conf.CloseConn
	conf.Ping=conf.PingConn
	return conf
}
func (c *Config)GetParams()*PoolOptions{
	p :=&PoolOptions{ServerList: make([]string, 0), ExternServerList: make([]string, 0)}
	p.ServerList = config.GetStringList("tdx.serverList")
	p.ExternServerList = config.GetStringList("tdx.externserverList")
	return p
}
func (c *Config)CloseConn(v interface{}) error{
	if ncext,ok := v.(*SyncExternClient);ok{
		return ncext.conn.Close()
	}else if nc,ok:=v.(*SyncQuoteClient);ok{
		return nc.conn.Close()
	}
	return nil
}
func (c *Config)CreateConn(v interface{})(interface{},error){
	var(
		host string
		p IPool
	)
	if sc,ok:=v.(*SyncExternClient);ok{
		fmt.Println("链接超时，关闭后，重新生成链接")
		host=sc.host
		p=sc.Pool
		conn, err := NewSyncExternClient(host, time.Second*1,c.tdx,c.Ctx,p)
		if err != nil {
			log.Logger.Error(err.Error())
		}else{
			return  &idleConn{conn: conn, t: time.Now()},nil
		}
	}else if sc,ok:=v.(*SyncQuoteClient);ok{
		fmt.Println("链接超时，关闭后，重新生成链接")
		host=sc.host
		p=sc.Pool
		conn, err := NewSyncQuoteClient(host, time.Second*1,c.tdx,c.Ctx,p)
		if err != nil {
			log.Logger.Error(err.Error())
		}else{
			return  &idleConn{conn: conn, t: time.Now()},nil
		}
	}
	return nil,errors.New("生成conn失败")
}
func (c *Config)CreateConns(v interface{})(interface{},error){
	var (
		host string
		tdx *Tdx
		conns chan *idleConn

	)
	if p,ok:=v.(IPool);ok{
		op:=c.GetParams()
		if c.CfgType=="EXT"{
			conns=make(chan *idleConn, len(op.ExternServerList))
			for i := range op.ExternServerList {
				server:=strings.Split(op.ExternServerList[i],"|")
				host=server[0]
				conn, err := NewSyncExternClient(host, time.Second*1,tdx,c.Ctx,p)
				if err != nil {
					log.Logger.Error(err.Error())
				}else{
					conns <- &idleConn{conn: conn, t: time.Now()}
				}
			}
			return conns,nil
		}else if c.CfgType=="STD"{
			conns=make(chan *idleConn, len(op.ServerList))
			for i := range op.ServerList {
				server:=strings.Split(op.ServerList[i],"|")
				host=server[0]
				conn, err := NewSyncQuoteClient(host, time.Second*1,tdx,c.Ctx,p)
				if err != nil {
					log.Logger.Error(err.Error())
				}else{
					conns <- &idleConn{conn: conn, t: time.Now()}
				}
			}
			return conns,nil
		}

	}else{
		log.Logger.Error("传入的IPOOL接口为空")
	}
	return nil,nil
}
func (c *Config)PingConn(v interface{})error{
	const(
		Ready =1
		NotReady=0
		Active = 1
		Free =0
	)
	if c,ok:=v.(*SyncExternClient);ok{
		if c.GetReady()==NotReady{
			return errors.New("该conn未准备好")
		}else{
			return nil
		}

	}else if c,ok:=v.(*SyncQuoteClient);ok{
		if c.GetReady()==NotReady{
			return errors.New("该conn未准备好")
		}else{
			return nil
		}

	}
	return errors.New("传入的对象不是net.Conn")
}
type connReq struct {
	idleConn *idleConn
}

// channelPool 存放连接信息
type channelPool struct {
	mu                       sync.RWMutex
	conns                    chan *idleConn
	factory                  func(interface{}) (interface{}, error)
	close                    func(interface{}) error
	ping                     func(interface{}) error
	idleTimeout, waitTimeOut time.Duration
	maxActive                int
	openingConns             int
	connReqs                 []chan connReq
	Errorchan                chan *ErrorConn
	Ctx                *pub.Context
}

type idleConn struct {
	conn interface{}
	t    time.Time
}
type ErrorConn struct {
	conn interface{}
	t    time.Time
}
// NewChannelPool 初始化连接
func NewChannelPool(poolConfig *Config) (IPool, error) {
	//if !(poolConfig.InitialCap <= poolConfig.MaxIdle && poolConfig.MaxCap >= poolConfig.MaxIdle && poolConfig.InitialCap >= 0) {
	//	return nil, errors.New("invalid capacity settings")
	//}
	if poolConfig.Factory == nil {
		return nil, errors.New("invalid factory func settings")
	}
	if poolConfig.Close == nil {
		return nil, errors.New("invalid close func settings")
	}
	c := &channelPool{
		factory:      poolConfig.Factory,
		close:        poolConfig.Close,
		idleTimeout:  poolConfig.IdleTimeout,
		maxActive:    poolConfig.MaxCap,
		openingConns: poolConfig.InitialCap,
		Errorchan:make(chan *ErrorConn,1),
		Ctx: poolConfig.Ctx,
	}

	if poolConfig.Ping != nil {
		c.ping = poolConfig.Ping
	}
	if v,err:=poolConfig.CreateConns(c);err==nil{
		if conns,ok:=v.(chan *idleConn);ok{
			c.conns=conns
		}
	}
	go c.CheckErrorChan()
	/*for i := 0; i < poolConfig.InitialCap; i++ {
		conn, err := c.factory(poolConfig.Params)
		if err != nil {
			c.Release()
			return nil, fmt.Errorf("factory is not able to fill the pool: %s", err)
		}
		c.conns <- &idleConn{conn: conn, t: time.Now()}
	}
*/
	return c, nil
}

// getConns 获取所有连接
func (c *channelPool) getConns() chan *idleConn {
	c.mu.Lock()
	conns := c.conns
	c.mu.Unlock()
	return conns
}

// Get 从pool中取一个连接

func (c *channelPool) CheckErrorChan(){
	for{
		select{
		case wrapConn,ok:=<-c.Errorchan:
			if ok{
				log.Logger.Error("捕捉到错误的error chan")
				c.Close(wrapConn.conn)
				if v,err:=c.factory(wrapConn.conn);err==nil&&v!=nil{
					if newConn,ok:=v.(*idleConn);ok{
						fmt.Println("开始处理出现错误的chan")
						c.Put(newConn.conn)
						log.Logger.Panic("捕捉到错误的error chan，程序终止")
					}
				}
			}
		case <-c.Ctx.Ctx.Done():
				return
		}
	}
}
func (c *channelPool) Get() (interface{}, error) {
	conns := c.getConns()
	if conns == nil {
		return nil, ErrClosed
	}
	for {
		select {
		case wrapConn := <-conns:
			if wrapConn == nil {
				return nil, ErrClosed
			}
			//判断是否超时，超时则丢弃
			if timeout := c.idleTimeout; timeout > 0 {
				if wrapConn.t.Add(timeout).Before(time.Now()) {
					//丢弃并关闭该连接
					c.Close(wrapConn.conn)
					if v,err:=c.factory(wrapConn.conn);err==nil&&v!=nil{
						if newConn,ok:=v.(*idleConn);ok{
							return newConn.conn,nil
						}
					}
					continue
				}
			}
			//判断是否失效，失效则丢弃，如果用户没有设定 ping 方法，就不检查
			if c.ping != nil {
				if err := c.ping(wrapConn.conn); err != nil {
					c.Close(wrapConn.conn)
					continue
				}
			}
			return wrapConn.conn, nil
		case <-time.After(time.Second*10):
			return nil,errors.New("连接池获取链接超时")
		case <-c.Ctx.Ctx.Done():
				return nil,errors.New("被主程序CTX强制退出")
		/*default:
			c.mu.Lock()
			log.Logger.Debug("openConn %v %v", c.openingConns, c.maxActive)
			if c.openingConns >= c.maxActive {
				req := make(chan connReq, 1)
				c.connReqs = append(c.connReqs, req)
				c.mu.Unlock()
				ret, ok := <-req
				if !ok {
					return nil, ErrMaxActiveConnReached
				}
				if timeout := c.idleTimeout; timeout > 0 {
					if ret.idleConn.t.Add(timeout).Before(time.Now()) {
						//丢弃并关闭该连接
						c.Close(ret.idleConn.conn)
						continue
					}
				}
				return ret.idleConn.conn, nil
			}
			if c.factory == nil {
				c.mu.Unlock()
				return nil, ErrClosed
			}
			conn, err := c.factory()
			if err != nil {
				c.mu.Unlock()
				return nil, err
			}
			c.openingConns++
			c.mu.Unlock()
			return conn, nil*/
		}
	}
}
func (c *channelPool)PutError(errorchan interface{})error{
	if errorchan == nil {
		return errors.New("connection is nil. rejecting")
	}
	if c.Errorchan == nil {
		return errors.New("error chan为空")
	}
	select {
		case c.Errorchan <-&ErrorConn{conn: errorchan, t: time.Now()}:
			fmt.Println("写入错误的chan成功")
			return nil
		case <-c.Ctx.Ctx.Done():
			return nil
		case <-time.After(500 * time.Millisecond):
			return errors.New("PutError写入超时")
	}
}

// Put 将连接放回pool中
func (c *channelPool) Put(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}
	c.mu.Lock()
	if c.conns == nil {
		c.mu.Unlock()
		return c.Close(conn)
	}
	if l := len(c.connReqs); l > 0 {
		req := c.connReqs[0]
		copy(c.connReqs, c.connReqs[1:])
		c.connReqs = c.connReqs[:l-1]
		req <- connReq{
			idleConn: &idleConn{conn: conn, t: time.Now()},
		}
		c.mu.Unlock()
		return nil
	} else {
		select {
		case c.conns <- &idleConn{conn: conn, t: time.Now()}:
			c.mu.Unlock()
			return nil
		default:
			c.mu.Unlock()
			//连接池已满，直接关闭该连接
			fmt.Print("关闭满连接")
			return c.Close(conn)
		}
	}
}

// Close 关闭单条连接
func (c *channelPool) Close(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.close == nil {
		return nil
	}
	c.openingConns--
	return c.close(conn)
}

// Ping 检查单条连接是否有效
func (c *channelPool) Ping(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}
	return c.ping(conn)
}

// Release 释放连接池中所有连接
func (c *channelPool) Release() {
	c.mu.Lock()
	conns := c.conns
	c.conns = nil
	c.factory = nil
	c.ping = nil
	closeFun := c.close
	c.close = nil
	c.mu.Unlock()

	if conns == nil {
		return
	}

	close(conns)
	for wrapConn := range conns {
		//log.Printf("Type %v\n",reflect.TypeOf(wrapConn.conn))
		closeFun(wrapConn.conn)
	}
}

// Len 连接池中已有的连接
func (c *channelPool) Len() int {
	return len(c.getConns())
}