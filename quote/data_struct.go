package quote

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"opquotes/pub"
	"opquotes/utils"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 通达信返回消息的header
type header struct {
	F1        uint32
	F2        uint32 // 一般作为查询功能号
	F3        uint32
	ZipSize   uint16 // 压缩大小
	UnzipSize uint16 // 非压缩大小.如果 ZipSize == UnzipSize 不需要解压
}

func fixedWrite(buf *bytes.Buffer, source string, fixedLen int) {
	if len(source) < fixedLen {
		buf.WriteString(source)
		for i := 0; i < fixedLen-len(source); i++ {
			buf.WriteByte(0)
		}
		return
	}
	buf.WriteString(source[:fixedLen])
}

func fixedRead(reader *bytes.Reader, fixedLen int) (string, error) {
	buf := make([]byte, fixedLen)
	n, err := reader.Read(buf)
	if err != nil || n != fixedLen {
		// log.Println(n, fixedLen, err)
		return strings.Trim(string(buf), "\u0000"), err
	}
	return strings.Trim(string(buf), "\u0000"), nil
}

// unmarshal 反序列化
func unmarshal(d []byte, input interface{}) error {
	if d == nil {
		return errors.New("nil input")
	}
	reader := bytes.NewReader(d)
	rv := reflect.ValueOf(input)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	t := rv.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		v := rv.Field(i)
		tag := f.Tag.Get("xlen")
		typeName := f.Type.Name()
		// log.Printf("name[%s] type[%s] tag[%s] value[%v]", f.Name, f.Type, tag, v)
		switch typeName {
		case "string":
			if tag == "" {
				//log.Printf("name[%s] type[%s] tag[%s] value[%v]", f.Name, f.Type, tag, v)
				continue
				//panic("string fixed length not set")
			}
			xlen, err := strconv.Atoi(tag)
			if err != nil {
				panic(err)
			}
			xstr, err := fixedRead(reader, xlen)
			if err != nil {
				// log.Println("fixedRead", xstr, err)
			}
			v.SetString(xstr)
		case "float32":
			var fv float32
			binary.Read(reader, binary.LittleEndian, &fv)
			v.SetFloat(float64(fv))
		case "float64":
			var fv float64
			binary.Read(reader, binary.LittleEndian, &fv)
			v.SetFloat(fv)
		case "uint8":
			var uv uint8
			binary.Read(reader, binary.LittleEndian, &uv)
			v.SetUint(uint64(uv))
		case "uint16":
			var uv uint16
			binary.Read(reader, binary.LittleEndian, &uv)
			v.SetUint(uint64(uv))
		case "uint32":
			var uv uint32
			binary.Read(reader, binary.LittleEndian, &uv)
			v.SetUint(uint64(uv))
		case "uint64":
			var uv uint64
			binary.Read(reader, binary.LittleEndian, &uv)
			v.SetUint(uint64(uv))
		case "int8":
			var iv int8
			binary.Read(reader, binary.LittleEndian, &iv)
			v.SetInt(int64(iv))
		case "int16":
			var iv int16
			binary.Read(reader, binary.LittleEndian, &iv)
			v.SetInt(int64(iv))
		case "int32":
			var iv int32
			binary.Read(reader, binary.LittleEndian, &iv)
			v.SetInt(int64(iv))
		case "int64":
			var iv int64
			binary.Read(reader, binary.LittleEndian, &iv)
			v.SetInt(int64(iv))
		default:
			log.Printf("name[%s] type[%s] tag[%s] value[%v]", f.Name, f.Type, tag, v)
			panic("field not specified")
		}
		// log.Printf("name[%s] type[%s] tag[%s] value[%v]", f.Name, f.Type, tag, v)
	}
	// log.Println("未读取", t.Name(), reader.Len())
	return nil
}

