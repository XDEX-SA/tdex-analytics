// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: v1/analytics.proto

package tdexav1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PredefinedPeriod int32

const (
	PredefinedPeriod_NULL          PredefinedPeriod = 0
	PredefinedPeriod_LAST_HOUR     PredefinedPeriod = 1
	PredefinedPeriod_LAST_DAY      PredefinedPeriod = 2
	PredefinedPeriod_LAST_MONTH    PredefinedPeriod = 3
	PredefinedPeriod_LAST_3_MONTHS PredefinedPeriod = 4
	PredefinedPeriod_YEAR_TO_DATE  PredefinedPeriod = 5
	PredefinedPeriod_ALL           PredefinedPeriod = 6
)

// Enum value maps for PredefinedPeriod.
var (
	PredefinedPeriod_name = map[int32]string{
		0: "NULL",
		1: "LAST_HOUR",
		2: "LAST_DAY",
		3: "LAST_MONTH",
		4: "LAST_3_MONTHS",
		5: "YEAR_TO_DATE",
		6: "ALL",
	}
	PredefinedPeriod_value = map[string]int32{
		"NULL":          0,
		"LAST_HOUR":     1,
		"LAST_DAY":      2,
		"LAST_MONTH":    3,
		"LAST_3_MONTHS": 4,
		"YEAR_TO_DATE":  5,
		"ALL":           6,
	}
)

func (x PredefinedPeriod) Enum() *PredefinedPeriod {
	p := new(PredefinedPeriod)
	*p = x
	return p
}

func (x PredefinedPeriod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PredefinedPeriod) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_analytics_proto_enumTypes[0].Descriptor()
}

func (PredefinedPeriod) Type() protoreflect.EnumType {
	return &file_v1_analytics_proto_enumTypes[0]
}

func (x PredefinedPeriod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PredefinedPeriod.Descriptor instead.
func (PredefinedPeriod) EnumDescriptor() ([]byte, []int) {
	return file_v1_analytics_proto_rawDescGZIP(), []int{0}
}

type MarketsBalancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// time_range fetch balances from time range
	TimeRange *TimeRange `protobuf:"bytes,1,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	// fetch balances for specific one or more market's, if no market_id is passed balances will be fetched for all
	MarketIds []string `protobuf:"bytes,2,rep,name=market_ids,json=marketIds,proto3" json:"market_ids,omitempty"`
}

func (x *MarketsBalancesRequest) Reset() {
	*x = MarketsBalancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketsBalancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketsBalancesRequest) ProtoMessage() {}

func (x *MarketsBalancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_analytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketsBalancesRequest.ProtoReflect.Descriptor instead.
func (*MarketsBalancesRequest) Descriptor() ([]byte, []int) {
	return file_v1_analytics_proto_rawDescGZIP(), []int{0}
}

func (x *MarketsBalancesRequest) GetTimeRange() *TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *MarketsBalancesRequest) GetMarketIds() []string {
	if x != nil {
		return x.MarketIds
	}
	return nil
}

type MarketsBalancesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// returns map of market_id and its balances sorted by time ASC
	MarketsBalances map[string]*MarketBalances `protobuf:"bytes,1,rep,name=markets_balances,json=marketsBalances,proto3" json:"markets_balances,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MarketsBalancesReply) Reset() {
	*x = MarketsBalancesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketsBalancesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketsBalancesReply) ProtoMessage() {}

func (x *MarketsBalancesReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketsBalancesReply.ProtoReflect.Descriptor instead.
func (*MarketsBalancesReply) Descriptor() ([]byte, []int) {
	return file_v1_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *MarketsBalancesReply) GetMarketsBalances() map[string]*MarketBalances {
	if x != nil {
		return x.MarketsBalances
	}
	return nil
}

type MarketBalances struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MarketBalance []*MarketBalance `protobuf:"bytes,1,rep,name=market_balance,json=marketBalance,proto3" json:"market_balance,omitempty"`
}

func (x *MarketBalances) Reset() {
	*x = MarketBalances{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_analytics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketBalances) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketBalances) ProtoMessage() {}

