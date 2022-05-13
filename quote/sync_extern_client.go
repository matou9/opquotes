package quote

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
	"opquotes/pub"
	util "opquotes/utils"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	exGetMarketsCmd            = []byte{0x01, 0x02, 0x48, 0x69, 0x00, 0x01, 0x02, 0x00, 0x02, 0x00, 0xf4, 0x23} // 查询市场列表
	exGetInstrumentBarCountCmd = []byte{0x01, 0x03, 0x48, 0x66, 0x00, 0x01, 0x02, 0x00, 0x02, 0x00, 0xf0, 0x23} // 查询合约数量
	exSetupCmd1                = []byte{0x01, 0x01, 0x48, 0x65, 0x00, 0x01, 0x52, 0x00, 0x52, 0x00, 0x54, 0x24, // 初始化连接
		0x1f, 0x32, 0xc6, 0xe5, 0xd5, 0x3d, 0xfb, 0x41, 0x1f, 0x32, 0xc6, 0xe5, 0xd5, 0x3d, 0xfb, 0x41, 0x1f, 0x32,
		0xc6, 0xe5, 0xd5, 0x3d, 0xfb, 0x41, 0x1f, 0x32, 0xc6, 0xe5, 0xd5, 0x3d, 0xfb, 0x41, 0x1f, 0x32, 0xc6, 0xe5,
		0xd5, 0x3d, 0xfb, 0x41, 0x1f, 0x32, 0xc6, 0xe5, 0xd5, 0x3d, 0xfb, 0x41, 0x1f, 0x32, 0xc6, 0xe5, 0xd5, 0x3d,
		0xfb, 0x41, 0x1f, 0x32, 0xc6, 0xe5, 0xd5, 0x3d, 0xfb, 0x41, 0xcc, 0xe1, 0x6d, 0xff, 0xd5, 0xba, 0x3f, 0xb8,
		0xcb, 0xc5, 0x7a, 0x05, 0x4f, 0x77, 0x48, 0xea}
)

type requestType struct {
	klineType   TdxKlineType
	requestID   int64
	handlerFunc interface{}
}

type tdxResponse struct {
	h    header
	body []byte
}

// SyncExternClient SyncExternClient
type SyncExternClient struct {
	mu             sync.RWMutex
	ready          int32
	host           string
	conn           net.Conn
	marketList     []ExtRspQryMarket
	referenceCount int
	Wg      *sync.WaitGroup
	Tdx      *Tdx
	Status         int32 //1：正在连接中 0：空闲
	Pool    IPool
	Ctx    *pub.Context
}
func NewSyncExternClient(host string, timeout time.Duration,tdx *Tdx,Ctx *pub.Context,Pool IPool) (*SyncExternClient, error) {
	c := &SyncExternClient{host: host,Pool:Pool,Ctx:Ctx}
	conn, err := net.DialTimeout("tcp", c.host, timeout)
	if err != nil {
		log.Printf("[%s] 服务器连接失败 [%v]", c.host, err)
		return c, err
	}
	_, err = conn.Write(exSetupCmd1)
	if err != nil {
		log.Println("获取扩展行情发生错误=",err)
		return c, err
	}
	c.conn = conn
	resp, err := read(conn)
	if err != nil {
		log.Println("获取扩展行情发生错误=",err)
		return c, err
	}
	l, err := c.GetMarketList()
	if err != nil {
		log.Println("获取扩展行情发生错误=",err)
		return c, err
	}
	log.Printf("[%s]连接成功,funcID:[0x%x], %v", c.host, resp.h.F2, err)
	c.marketList = l
	c.Tdx=tdx
	c.SetReady(Ready)
	return c, nil
}
/*func (c *SyncExternClient)SetStatus(flag int32){
	atomic.StoreInt32(&c.Status,flag)
}
func (c *SyncExternClient)GetStatus()int32{
	return atomic.LoadInt32(&c.Status)
}*/
func (c *SyncExternClient)SetReady(flag int32){
	atomic.StoreInt32(&c.ready,flag)
}
func (c *SyncExternClient)GetReady()int32{
	return atomic.LoadInt32(&c.ready)
}
func (c *SyncExternClient)ReConnected(timeout time.Duration)error{
	var(
		conn net.Conn
		err error
	    l	[]ExtRspQryMarket
		resp *tdxResponse
	)
	defer func(){
		if err!=nil{
			c.SetReady(NotReady)
		}
	}()
	conn, err = net.DialTimeout("tcp", c.host, timeout)
	if err != nil {
		log.Printf("[%s] 我这边连接失败 [%v]", c.host, err)
		return err
	}
	_, err = conn.Write(exSetupCmd1)
	if err != nil {
		return err
	}
	c.conn = conn
	resp, err = read(conn)
	if err != nil {
		return err
	}
	//c.ready = true
	log.Printf("[%s]连接成功,funcID:[0x%x], %v", c.host, resp.h.F2, err)
	l, err = c.GetMarketList()
	if err != nil {
		return err
	}
	c.marketList = l
	c.SetReady(Ready)
	return nil
}
// GetInstrumentCount 合约数量
func (c *SyncExternClient) GetInstrumentCount() (int, error) {
	var(
		err error
		resp *tdxResponse
	)
	defer func(){
		if err!=nil{
			c.SetReady(NotReady)
		}
		c.Pool.Put(c)
	}()
	c.mu.Lock()
	defer c.mu.Unlock()
	// 前四个字段等于F2 小端
	if c.conn==nil{
		return 0,errors.New("conn无效")
	}
	c.conn.Write(exGetInstrumentBarCountCmd)
	resp, err = read(c.conn)
	if resp!=nil{
		if resp.h.F2 != 0x66480301 {
			return 0, errors.New("resp.h.F2 != 0x66480301")
		}
		// 股票数量
		var rsp RspGetInstrumentCount
		err = unmarshal(resp.body[19:], &rsp)
		if err != nil {
			return 0, err
		}
		// log.Printf("[%s]查询合约数量[%d]", c.host, rsp.Count)
		return int(rsp.Count), nil
	}else{
		return 0,errors.New("返回resp为空")
	}

}

