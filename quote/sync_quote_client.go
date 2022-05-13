package quote

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
	"opquotes/basedata"
	"opquotes/pub"
	util "opquotes/utils"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// SyncQuoteClient SyncQuoteClient
const(
	Ready =1
	NotReady=0
	Active = 1
	Free =0
)
type SyncQuoteClient struct {
	lock sync.RWMutex
	ready          int32  //0:不可用 1：可用
	host           string
	conn           net.Conn
	referenceCount int
	Status         int32 //1：正在连接中 0：空闲
	Wg      *sync.WaitGroup
	Tdx      *Tdx
	Pool    IPool
	Ctx    *pub.Context
}

// NewSyncQuoteClient create sync quote client
func NewSyncQuoteClient(host string, timeout time.Duration,tdx *Tdx,Ctx *pub.Context,Pool IPool) (*SyncQuoteClient, error) {
	var err error
	var conn net.Conn
	var resp *tdxResponse
	var c *SyncQuoteClient
	defer func(){
		if err!=nil{
			c.SetReady(NotReady)
		}else{
			c.SetReady(Ready)
		}
	}()
	c = &SyncQuoteClient{host: host,Pool: Pool,Ctx: Ctx,Tdx: tdx}
	conn, err = net.DialTimeout("tcp", host, timeout)
	if err != nil {
		log.Printf("[%s] NewSyncQuoteClient中dial连接失败 [%v]", host, err)
		return c, err
	}
	pkgSetupCmd1 := []byte{0x0c, 0x02, 0x18, 0x93, 0x00, 0x01, 0x03, 0x00, 0x03, 0x00, 0x0d, 0x00, 0x01}
	pkgSetupCmd2 := []byte{0x0c, 0x02, 0x18, 0x94, 0x00, 0x01, 0x03, 0x00, 0x03, 0x00, 0x0d, 0x00, 0x02}
	pkgSetupCmd3 := []byte{0x0c, 0x03, 0x18, 0x99, 0x00, 0x01, 0x20, 0x00, 0x20, 0x00, 0xdb, 0x0f, 0xd5,
		0xd0, 0xc9, 0xcc, 0xd6, 0xa4, 0xa8, 0xaf, 0x00, 0x00, 0x00, 0x8f, 0xc2, 0x25, 0x40, 0x13, 0x00, 0x00,
		0xd5, 0x00, 0xc9, 0xcc, 0xbd, 0xf0, 0xd7, 0xea, 0x00, 0x00, 0x00, 0x02}
	_, err = conn.Write(pkgSetupCmd1)
	resp, err = read(conn)
	if err != nil {
		return c, err
	}

	_, err = conn.Write(pkgSetupCmd2)
	resp, err = read(conn)
	if err != nil {
		return c, err
	}

	_, err = conn.Write(pkgSetupCmd3)
	resp, err = read(conn)
	if err != nil {
		return c, err
	}
	c.conn = conn
	log.Printf("[%s]连接成功,funcID:[0x%x], %v", host, resp.h.F2, err)
	return c, nil
}

// ReqQryStockCount 查询股票数量
func (c *SyncQuoteClient)GetReady()int32{
	return atomic.LoadInt32(&c.ready)
}
func (c *SyncQuoteClient)SetReady(flag int32){
	atomic.StoreInt32(&c.ready,flag)
}