func (x *MarketBalances) ProtoReflect() protoreflect.Message {
	mi := &file_v1_analytics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketBalances.ProtoReflect.Descriptor instead.
func (*MarketBalances) Descriptor() ([]byte, []int) {
	return file_v1_analytics_proto_rawDescGZIP(), []int{2}
}

func (x *MarketBalances) GetMarketBalance() []*MarketBalance {
	if x != nil {
		return x.MarketBalance
	}
	return nil
}

type MarketBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// base balance
	BaseBalance int64 `protobuf:"varint,1,opt,name=base_balance,json=baseBalance,proto3" json:"base_balance,omitempty"`
	// quote balance
	QuoteBalance int64 `protobuf:"varint,2,opt,name=quote_balance,json=quoteBalance,proto3" json:"quote_balance,omitempty"`
	// point in time when market had this balance
	Time string `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *MarketBalance) Reset() {
	*x = MarketBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_analytics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketBalance) ProtoMessage() {}

func (x *MarketBalance) ProtoReflect() protoreflect.Message {
	mi := &file_v1_analytics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketBalance.ProtoReflect.Descriptor instead.
func (*MarketBalance) Descriptor() ([]byte, []int) {
	return file_v1_analytics_proto_rawDescGZIP(), []int{3}
}

func (x *MarketBalance) GetBaseBalance() int64 {
	if x != nil {
		return x.BaseBalance
	}
	return 0
}

func (x *MarketBalance) GetQuoteBalance() int64 {
	if x != nil {
		return x.QuoteBalance
	}
	return 0
}

func (x *MarketBalance) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type MarketsPricesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// time_range fetch prices for time range
	TimeRange *TimeRange `protobuf:"bytes,1,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	// fetch prices for specific one or more market's, if no market_id is passed balances will be fetched for all
	MarketIds []string `protobuf:"bytes,2,rep,name=market_ids,json=marketIds,proto3" json:"market_ids,omitempty"`
	// reference fiat currency to which base and quote will be converted
	ReferenceCurrency string `protobuf:"bytes,3,opt,name=reference_currency,json=referenceCurrency,proto3" json:"reference_currency,omitempty"`
}

func (x *MarketsPricesRequest) Reset() {
	*x = MarketsPricesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_analytics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketsPricesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketsPricesRequest) ProtoMessage() {}

func (x *MarketsPricesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_analytics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketsPricesRequest.ProtoReflect.Descriptor instead.
func (*MarketsPricesRequest) Descriptor() ([]byte, []int) {
	return file_v1_analytics_proto_rawDescGZIP(), []int{4}
}

func (x *MarketsPricesRequest) GetTimeRange() *TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *MarketsPricesRequest) GetMarketIds() []string {
	if x != nil {
		return x.MarketIds
	}
	return nil
}

func (x *MarketsPricesRequest) GetReferenceCurrency() string {
	if x != nil {
		return x.ReferenceCurrency
	}
	return ""
}

type MarketsPricesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// returns map of market_id and its prices sorted by time ASC
	MarketsPrices map[string]*MarketPrices `protobuf:"bytes,1,rep,name=markets_prices,json=marketsPrices,proto3" json:"markets_prices,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MarketsPricesReply) Reset() {
	*x = MarketsPricesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_analytics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketsPricesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketsPricesReply) ProtoMessage() {}

func (x *MarketsPricesReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_analytics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketsPricesReply.ProtoReflect.Descriptor instead.
func (*MarketsPricesReply) Descriptor() ([]byte, []int) {
	return file_v1_analytics_proto_rawDescGZIP(), []int{5}
}

func (x *MarketsPricesReply) GetMarketsPrices() map[string]*MarketPrices {
	if x != nil {
		return x.MarketsPrices
	}
	return nil
}

type MarketPrices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// returns list of Market's and their prices
	MarketPrice []*MarketPrice `protobuf:"bytes,1,rep,name=market_price,json=marketPrice,proto3" json:"market_price,omitempty"`
}

func (x *MarketPrices) Reset() {
	*x = MarketPrices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_analytics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketPrices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketPrices) ProtoMessage() {}

func (x *MarketPrices) ProtoReflect() protoreflect.Message {
	mi := &file_v1_analytics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketPrices.ProtoReflect.Descriptor instead.
func (*MarketPrices) Descriptor() ([]byte, []int) {
	return file_v1_analytics_proto_rawDescGZIP(), []int{6}
}

func (x *MarketPrices) GetMarketPrice() []*MarketPrice {
	if x != nil {
		return x.MarketPrice
	}
	return nil
}

type MarketPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// base is the price of one unit of the base asset, priced in quote asset assuming the trade amount was, ie. 100 unit of that asset
	BasePrice float32 `protobuf:"fixed32,1,opt,name=base_price,json=basePrice,proto3" json:"base_price,omitempty"`
	// quote amount
	QuotePrice float32 `protobuf:"fixed32,2,opt,name=quote_price,json=quotePrice,proto3" json:"quote_price,omitempty"`
	// base price converted to reference one
	BaseReferencePrice float32 `protobuf:"fixed32,3,opt,name=base_reference_price,json=baseReferencePrice,proto3" json:"base_reference_price,omitempty"`
	// quote price converted to reference one
	QuoteReferencePrice float32 `protobuf:"fixed32,4,opt,name=quote_reference_price,json=quoteReferencePrice,proto3" json:"quote_reference_price,omitempty"`
	// point in time when market had this price
	Time string `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *MarketPrice) Reset() {
	*x = MarketPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_analytics_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketPrice) ProtoMessage() {}