func (c *SyncExternClient)IsConnected(timeout time.Duration)error{
	var(
		err error
	)
	defer func(){
		if err!=nil{
			c.SetReady(NotReady)
		}
	}()
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err = net.DialTimeout("tcp", c.host, timeout)
	if err != nil {
		log.Printf("[%s] 连接失败 [%v]", c.host, err)

	}
	return err
}
// GetMarketList 查询市场列表
func (c *SyncExternClient) GetMarketList() ([]ExtRspQryMarket, error) {
	var(
		err error
		resp *tdxResponse
	)
	defer func(){
		if err!=nil{
			c.SetReady(NotReady)
		}
	}()
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err = c.conn.Write(exGetMarketsCmd)
	if err != nil {
		return nil, err
	}
	resp, err = read(c.conn)
	if err != nil {
		return nil, err
	}
	if resp.h.F2 != 1766326801 {
		log.Println("resp.h.F2", resp.h.F2)
		return nil, errors.New("resp.h.F2 != 1766326801")
	}
	var num uint16
	var l []ExtRspQryMarket
	reader := bytes.NewReader(resp.body)
	binary.Read(reader, binary.LittleEndian, &num)
	pos := 2
	for i := 0; i < int(num); i++ {
		var item ExtRspQryMarket
		err = unmarshal(resp.body[pos:], &item)
		item.Name = util.Decode(item.Name)
		item.ShortName = util.Decode(item.ShortName)
		if err != nil {
			return nil, err
		}
		if item.Name != "" {
			l = append(l, item)
		}
		// log.Println("扩展行情查询市场", item, num, i)
		pos += 64
	}
	return l, nil
}

