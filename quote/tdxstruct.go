package quote

import (
fmt "fmt"
proto "github.com/golang/protobuf/proto"
math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type OrderBook struct {
	Ask                  float64  `protobuf:"fixed64,1,opt,name=ask,proto3" json:"ask"`
	AskVolume            int32    `protobuf:"varint,2,opt,name=ask_volume,json=askVolume,proto3" json:"askVolume"`
	Bid                  float64  `protobuf:"fixed64,3,opt,name=bid,proto3" json:"bid"`
	BidVolume            int32    `protobuf:"varint,4,opt,name=bid_volume,json=bidVolume,proto3" json:"bidVolume"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderBook) Reset()         { *m = OrderBook{} }
func (m *OrderBook) String() string { return proto.CompactTextString(m) }
func (*OrderBook) ProtoMessage()    {}
func (*OrderBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{0}
}

func (m *OrderBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderBook.Unmarshal(m, b)
}
func (m *OrderBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderBook.Marshal(b, m, deterministic)
}
func (m *OrderBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBook.Merge(m, src)
}
func (m *OrderBook) XXX_Size() int {
	return xxx_messageInfo_OrderBook.Size(m)
}
func (m *OrderBook) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBook.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBook proto.InternalMessageInfo

func (m *OrderBook) GetAsk() float64 {
	if m != nil {
		return m.Ask
	}
	return 0
}

func (m *OrderBook) GetAskVolume() int32 {
	if m != nil {
		return m.AskVolume
	}
	return 0
}

func (m *OrderBook) GetBid() float64 {
	if m != nil {
		return m.Bid
	}
	return 0
}

func (m *OrderBook) GetBidVolume() int32 {
	if m != nil {
		return m.BidVolume
	}
	return 0
}

// 深度行情
type MarketDataSnapshot struct {
	Exchange             string       `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange"`
	Symbol               string       `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol"`
	Time                 int64        `protobuf:"varint,3,opt,name=time,proto3" json:"time"`
	Milliseconds         int32        `protobuf:"varint,4,opt,name=milliseconds,proto3" json:"milliseconds"`
	Open                 float64      `protobuf:"fixed64,5,opt,name=open,proto3" json:"open"`
	High                 float64      `protobuf:"fixed64,6,opt,name=high,proto3" json:"high"`
	Low                  float64      `protobuf:"fixed64,7,opt,name=low,proto3" json:"low"`
	Close                float64      `protobuf:"fixed64,8,opt,name=close,proto3" json:"close"`
	Volume               int32        `protobuf:"varint,9,opt,name=volume,proto3" json:"volume"`
	Amount               float64      `protobuf:"fixed64,10,opt,name=amount,proto3" json:"amount"`
	Position             int32        `protobuf:"varint,11,opt,name=position,proto3" json:"position"`
	Price                float64      `protobuf:"fixed64,12,opt,name=price,proto3" json:"price"`
	PreClose             float64      `protobuf:"fixed64,13,opt,name=pre_close,json=preClose,proto3" json:"preClose"`
	PreSettlement        float64      `protobuf:"fixed64,14,opt,name=pre_settlement,json=preSettlement,proto3" json:"preSettlement"`
	PrePosition          int32        `protobuf:"varint,15,opt,name=pre_position,json=prePosition,proto3" json:"prePosition"`
	SettlementPrice      float64      `protobuf:"fixed64,16,opt,name=settlement_price,json=settlementPrice,proto3" json:"settlementPrice"`
	UpperLimit           float64      `protobuf:"fixed64,17,opt,name=upper_limit,json=upperLimit,proto3" json:"upperLimit"`
	LowerLimit           float64      `protobuf:"fixed64,18,opt,name=lower_limit,json=lowerLimit,proto3" json:"lowerLimit"`
	PreDelta             float64      `protobuf:"fixed64,19,opt,name=pre_delta,json=preDelta,proto3" json:"preDelta"`
	Delta                float64      `protobuf:"fixed64,20,opt,name=delta,proto3" json:"delta"`
	AveragePrice         float64      `protobuf:"fixed64,21,opt,name=average_price,json=averagePrice,proto3" json:"averagePrice"`
	TradingDay           int32        `protobuf:"varint,22,opt,name=trading_day,json=tradingDay,proto3" json:"tradingDay"`
	Name                 string       `protobuf:"bytes,23,opt,name=name,proto3" json:"name"`
	ExercisePrice        float64      `protobuf:"fixed64,24,opt,name=exercise_price,json=exercisePrice,proto3" json:"exercisePrice"`
	VolumeDelta          int32        `protobuf:"varint,25,opt,name=volume_delta,json=volumeDelta,proto3" json:"volumeDelta"`
	ActionDay            int32        `protobuf:"varint,27,opt,name=action_day,json=actionDay,proto3" json:"actionDay"`
	Multiple             int32        `protobuf:"varint,28,opt,name=multiple,proto3" json:"multiple"`
	PriceTick            float64      `protobuf:"fixed64,39,opt,name=price_tick,json=priceTick,proto3" json:"priceTick"`
	Depths               []*OrderBook `protobuf:"bytes,30,rep,name=depths,proto3" json:"depths"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MarketDataSnapshot) Reset()         { *m = MarketDataSnapshot{} }
func (m *MarketDataSnapshot) String() string { return proto.CompactTextString(m) }
func (*MarketDataSnapshot) ProtoMessage()    {}
func (*MarketDataSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{1}
}

func (m *MarketDataSnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketDataSnapshot.Unmarshal(m, b)
}
func (m *MarketDataSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketDataSnapshot.Marshal(b, m, deterministic)
}
func (m *MarketDataSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketDataSnapshot.Merge(m, src)
}
func (m *MarketDataSnapshot) XXX_Size() int {
	return xxx_messageInfo_MarketDataSnapshot.Size(m)
}
func (m *MarketDataSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketDataSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_MarketDataSnapshot proto.InternalMessageInfo

func (m *MarketDataSnapshot) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *MarketDataSnapshot) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *MarketDataSnapshot) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *MarketDataSnapshot) GetMilliseconds() int32 {
	if m != nil {
		return m.Milliseconds
	}
	return 0
}

func (m *MarketDataSnapshot) GetOpen() float64 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *MarketDataSnapshot) GetHigh() float64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *MarketDataSnapshot) GetLow() float64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *MarketDataSnapshot) GetClose() float64 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *MarketDataSnapshot) GetVolume() int32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *MarketDataSnapshot) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *MarketDataSnapshot) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *MarketDataSnapshot) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *MarketDataSnapshot) GetPreClose() float64 {
	if m != nil {
		return m.PreClose
	}
	return 0
}

func (m *MarketDataSnapshot) GetPreSettlement() float64 {
	if m != nil {
		return m.PreSettlement
	}
	return 0
}

func (m *MarketDataSnapshot) GetPrePosition() int32 {
	if m != nil {
		return m.PrePosition
	}
	return 0
}

func (m *MarketDataSnapshot) GetSettlementPrice() float64 {
	if m != nil {
		return m.SettlementPrice
	}
	return 0
}

func (m *MarketDataSnapshot) GetUpperLimit() float64 {
	if m != nil {
		return m.UpperLimit
	}
	return 0
}

func (m *MarketDataSnapshot) GetLowerLimit() float64 {
	if m != nil {
		return m.LowerLimit
	}
	return 0
}

func (m *MarketDataSnapshot) GetPreDelta() float64 {
	if m != nil {
		return m.PreDelta
	}
	return 0
}

func (m *MarketDataSnapshot) GetDelta() float64 {
	if m != nil {
		return m.Delta
	}
	return 0
}

func (m *MarketDataSnapshot) GetAveragePrice() float64 {
	if m != nil {
		return m.AveragePrice
	}
	return 0
}

func (m *MarketDataSnapshot) GetTradingDay() int32 {
	if m != nil {
		return m.TradingDay
	}
	return 0
}

func (m *MarketDataSnapshot) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MarketDataSnapshot) GetExercisePrice() float64 {
	if m != nil {
		return m.ExercisePrice
	}
	return 0
}

func (m *MarketDataSnapshot) GetVolumeDelta() int32 {
	if m != nil {
		return m.VolumeDelta
	}
	return 0
}

func (m *MarketDataSnapshot) GetActionDay() int32 {
	if m != nil {
		return m.ActionDay
	}
	return 0
}

func (m *MarketDataSnapshot) GetMultiple() int32 {
	if m != nil {
		return m.Multiple
	}
	return 0
}

func (m *MarketDataSnapshot) GetPriceTick() float64 {
	if m != nil {
		return m.PriceTick
	}
	return 0
}

func (m *MarketDataSnapshot) GetDepths() []*OrderBook {
	if m != nil {
		return m.Depths
	}
	return nil
}

// MdsList 行情列表
type MdsList struct {
	List                 []*MarketDataSnapshot `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MdsList) Reset()         { *m = MdsList{} }
func (m *MdsList) String() string { return proto.CompactTextString(m) }
func (*MdsList) ProtoMessage()    {}
func (*MdsList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{2}
}

func (m *MdsList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MdsList.Unmarshal(m, b)
}
func (m *MdsList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MdsList.Marshal(b, m, deterministic)
}
func (m *MdsList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MdsList.Merge(m, src)
}
func (m *MdsList) XXX_Size() int {
	return xxx_messageInfo_MdsList.Size(m)
}
func (m *MdsList) XXX_DiscardUnknown() {
	xxx_messageInfo_MdsList.DiscardUnknown(m)
}

var xxx_messageInfo_MdsList proto.InternalMessageInfo

func (m *MdsList) GetList() []*MarketDataSnapshot {
	if m != nil {
		return m.List
	}
	return nil
}

//期权T型
type OptionTMarket struct {
	CallTk               *MarketDataSnapshot `protobuf:"bytes,1,opt,name=callTk,proto3" json:"callTk"`
	PutTk                *MarketDataSnapshot `protobuf:"bytes,2,opt,name=putTk,proto3" json:"putTk"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *OptionTMarket) Reset()         { *m = OptionTMarket{} }
func (m *OptionTMarket) String() string { return proto.CompactTextString(m) }
func (*OptionTMarket) ProtoMessage()    {}
func (*OptionTMarket) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{3}
}

func (m *OptionTMarket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptionTMarket.Unmarshal(m, b)
}
func (m *OptionTMarket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptionTMarket.Marshal(b, m, deterministic)
}
func (m *OptionTMarket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptionTMarket.Merge(m, src)
}
func (m *OptionTMarket) XXX_Size() int {
	return xxx_messageInfo_OptionTMarket.Size(m)
}
func (m *OptionTMarket) XXX_DiscardUnknown() {
	xxx_messageInfo_OptionTMarket.DiscardUnknown(m)
}

var xxx_messageInfo_OptionTMarket proto.InternalMessageInfo

func (m *OptionTMarket) GetCallTk() *MarketDataSnapshot {
	if m != nil {
		return m.CallTk
	}
	return nil
}

func (m *OptionTMarket) GetPutTk() *MarketDataSnapshot {
	if m != nil {
		return m.PutTk
	}
	return nil
}

// 简易期权T型报价
type SimpleTickForTQuote struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price"`
	UpDownRatio          float64  `protobuf:"fixed64,4,opt,name=up_down_ratio,json=upDownRatio,proto3" json:"upDownRatio"`
	PreSettlementPrice   float64  `protobuf:"fixed64,5,opt,name=pre_settlement_price,json=preSettlementPrice,proto3" json:"preSettlementPrice"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleTickForTQuote) Reset()         { *m = SimpleTickForTQuote{} }
func (m *SimpleTickForTQuote) String() string { return proto.CompactTextString(m) }
func (*SimpleTickForTQuote) ProtoMessage()    {}
func (*SimpleTickForTQuote) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{4}
}

func (m *SimpleTickForTQuote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleTickForTQuote.Unmarshal(m, b)
}
func (m *SimpleTickForTQuote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleTickForTQuote.Marshal(b, m, deterministic)
}
func (m *SimpleTickForTQuote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleTickForTQuote.Merge(m, src)
}
func (m *SimpleTickForTQuote) XXX_Size() int {
	return xxx_messageInfo_SimpleTickForTQuote.Size(m)
}
func (m *SimpleTickForTQuote) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleTickForTQuote.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleTickForTQuote proto.InternalMessageInfo

func (m *SimpleTickForTQuote) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *SimpleTickForTQuote) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *SimpleTickForTQuote) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *SimpleTickForTQuote) GetUpDownRatio() float64 {
	if m != nil {
		return m.UpDownRatio
	}
	return 0
}

func (m *SimpleTickForTQuote) GetPreSettlementPrice() float64 {
	if m != nil {
		return m.PreSettlementPrice
	}
	return 0
}

func (m *SimpleTickForTQuote) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// T型报价的一行
type OptionTQuoteItem struct {
	ExercisePrice        float64              `protobuf:"fixed64,1,opt,name=exercise_price,json=exercisePrice,proto3" json:"exercisePrice"`
	Call                 *SimpleTickForTQuote `protobuf:"bytes,2,opt,name=call,proto3" json:"call"`
	Put                  *SimpleTickForTQuote `protobuf:"bytes,3,opt,name=put,proto3" json:"put"`
	ExercisePriceFlag    string               `protobuf:"bytes,4,opt,name=exercise_price_flag,json=exercisePriceFlag,proto3" json:"exercisePriceFlag"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OptionTQuoteItem) Reset()         { *m = OptionTQuoteItem{} }
func (m *OptionTQuoteItem) String() string { return proto.CompactTextString(m) }
func (*OptionTQuoteItem) ProtoMessage()    {}
func (*OptionTQuoteItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{5}
}

func (m *OptionTQuoteItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptionTQuoteItem.Unmarshal(m, b)
}
func (m *OptionTQuoteItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptionTQuoteItem.Marshal(b, m, deterministic)
}
func (m *OptionTQuoteItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptionTQuoteItem.Merge(m, src)
}
func (m *OptionTQuoteItem) XXX_Size() int {
	return xxx_messageInfo_OptionTQuoteItem.Size(m)
}
func (m *OptionTQuoteItem) XXX_DiscardUnknown() {
	xxx_messageInfo_OptionTQuoteItem.DiscardUnknown(m)
}

var xxx_messageInfo_OptionTQuoteItem proto.InternalMessageInfo

func (m *OptionTQuoteItem) GetExercisePrice() float64 {
	if m != nil {
		return m.ExercisePrice
	}
	return 0
}

func (m *OptionTQuoteItem) GetCall() *SimpleTickForTQuote {
	if m != nil {
		return m.Call
	}
	return nil
}

func (m *OptionTQuoteItem) GetPut() *SimpleTickForTQuote {
	if m != nil {
		return m.Put
	}
	return nil
}

func (m *OptionTQuoteItem) GetExercisePriceFlag() string {
	if m != nil {
		return m.ExercisePriceFlag
	}
	return ""
}

// OptionTQuoteItemList 列表
type OptionTQuoteItemList struct {
	Exchange             int32               `protobuf:"varint,1,opt,name=exchange,proto3" json:"exchange"`
	StrikeSymbol         string              `protobuf:"bytes,2,opt,name=strike_symbol,json=strikeSymbol,proto3" json:"strikeSymbol"`
	Month                string              `protobuf:"bytes,3,opt,name=month,proto3" json:"month"`
	List                 []*OptionTQuoteItem `protobuf:"bytes,4,rep,name=list,proto3" json:"list"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *OptionTQuoteItemList) Reset()         { *m = OptionTQuoteItemList{} }
func (m *OptionTQuoteItemList) String() string { return proto.CompactTextString(m) }
func (*OptionTQuoteItemList) ProtoMessage()    {}
func (*OptionTQuoteItemList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{6}
}

func (m *OptionTQuoteItemList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptionTQuoteItemList.Unmarshal(m, b)
}
func (m *OptionTQuoteItemList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptionTQuoteItemList.Marshal(b, m, deterministic)
}
func (m *OptionTQuoteItemList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptionTQuoteItemList.Merge(m, src)
}
func (m *OptionTQuoteItemList) XXX_Size() int {
	return xxx_messageInfo_OptionTQuoteItemList.Size(m)
}
func (m *OptionTQuoteItemList) XXX_DiscardUnknown() {
	xxx_messageInfo_OptionTQuoteItemList.DiscardUnknown(m)
}

var xxx_messageInfo_OptionTQuoteItemList proto.InternalMessageInfo

func (m *OptionTQuoteItemList) GetExchange() int32 {
	if m != nil {
		return m.Exchange
	}
	return 0
}

func (m *OptionTQuoteItemList) GetStrikeSymbol() string {
	if m != nil {
		return m.StrikeSymbol
	}
	return ""
}

func (m *OptionTQuoteItemList) GetMonth() string {
	if m != nil {
		return m.Month
	}
	return ""
}

func (m *OptionTQuoteItemList) GetList() []*OptionTQuoteItem {
	if m != nil {
		return m.List
	}
	return nil
}

// K线
type Kline struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time"`
	Open                 float64  `protobuf:"fixed64,2,opt,name=open,proto3" json:"open"`
	High                 float64  `protobuf:"fixed64,3,opt,name=high,proto3" json:"high"`
	Low                  float64  `protobuf:"fixed64,4,opt,name=low,proto3" json:"low"`
	Close                float64  `protobuf:"fixed64,5,opt,name=close,proto3" json:"close"`
	Volume               int32    `protobuf:"varint,6,opt,name=volume,proto3" json:"volume"`
	Amount               float64  `protobuf:"fixed64,7,opt,name=amount,proto3" json:"amount"`
	Position             int32    `protobuf:"varint,8,opt,name=position,proto3" json:"position"`
	TradingDay           int32    `protobuf:"varint,9,opt,name=trading_day,json=tradingDay,proto3" json:"tradingDay"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Kline) Reset()         { *m = Kline{} }
func (m *Kline) String() string { return proto.CompactTextString(m) }
func (*Kline) ProtoMessage()    {}
func (*Kline) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{7}
}

func (m *Kline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Kline.Unmarshal(m, b)
}
func (m *Kline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Kline.Marshal(b, m, deterministic)
}
func (m *Kline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Kline.Merge(m, src)
}
func (m *Kline) XXX_Size() int {
	return xxx_messageInfo_Kline.Size(m)
}
func (m *Kline) XXX_DiscardUnknown() {
	xxx_messageInfo_Kline.DiscardUnknown(m)
}

var xxx_messageInfo_Kline proto.InternalMessageInfo

func (m *Kline) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Kline) GetOpen() float64 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *Kline) GetHigh() float64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *Kline) GetLow() float64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *Kline) GetClose() float64 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *Kline) GetVolume() int32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Kline) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Kline) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *Kline) GetTradingDay() int32 {
	if m != nil {
		return m.TradingDay
	}
	return 0
}

// KlineSeries K线序列
type KlineSeries struct {
	Exchange             string     `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange"`
	Symbol               string     `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol"`
	Period               PeriodType `protobuf:"varint,3,opt,name=period,proto3,enum=goshare.PeriodType" json:"period"`
	PeriodInSeconds      int32      `protobuf:"varint,4,opt,name=period_in_seconds,json=periodInSeconds,proto3" json:"periodInSeconds"`
	List                 []*Kline   `protobuf:"bytes,5,rep,name=list,proto3" json:"list"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *KlineSeries) Reset()         { *m = KlineSeries{} }
func (m *KlineSeries) String() string { return proto.CompactTextString(m) }
func (*KlineSeries) ProtoMessage()    {}
func (*KlineSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{8}
}

func (m *KlineSeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KlineSeries.Unmarshal(m, b)
}
func (m *KlineSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KlineSeries.Marshal(b, m, deterministic)
}
func (m *KlineSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KlineSeries.Merge(m, src)
}
func (m *KlineSeries) XXX_Size() int {
	return xxx_messageInfo_KlineSeries.Size(m)
}
func (m *KlineSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_KlineSeries.DiscardUnknown(m)
}

var xxx_messageInfo_KlineSeries proto.InternalMessageInfo

func (m *KlineSeries) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *KlineSeries) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *KlineSeries) GetPeriod() PeriodType {
	if m != nil {
		return m.Period
	}
	return PeriodType_TICK
}

func (m *KlineSeries) GetPeriodInSeconds() int32 {
	if m != nil {
		return m.PeriodInSeconds
	}
	return 0
}

func (m *KlineSeries) GetList() []*Kline {
	if m != nil {
		return m.List
	}
	return nil
}

// 订阅行情
type ReqSubscribeMarketData struct {
	Exchange             string       `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange"`
	Symbol               string       `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol"`
	PeriodList           []PeriodType `protobuf:"varint,3,rep,packed,name=periodList,proto3,enum=goshare.PeriodType" json:"periodList"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReqSubscribeMarketData) Reset()         { *m = ReqSubscribeMarketData{} }
func (m *ReqSubscribeMarketData) String() string { return proto.CompactTextString(m) }
func (*ReqSubscribeMarketData) ProtoMessage()    {}
func (*ReqSubscribeMarketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{9}
}

func (m *ReqSubscribeMarketData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSubscribeMarketData.Unmarshal(m, b)
}
func (m *ReqSubscribeMarketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSubscribeMarketData.Marshal(b, m, deterministic)
}
func (m *ReqSubscribeMarketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSubscribeMarketData.Merge(m, src)
}
func (m *ReqSubscribeMarketData) XXX_Size() int {
	return xxx_messageInfo_ReqSubscribeMarketData.Size(m)
}
func (m *ReqSubscribeMarketData) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSubscribeMarketData.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSubscribeMarketData proto.InternalMessageInfo

func (m *ReqSubscribeMarketData) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *ReqSubscribeMarketData) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqSubscribeMarketData) GetPeriodList() []PeriodType {
	if m != nil {
		return m.PeriodList
	}
	return nil
}

// 返回订阅行情
type RspSubscribeMarketData struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RspSubscribeMarketData) Reset()         { *m = RspSubscribeMarketData{} }
func (m *RspSubscribeMarketData) String() string { return proto.CompactTextString(m) }
func (*RspSubscribeMarketData) ProtoMessage()    {}
func (*RspSubscribeMarketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{10}
}

func (m *RspSubscribeMarketData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspSubscribeMarketData.Unmarshal(m, b)
}
func (m *RspSubscribeMarketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspSubscribeMarketData.Marshal(b, m, deterministic)
}
func (m *RspSubscribeMarketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspSubscribeMarketData.Merge(m, src)
}
func (m *RspSubscribeMarketData) XXX_Size() int {
	return xxx_messageInfo_RspSubscribeMarketData.Size(m)
}
func (m *RspSubscribeMarketData) XXX_DiscardUnknown() {
	xxx_messageInfo_RspSubscribeMarketData.DiscardUnknown(m)
}

var xxx_messageInfo_RspSubscribeMarketData proto.InternalMessageInfo

// 推送订阅行情更新事件
type RtnMarketDataUpdate struct {
	Tick                 *MarketDataSnapshot `protobuf:"bytes,1,opt,name=tick,proto3" json:"tick"`
	KlineList            []*Kline            `protobuf:"bytes,2,rep,name=kline_list,json=klineList,proto3" json:"klineList"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RtnMarketDataUpdate) Reset()         { *m = RtnMarketDataUpdate{} }
func (m *RtnMarketDataUpdate) String() string { return proto.CompactTextString(m) }
func (*RtnMarketDataUpdate) ProtoMessage()    {}
func (*RtnMarketDataUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{11}
}

func (m *RtnMarketDataUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RtnMarketDataUpdate.Unmarshal(m, b)
}
func (m *RtnMarketDataUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RtnMarketDataUpdate.Marshal(b, m, deterministic)
}
func (m *RtnMarketDataUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RtnMarketDataUpdate.Merge(m, src)
}
func (m *RtnMarketDataUpdate) XXX_Size() int {
	return xxx_messageInfo_RtnMarketDataUpdate.Size(m)
}
func (m *RtnMarketDataUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_RtnMarketDataUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_RtnMarketDataUpdate proto.InternalMessageInfo

func (m *RtnMarketDataUpdate) GetTick() *MarketDataSnapshot {
	if m != nil {
		return m.Tick
	}
	return nil
}

func (m *RtnMarketDataUpdate) GetKlineList() []*Kline {
	if m != nil {
		return m.KlineList
	}
	return nil
}

// tick序列
type TickSeries struct {
	Exchange             string                `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange"`
	Symbol               string                `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol"`
	TradingDay           int32                 `protobuf:"varint,3,opt,name=trading_day,json=tradingDay,proto3" json:"tradingDay"`
	List                 []*MarketDataSnapshot `protobuf:"bytes,4,rep,name=list,proto3" json:"list"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TickSeries) Reset()         { *m = TickSeries{} }
func (m *TickSeries) String() string { return proto.CompactTextString(m) }
func (*TickSeries) ProtoMessage()    {}
func (*TickSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{12}
}

func (m *TickSeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TickSeries.Unmarshal(m, b)
}
func (m *TickSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TickSeries.Marshal(b, m, deterministic)
}
func (m *TickSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickSeries.Merge(m, src)
}
func (m *TickSeries) XXX_Size() int {
	return xxx_messageInfo_TickSeries.Size(m)
}
func (m *TickSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_TickSeries.DiscardUnknown(m)
}

var xxx_messageInfo_TickSeries proto.InternalMessageInfo

func (m *TickSeries) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *TickSeries) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *TickSeries) GetTradingDay() int32 {
	if m != nil {
		return m.TradingDay
	}
	return 0
}

func (m *TickSeries) GetList() []*MarketDataSnapshot {
	if m != nil {
		return m.List
	}
	return nil
}

type SimpleTick struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time"`
	Price                float64  `protobuf:"fixed64,2,opt,name=price,proto3" json:"price"`
	Volume               int32    `protobuf:"varint,3,opt,name=volume,proto3" json:"volume"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleTick) Reset()         { *m = SimpleTick{} }
func (m *SimpleTick) String() string { return proto.CompactTextString(m) }
func (*SimpleTick) ProtoMessage()    {}
func (*SimpleTick) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{13}
}

func (m *SimpleTick) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleTick.Unmarshal(m, b)
}
func (m *SimpleTick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleTick.Marshal(b, m, deterministic)
}
func (m *SimpleTick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleTick.Merge(m, src)
}
func (m *SimpleTick) XXX_Size() int {
	return xxx_messageInfo_SimpleTick.Size(m)
}
func (m *SimpleTick) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleTick.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleTick proto.InternalMessageInfo

func (m *SimpleTick) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *SimpleTick) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *SimpleTick) GetVolume() int32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type SimpleTickSeries struct {
	Exchange             string        `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange"`
	Symbol               string        `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol"`
	TradingDay           int32         `protobuf:"varint,3,opt,name=trading_day,json=tradingDay,proto3" json:"tradingDay"`
	List                 []*SimpleTick `protobuf:"bytes,4,rep,name=list,proto3" json:"list"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SimpleTickSeries) Reset()         { *m = SimpleTickSeries{} }
func (m *SimpleTickSeries) String() string { return proto.CompactTextString(m) }
func (*SimpleTickSeries) ProtoMessage()    {}
func (*SimpleTickSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cf990f4c5aa8a63, []int{14}
}

func (m *SimpleTickSeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleTickSeries.Unmarshal(m, b)
}
func (m *SimpleTickSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleTickSeries.Marshal(b, m, deterministic)
}
func (m *SimpleTickSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleTickSeries.Merge(m, src)
}
func (m *SimpleTickSeries) XXX_Size() int {
	return xxx_messageInfo_SimpleTickSeries.Size(m)
}
func (m *SimpleTickSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleTickSeries.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleTickSeries proto.InternalMessageInfo

func (m *SimpleTickSeries) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *SimpleTickSeries) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *SimpleTickSeries) GetTradingDay() int32 {
	if m != nil {
		return m.TradingDay
	}
	return 0
}

func (m *SimpleTickSeries) GetList() []*SimpleTick {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderBook)(nil), "goshare.OrderBook")
	proto.RegisterType((*MarketDataSnapshot)(nil), "goshare.MarketDataSnapshot")
	proto.RegisterType((*MdsList)(nil), "goshare.MdsList")
	proto.RegisterType((*OptionTMarket)(nil), "goshare.OptionTMarket")
	proto.RegisterType((*SimpleTickForTQuote)(nil), "goshare.SimpleTickForTQuote")
	proto.RegisterType((*OptionTQuoteItem)(nil), "goshare.OptionTQuoteItem")
	proto.RegisterType((*OptionTQuoteItemList)(nil), "goshare.OptionTQuoteItemList")
	proto.RegisterType((*Kline)(nil), "goshare.Kline")
	proto.RegisterType((*KlineSeries)(nil), "goshare.KlineSeries")
	proto.RegisterType((*ReqSubscribeMarketData)(nil), "goshare.ReqSubscribeMarketData")
	proto.RegisterType((*RspSubscribeMarketData)(nil), "goshare.RspSubscribeMarketData")
	proto.RegisterType((*RtnMarketDataUpdate)(nil), "goshare.RtnMarketDataUpdate")
	proto.RegisterType((*TickSeries)(nil), "goshare.TickSeries")
	proto.RegisterType((*SimpleTick)(nil), "goshare.SimpleTick")
	proto.RegisterType((*SimpleTickSeries)(nil), "goshare.SimpleTickSeries")
}

func init() { proto.RegisterFile("goshare/market_data.proto", fileDescriptor_2cf990f4c5aa8a63) }

var fileDescriptor_2cf990f4c5aa8a63 = []byte{
	// 1119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xd6, 0xc6, 0x3f, 0x89, 0x8f, 0xed, 0x34, 0x19, 0x87, 0xb0, 0x49, 0x0a, 0xa4, 0x5b, 0xa1,
	0x86, 0x42, 0xed, 0xd2, 0xdc, 0x71, 0x59, 0xa2, 0x4a, 0x15, 0x2d, 0x0d, 0x6b, 0xc3, 0x05, 0x37,
	0xab, 0xf1, 0xee, 0x60, 0x8f, 0xf6, 0x67, 0x86, 0xdd, 0xd9, 0xba, 0xb9, 0xe4, 0x0d, 0x10, 0xd7,
	0x3c, 0x0b, 0xf7, 0x88, 0x17, 0xe0, 0x9e, 0x17, 0x41, 0x73, 0x66, 0xbc, 0xfe, 0x89, 0xa3, 0x54,
	0x41, 0xdc, 0xed, 0xf9, 0xe6, 0x9c, 0x39, 0xff, 0xdf, 0x2c, 0x1c, 0x4d, 0x44, 0x31, 0xa5, 0x39,
	0x1b, 0xa4, 0x34, 0x8f, 0x99, 0x0a, 0x22, 0xaa, 0x68, 0x5f, 0xe6, 0x42, 0x09, 0xb2, 0x6d, 0x8f,
	0x8e, 0x0f, 0xe6, 0x3a, 0xa1, 0x48, 0x53, 0x91, 0x99, 0x63, 0x2f, 0x85, 0xd6, 0x9b, 0x3c, 0x62,
	0xf9, 0x73, 0x21, 0x62, 0xb2, 0x07, 0x35, 0x5a, 0xc4, 0xae, 0x73, 0xea, 0x9c, 0x39, 0xbe, 0xfe,
	0x24, 0x1f, 0x01, 0xd0, 0x22, 0x0e, 0xde, 0x8a, 0xa4, 0x4c, 0x99, 0xbb, 0x75, 0xea, 0x9c, 0x35,
	0xfc, 0x16, 0x2d, 0xe2, 0x1f, 0x10, 0xd0, 0x06, 0x63, 0x1e, 0xb9, 0x35, 0x63, 0x30, 0xe6, 0x91,
	0x36, 0x18, 0xf3, 0x68, 0x6e, 0x50, 0x37, 0x06, 0x63, 0x1e, 0x19, 0x03, 0xef, 0x9f, 0x26, 0x90,
	0xd7, 0x18, 0xe3, 0x05, 0x55, 0x74, 0x98, 0x51, 0x59, 0x4c, 0x85, 0x22, 0xc7, 0xb0, 0xc3, 0xde,
	0x85, 0x53, 0x9a, 0x4d, 0x18, 0x7a, 0x6f, 0xf9, 0x95, 0x4c, 0x0e, 0xa1, 0x59, 0x5c, 0xa5, 0x63,
	0x91, 0xa0, 0xfb, 0x96, 0x6f, 0x25, 0x42, 0xa0, 0xae, 0x78, 0xca, 0xd0, 0x79, 0xcd, 0xc7, 0x6f,
	0xe2, 0x41, 0x27, 0xe5, 0x49, 0xc2, 0x0b, 0x16, 0x8a, 0x2c, 0x2a, 0xac, 0xff, 0x15, 0x4c, 0xdb,
	0x09, 0xc9, 0x32, 0xb7, 0x81, 0x41, 0xe3, 0xb7, 0xc6, 0xa6, 0x7c, 0x32, 0x75, 0x9b, 0x06, 0xd3,
	0xdf, 0x3a, 0xb7, 0x44, 0xcc, 0xdc, 0x6d, 0x93, 0x5b, 0x22, 0x66, 0xe4, 0x00, 0x1a, 0x61, 0x22,
	0x0a, 0xe6, 0xee, 0x20, 0x66, 0x04, 0x1d, 0x9f, 0xcd, 0xb6, 0x85, 0xde, 0xac, 0xa4, 0x71, 0x9a,
	0x8a, 0x32, 0x53, 0x2e, 0xa0, 0xba, 0x95, 0x74, 0xae, 0x52, 0x14, 0x5c, 0x71, 0x91, 0xb9, 0x6d,
	0xb4, 0xa8, 0x64, 0xed, 0x41, 0xe6, 0x3c, 0x64, 0x6e, 0xc7, 0x78, 0x40, 0x81, 0x9c, 0x40, 0x4b,
	0xe6, 0x2c, 0x30, 0xbe, 0xbb, 0x78, 0xb2, 0x23, 0x73, 0xf6, 0x35, 0xba, 0xff, 0x14, 0x76, 0xf5,
	0x61, 0xc1, 0x94, 0x4a, 0x58, 0xca, 0x32, 0xe5, 0xee, 0xa2, 0x46, 0x57, 0xe6, 0x6c, 0x58, 0x81,
	0xe4, 0x01, 0x74, 0xb4, 0x5a, 0xe5, 0xf9, 0x1e, 0x7a, 0x6e, 0xcb, 0x9c, 0x5d, 0xce, 0x9d, 0x7f,
	0x06, 0x7b, 0x8b, 0x5b, 0x02, 0x13, 0xc7, 0x1e, 0xde, 0x75, 0x6f, 0x81, 0x5f, 0x62, 0x44, 0x9f,
	0x40, 0xbb, 0x94, 0x92, 0xe5, 0x41, 0xc2, 0x53, 0xae, 0xdc, 0x7d, 0xd4, 0x02, 0x84, 0x5e, 0x69,
	0x44, 0x2b, 0x24, 0x62, 0x56, 0x29, 0x10, 0xa3, 0x80, 0x90, 0x51, 0xb0, 0x39, 0x45, 0x2c, 0x51,
	0xd4, 0xed, 0x55, 0x39, 0x5d, 0x68, 0x59, 0x97, 0xc1, 0x1c, 0x1c, 0x98, 0x32, 0xa0, 0x40, 0x1e,
	0x42, 0x97, 0xbe, 0x65, 0x39, 0x9d, 0x30, 0x1b, 0xdc, 0x07, 0x78, 0xda, 0xb1, 0x60, 0x15, 0x99,
	0xca, 0x69, 0xc4, 0xb3, 0x49, 0x10, 0xd1, 0x2b, 0xf7, 0x10, 0xd3, 0x04, 0x0b, 0x5d, 0xd0, 0x2b,
	0xdd, 0xea, 0x8c, 0xa6, 0xcc, 0xfd, 0x10, 0x87, 0x09, 0xbf, 0x75, 0x0d, 0xd9, 0x3b, 0x96, 0x87,
	0xbc, 0x98, 0x5f, 0xed, 0x9a, 0x1a, 0xce, 0x51, 0x73, 0xf7, 0x03, 0xe8, 0x98, 0xde, 0xda, 0xb0,
	0x8f, 0x4c, 0x0d, 0x0d, 0x66, 0x22, 0xd7, 0xfb, 0x12, 0xea, 0x6a, 0xa2, 0xf7, 0x13, 0xbb, 0x2f,
	0x88, 0x68, 0xe7, 0xc7, 0xb0, 0x93, 0x96, 0x89, 0xe2, 0x32, 0x61, 0xee, 0x7d, 0xd3, 0xfb, 0xb9,
	0xac, 0x4d, 0xd1, 0x77, 0xa0, 0x78, 0x18, 0xbb, 0x8f, 0x30, 0x80, 0x16, 0x22, 0x23, 0x1e, 0xc6,
	0xe4, 0x31, 0x34, 0x23, 0x26, 0xd5, 0xb4, 0x70, 0x3f, 0x3e, 0xad, 0x9d, 0xb5, 0x9f, 0x91, 0xbe,
	0xdd, 0xe7, 0x7e, 0xb5, 0xbf, 0xbe, 0xd5, 0xf0, 0xbe, 0x82, 0xed, 0xd7, 0x51, 0xf1, 0x8a, 0x17,
	0x8a, 0x0c, 0xa0, 0x9e, 0xf0, 0x42, 0xb9, 0x0e, 0x1a, 0x9d, 0x54, 0x46, 0xd7, 0x97, 0xd0, 0x47,
	0x45, 0x6f, 0x06, 0xdd, 0x37, 0x52, 0xc7, 0x3b, 0x32, 0x2a, 0xe4, 0x1c, 0x9a, 0x21, 0x4d, 0x92,
	0x91, 0xe1, 0x85, 0x5b, 0xee, 0xb0, 0xaa, 0xe4, 0x4b, 0x68, 0xc8, 0x52, 0x8d, 0x62, 0xdc, 0xd9,
	0x5b, 0x6c, 0x8c, 0xa6, 0xf7, 0xa7, 0x03, 0xbd, 0x21, 0x4f, 0x65, 0x82, 0xf9, 0xbe, 0x10, 0xf9,
	0xe8, 0xbb, 0x52, 0x28, 0x76, 0x27, 0x6e, 0xa8, 0xf6, 0xa8, 0xb6, 0xbc, 0x47, 0x1e, 0x74, 0x4b,
	0x19, 0x44, 0x62, 0x96, 0x05, 0x39, 0x55, 0x5c, 0x20, 0x3d, 0x38, 0x7e, 0xbb, 0x94, 0x17, 0x62,
	0x96, 0xf9, 0x1a, 0x22, 0x4f, 0xe1, 0x60, 0x75, 0x9d, 0xec, 0x40, 0x18, 0xb6, 0x20, 0x2b, 0x4b,
	0x65, 0xa6, 0x62, 0x3e, 0x50, 0xcd, 0xc5, 0x40, 0x79, 0x7f, 0x39, 0xb0, 0x67, 0xab, 0x88, 0x49,
	0xbc, 0x54, 0x2c, 0xdd, 0x30, 0x65, 0xce, 0xa6, 0x29, 0x7b, 0x0a, 0x75, 0x5d, 0x44, 0x5b, 0xb9,
	0xfb, 0x55, 0xe5, 0x36, 0xd4, 0xc6, 0x47, 0x4d, 0xd2, 0x87, 0x9a, 0x2c, 0x15, 0xe6, 0x7a, 0x9b,
	0x81, 0x56, 0x24, 0x7d, 0xe8, 0xad, 0x06, 0x12, 0xfc, 0x94, 0xd0, 0x09, 0x56, 0xa3, 0xe5, 0xef,
	0xaf, 0x44, 0xf3, 0x22, 0xa1, 0x13, 0xef, 0x77, 0x07, 0x0e, 0xd6, 0xb3, 0xc1, 0xe1, 0x5a, 0x6f,
	0x4d, 0x63, 0xa9, 0x35, 0x0f, 0xa1, 0x5b, 0xa8, 0x9c, 0xc7, 0x2c, 0x58, 0xe9, 0x50, 0xc7, 0x80,
	0xc3, 0xaa, 0x4f, 0xa9, 0xc8, 0xd4, 0x14, 0x63, 0x6f, 0xf9, 0x46, 0x20, 0x4f, 0xec, 0xcc, 0xd6,
	0x71, 0x66, 0x8f, 0x16, 0x83, 0xbe, 0x16, 0x83, 0x9d, 0xd8, 0xbf, 0x1d, 0x68, 0x7c, 0x93, 0xf0,
	0x8c, 0x55, 0x4f, 0x82, 0xb3, 0xf4, 0x24, 0xcc, 0xe9, 0x7e, 0x6b, 0x03, 0xdd, 0xd7, 0xae, 0xd3,
	0x7d, 0x7d, 0x03, 0xdd, 0x37, 0x36, 0xd3, 0x7d, 0xf3, 0x06, 0xba, 0xdf, 0xbe, 0x91, 0xee, 0x77,
	0xd6, 0xe8, 0x7e, 0x8d, 0xac, 0x5a, 0xeb, 0x64, 0xe5, 0xfd, 0xe1, 0x40, 0x1b, 0x53, 0x1b, 0xb2,
	0x9c, 0xb3, 0xe2, 0x4e, 0xbb, 0xf0, 0x39, 0x34, 0x25, 0xcb, 0xb9, 0x30, 0xcf, 0xf4, 0xee, 0xb3,
	0x5e, 0x55, 0xcf, 0x4b, 0x84, 0x47, 0x57, 0x92, 0xf9, 0x56, 0x85, 0x3c, 0x86, 0x7d, 0xf3, 0x15,
	0xf0, 0x2c, 0x58, 0x7d, 0x45, 0xef, 0x99, 0x83, 0x97, 0xd9, 0xd0, 0x3e, 0xa4, 0x9e, 0x6d, 0x53,
	0x03, 0xdb, 0xb4, 0x5b, 0x5d, 0x8b, 0x01, 0xdb, 0xde, 0xfc, 0xe2, 0xc0, 0xa1, 0xcf, 0x7e, 0x1e,
	0x96, 0xe3, 0x22, 0xcc, 0xf9, 0x98, 0x2d, 0xd6, 0xff, 0x4e, 0xb9, 0x9c, 0x03, 0x98, 0x28, 0xf4,
	0xf8, 0xb9, 0xb5, 0xd3, 0xda, 0x4d, 0xf9, 0x2c, 0xa9, 0x79, 0x2e, 0x1c, 0xfa, 0x85, 0xdc, 0x10,
	0x82, 0x57, 0x42, 0xcf, 0x57, 0xd9, 0x02, 0xf8, 0x5e, 0x46, 0x54, 0x31, 0xcd, 0x99, 0xc8, 0xc1,
	0xef, 0xc1, 0x77, 0xa8, 0x48, 0x9e, 0x00, 0xc4, 0x3a, 0xe9, 0x00, 0xeb, 0xb1, 0xb5, 0xb1, 0x1e,
	0x2d, 0xd4, 0xc0, 0x80, 0x7e, 0x73, 0x00, 0xf4, 0x5a, 0xfe, 0x87, 0xa6, 0xae, 0x4d, 0x4e, 0xed,
	0xda, 0x33, 0x37, 0x58, 0xd9, 0xa1, 0xf7, 0xe0, 0xfd, 0x6f, 0x01, 0x16, 0x84, 0xb1, 0x71, 0x93,
	0x2a, 0x52, 0xdd, 0x5a, 0x26, 0xd5, 0xc5, 0x3e, 0xd4, 0x96, 0xf7, 0xc1, 0xfb, 0xd5, 0x81, 0xbd,
	0xc5, 0x85, 0xff, 0x67, 0xaa, 0x8f, 0x56, 0x52, 0xed, 0x6d, 0xe0, 0x3f, 0x93, 0xe2, 0xf3, 0xfe,
	0x8f, 0x5f, 0x4c, 0xb8, 0x9a, 0x96, 0xe3, 0x7e, 0x28, 0xd2, 0x41, 0xca, 0x33, 0x96, 0xd3, 0x24,
	0x67, 0xc5, 0x60, 0xfe, 0x67, 0x2c, 0xe3, 0xc9, 0x40, 0x8e, 0xe7, 0xe2, 0xb8, 0x89, 0xbf, 0xc8,
	0xe7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x0b, 0xdc, 0x7c, 0x5e, 0x0b, 0x00, 0x00,
}