func (c *SyncQuoteClient)ReConnected(timeout time.Duration)error{
	c.lock.Lock()
	defer c.lock.Unlock()
	conn, err := net.DialTimeout("tcp", c.host, timeout)
	if err != nil {
		log.Printf("[%s] SyncQuoteClient ReConnected连接失败 [%v]", c.host, err)
		c.SetReady(NotReady)
		return err
	}
	pkgSetupCmd1 := []byte{0x0c, 0x02, 0x18, 0x93, 0x00, 0x01, 0x03, 0x00, 0x03, 0x00, 0x0d, 0x00, 0x01}
	pkgSetupCmd2 := []byte{0x0c, 0x02, 0x18, 0x94, 0x00, 0x01, 0x03, 0x00, 0x03, 0x00, 0x0d, 0x00, 0x02}
	pkgSetupCmd3 := []byte{0x0c, 0x03, 0x18, 0x99, 0x00, 0x01, 0x20, 0x00, 0x20, 0x00, 0xdb, 0x0f, 0xd5,
		0xd0, 0xc9, 0xcc, 0xd6, 0xa4, 0xa8, 0xaf, 0x00, 0x00, 0x00, 0x8f, 0xc2, 0x25, 0x40, 0x13, 0x00, 0x00,
		0xd5, 0x00, 0xc9, 0xcc, 0xbd, 0xf0, 0xd7, 0xea, 0x00, 0x00, 0x00, 0x02}
	_, err = conn.Write(pkgSetupCmd1)
	_, err = read(conn)
	if err != nil {
		c.SetReady(NotReady)
		return err
	}

	_, err = conn.Write(pkgSetupCmd2)
	_, err = read(conn)
	if err != nil {
		c.SetReady(NotReady)
		return err
	}

	_, err = conn.Write(pkgSetupCmd3)
	_, err = read(conn)
	if err != nil {
		c.SetReady(NotReady)
		return err
	}
	c.conn = conn
	c.SetReady(Ready)
	return nil
}
func (c *SyncQuoteClient) ReqQryStockCount() (uint16, error) {
	var(
		err error
		resp *tdxResponse
	)
	defer func(){
		if err!=nil{
			c.SetReady(NotReady)
		}
	}()
	c.lock.Lock()
	defer c.lock.Unlock()
	cmd := []byte{0x0c, 0x0c, 0x18, 0x6c, 0x00, 0x01, 0x08, 0x00, 0x08, 0x00, 0x4e, 0x04, 0x00, 0x00, 0x75, 0xc7, 0x33, 0x01}
	if c.conn==nil{
		return 0,errors.New("conn连接无效")

	}
	_, err = c.conn.Write(cmd)
	if err != nil {
		fmt.Println("ReqQryStockCount写入错误=",err)
		return 0, err
	}
	resp, err = read(c.conn)
	if err != nil {
		fmt.Println("ReqQryStockCount读取错误=",err)
		return 0, err
	}
	var rsp RspQryStockCount
	err = unmarshal(resp.body, &rsp)
	return rsp.Count, err
}

// ReqGetSecurityList 查询证券列表
func (c *SyncQuoteClient) ReqGetSecurityList(market, start uint16) ([]*RspQrySecurity, error) {
	var(
		err error
		resp *tdxResponse
	)
	defer func(){
		if err!=nil{
			c.SetReady(NotReady)
		}
	}()
	c.lock.Lock()
	defer c.lock.Unlock()
	var ret []*RspQrySecurity
	buf := &bytes.Buffer{}
	// pkg = bytearray.fromhex(u'0c 01 18 64 01 01 06 00 06 00 50 04')
	cmd := []byte{0x1c, 0x01, 0x18, 0x64, 0x01, 0x01, 0x06, 0x00, 0x06, 0x00, 0x50, 0x04}
	buf.Write(cmd)
	binary.Write(buf, binary.LittleEndian, &market)
	binary.Write(buf, binary.LittleEndian, &start)
	_, err = c.conn.Write(buf.Bytes())
	if err != nil {
		return ret, err
	}
	resp, err = read(c.conn)
	if err != nil {
		return ret, err
	}

	var num uint16
	reader := bytes.NewReader(resp.body)
	binary.Read(reader, binary.LittleEndian, &num)
	pos := 2
	for i := 0; i < int(num); i++ {
		var item RspQrySecurity
		err := unmarshal(resp.body[pos:], &item)
		if err == nil {
			item.Name = util.Decode(item.Name)
		}
		pos += 29
		ret = append(ret, &item)
	}
	return ret, nil
}