func marshal(msg interface{}) []byte {
	if msg == nil {
		panic("nil ctp message")
	}
	rv := reflect.ValueOf(msg)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	t := rv.Type()
	buf := &bytes.Buffer{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		v := rv.Field(i)
		tag := f.Tag.Get("xlen")
		kind := v.Kind()
		// log.Printf("name[%s] type[%s] tag[%s] value[%v]", f.Name, f.Type, tag, v)
		switch kind {
		case reflect.String:
			if tag == "" {
				log.Printf("name[%s] type[%s] tag[%s] value[%v]", f.Name, f.Type, tag, v)
				panic("string fixed length not set")
			}
			xlen, err := strconv.Atoi(tag)
			if err != nil {
				panic(err)
			}
			fixedWrite(buf, v.String(), xlen)
		case reflect.Float32:
			uv := float32(v.Float())
			binary.Write(buf, binary.LittleEndian, uv)
		case reflect.Float64:
			binary.Write(buf, binary.LittleEndian, v.Float())
		case reflect.Uint8:
			uv := uint8(v.Uint())
			binary.Write(buf, binary.LittleEndian, uv)
		case reflect.Uint16:
			uv := uint16(v.Uint())
			binary.Write(buf, binary.LittleEndian, uv)
		case reflect.Uint32:
			uv := uint32(v.Uint())
			binary.Write(buf, binary.LittleEndian, uv)
		case reflect.Uint64:
			uv := uint64(v.Uint())
			binary.Write(buf, binary.LittleEndian, uv)
		case reflect.Int8:
			uv := int8(v.Int())
			binary.Write(buf, binary.LittleEndian, uv)
		case reflect.Int16:
			uv := int16(v.Int())
			binary.Write(buf, binary.LittleEndian, uv)
		case reflect.Int32:
			uv := int32(v.Int())
			binary.Write(buf, binary.LittleEndian, uv)
		case reflect.Int64:
			uv := int64(v.Int())
			binary.Write(buf, binary.LittleEndian, uv)
		default:
			log.Printf("name[%s] type[%s] tag[%s] value[%v]", f.Name, f.Type, tag, v)
			panic("field not specified")
		}
		// log.Printf("name[%s] type[%s] tag[%s] value[%v]", f.Name, f.Type, tag, v)
	}
	return buf.Bytes()
}

// RspQryStockCount 返回股票数量
type RspQryStockCount struct {
	Count uint16
}

// RspGetInstrumentCount 扩展行情合约数量
type RspGetInstrumentCount struct {
	Count uint32
}

// RspQrySecurity 返回证券列表
type RspQrySecurity struct {
	Code         string `xlen:"6"`
	VolumeUnit   uint16
	Name         string `xlen:"8"`
	Reserve1     string `xlen:"4"`
	DecimalPoint uint8
	PreCloseRaw  uint32
	Reserve2     string `xlen:"4"`
}

// ExtRspQryMarket 扩展行情返回市场列表
type ExtRspQryMarket struct {
	Category  uint8
	Name      string `xlen:"32"`
	Market    uint8
	ShortName string `xlen:"2"`
	F1        string `xlen:"26"`
	F2        string `xlen:"2"`
}

// ReqGetInstrumentBars 扩展查询K线
type ReqGetInstrumentBars struct {
	Market   uint8
	Code     string `xlen:"9"`
	Category TdxKlineType
	F1       uint16
	Start    uint32
	Count    uint16
}

type TdxKlineType uint16

const (
	TdxKlineType_5MIN      = 0
	TdxKlineType_15MIN     = 1
	TdxKlineType_30MIN     = 2
	TdxKlineType_1HOUR     = 3
	TdxKlineType_DAILY     = 4
	TdxKlineType_WEEKLY    = 5
	TdxKlineType_MONTHLY   = 6
	TdxKlineType_EXHQ_1MIN = 7
	TdxKlineType_1MIN      = 8
	TdxKlineType_RI_K      = 9
	TdxKlineType_3MONTH    = 10
	TdxKlineType_YEARLY    = 11
)

// TdxKline 通达信K线
type TdxKline struct {
	Code    string
	Open     float32
	High     float32
	Low      float32
	Close    float32
	Position uint32
	Volume   uint32
	Price    float32
	Time     int64
	Category uint16
}

// ReqGetInstrumentQuote 请求行情
type ReqGetInstrumentQuote struct {
	Market uint8
	Code   string `xlen:"9"`
}

// ReqGetMinuteTimeData 查询分时
type ReqGetMinuteTimeData ReqGetInstrumentQuote

// RspGetInstrumentQuote 返回查询行情
type RspGetInstrumentQuote struct {
	Market       uint8
	Code         string `xlen:"9"`
	F1           string `xlen:"4"`
	PreClose     float32
	Open         float32
	High         float32
	Low          float32
	Price        float32
	OpenInterest int32
	F2           int32
	Volume       int32
	LastVolume   int32
	F3           int32
	Neipan       int32
	Waipan       int32
	F4           int32
	Position     int32
	Bid1         float32
	Bid2         float32
	Bid3         float32
	Bid4         float32
	Bid5         float32
	BidVolume1   int32
	BidVolume2   int32
	BidVolume3   int32
	BidVolume4   int32
	BidVolume5   int32
	Ask1         float32
	Ask2         float32
	Ask3         float32
	Ask4         float32
	Ask5         float32
	AskVolume1   int32
	AskVolume2   int32
	AskVolume3   int32
	AskVolume4   int32
	AskVolume5   int32
}

// ReqGetInstrumentInfo 查询合约信息
type ReqGetInstrumentInfo struct {
	Start uint32
	Count uint16
}

// RspGetInstrumentInfo 查询合约信息返回
type RspGetInstrumentInfo struct {
	Category    uint8
	Market      uint8
	F1          string `xlen:"3"`
	Code        string `xlen:"9"`
	Name        string `xlen:"17"`
	Discription string `xlen:"9"`
	F2          string `xlen:"24"`
}