// GetInstrumentInfo 合约信息
func (c *SyncExternClient) GetInstrumentInfo(start uint32, count uint16) ([]RspGetInstrumentInfo, error) {
	var(
		err error
		resp *tdxResponse
	)
	defer func(){
		if err!=nil{
			c.SetReady(NotReady)
		}
		c.Pool.Put(c)
	}()
	c.mu.Lock()
	defer c.mu.Unlock()
	var req ReqGetInstrumentInfo
	req.Start = start
	req.Count = count
	cmd := []byte{0x01, 0x04, 0x48, 0x67, 0x00, 0x01, 0x08, 0x00, 0x08, 0x00, 0xf5, 0x23}
	cmd = append(cmd, marshal(req)...)
	c.conn.SetWriteDeadline(time.Now().Add(time.Second*1))
	_, err = c.conn.Write(cmd)
	if err != nil {
		return nil, err
	}
	resp, err = read(c.conn)
	if err != nil {
		return nil, err
	}
	var l []RspGetInstrumentInfo
	var req2 ReqGetInstrumentInfo
	unmarshal(resp.body, &req2)
	pos := 6
	for i := 0; i < int(req2.Count); i++ {
		var rsp RspGetInstrumentInfo
		err = unmarshal(resp.body[pos:], &rsp)
		rsp.Code = util.Decode(rsp.Code)
		rsp.Name = util.Decode(rsp.Name)
		rsp.F1 = ""
		rsp.F2 = ""
		rsp.Discription = util.Decode(rsp.Discription)
		l = append(l, rsp)
		log.Println(rsp, err, req2.Count, i)
		pos += 64
	}
	return l, nil
}
func (c *SyncExternClient)Connect()(error) {
	var(
		err error
		conn net.Conn
		resp *tdxResponse
	)
	defer func(){
		if err!=nil{
			c.SetReady(NotReady)
		}
	}()
	conn, err = net.DialTimeout("tcp", c.host, time.Second*1)
	if err != nil {
		log.Printf("[%s] net.DialTimeout连接失败 [%v]", c.host, err)
		return  err
	}
	pkgSetupCmd1 := []byte{0x0c, 0x02, 0x18, 0x93, 0x00, 0x01, 0x03, 0x00, 0x03, 0x00, 0x0d, 0x00, 0x01}
	pkgSetupCmd2 := []byte{0x0c, 0x02, 0x18, 0x94, 0x00, 0x01, 0x03, 0x00, 0x03, 0x00, 0x0d, 0x00, 0x02}
	pkgSetupCmd3 := []byte{0x0c, 0x03, 0x18, 0x99, 0x00, 0x01, 0x20, 0x00, 0x20, 0x00, 0xdb, 0x0f, 0xd5,
		0xd0, 0xc9, 0xcc, 0xd6, 0xa4, 0xa8, 0xaf, 0x00, 0x00, 0x00, 0x8f, 0xc2, 0x25, 0x40, 0x13, 0x00, 0x00,
		0xd5, 0x00, 0xc9, 0xcc, 0xbd, 0xf0, 0xd7, 0xea, 0x00, 0x00, 0x00, 0x02}
	_, err = conn.Write(pkgSetupCmd1)
	resp, err = read(conn)
	if err != nil {
		return  err
	}

	_, err = conn.Write(pkgSetupCmd2)
	resp, err = read(conn)
	if err != nil {
		return  err
	}

	_, err = conn.Write(pkgSetupCmd3)
	resp, err = read(conn)
	if err != nil {
		return  err
	}
	c.conn = conn
	log.Printf("[%s]连接成功,funcID:[0x%x], %v", c.host, resp.h.F2, err)
	return nil
}
// GetLastTick 查询行情
func (c *SyncExternClient) GetLastTick(ex, symbol string,wg *sync.WaitGroup) (error) {
	var(
		err error
		resp *tdxResponse
	)
	defer func(){
		wg.Done()
		if err!=nil{
			c.SetReady(NotReady)
			if strings.Contains(err.Error(),"was forcibly closed")==true{
				c.Pool.PutError(c)
				return
			}
		}
		c.Pool.Put(c)
	}()
	defer c.mu.Unlock()
	c.mu.Lock()
	var cmd []byte
	if ex == "INE" {
		return  errors.New("当前TDX的SC合约是用SHFE交易所代码")
	}
	var req ReqGetInstrumentQuote
	req.Code = strings.ToUpper(symbol)
	req.Market = ToTdxMarket(ex)
	cmd = []byte{0x01, 0x01, 0x08, 0x02, 0x02, 0x01, 0x0c, 0x00, 0x0c, 0x00, 0xfa, 0x23}
	cmd = append(cmd, marshal(req)...)
	if c.conn==nil{
		return errors.New("获取GetLastTick的conn无效")
	}
	c.conn.SetWriteDeadline(time.Now().Add(time.Second*1))
	_, err = c.conn.Write(cmd)
	if err != nil {
		log.Println("GetLastTick Write",err)
		return  err
	}

	resp, err = read(c.conn)
	if err != nil {
		log.Println("GetLastTick Read",err)
		//c.ready=false
		return  err
	}
	var rsp RspGetInstrumentQuote //=&RspGetInstrumentQuote{}
	err = unmarshal(resp.body, &rsp)
	if err != nil {
		fmt.Println("gettick err",err)
		return  err
	}
	err=c.Tdx.SaveTick(&rsp)
	if err!=nil{
		fmt.Println("保存tick数据错误",err)
		return err
	}
	return err
}