// ReqGetSecurityQuotes quotes GetTransactionData(market uint8, code string, start int32, count uint16,wg *sync.WaitGroup) ([]*TransactionData, error)
func (c *SyncQuoteClient) GetTransactionData(market uint8, code string, start int32, count uint16,wg *sync.WaitGroup) ([]*RspGetTransactionData, error) {
	var(
		err error
		resp *tdxResponse
	)
	defer func(){
		wg.Done()
		if err!=nil{
			log.Println("SyncQuoteClient GetTransactionData开始获取交易明细数据错误",err)
			c.SetReady(NotReady)
			if strings.Contains(err.Error(),"was forcibly closed")==true{
				c.Pool.PutError(c)
				return
			}
		}
		c.Pool.Put(c)

	}()
	c.lock.Lock()
	defer c.lock.Unlock()
	var ret []*RspGetTransactionData
	req := ReqGetTransactionData{Market: market, Code: code, Start: start, Count: count}
	cmd := []byte{0x0c, 0x17, 0x08, 0x01, 0x01, 0x01, 0x0e, 0x00, 0x0e, 0x00, 0xc5, 0x0f}
	cmd = append(cmd, marshal(req)...)
	_, err = c.conn.Write(cmd)
	if err != nil {
		return ret, err
	}
	resp, err= read(c.conn)
	if err != nil {
		return ret, err
	}
	var rsp RspGetTransactionDataHeader
	err = unmarshal(resp.body, &rsp)
	if err != nil {
		log.Println(err)
		return ret, err
	}
	log.Println("resp: ", resp)
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
		log.Println("分笔成交", hour, minute, second, item, i, rsp.Count)
		pos += 16
	}
	return ret, nil
}
func (c *SyncQuoteClient) UnderlyingTick(){
	var err error
	var ticks []*SecurityQuote
	defer func(){
		if err!=nil{
			log.Println("SyncQuoteClient GUnderlyingTick开始获取交易明细数据错误",err)
			c.SetReady(NotReady)
		}
	}()
	Underlying:=[]string{"510050","510300"}
	var req =make([]*ReqGetInstrumentQuote,len(Underlying))
	for i,code :=range Underlying{
		req[i]=new(ReqGetInstrumentQuote)
		req[i].Code=code
		req[i].Market=1
	}
	if ticks,err=c.ReqGetSecurityQuotes(req);err!=nil{
		fmt.Println(err)
	}else{
		for _, t := range ticks {
			Stocktick, _ := toTick(t)
			if err:=Stocktick.Save();err!=nil{
				fmt.Println("optiontick保存错误，错误代码=",err)
			}
		}
	}
}
func (c *SyncQuoteClient) ReqGetSecurityQuotes(stockList []*ReqGetInstrumentQuote) ([]*SecurityQuote, error) {
	var(
		err error
		resp *tdxResponse
		ret []*SecurityQuote
	)
	defer func(){
		c.lock.Unlock()
		if err!=nil{
			log.Println("SyncQuoteClient ReqGetSecurityQuotes开始获取交易明细数据错误",err)
			c.SetReady(NotReady)
			if strings.Contains(err.Error(),"was forcibly closed")==true{
				c.Pool.PutError(c)
				return
			}
		}
		c.Pool.Put(c)
	}()
	c.lock.Lock()
	c.conn.SetDeadline(time.Now().Add(time.Second * 10))
	stockLen := len(stockList)
	cmd := []byte{}
	var req struct {
		F1          uint16
		F2          uint32
		PkgDataLen  uint16
		PkgDataLen2 uint16
		F3          uint32
		F4          uint32
		F5          uint16
		StockLen    uint16
	}
	// pkg_header = struct.pack("<HIHHIIHH", *values)

	req.F1 = 0x10c
	req.F2 = 0x02006320
	req.PkgDataLen = uint16(stockLen*7 + 12)
	req.PkgDataLen2 = req.PkgDataLen
	req.F3 = 0x5053e
	req.StockLen = uint16(stockLen)
	cmd = append(cmd, marshal(&req)...)
	for _, req := range stockList {
		cmd = append(cmd, req.Market)
		if len(req.Code) > 6 {
			cmd = append(cmd, req.Code[:6]...)
		} else if len(req.Code) <= 6 {
			cmd = append(cmd, req.Code...)
			for i := 0; i < 6-len(req.Code); i++ {
				cmd = append(cmd, 0)
			}
		}
	}
	// log.Println(cmd)
	c.conn.SetWriteDeadline(time.Now().Add(time.Second*1))
	_, err = c.conn.Write(cmd)
	if err != nil {
		return ret, err
	}
	resp, err = read(c.conn)
	if err != nil {
		return ret, err
	}
	var num uint16
	pos := 2
	d := resp.body
	reader := bytes.NewReader(d[pos:])
	binary.Read(reader, binary.LittleEndian, &num)
	pos += 2

	for i := 0; i < int(num); i++ {
		var item SecurityQuote
		item.Market = d[pos]
		item.Code = string(d[pos+1 : pos+7])
		pos += 9
		item.Price, pos = getPrice(d, pos)
		item.LastCloseDiff, pos = getPrice(d, pos)
		item.OpenDiff, pos = getPrice(d, pos)
		item.HighDiff, pos = getPrice(d, pos)
		item.LowDiff, pos = getPrice(d, pos)
		_, pos = getPrice(d, pos)
		_, pos = getPrice(d, pos)
		item.Vol, pos = getPrice(d, pos)
		item.CurVol, pos = getPrice(d, pos)
		var rawAmount uint32
		binary.Read(bytes.NewReader(d[pos:pos+4]), binary.LittleEndian, &rawAmount)
		pos += 4
		item.RawAmount = rawAmount
		item.Amount = GetVolume(rawAmount)
		item.SVol, pos = getPrice(d, pos)
		item.BVol, pos = getPrice(d, pos)
		_, pos = getPrice(d, pos)
		_, pos = getPrice(d, pos)

		item.Bid1, pos = getPrice(d, pos)
		item.Ask1, pos = getPrice(d, pos)
		item.BidVol1, pos = getPrice(d, pos)
		item.AskVol1, pos = getPrice(d, pos)

		item.Bid2, pos = getPrice(d, pos)
		item.Ask2, pos = getPrice(d, pos)
		item.BidVol2, pos = getPrice(d, pos)
		item.AskVol2, pos = getPrice(d, pos)

		item.Bid3, pos = getPrice(d, pos)
		item.Ask3, pos = getPrice(d, pos)
		item.BidVol3, pos = getPrice(d, pos)
		item.AskVol3, pos = getPrice(d, pos)

		item.Bid4, pos = getPrice(d, pos)
		item.Ask4, pos = getPrice(d, pos)
		item.BidVol4, pos = getPrice(d, pos)
		item.AskVol4, pos = getPrice(d, pos)

		item.Bid5, pos = getPrice(d, pos)
		item.Ask5, pos = getPrice(d, pos)
		item.BidVol5, pos = getPrice(d, pos)
		item.AskVol5, pos = getPrice(d, pos)

		// reversed_bytes4 = struct.unpack("<H", d[pos:pos+2])
		pos += 2
		_, pos = getPrice(d, pos)
		_, pos = getPrice(d, pos)
		_, pos = getPrice(d, pos)
		_, pos = getPrice(d, pos)
		// (reversed_bytes9, active2) = struct.unpack("<hH", d[pos: pos + 4])
		pos += 4

		ret = append(ret, &item)
	}
	//xret,err:=json.Marshal(ret)
	//if err!=nil{
	//	fmt.Println("json数据格式化出错",err)
	//}
	//var xxret interface{}
	//json.Unmarshal(xret,&xxret)
	//fmt.Println("股票行情=",xxret)
	return ret, nil
}

