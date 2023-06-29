// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: trade.proto

package db

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

type Trade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TradeId      string `protobuf:"bytes,1,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`
	AccountId    string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	TradeType    string `protobuf:"bytes,3,opt,name=trade_type,json=tradeType,proto3" json:"trade_type,omitempty"`
	TradeStatus  string `protobuf:"bytes,4,opt,name=trade_status,json=tradeStatus,proto3" json:"trade_status,omitempty"`
	FromCurrency string `protobuf:"bytes,5,opt,name=from_currency,json=fromCurrency,proto3" json:"from_currency,omitempty"`
	ToCurrency   string `protobuf:"bytes,6,opt,name=to_currency,json=toCurrency,proto3" json:"to_currency,omitempty"`
	FromAmount   int64  `protobuf:"varint,7,opt,name=from_amount,json=fromAmount,proto3" json:"from_amount,omitempty"`
	FinalAmount  int64  `protobuf:"varint,8,opt,name=final_amount,json=finalAmount,proto3" json:"final_amount,omitempty"`
}

func (x *Trade) Reset() {
	*x = Trade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trade) ProtoMessage() {}

func (x *Trade) ProtoReflect() protoreflect.Message {
	mi := &file_trade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trade.ProtoReflect.Descriptor instead.
func (*Trade) Descriptor() ([]byte, []int) {
	return file_trade_proto_rawDescGZIP(), []int{0}
}

func (x *Trade) GetTradeId() string {
	if x != nil {
		return x.TradeId
	}
	return ""
}

func (x *Trade) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Trade) GetTradeType() string {
	if x != nil {
		return x.TradeType
	}
	return ""
}

func (x *Trade) GetTradeStatus() string {
	if x != nil {
		return x.TradeStatus
	}
	return ""
}

func (x *Trade) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *Trade) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

func (x *Trade) GetFromAmount() int64 {
	if x != nil {
		return x.FromAmount
	}
	return 0
}

func (x *Trade) GetFinalAmount() int64 {
	if x != nil {
		return x.FinalAmount
	}
	return 0
}

type CreateTradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trade *Trade `protobuf:"bytes,1,opt,name=trade,proto3" json:"trade,omitempty"`
}

func (x *CreateTradeRequest) Reset() {
	*x = CreateTradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTradeRequest) ProtoMessage() {}

func (x *CreateTradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTradeRequest.ProtoReflect.Descriptor instead.
func (*CreateTradeRequest) Descriptor() ([]byte, []int) {
	return file_trade_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTradeRequest) GetTrade() *Trade {
	if x != nil {
		return x.Trade
	}
	return nil
}

type CreateTradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TradeId string `protobuf:"bytes,1,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`
}

func (x *CreateTradeResponse) Reset() {
	*x = CreateTradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTradeResponse) ProtoMessage() {}

func (x *CreateTradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trade_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTradeResponse.ProtoReflect.Descriptor instead.
func (*CreateTradeResponse) Descriptor() ([]byte, []int) {
	return file_trade_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTradeResponse) GetTradeId() string {
	if x != nil {
		return x.TradeId
	}
	return ""
}

type DeleteTradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TradeId string `protobuf:"bytes,1,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`
}

func (x *DeleteTradeRequest) Reset() {
	*x = DeleteTradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTradeRequest) ProtoMessage() {}

func (x *DeleteTradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTradeRequest.ProtoReflect.Descriptor instead.
func (*DeleteTradeRequest) Descriptor() ([]byte, []int) {
	return file_trade_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTradeRequest) GetTradeId() string {
	if x != nil {
		return x.TradeId
	}
	return ""
}

type UpdateTradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trade *Trade `protobuf:"bytes,1,opt,name=trade,proto3" json:"trade,omitempty"`
}

func (x *UpdateTradeRequest) Reset() {
	*x = UpdateTradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTradeRequest) ProtoMessage() {}

func (x *UpdateTradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTradeRequest.ProtoReflect.Descriptor instead.
func (*UpdateTradeRequest) Descriptor() ([]byte, []int) {
	return file_trade_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTradeRequest) GetTrade() *Trade {
	if x != nil {
		return x.Trade
	}
	return nil
}

type GetTradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TradeId string `protobuf:"bytes,1,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`
}

func (x *GetTradeRequest) Reset() {
	*x = GetTradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTradeRequest) ProtoMessage() {}