// GetInstrumentBars 查询K线
func (c *SyncExternClient) GetInstrumentBars(req *ReqGetInstrumentBars,wg *sync.WaitGroup) (error) {
	var(
		err error
		resp *tdxResponse
		v time.Time
	)
	defer func(){
		wg.Done()
		if err!=nil{
			c.SetReady(NotReady)
			if strings.Contains(err.Error(),"was forcibly closed")==true{
				c.Pool.PutError(c)
				return
			}
		}
		c.Pool.Put(c)
	}()
	c.mu.Lock()
	defer c.mu.Unlock()
	klinemap := map[TdxKlineType]string{TdxKlineType_EXHQ_1MIN: "min1" , TdxKlineType_5MIN:"min5" ,  TdxKlineType_15MIN:"min15",
		 TdxKlineType_30MIN: "min30", TdxKlineType_1HOUR:"min60", TdxKlineType_DAILY:"day",  TdxKlineType_WEEKLY:"WEEK",
		TdxKlineType_MONTHLY: "MONTH", TdxKlineType_YEARLY:"YEAR"}

	if req.Code == "" {
		err = errors.New("代码为空无效代码")
		log.Println("出现错误=",err)
		return err
	}
	if req.Category > TdxKlineType_DAILY && req.Category != TdxKlineType_EXHQ_1MIN {
		err = errors.New("不支持的周期")
		log.Println("出现错误=",err)
		return err
	}
	cmd := []byte{0x01, 0x01, 0x08, 0x6a, 0x01, 0x01, 0x16, 0x00, 0x16, 0x00, 0xff, 0x23}
	cmd = append(cmd, marshal(req)...)
	c.conn.SetWriteDeadline(time.Now().Add(time.Second*1))
	_, err = c.conn.Write(cmd)
	if err != nil {
		log.Println("出现错误=",err)
		c.Pool.PutError(c)
		return err
	}
	resp, err = read(c.conn)
	// 前18个字节没解析
	if resp==nil||len(resp.body)<18{
		log.Println("出现错误=","返回的的数据异常")
		return errors.New("返回的的数据异常")
	}
	body := resp.body[18:]
	var num uint16
	reader := bytes.NewReader(body)
	binary.Read(reader, binary.LittleEndian, &num)
	pos := 2
	k:=0
	for i := 0; i < int(num); i++ {
		y, m, d, h, min := getDateTime(req.Category, body[pos:])
		var kline TdxKline
		kline.Code=req.Code
		err := unmarshal(body[pos+4:], &kline)
		if err != nil {
			panic(err)
		}
		if req.Category>3 && req.Category!=7 && req.Category!=8{
			kline.Category=1
		}else{
			kline.Category=0
		}
		v, err = time.ParseInLocation("20060102 15:04", fmt.Sprintf("%04d%02d%02d %02d:%02d", y, m, d, h, min),time.Local)
		kline.Time = v.Unix() //- 3600*8
		//}
		pos += 32
		k+=1

		if tablename,ok:=klinemap[req.Category];ok{
			err=c.Tdx.SaveBar(&kline,tablename)
			if err!=nil{
				fmt.Println("插入bar数据错误=",err)
				return err
			}
		}
	}
	return nil
}

// GetMinuteTimeData 查询分时数据

