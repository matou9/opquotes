package quote

import (
	"errors"
	"log"
	"opquotes/config"
	"strings"
	"sync"
	"time"
)

// PoolOptions PoolOptions
type PoolOptions struct {
	ServerList       []string `json:"serverList"`
	ExternServerList []string `json:"externServerList"`
}

// Pool Pool
type TdxPool struct {
	clients   sync.Map//[]*SyncQuoteClient
	exclients sync.Map//[]*SyncExternClient
	mu        sync.RWMutex
	Status    int32 //0:代表未运行 1代表运行
	Tdx   *Tdx
}
func GetMapCount(src sync.Map)int64{
	var count int64 =0
	src.Range(func(k,v interface{})bool{
		count+=1
		return true
	})
	return count
}
// NewPool NewPool
func NewTdxPool(tdx *Tdx) *TdxPool {
	var host string
	//var Type int
	if tdx==nil{
		log.Panic("传入的tdx地址为空，不能生成tdxpool对象")
	}
	p := &TdxPool{}
	op := &PoolOptions{ServerList: make([]string, 0), ExternServerList: make([]string, 0)}
	op.ServerList = config.GetStringList("tdx.serverList")
	op.ExternServerList = config.GetStringList("tdx.externserverList")
	for i := range op.ServerList {
		cl, err := NewSyncQuoteClient(op.ServerList[i], time.Second*1,tdx,nil,nil)
		if err != nil {
			log.Println(err)
		}
		p.clients.Store(cl.host,cl)
	}
	for i := range op.ExternServerList {
		server:=strings.Split(op.ExternServerList[i],"|")
		host=server[0]
		/*if len(server)==2{
			Type=utils.ParseInt(server[1])
		}*/
		cl, err := NewSyncExternClient(host, time.Second*1,tdx,nil,nil)
		if err != nil {
			log.Println(err)
		}
		p.exclients.Store(cl.host,cl)
	}

	p.Status=1
	go p.checker()//这里发现异常，需要处理
	return p
}
func (p *TdxPool)Stop(){
	p.clients.Range(func(k,v interface{})bool{
		if c,ok:=v.(*SyncQuoteClient);ok{
			if c.conn!=nil{
				c.conn.Close()
			}
		}
		return true
	})
	p.exclients.Range(func(k,v interface{})bool{
		if c,ok:=v.(*SyncExternClient);ok{
			if c.conn!=nil{
				c.conn.Close()
			}
		}
		return true
	})
}
func (p *TdxPool) checker() {
	timer := time.NewTicker(time.Second * 30)
	defer timer.Stop()
	for {
		var clientCount, exClientCount, failedClientCount, failedExClientCount int
		select {
		case <-timer.C:
			if p.Status==0{
				return
			}
			p.clients.Range(func(k,v interface{})bool{
				if c,ok:=v.(*SyncQuoteClient);ok{
					if c.GetReady()==NotReady {
						c.ReConnected(time.Second)
						log.Println("检测行情服务器，地址=",c.host,c.ready)
						failedClientCount++
					}
				}
				return true
			})
			p.exclients.Range(func(k,v interface{})bool{
				if c,ok:=v.(*SyncExternClient);ok{
					if c.GetReady()==NotReady {//&&c.GetStatus()==0
						c.ReConnected(time.Second)
						log.Println("检测扩展行情服务器，地址=",c.host,c.ready)
						failedClientCount++
					}
				}
				return true
			})

			log.Print("有效quoteclient数量=",GetMapCount(p.clients))
			log.Print("有效extquoteclient数量=",GetMapCount(p.exclients))
		}
		log.Printf("检查连接结果, 普通连接[%d]个,失败[%d]个, 扩展连接[%d]个, 失败[%d]个", clientCount, failedClientCount, exClientCount, failedExClientCount)
	}
}

// GetExternClient get extern client
func (p *TdxPool) GetExternClient(Type int) (*SyncExternClient, error) {
	var extclient *SyncExternClient
	p.exclients.Range(func(k,v interface{})bool{
		if c,ok:=v.(*SyncExternClient);ok{
			if c.GetReady()==Ready {
				extclient=c
				return false
			}
		}
		return true
	})
	if extclient==nil{
		log.Print("没有找到合适的客户端")
		return nil, errors.New("无有效的扩展连接客户端返回")
	}
	return extclient,nil

}

// GetQuoteClient get quote client
func (p *TdxPool) GetQuoteClient() (*SyncQuoteClient, error) {
	var client *SyncQuoteClient
	p.clients.Range(func(k,v interface{})bool{
		if c,ok:=v.(*SyncQuoteClient);ok{
			if c.GetReady()==Ready{
				client=c
				return false
			}
		}
		return true
	})
	if client==nil{
		return nil, errors.New("no ex client valid")
	}
	return client,nil
}