// ReqGetSecurityBars get kline
func (c *SyncQuoteClient) ReqGetSecurityBars(category, market uint16, code string, start, count uint16,wg *sync.WaitGroup) ([]*SecurityBar, error) {
	var(
		err error
		resp *tdxResponse
		ret []*SecurityBar
	)
	defer func(){
		wg.Done()
		c.lock.Unlock()
		if err!=nil{
			log.Println("SyncQuoteClient GetTransactionData开始获取交易明细数据错误",err)
			c.SetReady(NotReady)
			if strings.Contains(err.Error(),"was forcibly closed")==true||strings.Contains(err.Error(),"use of closed network connection")==true{
				c.Pool.PutError(c)
				return
			}
		}
		c.Pool.Put(c)
	}()
	c.lock.Lock()

	var req struct {
		F1       uint16
		F2       uint32
		F3       uint16
		F4       uint16
		F5       uint16
		Market   uint16
		Code     string `xlen:"6"`
		Category uint16
		F6       uint16
		Start    uint16
		Count    uint16
		F7       uint32
		F8       uint32
		F9       uint16
	}
	req.F1 = 0x10c
	req.F2 = 0x01016408
	req.F3 = 0x1c
	req.F4 = 0x1c
	req.F5 = 0x052d
	req.F6 = 1
	req.Category = category
	req.Market = market
	req.Code = code
	req.Start = start
	req.Count = count
	c.conn.SetWriteDeadline(time.Now().Add(time.Second*1))
	_, err = c.conn.Write(marshal(&req))
	if err != nil {
		return ret, err
	}
	resp, err = read(c.conn)
	if err != nil {
		return ret, err
	}
	d := resp.body
	var num uint16
	binary.Read(bytes.NewReader(d[0:2]), binary.LittleEndian, &num)
	pos := 2
	var preDiffBase int
	for i := 0; i < int(num); i++ {
		var item SecurityBar
		item.Year, item.Mon, item.Day, item.Hour, item.Minute = getDateTime(TdxKlineType(category), d[pos:])
		pos += 4
		item.PriceOpenDiff, pos = getPrice(d, pos)
		item.PriceCloseDiff, pos = getPrice(d, pos)

		item.PriceHighDiff, pos = getPrice(d, pos)
		item.PriceLowDiff, pos = getPrice(d, pos)

		var rawVol uint32
		binary.Read(bytes.NewReader(d[pos:pos+4]), binary.LittleEndian, &rawVol)
		vol := GetVolume(rawVol)
		pos += 4
		item.RawVol = rawVol
		item.Vol = int(vol)

		var rawDBVol uint32
		binary.Read(bytes.NewReader(d[pos:pos+4]), binary.LittleEndian, &rawDBVol)
		dbvol := GetVolume(rawDBVol)
		pos += 4
		item.RawDBVol = rawDBVol
		item.DBVol = int(dbvol)

		item.Open = calPrice1000(item.PriceOpenDiff, preDiffBase)

		item.PriceOpenDiff = item.PriceOpenDiff + preDiffBase

		item.Close = calPrice1000(item.PriceOpenDiff, item.PriceCloseDiff)
		item.High = calPrice1000(item.PriceOpenDiff, item.PriceHighDiff)
		item.Low = calPrice1000(item.PriceOpenDiff, item.PriceLowDiff)
		item.Code=code
		item.Market=int(market)

		preDiffBase = item.PriceOpenDiff + item.PriceCloseDiff

		ret = append(ret, &item)
	}
	return ret, nil
}