func (x *GetTradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTradeRequest.ProtoReflect.Descriptor instead.
func (*GetTradeRequest) Descriptor() ([]byte, []int) {
	return file_trade_proto_rawDescGZIP(), []int{5}
}

func (x *GetTradeRequest) GetTradeId() string {
	if x != nil {
		return x.TradeId
	}
	return ""
}

type GetTradeByAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *GetTradeByAccountRequest) Reset() {
	*x = GetTradeByAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTradeByAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTradeByAccountRequest) ProtoMessage() {}

func (x *GetTradeByAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTradeByAccountRequest.ProtoReflect.Descriptor instead.
func (*GetTradeByAccountRequest) Descriptor() ([]byte, []int) {
	return file_trade_proto_rawDescGZIP(), []int{6}
}

func (x *GetTradeByAccountRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type GetTradeByAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trades []*Trade `protobuf:"bytes,1,rep,name=trades,proto3" json:"trades,omitempty"`
}

func (x *GetTradeByAccountResponse) Reset() {
	*x = GetTradeByAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTradeByAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTradeByAccountResponse) ProtoMessage() {}

func (x *GetTradeByAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trade_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTradeByAccountResponse.ProtoReflect.Descriptor instead.
func (*GetTradeByAccountResponse) Descriptor() ([]byte, []int) {
	return file_trade_proto_rawDescGZIP(), []int{7}
}

func (x *GetTradeByAccountResponse) GetTrades() []*Trade {
	if x != nil {
		return x.Trades
	}
	return nil
}

var File_trade_proto protoreflect.FileDescriptor

var file_trade_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x64,
	0x62, 0x22, 0x8d, 0x02, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x66, 0x72, 0x6f, 0x6d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x35, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x64, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x05, 0x74, 0x72, 0x61, 0x64, 0x65, 0x22, 0x30, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x64, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x05, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64,
	0x22, 0x39, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x42, 0x79, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x64, 0x62, 0x2e, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x06, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x70, 0x70, 0x73, 0x4c, 0x61,
	0x62, 0x2d, 0x4b, 0x45, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x65, 0x76, 0x65,
	0x72, 0x79, 0x73, 0x68, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x64, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trade_proto_rawDescOnce sync.Once
	file_trade_proto_rawDescData = file_trade_proto_rawDesc
)

func file_trade_proto_rawDescGZIP() []byte {
	file_trade_proto_rawDescOnce.Do(func() {
		file_trade_proto_rawDescData = protoimpl.X.CompressGZIP(file_trade_proto_rawDescData)
	})
	return file_trade_proto_rawDescData
}

var file_trade_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_trade_proto_goTypes = []interface{}{
	(*Trade)(nil),                     // 0: db.Trade
	(*CreateTradeRequest)(nil),        // 1: db.CreateTradeRequest
	(*CreateTradeResponse)(nil),       // 2: db.CreateTradeResponse
	(*DeleteTradeRequest)(nil),        // 3: db.DeleteTradeRequest
	(*UpdateTradeRequest)(nil),        // 4: db.UpdateTradeRequest
	(*GetTradeRequest)(nil),           // 5: db.GetTradeRequest
	(*GetTradeByAccountRequest)(nil),  // 6: db.GetTradeByAccountRequest
	(*GetTradeByAccountResponse)(nil), // 7: db.GetTradeByAccountResponse
}
var file_trade_proto_depIdxs = []int32{
	0, // 0: db.CreateTradeRequest.trade:type_name -> db.Trade
	0, // 1: db.UpdateTradeRequest.trade:type_name -> db.Trade
	0, // 2: db.GetTradeByAccountResponse.trades:type_name -> db.Trade
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_trade_proto_init() }
func file_trade_proto_init() {
	if File_trade_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trade); i {
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
		file_trade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTradeRequest); i {
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
		file_trade_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTradeResponse); i {
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
		file_trade_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTradeRequest); i {
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
		file_trade_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTradeRequest); i {
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
		file_trade_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTradeRequest); i {
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
		file_trade_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTradeByAccountRequest); i {
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
		file_trade_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTradeByAccountResponse); i {
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
			RawDescriptor: file_trade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trade_proto_goTypes,
		DependencyIndexes: file_trade_proto_depIdxs,
		MessageInfos:      file_trade_proto_msgTypes,
	}.Build()
	File_trade_proto = out.File
	file_trade_proto_rawDesc = nil
	file_trade_proto_goTypes = nil
	file_trade_proto_depIdxs = nil
}
