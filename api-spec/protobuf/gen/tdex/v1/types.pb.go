// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tdex/v1/types.proto

package tdexv1

import (
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

type TradeType int32

const (
	TradeType_TRADE_TYPE_BUY  TradeType = 0
	TradeType_TRADE_TYPE_SELL TradeType = 1
)

// Enum value maps for TradeType.
var (
	TradeType_name = map[int32]string{
		0: "TRADE_TYPE_BUY",
		1: "TRADE_TYPE_SELL",
	}
	TradeType_value = map[string]int32{
		"TRADE_TYPE_BUY":  0,
		"TRADE_TYPE_SELL": 1,
	}
)

func (x TradeType) Enum() *TradeType {
	p := new(TradeType)
	*p = x
	return p
}

func (x TradeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TradeType) Descriptor() protoreflect.EnumDescriptor {
	return file_tdex_v1_types_proto_enumTypes[0].Descriptor()
}

func (TradeType) Type() protoreflect.EnumType {
	return &file_tdex_v1_types_proto_enumTypes[0]
}

func (x TradeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TradeType.Descriptor instead.
func (TradeType) EnumDescriptor() ([]byte, []int) {
	return file_tdex_v1_types_proto_rawDescGZIP(), []int{0}
}

type ContentType int32

const (
	ContentType_CONTENT_TYPE_JSON        ContentType = 0
	ContentType_CONTENT_TYPE_GRPC        ContentType = 1
	ContentType_CONTENT_TYPE_GRPCWEB     ContentType = 2
	ContentType_CONTENT_TYPE_GRPCWEBTEXT ContentType = 3
)

// Enum value maps for ContentType.
var (
	ContentType_name = map[int32]string{
		0: "CONTENT_TYPE_JSON",
		1: "CONTENT_TYPE_GRPC",
		2: "CONTENT_TYPE_GRPCWEB",
		3: "CONTENT_TYPE_GRPCWEBTEXT",
	}
	ContentType_value = map[string]int32{
		"CONTENT_TYPE_JSON":        0,
		"CONTENT_TYPE_GRPC":        1,
		"CONTENT_TYPE_GRPCWEB":     2,
		"CONTENT_TYPE_GRPCWEBTEXT": 3,
	}
)

func (x ContentType) Enum() *ContentType {
	p := new(ContentType)
	*p = x
	return p
}

func (x ContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_tdex_v1_types_proto_enumTypes[1].Descriptor()
}

func (ContentType) Type() protoreflect.EnumType {
	return &file_tdex_v1_types_proto_enumTypes[1]
}

func (x ContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContentType.Descriptor instead.
func (ContentType) EnumDescriptor() ([]byte, []int) {
	return file_tdex_v1_types_proto_rawDescGZIP(), []int{1}
}

// Custom Types
type Fee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasisPoint int64  `protobuf:"varint,1,opt,name=basis_point,json=basisPoint,proto3" json:"basis_point,omitempty"`
	Fixed      *Fixed `protobuf:"bytes,2,opt,name=fixed,proto3" json:"fixed,omitempty"`
}

func (x *Fee) Reset() {
	*x = Fee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fee) ProtoMessage() {}

func (x *Fee) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fee.ProtoReflect.Descriptor instead.
func (*Fee) Descriptor() ([]byte, []int) {
	return file_tdex_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *Fee) GetBasisPoint() int64 {
	if x != nil {
		return x.BasisPoint
	}
	return 0
}

func (x *Fee) GetFixed() *Fixed {
	if x != nil {
		return x.Fixed
	}
	return nil
}

type Fixed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseFee  int64 `protobuf:"varint,1,opt,name=base_fee,json=baseFee,proto3" json:"base_fee,omitempty"`
	QuoteFee int64 `protobuf:"varint,2,opt,name=quote_fee,json=quoteFee,proto3" json:"quote_fee,omitempty"`
}

func (x *Fixed) Reset() {
	*x = Fixed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_v1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fixed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fixed) ProtoMessage() {}

func (x *Fixed) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_v1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fixed.ProtoReflect.Descriptor instead.
func (*Fixed) Descriptor() ([]byte, []int) {
	return file_tdex_v1_types_proto_rawDescGZIP(), []int{1}
}

func (x *Fixed) GetBaseFee() int64 {
	if x != nil {
		return x.BaseFee
	}
	return 0
}

func (x *Fixed) GetQuoteFee() int64 {
	if x != nil {
		return x.QuoteFee
	}
	return 0
}

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAmount  uint64 `protobuf:"varint,1,opt,name=base_amount,json=baseAmount,proto3" json:"base_amount,omitempty"`
	QuoteAmount uint64 `protobuf:"varint,2,opt,name=quote_amount,json=quoteAmount,proto3" json:"quote_amount,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_v1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_v1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_tdex_v1_types_proto_rawDescGZIP(), []int{2}
}

func (x *Balance) GetBaseAmount() uint64 {
	if x != nil {
		return x.BaseAmount
	}
	return 0
}

func (x *Balance) GetQuoteAmount() uint64 {
	if x != nil {
		return x.QuoteAmount
	}
	return 0
}

type BalanceWithFee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance *Balance `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Fee     *Fee     `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (x *BalanceWithFee) Reset() {
	*x = BalanceWithFee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_v1_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceWithFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceWithFee) ProtoMessage() {}

func (x *BalanceWithFee) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_v1_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceWithFee.ProtoReflect.Descriptor instead.
func (*BalanceWithFee) Descriptor() ([]byte, []int) {
	return file_tdex_v1_types_proto_rawDescGZIP(), []int{3}
}

func (x *BalanceWithFee) GetBalance() *Balance {
	if x != nil {
		return x.Balance
	}
	return nil
}

func (x *BalanceWithFee) GetFee() *Fee {
	if x != nil {
		return x.Fee
	}
	return nil
}

type Market struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAsset  string `protobuf:"bytes,1,opt,name=base_asset,json=baseAsset,proto3" json:"base_asset,omitempty"`
	QuoteAsset string `protobuf:"bytes,2,opt,name=quote_asset,json=quoteAsset,proto3" json:"quote_asset,omitempty"`
}

func (x *Market) Reset() {
	*x = Market{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_v1_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Market) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Market) ProtoMessage() {}

func (x *Market) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_v1_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Market.ProtoReflect.Descriptor instead.
func (*Market) Descriptor() ([]byte, []int) {
	return file_tdex_v1_types_proto_rawDescGZIP(), []int{4}
}

func (x *Market) GetBaseAsset() string {
	if x != nil {
		return x.BaseAsset
	}
	return ""
}

func (x *Market) GetQuoteAsset() string {
	if x != nil {
		return x.QuoteAsset
	}
	return ""
}

type MarketWithFee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Market *Market `protobuf:"bytes,1,opt,name=market,proto3" json:"market,omitempty"`
	Fee    *Fee    `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (x *MarketWithFee) Reset() {
	*x = MarketWithFee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_v1_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketWithFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketWithFee) ProtoMessage() {}

func (x *MarketWithFee) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_v1_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketWithFee.ProtoReflect.Descriptor instead.
func (*MarketWithFee) Descriptor() ([]byte, []int) {
	return file_tdex_v1_types_proto_rawDescGZIP(), []int{5}
}

func (x *MarketWithFee) GetMarket() *Market {
	if x != nil {
		return x.Market
	}
	return nil
}

func (x *MarketWithFee) GetFee() *Fee {
	if x != nil {
		return x.Fee
	}
	return nil
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasePrice  float64 `protobuf:"fixed64,1,opt,name=base_price,json=basePrice,proto3" json:"base_price,omitempty"`
	QuotePrice float64 `protobuf:"fixed64,2,opt,name=quote_price,json=quotePrice,proto3" json:"quote_price,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_v1_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_v1_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_tdex_v1_types_proto_rawDescGZIP(), []int{6}
}

func (x *Price) GetBasePrice() float64 {
	if x != nil {
		return x.BasePrice
	}
	return 0
}

func (x *Price) GetQuotePrice() float64 {
	if x != nil {
		return x.QuotePrice
	}
	return 0
}

type Preview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price   *Price   `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	Fee     *Fee     `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
	Amount  uint64   `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Asset   string   `protobuf:"bytes,4,opt,name=asset,proto3" json:"asset,omitempty"`
	Balance *Balance `protobuf:"bytes,5,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *Preview) Reset() {
	*x = Preview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_v1_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Preview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Preview) ProtoMessage() {}

func (x *Preview) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_v1_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Preview.ProtoReflect.Descriptor instead.
func (*Preview) Descriptor() ([]byte, []int) {
	return file_tdex_v1_types_proto_rawDescGZIP(), []int{7}
}

func (x *Preview) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *Preview) GetFee() *Fee {
	if x != nil {
		return x.Fee
	}
	return nil
}

func (x *Preview) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Preview) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *Preview) GetBalance() *Balance {
	if x != nil {
		return x.Balance
	}
	return nil
}

var File_tdex_v1_types_proto protoreflect.FileDescriptor

var file_tdex_v1_types_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x64, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x22, 0x4c,
	0x0a, 0x03, 0x46, 0x65, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x69, 0x73, 0x5f, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x69,
	0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x66, 0x69, 0x78, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x69, 0x78, 0x65, 0x64, 0x52, 0x05, 0x66, 0x69, 0x78, 0x65, 0x64, 0x22, 0x3f, 0x0a, 0x05,
	0x46, 0x69, 0x78, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x66, 0x65,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x46, 0x65, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x46, 0x65, 0x65, 0x22, 0x4d, 0x0a,
	0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62,
	0x61, 0x73, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5c, 0x0a, 0x0e,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x74, 0x68, 0x46, 0x65, 0x65, 0x12, 0x2a,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x03, 0x66, 0x65,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x03, 0x66, 0x65, 0x65, 0x22, 0x48, 0x0a, 0x06, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x22, 0x58, 0x0a, 0x0d, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x57, 0x69,
	0x74, 0x68, 0x46, 0x65, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x1e,
	0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x64,
	0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x03, 0x66, 0x65, 0x65, 0x22, 0x47,
	0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x62, 0x61, 0x73,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x12, 0x24, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x03, 0x66, 0x65, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x65, 0x65, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x2a, 0x34, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x0e, 0x54, 0x52, 0x41, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42,
	0x55, 0x59, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x52, 0x41, 0x44, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x01, 0x2a, 0x73, 0x0a, 0x0b, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4e, 0x54,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x00, 0x12,
	0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x47, 0x52, 0x50, 0x43, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x52, 0x50, 0x43, 0x57, 0x45, 0x42, 0x10, 0x02,
	0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x47, 0x52, 0x50, 0x43, 0x57, 0x45, 0x42, 0x54, 0x45, 0x58, 0x54, 0x10, 0x03, 0x42, 0xa3,
	0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x42, 0x0a,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x64, 0x65, 0x78, 0x2d, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x64, 0x65, 0x78, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x64, 0x65, 0x78, 0x2f,
	0x76, 0x31, 0x3b, 0x74, 0x64, 0x65, 0x78, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa,
	0x02, 0x07, 0x54, 0x64, 0x65, 0x78, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x54, 0x64, 0x65, 0x78,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x54, 0x64, 0x65, 0x78, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x54, 0x64, 0x65, 0x78,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tdex_v1_types_proto_rawDescOnce sync.Once
	file_tdex_v1_types_proto_rawDescData = file_tdex_v1_types_proto_rawDesc
)

func file_tdex_v1_types_proto_rawDescGZIP() []byte {
	file_tdex_v1_types_proto_rawDescOnce.Do(func() {
		file_tdex_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_tdex_v1_types_proto_rawDescData)
	})
	return file_tdex_v1_types_proto_rawDescData
}

var file_tdex_v1_types_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_tdex_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tdex_v1_types_proto_goTypes = []interface{}{
	(TradeType)(0),         // 0: tdex.v1.TradeType
	(ContentType)(0),       // 1: tdex.v1.ContentType
	(*Fee)(nil),            // 2: tdex.v1.Fee
	(*Fixed)(nil),          // 3: tdex.v1.Fixed
	(*Balance)(nil),        // 4: tdex.v1.Balance
	(*BalanceWithFee)(nil), // 5: tdex.v1.BalanceWithFee
	(*Market)(nil),         // 6: tdex.v1.Market
	(*MarketWithFee)(nil),  // 7: tdex.v1.MarketWithFee
	(*Price)(nil),          // 8: tdex.v1.Price
	(*Preview)(nil),        // 9: tdex.v1.Preview
}
var file_tdex_v1_types_proto_depIdxs = []int32{
	3, // 0: tdex.v1.Fee.fixed:type_name -> tdex.v1.Fixed
	4, // 1: tdex.v1.BalanceWithFee.balance:type_name -> tdex.v1.Balance
	2, // 2: tdex.v1.BalanceWithFee.fee:type_name -> tdex.v1.Fee
	6, // 3: tdex.v1.MarketWithFee.market:type_name -> tdex.v1.Market
	2, // 4: tdex.v1.MarketWithFee.fee:type_name -> tdex.v1.Fee
	8, // 5: tdex.v1.Preview.price:type_name -> tdex.v1.Price
	2, // 6: tdex.v1.Preview.fee:type_name -> tdex.v1.Fee
	4, // 7: tdex.v1.Preview.balance:type_name -> tdex.v1.Balance
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_tdex_v1_types_proto_init() }
func file_tdex_v1_types_proto_init() {
	if File_tdex_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tdex_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fee); i {
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
		file_tdex_v1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fixed); i {
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
		file_tdex_v1_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
		file_tdex_v1_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceWithFee); i {
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
		file_tdex_v1_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Market); i {
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
		file_tdex_v1_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketWithFee); i {
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
		file_tdex_v1_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_tdex_v1_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Preview); i {
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
			RawDescriptor: file_tdex_v1_types_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tdex_v1_types_proto_goTypes,
		DependencyIndexes: file_tdex_v1_types_proto_depIdxs,
		EnumInfos:         file_tdex_v1_types_proto_enumTypes,
		MessageInfos:      file_tdex_v1_types_proto_msgTypes,
	}.Build()
	File_tdex_v1_types_proto = out.File
	file_tdex_v1_types_proto_rawDesc = nil
	file_tdex_v1_types_proto_goTypes = nil
	file_tdex_v1_types_proto_depIdxs = nil
}