func (c *SyncExternClient) GetMinuteTimeData(market uint8, code string,wg *sync.WaitGroup)(error) {
	var(
		err error
		resp *tdxResponse
	)
	defer func(){
		wg.Done()
		if err!=nil{
			c.SetReady(NotReady)
			if strings.Contains(err.Error(),"was forcibly closed")==true{
				c.Pool.PutError(c)
				return
			}

		}
		c.Pool.Put(c)
	}()
	c.mu.Lock()
	defer c.mu.Unlock()
	req := ReqGetMinuteTimeData{Market: market, Code: code}
	cmd := []byte{0x01, 0x07, 0x08, 0x00, 0x01, 0x01, 0x0c, 0x00, 0x0c, 0x00, 0x0b, 0x24}
	cmd = append(cmd, marshal(req)...)
	c.conn.SetWriteDeadline(time.Now().Add(time.Second*1))
	_, err = c.conn.Write(cmd)
	if err != nil {
		log.Println("GetMinuteTimeData Write发生错误=",err)
		c.Pool.PutError(c)
		return  err
	}
	resp, err = read(c.conn)
	if err != nil {
		log.Println("GetMinuteTimeData read发生错误=",err)
		return  err
	}
	var rsp RspGetMinuteTimeDataHeader
	err = unmarshal(resp.body, &rsp)
	if err != nil {
		log.Println("GetMinuteTimeData unmarshal发生错误=",err)
		return  err
	}
	pos := 12
	for i := 0; i < int(rsp.Count); i++ {
		var item RspGetMinuteTimeData
		err = unmarshal(resp.body[pos:], &item)
		item.Code=code
		pos += 18
		err=c.Tdx.SaveMindata(&item)
		if err!=nil{
			log.Println("GetMinuteTimeData SaveMindata发生错误=",err)
			return err
		}

		//ret = append(ret, &item)
	}

	//out_chan<-&ret
	return err
}

// GetHistoryMinuteTimeData 查询历史分时
func (c *SyncExternClient) GetHistoryMinuteTimeData(market uint8, code string, date uint32) ([]*RspGetMinuteTimeData, error) {
	var(
		err error
		resp *tdxResponse
	)
	defer func(){
		if err!=nil{
			c.SetReady(NotReady)
			if strings.Contains(err.Error(),"was forcibly closed")==true{
				c.Pool.PutError(c)
				return
			}

		}
		c.Pool.Put(c)
	}()
	c.mu.Lock()
	defer c.mu.Unlock()
	var ret []*RspGetMinuteTimeData
	req := ReqGetHistoryMinuteTimeData{Market: market, Code: code, Date: date}
	cmd := []byte{0x01, 0x01, 0x30, 0x00, 0x01, 0x01, 0x10, 0x00, 0x10, 0x00, 0x0c, 0x24}
	cmd = append(cmd, marshal(req)...)
	c.conn.SetWriteDeadline(time.Now().Add(time.Second*1))
	_, err = c.conn.Write(cmd)
	if err != nil {
		return ret, err
	}
	resp, err = read(c.conn)
	if err != nil {
		return ret, err
	}
	var rsp RspGetHistoryMinuteTimeDataHeader
	err = unmarshal(resp.body, &rsp)
	if err != nil {
		log.Println(err)
		return ret, err
	}
	pos := 20
	for i := 0; i < int(rsp.Count); i++ {
		var item RspGetMinuteTimeData
		err = unmarshal(resp.body[pos:], &item)
		hour := (item.Time / 60)
		minute := (item.Time % 60)

		log.Println("分时", hour, minute, item)
		pos += 18
	}
	return ret, nil
}
// GetTransactionData 查询分笔数据
func (c *SyncExternClient) GetTransactionData(market uint8, code string, start int32, count uint16,wg *sync.WaitGroup) ([]*TransactionData, error) {
	//      pkg = bytearray.fromhex('01 01 08 00 03 01 12 00 12 00 fc 23')
	var(
		err error
		resp *tdxResponse
		exchange string
		s_second,s_minute,s_hour string
	)
	defer func(){
		wg.Done()
		if err!=nil{
			log.Println("GetTransactionData开始获取交易明细数据错误",err)
			c.SetReady(NotReady)
			if strings.Contains(err.Error(),"was forcibly closed")==true{
				c.Pool.PutError(c)
				return
			}
		}
		c.Pool.Put(c)

	}()
	c.mu.Lock()
	defer c.mu.Unlock()
	var ret =make([]*TransactionData,0)
	req := ReqGetTransactionData{Market: market, Code: code, Start: start, Count: count}
	cmd := []byte{0x01, 0x01, 0x08, 0x00, 0x03, 0x01, 0x12, 0x00, 0x12, 0x00, 0xfc, 0x23}
	cmd = append(cmd, marshal(req)...)
	c.conn.SetWriteDeadline(time.Now().Add(time.Second*1))
	_, err = c.conn.Write(cmd)
	if err != nil {
		c.Pool.PutError(c)
		log.Println("GetTransactionData Write发生错误=",err)
		return ret, err
	}
	resp, err = read(c.conn)
	if err != nil {
		log.Println("GetTransactionData read发生错误=",err)
		return ret, err
	}
	var rsp RspGetTransactionDataHeader
	err = unmarshal(resp.body, &rsp)
	if err != nil {
		log.Println(err)
		return ret, err
	}
	pos := 16
	for i := 0; i < int(rsp.Count); i++ {
		var item RspGetTransactionData
		err = unmarshal(resp.body[pos:], &item)
		hour := (item.Time / 60)
		minute := (item.Time % 60)
		second := (item.Direction % 10000)
		if second > 59 {
			second = 0
		}
		exchange="SH"

		trans:= NewTransactionData(code,exchange)
		trans.Volume=int64(item.Volume)
		trans.Price=float64(item.Price)/10000
		trans.Posdir=int64(item.Posdir)
		trans.Direction=int64(item.Direction)
		if second<10 {
			s_second = "0" + util.ToString(second)
		}else{
			s_second=util.ToString(second)
		}
		if minute<10{
			s_minute="0"+util.ToString(minute)
		}else{
			s_minute=util.ToString(minute)
		}
		if hour<10{
			s_hour="0"+util.ToString(hour)
		}else{
			s_hour=util.ToString(hour)
		}

		trans.Datetime=util.Strtotime(util.Time2StrDate(time.Now())+" "+s_hour+":"+s_minute+":"+s_second)
		if trans.Datetime.IsZero(){
			log.Panic("错误，code=",code,trans,hour,minute,second)
		}
		trans.Save()
		ret=append(ret,trans)
		pos += 16
	}
	return ret, nil
}