func (x *MarketPrice) ProtoReflect() protoreflect.Message {
	mi := &file_v1_analytics_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketPrice.ProtoReflect.Descriptor instead.
func (*MarketPrice) Descriptor() ([]byte, []int) {
	return file_v1_analytics_proto_rawDescGZIP(), []int{7}
}

func (x *MarketPrice) GetBasePrice() float32 {
	if x != nil {
		return x.BasePrice
	}
	return 0
}

func (x *MarketPrice) GetQuotePrice() float32 {
	if x != nil {
		return x.QuotePrice
	}
	return 0
}

func (x *MarketPrice) GetBaseReferencePrice() float32 {
	if x != nil {
		return x.BaseReferencePrice
	}
	return 0
}

func (x *MarketPrice) GetQuoteReferencePrice() float32 {
	if x != nil {
		return x.QuoteReferencePrice
	}
	return 0
}

func (x *MarketPrice) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

// TimeRange is flexible type used to determine time span for which specific
// api will fetch data, either one of predefined_period or custom_period should be provided.
type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// predefined time_period till now
	PredefinedPeriod PredefinedPeriod `protobuf:"varint,1,opt,name=predefined_period,json=predefinedPeriod,proto3,enum=tdexa.v1.PredefinedPeriod" json:"predefined_period,omitempty"`
	// granular time range
	CustomPeriod *CustomPeriod `protobuf:"bytes,2,opt,name=custom_period,json=customPeriod,proto3" json:"custom_period,omitempty"`
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_analytics_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_v1_analytics_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_v1_analytics_proto_rawDescGZIP(), []int{8}
}

func (x *TimeRange) GetPredefinedPeriod() PredefinedPeriod {
	if x != nil {
		return x.PredefinedPeriod
	}
	return PredefinedPeriod_NULL
}

func (x *TimeRange) GetCustomPeriod() *CustomPeriod {
	if x != nil {
		return x.CustomPeriod
	}
	return nil
}

type CustomPeriod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// start_date in RFC3339 format
	StartDate string `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// end_date in RFC3339 format
	EndDate string `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *CustomPeriod) Reset() {
	*x = CustomPeriod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_analytics_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomPeriod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomPeriod) ProtoMessage() {}

func (x *CustomPeriod) ProtoReflect() protoreflect.Message {
	mi := &file_v1_analytics_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomPeriod.ProtoReflect.Descriptor instead.
func (*CustomPeriod) Descriptor() ([]byte, []int) {
	return file_v1_analytics_proto_rawDescGZIP(), []int{9}
}

func (x *CustomPeriod) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *CustomPeriod) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

var File_v1_analytics_proto protoreflect.FileDescriptor

var file_v1_analytics_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x64, 0x65, 0x78, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x16,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x64, 0x65,
	0x78, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x73, 0x22, 0xd4, 0x01, 0x0a, 0x14, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x5e, 0x0a, 0x10, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74,
	0x64, 0x65, 0x78, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x1a, 0x5c, 0x0a, 0x14, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x64,
	0x65, 0x78, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x50, 0x0a, 0x0e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x64, 0x65,
	0x78, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x0d, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0x6b, 0x0a, 0x0d, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22,
	0x98, 0x01, 0x0a, 0x14, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74,
	0x64, 0x65, 0x78, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0xc6, 0x01, 0x0a, 0x12, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x56, 0x0a, 0x0e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x74, 0x64, 0x65, 0x78,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x58, 0x0a, 0x12, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x48, 0x0a, 0x0c, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0c, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x64, 0x65, 0x78,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0xc7, 0x01,
	0x0a, 0x0b, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a,
	0x14, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x32, 0x0a, 0x15, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x10, 0x70, 0x72,
	0x65, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x3b,
	0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x0c, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x48, 0x0a, 0x0c, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x2a, 0x77, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x64, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c,
	0x4c, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x48, 0x4f, 0x55, 0x52,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x44, 0x41, 0x59, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x10, 0x03,
	0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x33, 0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48,
	0x53, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x59, 0x45, 0x41, 0x52, 0x5f, 0x54, 0x4f, 0x5f, 0x44,
	0x41, 0x54, 0x45, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x06, 0x32, 0xdf,
	0x01, 0x0a, 0x09, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x12, 0x6c, 0x0a, 0x0f,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12,
	0x20, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x64, 0x0a, 0x0d, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x74, 0x64,
	0x65, 0x78, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x64,
	0x65, 0x78, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73,
	0x42, 0xa8, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x61, 0x2e, 0x76,
	0x31, 0x42, 0x0e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x64, 0x65, 0x78, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x64, 0x65,
	0x78, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x64, 0x65, 0x78, 0x61, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54,
	0x58, 0x58, 0xaa, 0x02, 0x08, 0x54, 0x64, 0x65, 0x78, 0x61, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08,
	0x54, 0x64, 0x65, 0x78, 0x61, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x54, 0x64, 0x65, 0x78, 0x61,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x09, 0x54, 0x64, 0x65, 0x78, 0x61, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v1_analytics_proto_rawDescOnce sync.Once
	file_v1_analytics_proto_rawDescData = file_v1_analytics_proto_rawDesc
)

