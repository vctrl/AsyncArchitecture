// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: billing/billing.proto

package billing

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreatePlusTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	TaskId string `protobuf:"bytes,2,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
	Amount int64  `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *CreatePlusTransactionRequest) Reset() {
	*x = CreatePlusTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_billing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlusTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlusTransactionRequest) ProtoMessage() {}

func (x *CreatePlusTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_billing_billing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlusTransactionRequest.ProtoReflect.Descriptor instead.
func (*CreatePlusTransactionRequest) Descriptor() ([]byte, []int) {
	return file_billing_billing_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePlusTransactionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreatePlusTransactionRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *CreatePlusTransactionRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreatePlusTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreatePlusTransactionResponse) Reset() {
	*x = CreatePlusTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_billing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlusTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlusTransactionResponse) ProtoMessage() {}

func (x *CreatePlusTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_billing_billing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlusTransactionResponse.ProtoReflect.Descriptor instead.
func (*CreatePlusTransactionResponse) Descriptor() ([]byte, []int) {
	return file_billing_billing_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePlusTransactionResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type CreateMinusTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	TaskId string `protobuf:"bytes,2,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
	Amount int64  `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *CreateMinusTransactionRequest) Reset() {
	*x = CreateMinusTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_billing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMinusTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMinusTransactionRequest) ProtoMessage() {}

func (x *CreateMinusTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_billing_billing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMinusTransactionRequest.ProtoReflect.Descriptor instead.
func (*CreateMinusTransactionRequest) Descriptor() ([]byte, []int) {
	return file_billing_billing_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMinusTransactionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateMinusTransactionRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *CreateMinusTransactionRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreateMinusTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateMinusTransactionResponse) Reset() {
	*x = CreateMinusTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_billing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMinusTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMinusTransactionResponse) ProtoMessage() {}

func (x *CreateMinusTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_billing_billing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMinusTransactionResponse.ProtoReflect.Descriptor instead.
func (*CreateMinusTransactionResponse) Descriptor() ([]byte, []int) {
	return file_billing_billing_proto_rawDescGZIP(), []int{3}
}

func (x *CreateMinusTransactionResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type CloseBillingCycleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloseBillingCycleRequest) Reset() {
	*x = CloseBillingCycleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_billing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseBillingCycleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseBillingCycleRequest) ProtoMessage() {}

func (x *CloseBillingCycleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_billing_billing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseBillingCycleRequest.ProtoReflect.Descriptor instead.
func (*CloseBillingCycleRequest) Descriptor() ([]byte, []int) {
	return file_billing_billing_proto_rawDescGZIP(), []int{4}
}

type CloseBillingCycleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CloseBillingCycleResponse) Reset() {
	*x = CloseBillingCycleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_billing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseBillingCycleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseBillingCycleResponse) ProtoMessage() {}

func (x *CloseBillingCycleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_billing_billing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseBillingCycleResponse.ProtoReflect.Descriptor instead.
func (*CloseBillingCycleResponse) Descriptor() ([]byte, []int) {
	return file_billing_billing_proto_rawDescGZIP(), []int{5}
}

func (x *CloseBillingCycleResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_billing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_billing_billing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_billing_billing_proto_rawDescGZIP(), []int{6}
}

func (x *Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Status) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_billing_billing_proto protoreflect.FileDescriptor

var file_billing_billing_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6c, 0x75, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x40, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x75, 0x73, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x67, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x6e, 0x75, 0x73,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x41, 0x0a, 0x1e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1a, 0x0a,
	0x18, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x79, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x19, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x88, 0x02, 0x0a, 0x07, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x12, 0x56, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x75,
	0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x75, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x75, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69,
	0x6e, 0x75, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69,
	0x6e, 0x75, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x11, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x42,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_billing_billing_proto_rawDescOnce sync.Once
	file_billing_billing_proto_rawDescData = file_billing_billing_proto_rawDesc
)

func file_billing_billing_proto_rawDescGZIP() []byte {
	file_billing_billing_proto_rawDescOnce.Do(func() {
		file_billing_billing_proto_rawDescData = protoimpl.X.CompressGZIP(file_billing_billing_proto_rawDescData)
	})
	return file_billing_billing_proto_rawDescData
}

var file_billing_billing_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_billing_billing_proto_goTypes = []interface{}{
	(*CreatePlusTransactionRequest)(nil),   // 0: CreatePlusTransactionRequest
	(*CreatePlusTransactionResponse)(nil),  // 1: CreatePlusTransactionResponse
	(*CreateMinusTransactionRequest)(nil),  // 2: CreateMinusTransactionRequest
	(*CreateMinusTransactionResponse)(nil), // 3: CreateMinusTransactionResponse
	(*CloseBillingCycleRequest)(nil),       // 4: CloseBillingCycleRequest
	(*CloseBillingCycleResponse)(nil),      // 5: CloseBillingCycleResponse
	(*Status)(nil),                         // 6: Status
}
var file_billing_billing_proto_depIdxs = []int32{
	6, // 0: CreatePlusTransactionResponse.status:type_name -> Status
	6, // 1: CreateMinusTransactionResponse.status:type_name -> Status
	6, // 2: CloseBillingCycleResponse.status:type_name -> Status
	0, // 3: Billing.CreatePlusTransaction:input_type -> CreatePlusTransactionRequest
	2, // 4: Billing.CreateMinusTransaction:input_type -> CreateMinusTransactionRequest
	4, // 5: Billing.CloseBillingCycle:input_type -> CloseBillingCycleRequest
	1, // 6: Billing.CreatePlusTransaction:output_type -> CreatePlusTransactionResponse
	3, // 7: Billing.CreateMinusTransaction:output_type -> CreateMinusTransactionResponse
	5, // 8: Billing.CloseBillingCycle:output_type -> CloseBillingCycleResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_billing_billing_proto_init() }
func file_billing_billing_proto_init() {
	if File_billing_billing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_billing_billing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlusTransactionRequest); i {
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
		file_billing_billing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlusTransactionResponse); i {
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
		file_billing_billing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMinusTransactionRequest); i {
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
		file_billing_billing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMinusTransactionResponse); i {
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
		file_billing_billing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseBillingCycleRequest); i {
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
		file_billing_billing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseBillingCycleResponse); i {
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
		file_billing_billing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_billing_billing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_billing_billing_proto_goTypes,
		DependencyIndexes: file_billing_billing_proto_depIdxs,
		MessageInfos:      file_billing_billing_proto_msgTypes,
	}.Build()
	File_billing_billing_proto = out.File
	file_billing_billing_proto_rawDesc = nil
	file_billing_billing_proto_goTypes = nil
	file_billing_billing_proto_depIdxs = nil
}