// RspGetMinuteTimeDataHeader 返回查询分时图结果
type RspGetMinuteTimeDataHeader struct {
	Market uint8
	Code   string `xlen:"9"`
	Count  uint16
}

// RspGetMinuteTimeData 查询分时图返回
type RspGetMinuteTimeData struct {
	Time         uint16
	Price        float32
	AveragePrice float32
	Volume       uint32
	Amount       uint32
	Code         string
}

// ReqGetHistoryMinuteTimeData 请求查询
type ReqGetHistoryMinuteTimeData struct {
	Date   uint32
	Market uint8
	Code   string `xlen:"9"`
}

// RspGetHistoryMinuteTimeDataHeader 查询历史分时响应头
type RspGetHistoryMinuteTimeDataHeader struct {
	Market uint8
	Code   string `xlen:"9"`
	F1     string `xlen:"8"`
	Count  uint16
}

// ReqGetTransactionData 查询分笔成交
type ReqGetTransactionData struct {
	Market uint8
	Code   string `xlen:"9"`
	Start  int32
	Count  uint16
}

// RspGetTransactionDataHeader 返回查询分笔成交头
type RspGetTransactionDataHeader struct {
	Market uint8
	Code   string `xlen:"9"`
	F1     string `xlen:"4"`
	Count  uint16
}

// RspGetTransactionData 查询分笔成交

type TransactionData struct{
	Code string
	Exchange string
	Datetime time.Time
	Price float64
	Volume int64
	Mindatetime time.Time
	Posdir int64
	Direction int64
	Pre_close  float64
}
func NewTransactionData(code,exchange string)*TransactionData{
	return &TransactionData{
		Code:     code,
		Exchange: exchange,
		Datetime:     time.Time{},
		Price:    0,
		Volume:   0,
		Posdir:0,
		Direction:0,
		Pre_close:0,
	}
}
func (t *TransactionData)Save()error{
	var mindatetime time.Time
	if t.Datetime.IsZero(){
		return errors.New("时间为空")
	}
	if t.Datetime.Hour()==15 && t.Datetime.Minute()==0 {
		mindatetime = utils.Str2Time(utils.Time2StrDate(time.Now()) + " 14:59:59")
	}else if t.Datetime.Hour()==9&&t.Datetime.Minute()<30{
		mindatetime = utils.Str2Time(utils.Time2StrDate(time.Now()) + " 09:30:00")
	}else{
		mindatetime= t.Datetime
	}
	t.Pre_close= OptionMgr.Items[t.Code].Preclose
	sql:=fmt.Sprintf(`insert into transactiondata(code,exchange,datetime,price,volume,mindatetime,posdir,direction,pre_close)values('%s','%s','%s',%f,%d,'%s',%d,%d,%f)  ON CONFLICT (code, datetime,posdir,direction) do nothing`,t.Code,t.Exchange,t.Datetime.Format("2006-01-02 15:04:05"),t.Price,t.Volume,mindatetime.Format("2006-01-02 15:04:00"),t.Posdir,t.Direction,t.Pre_close)
	_,err:=pub.DB.Exec(sql)
	if err!=nil{
		log.Println("通达信接口分笔数据插入失败，错误代码=",err.Error())
		return err
	}
	return nil

}
type RspGetTransactionData struct {
	Time                 uint16
	Price                uint32
	Volume               uint32
	Posdir 				 int32
	Direction            uint16
}

// ReqGetHistoryTransactionData 查询分笔成交
type ReqGetHistoryTransactionData struct {
	Date   uint32
	Market uint8
	Code   string `xlen:"9"`
	Start  int32
	Count  uint16
}

// SecurityQuote 股票报价
type SecurityQuote struct {
	Market        uint8
	Code          string
	Active1       uint16
	Price         int
	LastCloseDiff int
	OpenDiff      int
	HighDiff      int
	LowDiff       int
	Vol           int
	CurVol        int
	RawAmount     uint32
	Amount        float64
	SVol          int
	BVol          int
	Bid1          int
	Ask1          int
	BidVol1       int
	AskVol1       int
	Bid2          int
	Ask2          int
	BidVol2       int
	AskVol2       int
	Bid3          int
	Ask3          int
	BidVol3       int
	AskVol3       int
	Bid4          int
	Ask4          int
	BidVol4       int
	AskVol4       int
	Bid5          int
	Ask5          int
	BidVol5       int
	AskVol5       int
}

// SecurityBar  bar
type SecurityBar struct {
	Code 		   string
	Market         int
	Year           int
	Mon            int
	Day            int
	Hour           int
	Minute         int
	PriceOpenDiff  int
	PriceCloseDiff int
	PriceHighDiff  int
	PriceLowDiff   int
	RawVol         uint32
	Vol            int
	RawDBVol       uint32
	DBVol          int
	Open           float64
	Close          float64
	High           float64
	Low            float64
	UpCount        uint16
	DownCount      uint16
}
