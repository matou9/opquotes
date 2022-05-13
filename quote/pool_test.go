package quote

import (
	"fmt"
	"opquotes/pub"
	"testing"
	"time"
)
var (
	InitialCap = 5
	MaxIdleCap = 10
	MaximumCap = 100
	network    = "tcp"
	address    = "127.0.0.1:7777"
	params =make(map[string]string,0)
	host string
	timeout time.Duration
	//factory    = func() (interface{}, error) { return net.Dial(network, address) }

	//factory = func(v interface{}) (interface{}, error){
	//	switch v.(type) {
	//	case string:
	//		if conn, err := NewSyncExternClient(tdx, s, time.Second*1, 0); err == nil {
	//			return conn, err
	//		}
	//	}
	//	return nil,errors.New("生成tdx连接对象失败")
	//}

	closeFac = func(v interface{}) error {
		nc := v.(*SyncExternClient)
		return nc.conn.Close()
	}


)
func Test_Pool1(t *testing.T){
	var x int
	defer func(){
		if x==1{
			fmt.Print("直接返回")
			return
		}
		fmt.Println("没有返回")
	}()
	x=1
}
func Test_Pool(t *testing.T){

	ctx := pub.NewContext()
	tushare:=&Tushare{}
	tdx:=NewTdx(ctx)
	OptionMgr = NewOptions(ctx)
	OptionMgr.QuoteFunc=tushare.GetDayData
	OptionMgr.Init()

	cext := NewPoolConfig(ctx,tdx,"EXT")
	cext.InitialCap = InitialCap
	cext.IdleTimeout = time.Second*60
	cext.Factory=cext.CreateConn
	cext.Close=cext.CloseConn
	cext.Ping=cext.PingConn
	if p, err := NewChannelPool(cext); err == nil {
		tdx.ExtPool=p
		go func(){
			for{
				begin:=time.Now()
				//tdx.GetMinData()
				////time.Sleep(time.Minute*2)
				////tdx.GetTickData()
				////time.Sleep(time.Minute*2)
				//tdx.GetTransDetailData(10)
				tdx.GetBarData("1min",300)
				fmt.Println("用时=",time.Now().Sub(begin).Seconds())
			}
		}()

	}
	cstd := NewPoolConfig(ctx,tdx,"STD")
	cstd.InitialCap = InitialCap
	cstd.IdleTimeout = time.Second*60
	cstd.Factory=cstd.CreateConn
	cstd.Close=cstd.CloseConn
	cstd.Ping=cstd.PingConn
	if p, err := NewChannelPool(cstd); err == nil {
		tdx.StdPool=p
		go func(){
			for{
				//tdx.UnderlyingTick()
				time.Sleep(time.Second)
			}

		}()
	}
	select{

	}

}