func file_v1_analytics_proto_rawDescGZIP() []byte {
	file_v1_analytics_proto_rawDescOnce.Do(func() {
		file_v1_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_analytics_proto_rawDescData)
	})
	return file_v1_analytics_proto_rawDescData
}

var file_v1_analytics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_v1_analytics_proto_goTypes = []interface{}{
	(PredefinedPeriod)(0),          // 0: tdexa.v1.PredefinedPeriod
	(*MarketsBalancesRequest)(nil), // 1: tdexa.v1.MarketsBalancesRequest
	(*MarketsBalancesReply)(nil),   // 2: tdexa.v1.MarketsBalancesReply
	(*MarketBalances)(nil),         // 3: tdexa.v1.MarketBalances
	(*MarketBalance)(nil),          // 4: tdexa.v1.MarketBalance
	(*MarketsPricesRequest)(nil),   // 5: tdexa.v1.MarketsPricesRequest
	(*MarketsPricesReply)(nil),     // 6: tdexa.v1.MarketsPricesReply
	(*MarketPrices)(nil),           // 7: tdexa.v1.MarketPrices
	(*MarketPrice)(nil),            // 8: tdexa.v1.MarketPrice
	(*TimeRange)(nil),              // 9: tdexa.v1.TimeRange
	(*CustomPeriod)(nil),           // 10: tdexa.v1.CustomPeriod
	nil,                            // 11: tdexa.v1.MarketsBalancesReply.MarketsBalancesEntry
	nil,                            // 12: tdexa.v1.MarketsPricesReply.MarketsPricesEntry
}
var file_v1_analytics_proto_depIdxs = []int32{
	9,  // 0: tdexa.v1.MarketsBalancesRequest.time_range:type_name -> tdexa.v1.TimeRange
	11, // 1: tdexa.v1.MarketsBalancesReply.markets_balances:type_name -> tdexa.v1.MarketsBalancesReply.MarketsBalancesEntry
	4,  // 2: tdexa.v1.MarketBalances.market_balance:type_name -> tdexa.v1.MarketBalance
	9,  // 3: tdexa.v1.MarketsPricesRequest.time_range:type_name -> tdexa.v1.TimeRange
	12, // 4: tdexa.v1.MarketsPricesReply.markets_prices:type_name -> tdexa.v1.MarketsPricesReply.MarketsPricesEntry
	8,  // 5: tdexa.v1.MarketPrices.market_price:type_name -> tdexa.v1.MarketPrice
	0,  // 6: tdexa.v1.TimeRange.predefined_period:type_name -> tdexa.v1.PredefinedPeriod
	10, // 7: tdexa.v1.TimeRange.custom_period:type_name -> tdexa.v1.CustomPeriod
	3,  // 8: tdexa.v1.MarketsBalancesReply.MarketsBalancesEntry.value:type_name -> tdexa.v1.MarketBalances
	7,  // 9: tdexa.v1.MarketsPricesReply.MarketsPricesEntry.value:type_name -> tdexa.v1.MarketPrices
	1,  // 10: tdexa.v1.Analytics.MarketsBalances:input_type -> tdexa.v1.MarketsBalancesRequest
	5,  // 11: tdexa.v1.Analytics.MarketsPrices:input_type -> tdexa.v1.MarketsPricesRequest
	2,  // 12: tdexa.v1.Analytics.MarketsBalances:output_type -> tdexa.v1.MarketsBalancesReply
	6,  // 13: tdexa.v1.Analytics.MarketsPrices:output_type -> tdexa.v1.MarketsPricesReply
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_v1_analytics_proto_init() }
func file_v1_analytics_proto_init() {
	if File_v1_analytics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_analytics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketsBalancesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_analytics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketsBalancesReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_analytics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketBalances); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_analytics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketBalance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_analytics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketsPricesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_analytics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketsPricesReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_analytics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketPrices); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_analytics_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketPrice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_analytics_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_analytics_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomPeriod); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_analytics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_analytics_proto_goTypes,
		DependencyIndexes: file_v1_analytics_proto_depIdxs,
		EnumInfos:         file_v1_analytics_proto_enumTypes,
		MessageInfos:      file_v1_analytics_proto_msgTypes,
	}.Build()
	File_v1_analytics_proto = out.File
	file_v1_analytics_proto_rawDesc = nil
	file_v1_analytics_proto_goTypes = nil
	file_v1_analytics_proto_depIdxs = nil
}