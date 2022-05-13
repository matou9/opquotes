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

type PeriodType int32

const (
	// Tick
	PeriodType_TICK PeriodType = 0
	// 分钟
	PeriodType_M1 PeriodType = 1
	// 3分钟
	PeriodType_M3 PeriodType = 2
	// 5分钟
	PeriodType_M5 PeriodType = 3
	// 10分钟
	PeriodType_M10 PeriodType = 4
	// 15分钟
	PeriodType_M15 PeriodType = 5
	// 30分钟
	PeriodType_M30 PeriodType = 6
	// 小时
	PeriodType_H1 PeriodType = 7
	// 3小时
	PeriodType_H3 PeriodType = 8
	// 日线
	PeriodType_D1 PeriodType = 9
	// 周线
	PeriodType_W1 PeriodType = 10
	// 月线
	PeriodType_MON1 PeriodType = 11
)

var PeriodType_name = map[int32]string{
	0:  "TICK",
	1:  "M1",
	2:  "M3",
	3:  "M5",
	4:  "M10",
	5:  "M15",
	6:  "M30",
	7:  "H1",
	8:  "H3",
	9:  "D1",
	10: "W1",
	11: "MON1",
}

var PeriodType_value = map[string]int32{
	"TICK": 0,
	"M1":   1,
	"M3":   2,
	"M5":   3,
	"M10":  4,
	"M15":  5,
	"M30":  6,
	"H1":   7,
	"H3":   8,
	"D1":   9,
	"W1":   10,
	"MON1": 11,
}

func (x PeriodType) String() string {
	return proto.EnumName(PeriodType_name, int32(x))
}

func (PeriodType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{0}
}

// OptionCallPutType 期权call or put 类型
type OptionCallPutType int32

const (
	OptionCallPutType_OCPT_CALL OptionCallPutType = 0
	OptionCallPutType_OCPT_PUT  OptionCallPutType = 1
)

var OptionCallPutType_name = map[int32]string{
	0: "OCPT_CALL",
	1: "OCPT_PUT",
}

var OptionCallPutType_value = map[string]int32{
	"OCPT_CALL": 0,
	"OCPT_PUT":  1,
}

func (x OptionCallPutType) String() string {
	return proto.EnumName(OptionCallPutType_name, int32(x))
}

func (OptionCallPutType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{1}
}

// 行权日类型
type OptionDeliveryDateType int32

const (
	OptionDeliveryDateType_ODDT_EUR OptionDeliveryDateType = 0
)

var OptionDeliveryDateType_name = map[int32]string{
	0: "ODDT_EUR",
}

var OptionDeliveryDateType_value = map[string]int32{
	"ODDT_EUR": 0,
}

func (x OptionDeliveryDateType) String() string {
	return proto.EnumName(OptionDeliveryDateType_name, int32(x))
}

func (OptionDeliveryDateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{2}
}

// 平仓手续费算法
type CloseCommissionAlgorithim int32

const (
	// 普通
	CloseCommissionAlgorithim_CCA_NORMAL CloseCommissionAlgorithim = 0
	// 收一次
	CloseCommissionAlgorithim_CCA_ONCE CloseCommissionAlgorithim = 1
	// 隔夜日期
	CloseCommissionAlgorithim_CCA_MULTIPLE_BY_DATE CloseCommissionAlgorithim = 2
)

var CloseCommissionAlgorithim_name = map[int32]string{
	0: "CCA_NORMAL",
	1: "CCA_ONCE",
	2: "CCA_MULTIPLE_BY_DATE",
}

var CloseCommissionAlgorithim_value = map[string]int32{
	"CCA_NORMAL":           0,
	"CCA_ONCE":             1,
	"CCA_MULTIPLE_BY_DATE": 2,
}

func (x CloseCommissionAlgorithim) String() string {
	return proto.EnumName(CloseCommissionAlgorithim_name, int32(x))
}

func (CloseCommissionAlgorithim) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{3}
}

type AccountType int32

const (
	AccountType_CN_FUTURE AccountType = 0
	AccountType_CN_STOCK  AccountType = 1
)

var AccountType_name = map[int32]string{
	0: "CN_FUTURE",
	1: "CN_STOCK",
}

var AccountType_value = map[string]int32{
	"CN_FUTURE": 0,
	"CN_STOCK":  1,
}

func (x AccountType) String() string {
	return proto.EnumName(AccountType_name, int32(x))
}

func (AccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{4}
}

type CurrencyType int32

const (
	CurrencyType_CNY CurrencyType = 0
	CurrencyType_USD CurrencyType = 1
)

var CurrencyType_name = map[int32]string{
	0: "CNY",
	1: "USD",
}

var CurrencyType_value = map[string]int32{
	"CNY": 0,
	"USD": 1,
}

func (x CurrencyType) String() string {
	return proto.EnumName(CurrencyType_name, int32(x))
}

func (CurrencyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{5}
}

// 交易账号类型
type TradingAccountType int32

const (
	TradingAccountType_TAT_NORMAL     TradingAccountType = 0
	TradingAccountType_TAT_TRAINNING  TradingAccountType = 1
	TradingAccountType_TAT_CTP_MIRROR TradingAccountType = 2
)

var TradingAccountType_name = map[int32]string{
	0: "TAT_NORMAL",
	1: "TAT_TRAINNING",
	2: "TAT_CTP_MIRROR",
}

var TradingAccountType_value = map[string]int32{
	"TAT_NORMAL":     0,
	"TAT_TRAINNING":  1,
	"TAT_CTP_MIRROR": 2,
}

func (x TradingAccountType) String() string {
	return proto.EnumName(TradingAccountType_name, int32(x))
}

func (TradingAccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{6}
}

// 经纪商通道
type BrokerRoute struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	MdFrontList          []string `protobuf:"bytes,4,rep,name=md_front_list,json=mdFrontList,proto3" json:"mdFrontList"`
	TradeFrontList       []string `protobuf:"bytes,5,rep,name=trade_front_list,json=tradeFrontList,proto3" json:"tradeFrontList"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BrokerRoute) Reset()         { *m = BrokerRoute{} }
func (m *BrokerRoute) String() string { return proto.CompactTextString(m) }
func (*BrokerRoute) ProtoMessage()    {}
func (*BrokerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{0}
}

func (m *BrokerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrokerRoute.Unmarshal(m, b)
}
func (m *BrokerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrokerRoute.Marshal(b, m, deterministic)
}
func (m *BrokerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrokerRoute.Merge(m, src)
}
func (m *BrokerRoute) XXX_Size() int {
	return xxx_messageInfo_BrokerRoute.Size(m)
}
func (m *BrokerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_BrokerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_BrokerRoute proto.InternalMessageInfo

func (m *BrokerRoute) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BrokerRoute) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *BrokerRoute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BrokerRoute) GetMdFrontList() []string {
	if m != nil {
		return m.MdFrontList
	}
	return nil
}

func (m *BrokerRoute) GetTradeFrontList() []string {
	if m != nil {
		return m.TradeFrontList
	}
	return nil
}

type BrokerRouteList struct {
	List                 []*BrokerRoute `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BrokerRouteList) Reset()         { *m = BrokerRouteList{} }
func (m *BrokerRouteList) String() string { return proto.CompactTextString(m) }
func (*BrokerRouteList) ProtoMessage()    {}
func (*BrokerRouteList) Descriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{1}
}

func (m *BrokerRouteList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrokerRouteList.Unmarshal(m, b)
}
func (m *BrokerRouteList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrokerRouteList.Marshal(b, m, deterministic)
}
func (m *BrokerRouteList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrokerRouteList.Merge(m, src)
}
func (m *BrokerRouteList) XXX_Size() int {
	return xxx_messageInfo_BrokerRouteList.Size(m)
}
func (m *BrokerRouteList) XXX_DiscardUnknown() {
	xxx_messageInfo_BrokerRouteList.DiscardUnknown(m)
}

var xxx_messageInfo_BrokerRouteList proto.InternalMessageInfo

func (m *BrokerRouteList) GetList() []*BrokerRoute {
	if m != nil {
		return m.List
	}
	return nil
}