// GetHistoryTransactionData 查询历史分笔成交
func (c *SyncExternClient) GetHistoryTransactionData(date uint32, market uint8, code string, start int32, count uint16) ([]*RspGetTransactionData, error) {
	var(
		err error
		resp *tdxResponse
	)
	defer func(){
		if err!=nil{
			c.SetReady(NotReady)
			if strings.Contains(err.Error(),"was forcibly closed")==true{
				c.Pool.PutError(c)
				return
			}

		}
		c.Pool.Put(c)
	}()
	c.mu.Lock()
	defer c.mu.Unlock()
	var ret []*RspGetTransactionData
	req := ReqGetHistoryTransactionData{Date: date, Market: market, Code: code, Start: start, Count: count}
	cmd := []byte{0x01, 0x02, 0x30, 0x00, 0x02, 0x01, 0x16, 0x00, 0x16, 0x00, 0x06, 0x24}
	cmd = append(cmd, marshal(req)...)
	c.conn.SetWriteDeadline(time.Now().Add(time.Second*1))
	_, err = c.conn.Write(cmd)
	if err != nil {
		return ret, err
	}
	resp, err = read(c.conn)
	if err != nil {
		return ret, err
	}
	var rsp RspGetTransactionDataHeader
	err = unmarshal(resp.body, &rsp)
	if err != nil {
		log.Println(err)
		return ret, err
	}
	pos := 16
	for i := 0; i < int(rsp.Count); i++ {
		var item RspGetTransactionData
		err = unmarshal(resp.body[pos:], &item)
		hour := (item.Time / 60)
		minute := (item.Time % 60)
		second := (item.Direction % 10000)
		if second > 59 {
			second = 0
		}
		log.Println("历史分笔成交", hour, minute, second, item, i, rsp.Count)
		pos += 16
	}
	return ret, nil
}