// ReqGetIndexBars req get index bar
func (c *SyncQuoteClient) ReqGetIndexBars(category, market uint16, code string, start, count uint16,wg *sync.WaitGroup) ([]*SecurityBar, error) {
	var(
		err error
		resp *tdxResponse
		ret []*SecurityBar
		req struct {
			F1       uint16
			F2       uint32
			F3       uint16
			F4       uint16
			F5       uint16
			Market   uint16
			Code     string `xlen:"6"`
			Category uint16
			F6       uint16
			Start    uint16
			Count    uint16
			F7       uint32
			F8       uint32
			F9       uint16
		}
	)
	defer func(){
		wg.Done()
		c.lock.Unlock()
		if err!=nil{
			log.Println("SyncQuoteClient GetTransactionData开始获取交易明细数据错误",err)
			c.SetReady(NotReady)
			if strings.Contains(err.Error(),"was forcibly closed")==true{
				c.Pool.PutError(c)
				return
			}
		}
		c.Pool.Put(c)


	}()
	c.lock.Lock()
	req.F1 = 0x10c
	req.F2 = 0x01016408
	req.F3 = 0x1c
	req.F4 = 0x1c
	req.F5 = 0x052d
	req.F6 = 1
	req.Category = category
	req.Market = market
	req.Code = code
	req.Start = start
	req.Count = count
	c.conn.SetWriteDeadline(time.Now().Add(time.Second*1))
	_, err = c.conn.Write(marshal(&req))
	if err != nil {
		return ret, err
	}
	resp, err = read(c.conn)
	if err != nil {
		return ret, err
	}
	d := resp.body
	var num uint16
	binary.Read(bytes.NewReader(d[0:2]), binary.LittleEndian, &num)
	pos := 2
	var preDiffBase int
	for i := 0; i < int(num); i++ {
		var item SecurityBar
		item.Year, item.Mon, item.Day, item.Hour, item.Minute = getDateTime(TdxKlineType(category), d[pos:])
		pos += 4
		item.PriceOpenDiff, pos = getPrice(d, pos)
		item.PriceCloseDiff, pos = getPrice(d, pos)

		item.PriceHighDiff, pos = getPrice(d, pos)
		item.PriceLowDiff, pos = getPrice(d, pos)

		var rawVol uint32
		binary.Read(bytes.NewReader(d[pos:pos+4]), binary.LittleEndian, &rawVol)
		vol := GetVolume(rawVol)
		pos += 4
		item.RawVol = rawVol
		item.Vol = int(vol)

		var rawDBVol uint32
		binary.Read(bytes.NewReader(d[pos:pos+4]), binary.LittleEndian, &rawDBVol)
		dbvol := GetVolume(rawDBVol)
		pos += 4
		item.RawDBVol = rawDBVol
		item.DBVol = int(dbvol)

		binary.Read(bytes.NewReader(d[pos:pos+2]), binary.LittleEndian, &item.UpCount)
		binary.Read(bytes.NewReader(d[pos+2:pos+4]), binary.LittleEndian, &item.DownCount)
		pos += 4

		item.Open = calPrice1000(item.PriceOpenDiff, preDiffBase)

		item.PriceOpenDiff = item.PriceOpenDiff + preDiffBase

		item.Close = calPrice1000(item.PriceOpenDiff, item.PriceCloseDiff)
		item.High = calPrice1000(item.PriceOpenDiff, item.PriceHighDiff)
		item.Low = calPrice1000(item.PriceOpenDiff, item.PriceLowDiff)

		preDiffBase = item.PriceOpenDiff + item.PriceCloseDiff

		ret = append(ret, &item)

	}
	return ret, nil
}