type ReqUpdateTIOpenDate struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol"`
	OpenDate             int32    `protobuf:"varint,3,opt,name=open_date,json=openDate,proto3" json:"openDate"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUpdateTIOpenDate) Reset()         { *m = ReqUpdateTIOpenDate{} }
func (m *ReqUpdateTIOpenDate) String() string { return proto.CompactTextString(m) }
func (*ReqUpdateTIOpenDate) ProtoMessage()    {}
func (*ReqUpdateTIOpenDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{2}
}

func (m *ReqUpdateTIOpenDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUpdateTIOpenDate.Unmarshal(m, b)
}
func (m *ReqUpdateTIOpenDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUpdateTIOpenDate.Marshal(b, m, deterministic)
}
func (m *ReqUpdateTIOpenDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUpdateTIOpenDate.Merge(m, src)
}
func (m *ReqUpdateTIOpenDate) XXX_Size() int {
	return xxx_messageInfo_ReqUpdateTIOpenDate.Size(m)
}
func (m *ReqUpdateTIOpenDate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUpdateTIOpenDate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUpdateTIOpenDate proto.InternalMessageInfo

func (m *ReqUpdateTIOpenDate) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *ReqUpdateTIOpenDate) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqUpdateTIOpenDate) GetOpenDate() int32 {
	if m != nil {
		return m.OpenDate
	}
	return 0
}

type ReqUpdateTIOpenDateList struct {
	Exchange             int32                  `protobuf:"varint,1,opt,name=exchange,proto3" json:"exchange"`
	List                 []*ReqUpdateTIOpenDate `protobuf:"bytes,2,rep,name=list,proto3" json:"list"`
	OpCode               string                 `protobuf:"bytes,3,opt,name=op_code,json=opCode,proto3" json:"opCode"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ReqUpdateTIOpenDateList) Reset()         { *m = ReqUpdateTIOpenDateList{} }
func (m *ReqUpdateTIOpenDateList) String() string { return proto.CompactTextString(m) }
func (*ReqUpdateTIOpenDateList) ProtoMessage()    {}
func (*ReqUpdateTIOpenDateList) Descriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{3}
}

func (m *ReqUpdateTIOpenDateList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUpdateTIOpenDateList.Unmarshal(m, b)
}
func (m *ReqUpdateTIOpenDateList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUpdateTIOpenDateList.Marshal(b, m, deterministic)
}
func (m *ReqUpdateTIOpenDateList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUpdateTIOpenDateList.Merge(m, src)
}
func (m *ReqUpdateTIOpenDateList) XXX_Size() int {
	return xxx_messageInfo_ReqUpdateTIOpenDateList.Size(m)
}
func (m *ReqUpdateTIOpenDateList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUpdateTIOpenDateList.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUpdateTIOpenDateList proto.InternalMessageInfo

func (m *ReqUpdateTIOpenDateList) GetExchange() int32 {
	if m != nil {
		return m.Exchange
	}
	return 0
}

func (m *ReqUpdateTIOpenDateList) GetList() []*ReqUpdateTIOpenDate {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReqUpdateTIOpenDateList) GetOpCode() string {
	if m != nil {
		return m.OpCode
	}
	return ""
}

// 结算账户资金信息
type AccountMoneySummary struct {
	// 账号
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 结算组（可能单个结算组给个摘要)
	Type AccountType `protobuf:"varint,2,opt,name=type,proto3,enum=goshare.AccountType" json:"type"`
	// 币种
	CurrencyType CurrencyType `protobuf:"varint,3,opt,name=currency_type,json=currencyType,proto3,enum=goshare.CurrencyType" json:"currencyType"`
	/// 余额
	Balance float64 `protobuf:"fixed64,4,opt,name=balance,proto3" json:"balance"`
	/// 手续费
	Commission float64 `protobuf:"fixed64,5,opt,name=commission,proto3" json:"commission"`
	/// 占用保证金
	CurMargin float64 `protobuf:"fixed64,6,opt,name=cur_margin,json=curMargin,proto3" json:"curMargin"`
	/// 持仓市值.
	NetPositionAmount float64 `protobuf:"fixed64,7,opt,name=net_position_amount,json=netPositionAmount,proto3" json:"netPositionAmount"`
	/// 上次余额
	PreBalance float64 `protobuf:"fixed64,8,opt,name=pre_balance,json=preBalance,proto3" json:"preBalance"`
	/// 利息收入
	InterestIn float64 `protobuf:"fixed64,9,opt,name=interest_in,json=interestIn,proto3" json:"interestIn"`
	/// 存款
	Deposit float64 `protobuf:"fixed64,10,opt,name=deposit,proto3" json:"deposit"`
	/// 取款
	Withdraw float64 `protobuf:"fixed64,11,opt,name=withdraw,proto3" json:"withdraw"`
	/// 冻结保证金
	FrozenMargin float64 `protobuf:"fixed64,12,opt,name=frozen_margin,json=frozenMargin,proto3" json:"frozenMargin"`
	/// 冻结手续费
	FrozenCommission float64 `protobuf:"fixed64,13,opt,name=frozen_commission,json=frozenCommission,proto3" json:"frozenCommission"`
	/// 可用资金
	Available float64 `protobuf:"fixed64,14,opt,name=available,proto3" json:"available"`
	/// 平仓盈亏
	CloseProfit float64 `protobuf:"fixed64,15,opt,name=close_profit,json=closeProfit,proto3" json:"closeProfit"`
	/// 持仓盈亏
	PositionProfit float64 `protobuf:"fixed64,16,opt,name=position_profit,json=positionProfit,proto3" json:"positionProfit"`
	// 上次质押金额
	PreMortgate float64 `protobuf:"fixed64,17,opt,name=pre_mortgate,json=preMortgate,proto3" json:"preMortgate"`
	// 上次信用额度
	PreCredit float64 `protobuf:"fixed64,18,opt,name=pre_credit,json=preCredit,proto3" json:"preCredit"`
	// 利息基数
	InterestBase float64 `protobuf:"fixed64,19,opt,name=interest_base,json=interestBase,proto3" json:"interestBase"`
	// 可取资金
	WithdrawAvailable float64 `protobuf:"fixed64,20,opt,name=withdraw_available,json=withdrawAvailable,proto3" json:"withdrawAvailable"`
	// 基本准备金
	Reserve float64 `protobuf:"fixed64,21,opt,name=reserve,proto3" json:"reserve"`
	// 交易日
	TradingDay int32 `protobuf:"varint,22,opt,name=trading_day,json=tradingDay,proto3" json:"tradingDay"`
	// 信用额度
	Credit float64 `protobuf:"fixed64,23,opt,name=credit,proto3" json:"credit"`
	// 质押金额
	Mortgate float64 `protobuf:"fixed64,24,opt,name=mortgate,proto3" json:"mortgate"`
	// 交易所保证金
	ExchangeMargin float64 `protobuf:"fixed64,25,opt,name=exchange_margin,json=exchangeMargin,proto3" json:"exchangeMargin"`
	// 交割保证金
	DeliveryMargin float64 `protobuf:"fixed64,26,opt,name=delivery_margin,json=deliveryMargin,proto3" json:"deliveryMargin"`
	// 交易所交割保证金
	ExchangeDeliveryMargin float64 `protobuf:"fixed64,27,opt,name=exchange_delivery_margin,json=exchangeDeliveryMargin,proto3" json:"exchangeDeliveryMargin"`
	// 保底期货结算准备金
	ReserveBalance float64 `protobuf:"fixed64,28,opt,name=reserve_balance,json=reserveBalance,proto3" json:"reserveBalance"`
	// 交割手续费
	DeliveryCommission float64 `protobuf:"fixed64,29,opt,name=delivery_commission,json=deliveryCommission,proto3" json:"deliveryCommission"`
	// 冻结过户费
	FrozenTransferFee float64 `protobuf:"fixed64,30,opt,name=frozen_transfer_fee,json=frozenTransferFee,proto3" json:"frozenTransferFee"`
	// 冻结的印花税
	FrozenStampTax float64 `protobuf:"fixed64,31,opt,name=frozen_stamp_tax,json=frozenStampTax,proto3" json:"frozenStampTax"`
	// 过户费
	TransferFee float64 `protobuf:"fixed64,32,opt,name=transfer_fee,json=transferFee,proto3" json:"transferFee"`
	// 印花税
	StampTax float64 `protobuf:"fixed64,33,opt,name=stamp_tax,json=stampTax,proto3" json:"stampTax"`
	/// 盯市盈亏
	MtmProfit float64 `protobuf:"fixed64,34,opt,name=mtm_profit,json=mtmProfit,proto3" json:"mtmProfit"`
	///授信额度
	PreMtmProfit float64 `protobuf:"fixed64,35,opt,name=pre_mtm_profit,json=preMtmProfit,proto3" json:"preMtmProfit"`
	///证券总价值
	StockValue float64 `protobuf:"fixed64,36,opt,name=stock_value,json=stockValue,proto3" json:"stockValue"`
	///国债回购占用资金
	BondRepurchaseAmount float64 `protobuf:"fixed64,37,opt,name=bond_repurchase_amount,json=bondRepurchaseAmount,proto3" json:"bondRepurchaseAmount"`
	///国债逆回购占用资金
	ReverseRepurchaseAmount float64 `protobuf:"fixed64,38,opt,name=reverse_repurchase_amount,json=reverseRepurchaseAmount,proto3" json:"reverseRepurchaseAmount"`
	///融资买入金额
	MarginTradeAmount float64 `protobuf:"fixed64,39,opt,name=margin_trade_amount,json=marginTradeAmount,proto3" json:"marginTradeAmount"`
	///融券卖出金额
	ShortSellAmount float64 `protobuf:"fixed64,40,opt,name=short_sell_amount,json=shortSellAmount,proto3" json:"shortSellAmount"`
	///融资持仓盈亏
	MarginTradeProfit float64 `protobuf:"fixed64,41,opt,name=margin_trade_profit,json=marginTradeProfit,proto3" json:"marginTradeProfit"`
	///融券持仓盈亏
	ShortSellProfit float64 `protobuf:"fixed64,42,opt,name=short_sell_profit,json=shortSellProfit,proto3" json:"shortSellProfit"`
	// 冻结平仓所需要费用
	FrozenCloseCommission float64 `protobuf:"fixed64,43,opt,name=frozen_close_commission,json=frozenCloseCommission,proto3" json:"frozenCloseCommission"`
	// 劣后
	CommonBalance float64 `protobuf:"fixed64,44,opt,name=common_balance,json=commonBalance,proto3" json:"commonBalance"`
	// 优先资金
	PreferedBalance float64 `protobuf:"fixed64,45,opt,name=prefered_balance,json=preferedBalance,proto3" json:"preferedBalance"`
	// 兑换基础货币汇率
	BaseExchangeRate float64 `protobuf:"fixed64,46,opt,name=base_exchange_rate,json=baseExchangeRate,proto3" json:"baseExchangeRate"`
	// 报警
	WarningLevel float64 `protobuf:"fixed64,47,opt,name=warning_level,json=warningLevel,proto3" json:"warningLevel"`
	// 强平
	ForceCloseLevel float64 `protobuf:"fixed64,48,opt,name=force_close_level,json=forceCloseLevel,proto3" json:"forceCloseLevel"`
	// 进入清算的平仓盈亏
	ClearedClosedProfit float64 `protobuf:"fixed64,49,opt,name=cleared_closed_profit,json=clearedClosedProfit,proto3" json:"clearedClosedProfit"`
	// 名称
	Name string `protobuf:"bytes,50,opt,name=name,proto3" json:"name"`
	// 利息
	Interest float64 `protobuf:"fixed64,51,opt,name=interest,proto3" json:"interest"`
	// 冻结利息
	FrozenInterest       float64  `protobuf:"fixed64,52,opt,name=frozen_interest,json=frozenInterest,proto3" json:"frozenInterest"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountMoneySummary) Reset()         { *m = AccountMoneySummary{} }
func (m *AccountMoneySummary) String() string { return proto.CompactTextString(m) }
func (*AccountMoneySummary) ProtoMessage()    {}
func (*AccountMoneySummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{4}
}

func (m *AccountMoneySummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountMoneySummary.Unmarshal(m, b)
}
func (m *AccountMoneySummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountMoneySummary.Marshal(b, m, deterministic)
}
func (m *AccountMoneySummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountMoneySummary.Merge(m, src)
}
func (m *AccountMoneySummary) XXX_Size() int {
	return xxx_messageInfo_AccountMoneySummary.Size(m)
}
func (m *AccountMoneySummary) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountMoneySummary.DiscardUnknown(m)
}

var xxx_messageInfo_AccountMoneySummary proto.InternalMessageInfo

func (m *AccountMoneySummary) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AccountMoneySummary) GetType() AccountType {
	if m != nil {
		return m.Type
	}
	return AccountType_CN_FUTURE
}

func (m *AccountMoneySummary) GetCurrencyType() CurrencyType {
	if m != nil {
		return m.CurrencyType
	}
	return CurrencyType_CNY
}

func (m *AccountMoneySummary) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *AccountMoneySummary) GetCommission() float64 {
	if m != nil {
		return m.Commission
	}
	return 0
}

func (m *AccountMoneySummary) GetCurMargin() float64 {
	if m != nil {
		return m.CurMargin
	}
	return 0
}

func (m *AccountMoneySummary) GetNetPositionAmount() float64 {
	if m != nil {
		return m.NetPositionAmount
	}
	return 0
}

func (m *AccountMoneySummary) GetPreBalance() float64 {
	if m != nil {
		return m.PreBalance
	}
	return 0
}

func (m *AccountMoneySummary) GetInterestIn() float64 {
	if m != nil {
		return m.InterestIn
	}
	return 0
}

func (m *AccountMoneySummary) GetDeposit() float64 {
	if m != nil {
		return m.Deposit
	}
	return 0
}

func (m *AccountMoneySummary) GetWithdraw() float64 {
	if m != nil {
		return m.Withdraw
	}
	return 0
}

func (m *AccountMoneySummary) GetFrozenMargin() float64 {
	if m != nil {
		return m.FrozenMargin
	}
	return 0
}

func (m *AccountMoneySummary) GetFrozenCommission() float64 {
	if m != nil {
		return m.FrozenCommission
	}
	return 0
}

func (m *AccountMoneySummary) GetAvailable() float64 {
	if m != nil {
		return m.Available
	}
	return 0
}

func (m *AccountMoneySummary) GetCloseProfit() float64 {
	if m != nil {
		return m.CloseProfit
	}
	return 0
}

func (m *AccountMoneySummary) GetPositionProfit() float64 {
	if m != nil {
		return m.PositionProfit
	}
	return 0
}

func (m *AccountMoneySummary) GetPreMortgate() float64 {
	if m != nil {
		return m.PreMortgate
	}
	return 0
}

func (m *AccountMoneySummary) GetPreCredit() float64 {
	if m != nil {
		return m.PreCredit
	}
	return 0
}

func (m *AccountMoneySummary) GetInterestBase() float64 {
	if m != nil {
		return m.InterestBase
	}
	return 0
}

func (m *AccountMoneySummary) GetWithdrawAvailable() float64 {
	if m != nil {
		return m.WithdrawAvailable
	}
	return 0
}

func (m *AccountMoneySummary) GetReserve() float64 {
	if m != nil {
		return m.Reserve
	}
	return 0
}

func (m *AccountMoneySummary) GetTradingDay() int32 {
	if m != nil {
		return m.TradingDay
	}
	return 0
}

func (m *AccountMoneySummary) GetCredit() float64 {
	if m != nil {
		return m.Credit
	}
	return 0
}

func (m *AccountMoneySummary) GetMortgate() float64 {
	if m != nil {
		return m.Mortgate
	}
	return 0
}

func (m *AccountMoneySummary) GetExchangeMargin() float64 {
	if m != nil {
		return m.ExchangeMargin
	}
	return 0
}

func (m *AccountMoneySummary) GetDeliveryMargin() float64 {
	if m != nil {
		return m.DeliveryMargin
	}
	return 0
}

func (m *AccountMoneySummary) GetExchangeDeliveryMargin() float64 {
	if m != nil {
		return m.ExchangeDeliveryMargin
	}
	return 0
}

func (m *AccountMoneySummary) GetReserveBalance() float64 {
	if m != nil {
		return m.ReserveBalance
	}
	return 0
}

func (m *AccountMoneySummary) GetDeliveryCommission() float64 {
	if m != nil {
		return m.DeliveryCommission
	}
	return 0
}

func (m *AccountMoneySummary) GetFrozenTransferFee() float64 {
	if m != nil {
		return m.FrozenTransferFee
	}
	return 0
}

func (m *AccountMoneySummary) GetFrozenStampTax() float64 {
	if m != nil {
		return m.FrozenStampTax
	}
	return 0
}

func (m *AccountMoneySummary) GetTransferFee() float64 {
	if m != nil {
		return m.TransferFee
	}
	return 0
}

func (m *AccountMoneySummary) GetStampTax() float64 {
	if m != nil {
		return m.StampTax
	}
	return 0
}

func (m *AccountMoneySummary) GetMtmProfit() float64 {
	if m != nil {
		return m.MtmProfit
	}
	return 0
}

func (m *AccountMoneySummary) GetPreMtmProfit() float64 {
	if m != nil {
		return m.PreMtmProfit
	}
	return 0
}

func (m *AccountMoneySummary) GetStockValue() float64 {
	if m != nil {
		return m.StockValue
	}
	return 0
}

func (m *AccountMoneySummary) GetBondRepurchaseAmount() float64 {
	if m != nil {
		return m.BondRepurchaseAmount
	}
	return 0
}

func (m *AccountMoneySummary) GetReverseRepurchaseAmount() float64 {
	if m != nil {
		return m.ReverseRepurchaseAmount
	}
	return 0
}

func (m *AccountMoneySummary) GetMarginTradeAmount() float64 {
	if m != nil {
		return m.MarginTradeAmount
	}
	return 0
}

func (m *AccountMoneySummary) GetShortSellAmount() float64 {
	if m != nil {
		return m.ShortSellAmount
	}
	return 0
}

func (m *AccountMoneySummary) GetMarginTradeProfit() float64 {
	if m != nil {
		return m.MarginTradeProfit
	}
	return 0
}

func (m *AccountMoneySummary) GetShortSellProfit() float64 {
	if m != nil {
		return m.ShortSellProfit
	}
	return 0
}

func (m *AccountMoneySummary) GetFrozenCloseCommission() float64 {
	if m != nil {
		return m.FrozenCloseCommission
	}
	return 0
}

func (m *AccountMoneySummary) GetCommonBalance() float64 {
	if m != nil {
		return m.CommonBalance
	}
	return 0
}

func (m *AccountMoneySummary) GetPreferedBalance() float64 {
	if m != nil {
		return m.PreferedBalance
	}
	return 0
}

func (m *AccountMoneySummary) GetBaseExchangeRate() float64 {
	if m != nil {
		return m.BaseExchangeRate
	}
	return 0
}

func (m *AccountMoneySummary) GetWarningLevel() float64 {
	if m != nil {
		return m.WarningLevel
	}
	return 0
}

func (m *AccountMoneySummary) GetForceCloseLevel() float64 {
	if m != nil {
		return m.ForceCloseLevel
	}
	return 0
}

func (m *AccountMoneySummary) GetClearedClosedProfit() float64 {
	if m != nil {
		return m.ClearedClosedProfit
	}
	return 0
}

func (m *AccountMoneySummary) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountMoneySummary) GetInterest() float64 {
	if m != nil {
		return m.Interest
	}
	return 0
}

func (m *AccountMoneySummary) GetFrozenInterest() float64 {
	if m != nil {
		return m.FrozenInterest
	}
	return 0
}

type AccountMoneySummaryList struct {
	List                 []*AccountMoneySummary `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AccountMoneySummaryList) Reset()         { *m = AccountMoneySummaryList{} }