// ReqGetMinuteTimeData get minute time data
func (c *SyncQuoteClient) ReqGetMinuteTimeData(market uint16, code string,wg *sync.WaitGroup) ([]*RspGetMinuteTimeData, error) {
	var(
		err error
		resp *tdxResponse
		ret []*RspGetMinuteTimeData
	)
	defer func(){
		wg.Done()
		c.lock.Unlock()
		if err!=nil{
			log.Println("SyncQuoteClient GetTransactionData开始获取交易明细数据错误",err)
			c.SetReady(NotReady)
			if strings.Contains(err.Error(),"was forcibly closed")==true{
				c.Pool.PutError(c)
				return
			}
		}
		c.Pool.Put(c)
	}()
	c.lock.Lock()
	if len(code) != 6 {
		return nil, errors.New("code len should be 6")
	}
	cmd := []byte{0x0c, 0x1b, 0x08, 0x00, 0x01, 0x01, 0x0e, 0x00, 0x0e, 0x00, 0x1d, 0x05}
	buf := &bytes.Buffer{}
	buf.Write(cmd)
	binary.Write(buf, binary.LittleEndian, &market)
	buf.Write([]byte(code))
	buf.Write([]byte{0, 0, 0, 0})
	c.conn.SetWriteDeadline(time.Now().Add(time.Second*1))
	_, err = c.conn.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}
	resp, err= read(c.conn)
	if err != nil {
		return nil, err
	}
	d := resp.body
	var num uint16
	binary.Read(bytes.NewReader(d[0:2]), binary.LittleEndian, &num)
	log.Println(num, len(resp.body))
	pos := 4
	var lastPrice int
	for i := 0; i < int(num); i++ {
		var item RspGetMinuteTimeData
		rawPrice, pos := getPrice(d, pos)
		_, pos = getPrice(d, pos)
		vol, pos := getPrice(d, pos)
		lastPrice += rawPrice
		item.Time = uint16(lastPrice)
		item.Volume = uint32(vol)
		item.Price = float32(rawPrice)
		ret = append(ret, &item)
	}
	return ret, nil
}
func toTick(s *SecurityQuote) (tick basedata.Optiontick, err error) {

	switch s.Market{
	case 1:
		tick.Exchange="SH"
	case 0:
		tick.Exchange="SZ"

	}
	tick.Code=s.Code
	tick.Name=s.Code+"ETF"
	tick.Datetime=time.Now()
	tick.Underlying=tick.Name
	tick.Unit=100
	tick.Flag="期权标的"
	tick.Lastprice = float64(s.Price) / 1000
	tick.Open = float64(s.OpenDiff+s.Price) / 1000
	tick.High = float64(s.HighDiff+s.Price) / 1000
	tick.Low = float64(s.LowDiff+s.Price) / 1000
	tick.Vol = int64(s.Vol)
	//tick.VolumeDelta = int32(q.CurVol)
	tick.Amount = s.Amount
	tick.Askprice1=float64(s.Price+s.Ask1) / 1000
	tick.Bidprice1=float64(s.Price+s.Bid1) / 1000
	tick.Askvolume1 = float64(s.AskVol1)
	tick.Bidvolume1 = float64(s.BidVol1)

	tick.Askprice2=float64(s.Price+s.Ask2) / 1000
	tick.Bidprice2=float64(s.Price+s.Bid2) / 1000
	tick.Askvolume2 = float64(s.AskVol2)
	tick.Bidvolume2 = float64(s.BidVol2)

	tick.Askprice3=float64(s.Price+s.Ask3) / 1000
	tick.Bidprice3=float64(s.Price+s.Bid3) / 1000
	tick.Askvolume3 = float64(s.AskVol3)
	tick.Bidvolume3 = float64(s.BidVol3)

	tick.Askprice4=float64(s.Price+s.Ask4) / 1000
	tick.Bidprice4=float64(s.Price+s.Bid4) / 1000
	tick.Askvolume4 = float64(s.AskVol4)
	tick.Bidvolume4 = float64(s.BidVol4)

	tick.Askprice5=float64(s.Price+s.Ask5) / 1000
	tick.Bidprice5=float64(s.Price+s.Bid5) / 1000
	tick.Askvolume5 = float64(s.AskVol5)
	tick.Bidvolume5 = float64(s.BidVol5)
	return tick,nil
}