func (m *AccountMoneySummaryList) String() string { return proto.CompactTextString(m) }
func (*AccountMoneySummaryList) ProtoMessage()    {}
func (*AccountMoneySummaryList) Descriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{5}
}

func (m *AccountMoneySummaryList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountMoneySummaryList.Unmarshal(m, b)
}
func (m *AccountMoneySummaryList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountMoneySummaryList.Marshal(b, m, deterministic)
}
func (m *AccountMoneySummaryList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountMoneySummaryList.Merge(m, src)
}
func (m *AccountMoneySummaryList) XXX_Size() int {
	return xxx_messageInfo_AccountMoneySummaryList.Size(m)
}
func (m *AccountMoneySummaryList) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountMoneySummaryList.DiscardUnknown(m)
}

var xxx_messageInfo_AccountMoneySummaryList proto.InternalMessageInfo

func (m *AccountMoneySummaryList) GetList() []*AccountMoneySummary {
	if m != nil {
		return m.List
	}
	return nil
}

// 出入金操作
type MoneyTransferRecord struct {
	// ID
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
	// 出入
	Direction int32 `protobuf:"varint,2,opt,name=direction,proto3" json:"direction"`
	// 币种
	CurrencyType int32 `protobuf:"varint,3,opt,name=currency_type,json=currencyType,proto3" json:"currencyType"`
	// 结算组
	SettlementGroup int32 `protobuf:"varint,4,opt,name=settlement_group,json=settlementGroup,proto3" json:"settlementGroup"`
	// 是否确认
	Confirmed int32 `protobuf:"varint,5,opt,name=confirmed,proto3" json:"confirmed"`
	// 优先
	IsPrefered int32 `protobuf:"varint,6,opt,name=is_prefered,json=isPrefered,proto3" json:"isPrefered"`
	// 金额
	Amount float64 `protobuf:"fixed64,7,opt,name=amount,proto3" json:"amount"`
	// 账户
	AccountId string `protobuf:"bytes,8,opt,name=account_id,json=accountId,proto3" json:"accountId"`
	// 操作员
	OperatorId string `protobuf:"bytes,9,opt,name=operator_id,json=operatorId,proto3" json:"operatorId"`
	// 会话编号
	SessionId int32 `protobuf:"varint,10,opt,name=session_id,json=sessionId,proto3" json:"sessionId"`
	// 请求时间
	Time int64 `protobuf:"varint,11,opt,name=time,proto3" json:"time"`
	// 备注
	Comment string `protobuf:"bytes,12,opt,name=comment,proto3" json:"comment"`
	// 第三方单号
	BankTradeId string `protobuf:"bytes,13,opt,name=bank_trade_id,json=bankTradeId,proto3" json:"bankTradeId"`
	// 配置编号，用于出金
	PaymentConfigId int64 `protobuf:"varint,14,opt,name=payment_config_id,json=paymentConfigId,proto3" json:"paymentConfigId"`
	// 原始金额
	OriginalAmount float64 `protobuf:"fixed64,15,opt,name=original_amount,json=originalAmount,proto3" json:"originalAmount"`
	// 费用
	Fee float64 `protobuf:"fixed64,16,opt,name=fee,proto3" json:"fee"`
	// 通道
	PaymentPath string `protobuf:"bytes,17,opt,name=payment_path,json=paymentPath,proto3" json:"paymentPath"`
	// 身份证号
	IdentityNumber string `protobuf:"bytes,18,opt,name=identity_number,json=identityNumber,proto3" json:"identityNumber"`
	// 电话
	PhoneNumber string `protobuf:"bytes,19,opt,name=phone_number,json=phoneNumber,proto3" json:"phoneNumber"`
	// 银行名称
	BankName string `protobuf:"bytes,20,opt,name=bank_name,json=bankName,proto3" json:"bankName"`
	// 支行名称
	BankBranchName string `protobuf:"bytes,21,opt,name=bank_branch_name,json=bankBranchName,proto3" json:"bankBranchName"`
	// 银行卡号
	BankCardNo string `protobuf:"bytes,22,opt,name=bank_card_no,json=bankCardNo,proto3" json:"bankCardNo"`
	// 省份
	Province string `protobuf:"bytes,23,opt,name=province,proto3" json:"province"`
	// 城市
	City string `protobuf:"bytes,24,opt,name=city,proto3" json:"city"`
	// 支付方式代码
	PayType string `protobuf:"bytes,25,opt,name=pay_type,json=payType,proto3" json:"payType"`
	// 支付时间
	PayedTime int64 `protobuf:"varint,26,opt,name=payed_time,json=payedTime,proto3" json:"payedTime"`
	// 确认时间
	ConfirmedTime int64 `protobuf:"varint,27,opt,name=confirmed_time,json=confirmedTime,proto3" json:"confirmedTime"`
	// 名称
	Name                 string   `protobuf:"bytes,28,opt,name=name,proto3" json:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoneyTransferRecord) Reset()         { *m = MoneyTransferRecord{} }
func (m *MoneyTransferRecord) String() string { return proto.CompactTextString(m) }
func (*MoneyTransferRecord) ProtoMessage()    {}
func (*MoneyTransferRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{6}
}

func (m *MoneyTransferRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoneyTransferRecord.Unmarshal(m, b)
}
func (m *MoneyTransferRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoneyTransferRecord.Marshal(b, m, deterministic)
}
func (m *MoneyTransferRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoneyTransferRecord.Merge(m, src)
}
func (m *MoneyTransferRecord) XXX_Size() int {
	return xxx_messageInfo_MoneyTransferRecord.Size(m)
}
func (m *MoneyTransferRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_MoneyTransferRecord.DiscardUnknown(m)
}

var xxx_messageInfo_MoneyTransferRecord proto.InternalMessageInfo

func (m *MoneyTransferRecord) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *MoneyTransferRecord) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *MoneyTransferRecord) GetCurrencyType() int32 {
	if m != nil {
		return m.CurrencyType
	}
	return 0
}

func (m *MoneyTransferRecord) GetSettlementGroup() int32 {
	if m != nil {
		return m.SettlementGroup
	}
	return 0
}

func (m *MoneyTransferRecord) GetConfirmed() int32 {
	if m != nil {
		return m.Confirmed
	}
	return 0
}

func (m *MoneyTransferRecord) GetIsPrefered() int32 {
	if m != nil {
		return m.IsPrefered
	}
	return 0
}

func (m *MoneyTransferRecord) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *MoneyTransferRecord) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *MoneyTransferRecord) GetOperatorId() string {
	if m != nil {
		return m.OperatorId
	}
	return ""
}

func (m *MoneyTransferRecord) GetSessionId() int32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *MoneyTransferRecord) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *MoneyTransferRecord) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *MoneyTransferRecord) GetBankTradeId() string {
	if m != nil {
		return m.BankTradeId
	}
	return ""
}

func (m *MoneyTransferRecord) GetPaymentConfigId() int64 {
	if m != nil {
		return m.PaymentConfigId
	}
	return 0
}

func (m *MoneyTransferRecord) GetOriginalAmount() float64 {
	if m != nil {
		return m.OriginalAmount
	}
	return 0
}

func (m *MoneyTransferRecord) GetFee() float64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *MoneyTransferRecord) GetPaymentPath() string {
	if m != nil {
		return m.PaymentPath
	}
	return ""
}

func (m *MoneyTransferRecord) GetIdentityNumber() string {
	if m != nil {
		return m.IdentityNumber
	}
	return ""
}

func (m *MoneyTransferRecord) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *MoneyTransferRecord) GetBankName() string {
	if m != nil {
		return m.BankName
	}
	return ""
}

func (m *MoneyTransferRecord) GetBankBranchName() string {
	if m != nil {
		return m.BankBranchName
	}
	return ""
}

func (m *MoneyTransferRecord) GetBankCardNo() string {
	if m != nil {
		return m.BankCardNo
	}
	return ""
}

func (m *MoneyTransferRecord) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *MoneyTransferRecord) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *MoneyTransferRecord) GetPayType() string {
	if m != nil {
		return m.PayType
	}
	return ""
}

func (m *MoneyTransferRecord) GetPayedTime() int64 {
	if m != nil {
		return m.PayedTime
	}
	return 0
}

func (m *MoneyTransferRecord) GetConfirmedTime() int64 {
	if m != nil {
		return m.ConfirmedTime
	}
	return 0
}

func (m *MoneyTransferRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MTRList struct {
	List                 []*MoneyTransferRecord `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MTRList) Reset()         { *m = MTRList{} }
func (m *MTRList) String() string { return proto.CompactTextString(m) }
func (*MTRList) ProtoMessage()    {}
func (*MTRList) Descriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{7}
}

func (m *MTRList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MTRList.Unmarshal(m, b)
}
func (m *MTRList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MTRList.Marshal(b, m, deterministic)
}
func (m *MTRList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MTRList.Merge(m, src)
}
func (m *MTRList) XXX_Size() int {
	return xxx_messageInfo_MTRList.Size(m)
}
func (m *MTRList) XXX_DiscardUnknown() {
	xxx_messageInfo_MTRList.DiscardUnknown(m)
}

var xxx_messageInfo_MTRList proto.InternalMessageInfo

func (m *MTRList) GetList() []*MoneyTransferRecord {
	if m != nil {
		return m.List
	}
	return nil
}

// 期权月份
type OptionMonth struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
	Short                string   `protobuf:"bytes,3,opt,name=short,proto3" json:"short"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OptionMonth) Reset()         { *m = OptionMonth{} }
func (m *OptionMonth) String() string { return proto.CompactTextString(m) }
func (*OptionMonth) ProtoMessage()    {}
func (*OptionMonth) Descriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{8}
}

func (m *OptionMonth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptionMonth.Unmarshal(m, b)
}
func (m *OptionMonth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptionMonth.Marshal(b, m, deterministic)
}
func (m *OptionMonth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptionMonth.Merge(m, src)
}
func (m *OptionMonth) XXX_Size() int {
	return xxx_messageInfo_OptionMonth.Size(m)
}
func (m *OptionMonth) XXX_DiscardUnknown() {
	xxx_messageInfo_OptionMonth.DiscardUnknown(m)
}

var xxx_messageInfo_OptionMonth proto.InternalMessageInfo

func (m *OptionMonth) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OptionMonth) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *OptionMonth) GetShort() string {
	if m != nil {
		return m.Short
	}
	return ""
}

// 期权月份列表
type OptionMonthList struct {
	List                 []*OptionMonth `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *OptionMonthList) Reset()         { *m = OptionMonthList{} }
func (m *OptionMonthList) String() string { return proto.CompactTextString(m) }
func (*OptionMonthList) ProtoMessage()    {}
func (*OptionMonthList) Descriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{9}
}

func (m *OptionMonthList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptionMonthList.Unmarshal(m, b)
}
func (m *OptionMonthList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptionMonthList.Marshal(b, m, deterministic)
}
func (m *OptionMonthList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptionMonthList.Merge(m, src)
}
func (m *OptionMonthList) XXX_Size() int {
	return xxx_messageInfo_OptionMonthList.Size(m)
}
func (m *OptionMonthList) XXX_DiscardUnknown() {
	xxx_messageInfo_OptionMonthList.DiscardUnknown(m)
}

var xxx_messageInfo_OptionMonthList proto.InternalMessageInfo

func (m *OptionMonthList) GetList() []*OptionMonth {
	if m != nil {
		return m.List
	}
	return nil
}

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{10}
}

func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (m *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(m, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{11}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type CommonRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonRequest) Reset()         { *m = CommonRequest{} }
func (m *CommonRequest) String() string { return proto.CompactTextString(m) }
func (*CommonRequest) ProtoMessage()    {}
func (*CommonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{12}
}

func (m *CommonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonRequest.Unmarshal(m, b)
}
func (m *CommonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonRequest.Marshal(b, m, deterministic)
}
func (m *CommonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonRequest.Merge(m, src)
}
func (m *CommonRequest) XXX_Size() int {
	return xxx_messageInfo_CommonRequest.Size(m)
}
func (m *CommonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommonRequest proto.InternalMessageInfo

type CommonResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonResponse) Reset()         { *m = CommonResponse{} }
func (m *CommonResponse) String() string { return proto.CompactTextString(m) }
func (*CommonResponse) ProtoMessage()    {}
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_239886423b810ded, []int{13}
}

func (m *CommonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonResponse.Unmarshal(m, b)
}
func (m *CommonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonResponse.Marshal(b, m, deterministic)
}
func (m *CommonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonResponse.Merge(m, src)
}
func (m *CommonResponse) XXX_Size() int {
	return xxx_messageInfo_CommonResponse.Size(m)
}
func (m *CommonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonResponse proto.InternalMessageInfo

func (m *CommonResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CommonResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CommonResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("goshare.PeriodType", PeriodType_name, PeriodType_value)
	proto.RegisterEnum("goshare.OptionCallPutType", OptionCallPutType_name, OptionCallPutType_value)
	proto.RegisterEnum("goshare.OptionDeliveryDateType", OptionDeliveryDateType_name, OptionDeliveryDateType_value)
	proto.RegisterEnum("goshare.CloseCommissionAlgorithim", CloseCommissionAlgorithim_name, CloseCommissionAlgorithim_value)
	proto.RegisterEnum("goshare.AccountType", AccountType_name, AccountType_value)
	proto.RegisterEnum("goshare.CurrencyType", CurrencyType_name, CurrencyType_value)
	proto.RegisterEnum("goshare.TradingAccountType", TradingAccountType_name, TradingAccountType_value)
	proto.RegisterType((*BrokerRoute)(nil), "goshare.BrokerRoute")
	proto.RegisterType((*BrokerRouteList)(nil), "goshare.BrokerRouteList")
	proto.RegisterType((*ReqUpdateTIOpenDate)(nil), "goshare.ReqUpdateTIOpenDate")
	proto.RegisterType((*ReqUpdateTIOpenDateList)(nil), "goshare.ReqUpdateTIOpenDateList")
	proto.RegisterType((*AccountMoneySummary)(nil), "goshare.AccountMoneySummary")
	proto.RegisterType((*AccountMoneySummaryList)(nil), "goshare.AccountMoneySummaryList")
	proto.RegisterType((*MoneyTransferRecord)(nil), "goshare.MoneyTransferRecord")
	proto.RegisterType((*MTRList)(nil), "goshare.MTRList")
	proto.RegisterType((*OptionMonth)(nil), "goshare.OptionMonth")
	proto.RegisterType((*OptionMonthList)(nil), "goshare.OptionMonthList")
	proto.RegisterType((*EmptyRequest)(nil), "goshare.EmptyRequest")
	proto.RegisterType((*EmptyResponse)(nil), "goshare.EmptyResponse")
	proto.RegisterType((*CommonRequest)(nil), "goshare.CommonRequest")
	proto.RegisterType((*CommonResponse)(nil), "goshare.CommonResponse")
}

func init() { proto.RegisterFile("goshare/common.proto", fileDescriptor_239886423b810ded) }

var fileDescriptor_239886423b810ded = []byte{
	// 1994 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x58, 0x5f, 0x53, 0xdb, 0x48,
	0x12, 0x8f, 0x01, 0x03, 0x6a, 0x83, 0x11, 0x82, 0x04, 0xe5, 0xcf, 0x6e, 0x08, 0xd9, 0xdd, 0x10,
	0x36, 0xcb, 0x9f, 0x64, 0xf7, 0xea, 0x6a, 0xf7, 0x09, 0x0c, 0xd9, 0x73, 0x05, 0x1b, 0x97, 0x30,
	0x77, 0xb7, 0xf7, 0xa2, 0x1a, 0x4b, 0x83, 0xad, 0x8a, 0xa4, 0xd1, 0x8e, 0xc6, 0x24, 0xbe, 0xa7,
	0xbb, 0xcf, 0x70, 0x1f, 0xf5, 0xbe, 0xc0, 0x55, 0xf7, 0xcc, 0x08, 0x85, 0xe4, 0xf6, 0xc9, 0xea,
	0x5f, 0xff, 0xa6, 0x67, 0xba, 0xa7, 0xa7, 0xa7, 0xc7, 0xb0, 0x39, 0x16, 0xe5, 0x84, 0x49, 0x7e,
	0x10, 0x89, 0x2c, 0x13, 0xf9, 0x7e, 0x21, 0x85, 0x12, 0xde, 0x92, 0x41, 0x77, 0xfe, 0xd3, 0x80,
	0xd6, 0x89, 0x14, 0xef, 0xb9, 0x0c, 0xc4, 0x54, 0x71, 0xaf, 0x0d, 0x73, 0x49, 0xec, 0x37, 0xb6,
	0x1b, 0xbb, 0x4e, 0x30, 0x97, 0xc4, 0x9e, 0x07, 0x0b, 0x6a, 0x56, 0x70, 0x7f, 0x6e, 0xbb, 0xb1,
	0xdb, 0x0c, 0xe8, 0x1b, 0xb1, 0x9c, 0x65, 0xdc, 0x9f, 0x27, 0x16, 0x7d, 0x7b, 0x3b, 0xb0, 0x9a,
	0xc5, 0xe1, 0xb5, 0x14, 0xb9, 0x0a, 0xd3, 0xa4, 0x54, 0xfe, 0xc2, 0xf6, 0xfc, 0xae, 0x13, 0xb4,
	0xb2, 0xf8, 0x2d, 0x62, 0xe7, 0x49, 0xa9, 0xbc, 0x5d, 0x70, 0x95, 0x64, 0x31, 0xaf, 0xd3, 0x9a,
	0x44, 0x6b, 0x13, 0x5e, 0x31, 0x77, 0x7e, 0x81, 0xb5, 0xda, 0xa2, 0xcc, 0xe0, 0x05, 0x1a, 0xd0,
	0xd8, 0x9e, 0xdf, 0x6d, 0xbd, 0xde, 0xdc, 0x37, 0x0e, 0xec, 0xd7, 0x78, 0x01, 0x31, 0x76, 0xae,
	0x61, 0x23, 0xe0, 0xbf, 0x5f, 0x15, 0x31, 0x53, 0x7c, 0xd8, 0xbd, 0x28, 0x78, 0x7e, 0xca, 0x14,
	0xf7, 0x1e, 0xc1, 0x32, 0xff, 0x18, 0x4d, 0x58, 0x3e, 0xe6, 0xc6, 0xbf, 0x4a, 0xf6, 0x1e, 0xc0,
	0x62, 0x39, 0xcb, 0x46, 0x22, 0x25, 0x3f, 0x9d, 0xc0, 0x48, 0xde, 0x63, 0x70, 0x44, 0xc1, 0xf3,
	0x10, 0x4d, 0x91, 0xbb, 0xcd, 0x60, 0x59, 0x18, 0x83, 0x3b, 0xff, 0x6a, 0xc0, 0xd6, 0x17, 0x26,
	0xa2, 0xd5, 0xde, 0x9d, 0xac, 0x59, 0x9b, 0xec, 0xd0, 0x78, 0x32, 0x47, 0x9e, 0x3c, 0xa9, 0x3c,
	0xf9, 0x82, 0x2d, 0xed, 0x91, 0xb7, 0x05, 0x4b, 0xa2, 0x08, 0x23, 0x11, 0xdb, 0x98, 0x2f, 0x8a,
	0xa2, 0x23, 0x62, 0xbe, 0xf3, 0xdf, 0x35, 0xd8, 0x38, 0x8e, 0x22, 0x31, 0xcd, 0x55, 0x4f, 0xe4,
	0x7c, 0x76, 0x39, 0xcd, 0x32, 0x26, 0x67, 0x9f, 0xed, 0xe2, 0x6e, 0x6d, 0x17, 0xdb, 0xb5, 0xe0,
	0x99, 0xb1, 0xc3, 0x59, 0xc1, 0xcd, 0xde, 0xfe, 0x0c, 0xab, 0xd1, 0x54, 0x4a, 0x9e, 0x47, 0xb3,
	0x90, 0x86, 0xcc, 0xd3, 0x90, 0xfb, 0xd5, 0x90, 0x8e, 0xd1, 0xd2, 0x98, 0x95, 0xa8, 0x26, 0x79,
	0x3e, 0x2c, 0x8d, 0x58, 0xca, 0xf2, 0x88, 0xfb, 0x0b, 0xdb, 0x8d, 0xdd, 0x46, 0x60, 0x45, 0xef,
	0x6b, 0x00, 0x4c, 0xbf, 0xa4, 0x2c, 0x13, 0x91, 0xfb, 0x4d, 0x52, 0xd6, 0x10, 0xef, 0x2b, 0x80,
	0x68, 0x2a, 0xc3, 0x8c, 0xc9, 0x71, 0x92, 0xfb, 0x8b, 0xa4, 0x77, 0xa2, 0xa9, 0xec, 0x11, 0xe0,
	0xed, 0xc3, 0x46, 0xce, 0x55, 0x58, 0x88, 0x32, 0x51, 0x89, 0xc8, 0x43, 0x96, 0xe1, 0xaa, 0xfd,
	0x25, 0xe2, 0xad, 0xe7, 0x5c, 0x0d, 0x8c, 0xe6, 0x98, 0x14, 0xde, 0x53, 0x68, 0x15, 0x92, 0x87,
	0x76, 0x31, 0xcb, 0x7a, 0xbe, 0x42, 0xf2, 0x13, 0xb3, 0x9e, 0xa7, 0xd0, 0x4a, 0x72, 0xc5, 0x25,
	0x2f, 0x55, 0x98, 0xe4, 0xbe, 0xa3, 0x09, 0x16, 0xea, 0xe6, 0xe8, 0x4a, 0xcc, 0x69, 0x3e, 0x1f,
	0xb4, 0x2b, 0x46, 0xc4, 0x9d, 0xfd, 0x90, 0xa8, 0x49, 0x2c, 0xd9, 0x07, 0xbf, 0x45, 0xaa, 0x4a,
	0xf6, 0x9e, 0xc3, 0xea, 0xb5, 0x14, 0xff, 0xe4, 0xb9, 0xf5, 0x64, 0x85, 0x08, 0x2b, 0x1a, 0x34,
	0xce, 0x7c, 0x0f, 0xeb, 0x86, 0x54, 0x0b, 0xc9, 0x2a, 0x11, 0x5d, 0xad, 0xe8, 0xdc, 0x06, 0xe6,
	0x09, 0x38, 0xec, 0x86, 0x25, 0x29, 0x1b, 0xa5, 0xdc, 0x6f, 0xeb, 0xb8, 0x54, 0x80, 0xf7, 0x0c,
	0x56, 0xa2, 0x54, 0x94, 0x3c, 0x2c, 0xa4, 0xb8, 0x4e, 0x94, 0xbf, 0x46, 0x84, 0x16, 0x61, 0x03,
	0x82, 0xbc, 0x17, 0xb0, 0x56, 0x85, 0xcd, 0xb0, 0x5c, 0x62, 0xb5, 0x2d, 0x6c, 0x88, 0xcf, 0x60,
	0x05, 0x63, 0x96, 0x09, 0xa9, 0xc6, 0x98, 0xed, 0xeb, 0xda, 0x56, 0x21, 0x79, 0xcf, 0x40, 0xb8,
	0x4b, 0x48, 0x89, 0x24, 0x8f, 0x13, 0xe5, 0x7b, 0x7a, 0x35, 0x85, 0xe4, 0x1d, 0x02, 0xd0, 0xfb,
	0x2a, 0xa8, 0x23, 0x56, 0x72, 0x7f, 0x43, 0x7b, 0x6f, 0xc1, 0x13, 0x56, 0x72, 0xef, 0x07, 0xf0,
	0x6c, 0xb8, 0xc2, 0x5b, 0xcf, 0x36, 0xf5, 0x4e, 0x5a, 0xcd, 0x71, 0xe5, 0xa1, 0x0f, 0x4b, 0x92,
	0x97, 0x5c, 0xde, 0x70, 0xff, 0xbe, 0xde, 0x07, 0x23, 0xe2, 0x16, 0x62, 0xd1, 0x48, 0xf2, 0x71,
	0x18, 0xb3, 0x99, 0xff, 0x80, 0x0e, 0x19, 0x18, 0xe8, 0x94, 0xcd, 0xf0, 0x4c, 0x9b, 0x95, 0x6e,
	0xd1, 0x48, 0x23, 0xe1, 0x06, 0x56, 0x4e, 0xfa, 0x7a, 0x03, 0xad, 0x8c, 0xd1, 0xb2, 0xc7, 0xd4,
	0x6e, 0xe1, 0x43, 0x1d, 0x2d, 0x0b, 0x9b, 0x4d, 0x7c, 0x01, 0x6b, 0x31, 0x4f, 0x93, 0x1b, 0x2e,
	0x67, 0x96, 0xf8, 0x48, 0x13, 0x2d, 0x6c, 0x88, 0x7f, 0x06, 0xbf, 0xb2, 0x78, 0x77, 0xc4, 0x63,
	0x1a, 0xf1, 0xc0, 0xea, 0x4f, 0x3f, 0x1d, 0xf9, 0x02, 0xd6, 0x8c, 0xaf, 0x55, 0x22, 0x3f, 0xd1,
	0x53, 0x18, 0xd8, 0x26, 0xf3, 0x01, 0x6c, 0x54, 0x96, 0x6b, 0x29, 0xf5, 0x15, 0x91, 0x3d, 0xab,
	0xaa, 0x25, 0xd5, 0x3e, 0x6c, 0x98, 0x0c, 0x54, 0x92, 0xe5, 0xe5, 0x35, 0x97, 0xe1, 0x35, 0xe7,
	0xfe, 0xd7, 0x7a, 0x13, 0xb4, 0x6a, 0x68, 0x34, 0x6f, 0x39, 0xc7, 0xba, 0x6d, 0xf8, 0xa5, 0x62,
	0x59, 0x11, 0x2a, 0xf6, 0xd1, 0x7f, 0xaa, 0x97, 0xa2, 0xf1, 0x4b, 0x84, 0x87, 0xec, 0x23, 0x26,
	0xd1, 0x27, 0x26, 0xb7, 0x75, 0x12, 0xa9, 0x9a, 0xb1, 0xc7, 0xe0, 0xdc, 0x5a, 0x79, 0xa6, 0xe3,
	0x5f, 0xda, 0xf1, 0x5f, 0x01, 0x64, 0x2a, 0xb3, 0x89, 0xba, 0xa3, 0x33, 0x2c, 0x53, 0x99, 0xc9,
	0xd1, 0x6f, 0xa0, 0x4d, 0x39, 0x7a, 0x4b, 0x79, 0xae, 0x53, 0x0c, 0xb3, 0xb4, 0x62, 0x3d, 0x85,
	0x56, 0xa9, 0x44, 0xf4, 0x3e, 0xbc, 0x61, 0xe9, 0x94, 0xfb, 0xdf, 0xe8, 0xc3, 0x4d, 0xd0, 0x5f,
	0x11, 0xf1, 0x7e, 0x84, 0x07, 0x23, 0x91, 0xc7, 0xa1, 0xe4, 0xc5, 0x54, 0x46, 0x13, 0x56, 0x72,
	0x5b, 0x51, 0xbe, 0x25, 0xee, 0x26, 0x6a, 0x83, 0x4a, 0x69, 0x8a, 0xca, 0xcf, 0xf0, 0x50, 0xf2,
	0x1b, 0x2e, 0x4b, 0xfe, 0x85, 0x81, 0xdf, 0xd1, 0xc0, 0x2d, 0x43, 0xf8, 0x6c, 0xec, 0x3e, 0x6c,
	0xe8, 0x3d, 0x0f, 0xf5, 0x05, 0x68, 0x46, 0xbd, 0xd0, 0x11, 0xd7, 0xaa, 0x21, 0x6a, 0x0c, 0x7f,
	0x0f, 0xd6, 0xcb, 0x89, 0x90, 0x2a, 0x2c, 0x79, 0x9a, 0x5a, 0xf6, 0x2e, 0xb1, 0xd7, 0x48, 0x71,
	0xc9, 0xd3, 0xf4, 0xff, 0xd8, 0x36, 0x91, 0x79, 0xf9, 0x99, 0x6d, 0x13, 0x9e, 0x4f, 0x6d, 0x1b,
	0xf6, 0xde, 0x1d, 0xdb, 0x86, 0xfb, 0x27, 0xd8, 0xb2, 0xb5, 0x8a, 0xea, 0x4c, 0x2d, 0xbd, 0xbe,
	0xa7, 0x11, 0xf7, 0x4d, 0xc5, 0x42, 0x6d, 0x2d, 0xc3, 0xbe, 0x85, 0xb6, 0x6e, 0x37, 0xaa, 0xd4,
	0x7d, 0x45, 0xf4, 0x55, 0x8d, 0xda, 0xcc, 0x7d, 0x09, 0x6e, 0x21, 0xf9, 0x35, 0x97, 0x3c, 0xae,
	0x88, 0x3f, 0xe8, 0x95, 0x58, 0xdc, 0x52, 0x5f, 0x81, 0x87, 0x35, 0x25, 0xac, 0x0e, 0x93, 0xc4,
	0xf3, 0xbb, 0xaf, 0xcb, 0x26, 0x6a, 0xce, 0x8c, 0x22, 0xc0, 0x73, 0xfc, 0x1c, 0x56, 0x3f, 0x30,
	0x99, 0x63, 0x71, 0x48, 0xf9, 0x0d, 0x4f, 0xfd, 0x03, 0x9d, 0x27, 0x06, 0x3c, 0x47, 0x0c, 0x03,
	0x71, 0x2d, 0x64, 0xc4, 0x8d, 0x6f, 0x9a, 0x78, 0xa8, 0xa7, 0x27, 0x05, 0x79, 0xa5, 0xb9, 0xaf,
	0xe1, 0x7e, 0x94, 0x72, 0x86, 0x0b, 0x25, 0x76, 0x6c, 0x03, 0x77, 0x44, 0xfc, 0x0d, 0xa3, 0xa4,
	0x11, 0xb1, 0x09, 0x9e, 0x6d, 0x93, 0x5e, 0xd7, 0xda, 0xa4, 0x47, 0xb0, 0x6c, 0xcb, 0xa1, 0xff,
	0x46, 0x27, 0xbf, 0x95, 0xf1, 0xc0, 0x9b, 0x60, 0x57, 0x94, 0x1f, 0xeb, 0xa7, 0xac, 0x6b, 0xd0,
	0x9d, 0x77, 0xb0, 0xf5, 0x85, 0x4b, 0x9f, 0xfa, 0x8e, 0xc3, 0x4f, 0xba, 0xa4, 0x27, 0x77, 0x2f,
	0xfa, 0x3a, 0xdf, 0x74, 0x4b, 0xff, 0x5e, 0x82, 0x0d, 0x82, 0xed, 0x89, 0x0f, 0x78, 0x24, 0x64,
	0xec, 0xb9, 0x30, 0x3f, 0x35, 0x3d, 0xc4, 0x7c, 0x80, 0x9f, 0x78, 0x17, 0xc5, 0x89, 0xe4, 0x11,
	0x5e, 0x1a, 0xa6, 0x1f, 0xbc, 0x05, 0x30, 0xe4, 0x9f, 0x37, 0x0e, 0xcd, 0x3b, 0x1d, 0xc2, 0x4b,
	0x70, 0x4b, 0xae, 0x54, 0xca, 0x33, 0x9e, 0xab, 0x70, 0x2c, 0xc5, 0xb4, 0xa0, 0x56, 0xa1, 0x19,
	0xac, 0xdd, 0xe2, 0xbf, 0x22, 0x8c, 0xb3, 0x45, 0x22, 0xbf, 0x4e, 0x64, 0xc6, 0x63, 0xea, 0x18,
	0x9a, 0xc1, 0x2d, 0x40, 0x17, 0x78, 0x19, 0xda, 0x24, 0xa1, 0x8e, 0xa1, 0x19, 0x40, 0x52, 0x0e,
	0x0c, 0x82, 0xd5, 0xff, 0x93, 0x2e, 0xc1, 0x48, 0x58, 0x61, 0x98, 0x8e, 0x45, 0x98, 0xc4, 0xd4,
	0x19, 0x38, 0x81, 0x63, 0x90, 0x2e, 0xd9, 0x15, 0x05, 0x97, 0x4c, 0x09, 0x89, 0x7a, 0x87, 0xf4,
	0x60, 0xa1, 0x6e, 0x8c, 0xe3, 0x4b, 0x4e, 0x49, 0x8e, 0x7a, 0xd0, 0xeb, 0x32, 0x48, 0x57, 0xb7,
	0xcb, 0x49, 0xc6, 0xa9, 0x33, 0x98, 0x0f, 0xe8, 0x1b, 0xef, 0x30, 0x4c, 0x7b, 0x9e, 0x2b, 0xea,
	0x07, 0x9c, 0xc0, 0x8a, 0xd8, 0x34, 0x8f, 0x58, 0xfe, 0xde, 0x1c, 0xdc, 0x24, 0xa6, 0x36, 0xc0,
	0x09, 0x5a, 0x08, 0xd2, 0x91, 0xed, 0xc6, 0x98, 0xa5, 0x05, 0x9b, 0x51, 0xbc, 0xc8, 0xfd, 0x31,
	0xf2, 0xda, 0x64, 0x7e, 0xcd, 0x28, 0x3a, 0x84, 0x77, 0x63, 0xcc, 0x20, 0x21, 0x93, 0x71, 0x92,
	0xb3, 0xaa, 0x68, 0xe8, 0x96, 0xa0, 0x6d, 0x61, 0x53, 0x33, 0x5c, 0x98, 0xc7, 0xf2, 0xac, 0x3b,
	0x01, 0xfc, 0xa4, 0xeb, 0xdf, 0x4c, 0x53, 0x30, 0x35, 0xa1, 0xeb, 0xdf, 0x09, 0x5a, 0x06, 0x1b,
	0x30, 0x35, 0x41, 0xeb, 0x49, 0xcc, 0x73, 0x95, 0xa8, 0x59, 0x98, 0x4f, 0xb3, 0x11, 0x97, 0xd4,
	0x03, 0x38, 0x41, 0xdb, 0xc2, 0x7d, 0x42, 0xc9, 0xd6, 0x44, 0xe4, 0xdc, 0xb2, 0x36, 0x8c, 0x2d,
	0xc4, 0x0c, 0xe5, 0x31, 0x38, 0xe4, 0x39, 0x1d, 0x90, 0x4d, 0xdd, 0x8d, 0x23, 0xd0, 0xc7, 0x43,
	0xb2, 0x0b, 0x2e, 0x29, 0x47, 0x92, 0xe5, 0xd1, 0x44, 0x73, 0xee, 0xeb, 0x99, 0x10, 0x3f, 0x21,
	0x98, 0x98, 0xdb, 0xb0, 0x42, 0xcc, 0x88, 0xc9, 0x38, 0xcc, 0x05, 0x75, 0x01, 0x4e, 0x00, 0x88,
	0x75, 0x98, 0x8c, 0xfb, 0x02, 0x0f, 0x5c, 0x21, 0xc5, 0x4d, 0x82, 0xa5, 0x65, 0x4b, 0xcf, 0x63,
	0x65, 0xdc, 0xac, 0x28, 0x51, 0x33, 0xea, 0x02, 0x9c, 0x80, 0xbe, 0xbd, 0x87, 0xb0, 0x5c, 0x30,
	0x93, 0xc1, 0x0f, 0xf5, 0x6e, 0x15, 0x4c, 0x27, 0x2f, 0xb6, 0x3f, 0x6c, 0xc6, 0xe3, 0x90, 0x76,
	0xf8, 0x11, 0x6d, 0x81, 0x43, 0xc8, 0x10, 0xb7, 0x99, 0x6a, 0x9e, 0xc9, 0x4f, 0x4d, 0x79, 0x4c,
	0x94, 0xd5, 0x0a, 0x25, 0x9a, 0xad, 0x0a, 0x4f, 0x6e, 0xab, 0xc2, 0xce, 0x2f, 0xb0, 0xd4, 0x1b,
	0x06, 0x7f, 0x78, 0x80, 0xbf, 0x70, 0x44, 0xcd, 0x01, 0xee, 0x41, 0xeb, 0xa2, 0xc0, 0x23, 0xd8,
	0x13, 0xb9, 0x9a, 0x54, 0xf6, 0x1b, 0xb5, 0xaa, 0xb3, 0x09, 0x4d, 0x7d, 0x17, 0xea, 0xd7, 0x8d,
	0x16, 0x10, 0xa5, 0x7a, 0x6f, 0xde, 0x14, 0x5a, 0xc0, 0xa7, 0x57, 0xcd, 0xdc, 0x1f, 0x3e, 0xbd,
	0x6a, 0x3c, 0xb3, 0x96, 0x36, 0xac, 0x9c, 0x65, 0x85, 0x9a, 0x05, 0xfc, 0xf7, 0x29, 0x56, 0xaa,
	0x35, 0x58, 0x35, 0x72, 0x59, 0x88, 0xbc, 0xe4, 0x08, 0x74, 0xe8, 0x0a, 0xb0, 0x8c, 0xbf, 0x43,
	0xdb, 0x02, 0x9a, 0x82, 0xc7, 0xa5, 0x9c, 0x46, 0x11, 0x2f, 0x4b, 0xf2, 0x61, 0x39, 0xb0, 0x22,
	0xed, 0x17, 0xbe, 0x81, 0xcc, 0x5b, 0x14, 0xbf, 0x91, 0x9d, 0xf1, 0xb2, 0x64, 0x63, 0xfb, 0x34,
	0xb2, 0xe2, 0x5e, 0x09, 0x30, 0xe0, 0x32, 0x11, 0x31, 0x6d, 0xde, 0x32, 0x2c, 0x0c, 0xbb, 0x9d,
	0x77, 0xee, 0x3d, 0x6f, 0x11, 0xe6, 0x7a, 0x47, 0x6e, 0x83, 0x7e, 0xdf, 0xb8, 0x73, 0xf4, 0xfb,
	0x93, 0x3b, 0xef, 0x2d, 0xc1, 0x7c, 0xef, 0xe8, 0xd0, 0x5d, 0xd0, 0x1f, 0x3f, 0xb9, 0x4d, 0xfa,
	0x78, 0x73, 0xe8, 0x2e, 0x22, 0xe5, 0x2f, 0x47, 0xee, 0x12, 0xfd, 0xbe, 0x71, 0x97, 0xf1, 0xf7,
	0xf4, 0xc8, 0x75, 0xf0, 0xf7, 0x6f, 0x47, 0x2e, 0xa0, 0xf1, 0xde, 0x45, 0xff, 0xc8, 0x6d, 0xed,
	0x1d, 0xc2, 0xba, 0x8e, 0x4a, 0x87, 0xa5, 0xe9, 0x60, 0x4a, 0x2f, 0x2b, 0x6f, 0x15, 0x9c, 0x8b,
	0xce, 0x60, 0x18, 0x76, 0x8e, 0xcf, 0xcf, 0xdd, 0x7b, 0xde, 0x0a, 0x2c, 0x93, 0x38, 0xb8, 0x1a,
	0xba, 0x8d, 0xbd, 0xef, 0xe0, 0x81, 0x1e, 0x61, 0xdb, 0x3f, 0x7c, 0xf7, 0xd1, 0x30, 0xe4, 0x9d,
	0x9e, 0x0e, 0xc3, 0xb3, 0xab, 0xc0, 0xbd, 0xb7, 0x77, 0x09, 0x0f, 0xef, 0xdc, 0xb2, 0xc7, 0xe9,
	0x58, 0xc8, 0x44, 0x4d, 0x92, 0xcc, 0x6b, 0x03, 0x74, 0x3a, 0xc7, 0x61, 0xff, 0x22, 0xe8, 0x1d,
	0x9b, 0x29, 0x50, 0xbe, 0xe8, 0x77, 0xce, 0xdc, 0x86, 0xe7, 0xc3, 0x26, 0x4a, 0xbd, 0xab, 0xf3,
	0x61, 0x77, 0x70, 0x7e, 0x16, 0x9e, 0xfc, 0x16, 0x9e, 0x1e, 0x0f, 0xcf, 0xdc, 0xb9, 0xbd, 0x3d,
	0x68, 0xd5, 0x9e, 0x80, 0xb8, 0xd0, 0x4e, 0x3f, 0x7c, 0x7b, 0x35, 0xbc, 0x0a, 0xce, 0x8c, 0x95,
	0x7e, 0x78, 0x39, 0xbc, 0xe8, 0xbc, 0x73, 0x1b, 0x7b, 0xdb, 0xb0, 0x52, 0x7f, 0xfb, 0x61, 0x74,
	0x3a, 0xfd, 0xdf, 0xdc, 0x7b, 0xf8, 0x71, 0x75, 0x79, 0xea, 0x36, 0xf6, 0xde, 0x81, 0x37, 0xd4,
	0xfd, 0x77, 0xdd, 0x68, 0x1b, 0x60, 0x78, 0x3c, 0xbc, 0x5d, 0xdb, 0x3a, 0xac, 0xa2, 0x3c, 0x0c,
	0x8e, 0xbb, 0xfd, 0x7e, 0xb7, 0xff, 0xab, 0xdb, 0xf0, 0x3c, 0x68, 0x23, 0xd4, 0x19, 0x0e, 0xc2,
	0x5e, 0x37, 0x08, 0x2e, 0x02, 0x77, 0xee, 0x64, 0xff, 0x1f, 0xaf, 0xc6, 0x89, 0x9a, 0x4c, 0x47,
	0xfb, 0x91, 0xc8, 0x0e, 0xb2, 0x24, 0xe7, 0x92, 0xa5, 0x92, 0x97, 0x07, 0xf6, 0xff, 0x8c, 0xe2,
	0xfd, 0xf8, 0xa0, 0x18, 0x59, 0x71, 0xb4, 0x48, 0x7f, 0x6c, 0xbc, 0xf9, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd5, 0x7d, 0x89, 0x6a, 0xf0, 0x10, 0x00, 0x00,
}

