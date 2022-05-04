// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: tinkoff/invest/grpc/sandbox.proto

package investapi

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

//Запрос открытия счёта в песочнице.
type OpenSandboxAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OpenSandboxAccountRequest) Reset() {
	*x = OpenSandboxAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_invest_grpc_sandbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenSandboxAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenSandboxAccountRequest) ProtoMessage() {}

func (x *OpenSandboxAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_invest_grpc_sandbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenSandboxAccountRequest.ProtoReflect.Descriptor instead.
func (*OpenSandboxAccountRequest) Descriptor() ([]byte, []int) {
	return file_tinkoff_invest_grpc_sandbox_proto_rawDescGZIP(), []int{0}
}

//Номер открытого счёта в песочнице.
type OpenSandboxAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"` //Номер счёта
}

func (x *OpenSandboxAccountResponse) Reset() {
	*x = OpenSandboxAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_invest_grpc_sandbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenSandboxAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenSandboxAccountResponse) ProtoMessage() {}

func (x *OpenSandboxAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_invest_grpc_sandbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenSandboxAccountResponse.ProtoReflect.Descriptor instead.
func (*OpenSandboxAccountResponse) Descriptor() ([]byte, []int) {
	return file_tinkoff_invest_grpc_sandbox_proto_rawDescGZIP(), []int{1}
}

func (x *OpenSandboxAccountResponse) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

//Запрос закрытия счёта в песочнице.
type CloseSandboxAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"` //Номер счёта
}

func (x *CloseSandboxAccountRequest) Reset() {
	*x = CloseSandboxAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_invest_grpc_sandbox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseSandboxAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseSandboxAccountRequest) ProtoMessage() {}

func (x *CloseSandboxAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_invest_grpc_sandbox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseSandboxAccountRequest.ProtoReflect.Descriptor instead.
func (*CloseSandboxAccountRequest) Descriptor() ([]byte, []int) {
	return file_tinkoff_invest_grpc_sandbox_proto_rawDescGZIP(), []int{2}
}

func (x *CloseSandboxAccountRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

//Результат закрытия счёта в песочнице.
type CloseSandboxAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloseSandboxAccountResponse) Reset() {
	*x = CloseSandboxAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_invest_grpc_sandbox_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseSandboxAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseSandboxAccountResponse) ProtoMessage() {}

func (x *CloseSandboxAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_invest_grpc_sandbox_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseSandboxAccountResponse.ProtoReflect.Descriptor instead.
func (*CloseSandboxAccountResponse) Descriptor() ([]byte, []int) {
	return file_tinkoff_invest_grpc_sandbox_proto_rawDescGZIP(), []int{3}
}

//Запрос пополнения счёта в песочнице.
type SandboxPayInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string      `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"` //Номер счёта
	Amount    *MoneyValue `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`                        //Сумма пополнения счёта в рублях
}

func (x *SandboxPayInRequest) Reset() {
	*x = SandboxPayInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_invest_grpc_sandbox_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxPayInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxPayInRequest) ProtoMessage() {}

func (x *SandboxPayInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_invest_grpc_sandbox_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxPayInRequest.ProtoReflect.Descriptor instead.
func (*SandboxPayInRequest) Descriptor() ([]byte, []int) {
	return file_tinkoff_invest_grpc_sandbox_proto_rawDescGZIP(), []int{4}
}

func (x *SandboxPayInRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *SandboxPayInRequest) GetAmount() *MoneyValue {
	if x != nil {
		return x.Amount
	}
	return nil
}

//Результат пополнения счёта, текущий баланс.
type SandboxPayInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance *MoneyValue `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"` //Текущий баланс счёта
}

func (x *SandboxPayInResponse) Reset() {
	*x = SandboxPayInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_invest_grpc_sandbox_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxPayInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxPayInResponse) ProtoMessage() {}

func (x *SandboxPayInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_invest_grpc_sandbox_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxPayInResponse.ProtoReflect.Descriptor instead.
func (*SandboxPayInResponse) Descriptor() ([]byte, []int) {
	return file_tinkoff_invest_grpc_sandbox_proto_rawDescGZIP(), []int{5}
}

func (x *SandboxPayInResponse) GetBalance() *MoneyValue {
	if x != nil {
		return x.Balance
	}
	return nil
}

var File_tinkoff_invest_grpc_sandbox_proto protoreflect.FileDescriptor

var file_tinkoff_invest_grpc_sandbox_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x25, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x74, 0x69, 0x6e, 0x6b,
	0x6f, 0x66, 0x66, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x74, 0x69,
	0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2f, 0x69, 0x6e,
	0x76, 0x65, 0x73, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x19, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x61, 0x6e,
	0x64, 0x62, 0x6f, 0x78, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x3b, 0x0a, 0x1a, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x78, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x3b, 0x0a, 0x1a, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x1d, 0x0a, 0x1b,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7f, 0x0a, 0x13, 0x53,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x50, 0x61, 0x79, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x49, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x63, 0x0a, 0x14,
	0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x50, 0x61, 0x79, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f,
	0x6e, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x32, 0xae, 0x0c, 0x0a, 0x0e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x12, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x61, 0x6e,
	0x64, 0x62, 0x6f, 0x78, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x40, 0x2e, 0x74, 0x69,
	0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e,
	0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x78, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x8b, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x39, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66,
	0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9c,
	0x01, 0x0a, 0x13, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x41, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x74, 0x69, 0x6e, 0x6b,
	0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x85, 0x01,
	0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x37, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x74, 0x69,
	0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6e,
	0x64, 0x62, 0x6f, 0x78, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x37, 0x2e, 0x74, 0x69, 0x6e,
	0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8b, 0x01,
	0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x39, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3a, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x31, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6e, 0x64,
	0x62, 0x6f, 0x78, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x37, 0x2e, 0x74,
	0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x8b, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f,
	0x66, 0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x39, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x88, 0x01,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x50, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x37, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38,
	0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x0c, 0x53, 0x61, 0x6e,
	0x64, 0x62, 0x6f, 0x78, 0x50, 0x61, 0x79, 0x49, 0x6e, 0x12, 0x3a, 0x2e, 0x74, 0x69, 0x6e, 0x6b,
	0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x50, 0x61, 0x79, 0x49, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61,
	0x6e, 0x64, 0x62, 0x6f, 0x78, 0x50, 0x61, 0x79, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x61, 0x0a, 0x1c, 0x72, 0x75, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66,
	0x2e, 0x70, 0x69, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x0c, 0x2e, 0x2f, 0x3b, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61,
	0x70, 0x69, 0xa2, 0x02, 0x05, 0x54, 0x49, 0x41, 0x50, 0x49, 0xaa, 0x02, 0x14, 0x54, 0x69, 0x6e,
	0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x41, 0x70, 0x69, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x11, 0x54, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x5c, 0x49, 0x6e, 0x76, 0x65,
	0x73, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tinkoff_invest_grpc_sandbox_proto_rawDescOnce sync.Once
	file_tinkoff_invest_grpc_sandbox_proto_rawDescData = file_tinkoff_invest_grpc_sandbox_proto_rawDesc
)

func file_tinkoff_invest_grpc_sandbox_proto_rawDescGZIP() []byte {
	file_tinkoff_invest_grpc_sandbox_proto_rawDescOnce.Do(func() {
		file_tinkoff_invest_grpc_sandbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_tinkoff_invest_grpc_sandbox_proto_rawDescData)
	})
	return file_tinkoff_invest_grpc_sandbox_proto_rawDescData
}

var file_tinkoff_invest_grpc_sandbox_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tinkoff_invest_grpc_sandbox_proto_goTypes = []interface{}{
	(*OpenSandboxAccountRequest)(nil),   // 0: tinkoff.public.invest.api.contract.v1.OpenSandboxAccountRequest
	(*OpenSandboxAccountResponse)(nil),  // 1: tinkoff.public.invest.api.contract.v1.OpenSandboxAccountResponse
	(*CloseSandboxAccountRequest)(nil),  // 2: tinkoff.public.invest.api.contract.v1.CloseSandboxAccountRequest
	(*CloseSandboxAccountResponse)(nil), // 3: tinkoff.public.invest.api.contract.v1.CloseSandboxAccountResponse
	(*SandboxPayInRequest)(nil),         // 4: tinkoff.public.invest.api.contract.v1.SandboxPayInRequest
	(*SandboxPayInResponse)(nil),        // 5: tinkoff.public.invest.api.contract.v1.SandboxPayInResponse
	(*MoneyValue)(nil),                  // 6: tinkoff.public.invest.api.contract.v1.MoneyValue
	(*GetAccountsRequest)(nil),          // 7: tinkoff.public.invest.api.contract.v1.GetAccountsRequest
	(*PostOrderRequest)(nil),            // 8: tinkoff.public.invest.api.contract.v1.PostOrderRequest
	(*GetOrdersRequest)(nil),            // 9: tinkoff.public.invest.api.contract.v1.GetOrdersRequest
	(*CancelOrderRequest)(nil),          // 10: tinkoff.public.invest.api.contract.v1.CancelOrderRequest
	(*GetOrderStateRequest)(nil),        // 11: tinkoff.public.invest.api.contract.v1.GetOrderStateRequest
	(*PositionsRequest)(nil),            // 12: tinkoff.public.invest.api.contract.v1.PositionsRequest
	(*OperationsRequest)(nil),           // 13: tinkoff.public.invest.api.contract.v1.OperationsRequest
	(*PortfolioRequest)(nil),            // 14: tinkoff.public.invest.api.contract.v1.PortfolioRequest
	(*GetAccountsResponse)(nil),         // 15: tinkoff.public.invest.api.contract.v1.GetAccountsResponse
	(*PostOrderResponse)(nil),           // 16: tinkoff.public.invest.api.contract.v1.PostOrderResponse
	(*GetOrdersResponse)(nil),           // 17: tinkoff.public.invest.api.contract.v1.GetOrdersResponse
	(*CancelOrderResponse)(nil),         // 18: tinkoff.public.invest.api.contract.v1.CancelOrderResponse
	(*OrderState)(nil),                  // 19: tinkoff.public.invest.api.contract.v1.OrderState
	(*PositionsResponse)(nil),           // 20: tinkoff.public.invest.api.contract.v1.PositionsResponse
	(*OperationsResponse)(nil),          // 21: tinkoff.public.invest.api.contract.v1.OperationsResponse
	(*PortfolioResponse)(nil),           // 22: tinkoff.public.invest.api.contract.v1.PortfolioResponse
}
var file_tinkoff_invest_grpc_sandbox_proto_depIdxs = []int32{
	6,  // 0: tinkoff.public.invest.api.contract.v1.SandboxPayInRequest.amount:type_name -> tinkoff.public.invest.api.contract.v1.MoneyValue
	6,  // 1: tinkoff.public.invest.api.contract.v1.SandboxPayInResponse.balance:type_name -> tinkoff.public.invest.api.contract.v1.MoneyValue
	0,  // 2: tinkoff.public.invest.api.contract.v1.SandboxService.OpenSandboxAccount:input_type -> tinkoff.public.invest.api.contract.v1.OpenSandboxAccountRequest
	7,  // 3: tinkoff.public.invest.api.contract.v1.SandboxService.GetSandboxAccounts:input_type -> tinkoff.public.invest.api.contract.v1.GetAccountsRequest
	2,  // 4: tinkoff.public.invest.api.contract.v1.SandboxService.CloseSandboxAccount:input_type -> tinkoff.public.invest.api.contract.v1.CloseSandboxAccountRequest
	8,  // 5: tinkoff.public.invest.api.contract.v1.SandboxService.PostSandboxOrder:input_type -> tinkoff.public.invest.api.contract.v1.PostOrderRequest
	9,  // 6: tinkoff.public.invest.api.contract.v1.SandboxService.GetSandboxOrders:input_type -> tinkoff.public.invest.api.contract.v1.GetOrdersRequest
	10, // 7: tinkoff.public.invest.api.contract.v1.SandboxService.CancelSandboxOrder:input_type -> tinkoff.public.invest.api.contract.v1.CancelOrderRequest
	11, // 8: tinkoff.public.invest.api.contract.v1.SandboxService.GetSandboxOrderState:input_type -> tinkoff.public.invest.api.contract.v1.GetOrderStateRequest
	12, // 9: tinkoff.public.invest.api.contract.v1.SandboxService.GetSandboxPositions:input_type -> tinkoff.public.invest.api.contract.v1.PositionsRequest
	13, // 10: tinkoff.public.invest.api.contract.v1.SandboxService.GetSandboxOperations:input_type -> tinkoff.public.invest.api.contract.v1.OperationsRequest
	14, // 11: tinkoff.public.invest.api.contract.v1.SandboxService.GetSandboxPortfolio:input_type -> tinkoff.public.invest.api.contract.v1.PortfolioRequest
	4,  // 12: tinkoff.public.invest.api.contract.v1.SandboxService.SandboxPayIn:input_type -> tinkoff.public.invest.api.contract.v1.SandboxPayInRequest
	1,  // 13: tinkoff.public.invest.api.contract.v1.SandboxService.OpenSandboxAccount:output_type -> tinkoff.public.invest.api.contract.v1.OpenSandboxAccountResponse
	15, // 14: tinkoff.public.invest.api.contract.v1.SandboxService.GetSandboxAccounts:output_type -> tinkoff.public.invest.api.contract.v1.GetAccountsResponse
	3,  // 15: tinkoff.public.invest.api.contract.v1.SandboxService.CloseSandboxAccount:output_type -> tinkoff.public.invest.api.contract.v1.CloseSandboxAccountResponse
	16, // 16: tinkoff.public.invest.api.contract.v1.SandboxService.PostSandboxOrder:output_type -> tinkoff.public.invest.api.contract.v1.PostOrderResponse
	17, // 17: tinkoff.public.invest.api.contract.v1.SandboxService.GetSandboxOrders:output_type -> tinkoff.public.invest.api.contract.v1.GetOrdersResponse
	18, // 18: tinkoff.public.invest.api.contract.v1.SandboxService.CancelSandboxOrder:output_type -> tinkoff.public.invest.api.contract.v1.CancelOrderResponse
	19, // 19: tinkoff.public.invest.api.contract.v1.SandboxService.GetSandboxOrderState:output_type -> tinkoff.public.invest.api.contract.v1.OrderState
	20, // 20: tinkoff.public.invest.api.contract.v1.SandboxService.GetSandboxPositions:output_type -> tinkoff.public.invest.api.contract.v1.PositionsResponse
	21, // 21: tinkoff.public.invest.api.contract.v1.SandboxService.GetSandboxOperations:output_type -> tinkoff.public.invest.api.contract.v1.OperationsResponse
	22, // 22: tinkoff.public.invest.api.contract.v1.SandboxService.GetSandboxPortfolio:output_type -> tinkoff.public.invest.api.contract.v1.PortfolioResponse
	5,  // 23: tinkoff.public.invest.api.contract.v1.SandboxService.SandboxPayIn:output_type -> tinkoff.public.invest.api.contract.v1.SandboxPayInResponse
	13, // [13:24] is the sub-list for method output_type
	2,  // [2:13] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_tinkoff_invest_grpc_sandbox_proto_init() }
func file_tinkoff_invest_grpc_sandbox_proto_init() {
	if File_tinkoff_invest_grpc_sandbox_proto != nil {
		return
	}
	file_tinkoff_invest_grpc_common_proto_init()
	file_tinkoff_invest_grpc_orders_proto_init()
	file_tinkoff_invest_grpc_operations_proto_init()
	file_tinkoff_invest_grpc_users_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tinkoff_invest_grpc_sandbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenSandboxAccountRequest); i {
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
		file_tinkoff_invest_grpc_sandbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenSandboxAccountResponse); i {
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
		file_tinkoff_invest_grpc_sandbox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseSandboxAccountRequest); i {
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
		file_tinkoff_invest_grpc_sandbox_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseSandboxAccountResponse); i {
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
		file_tinkoff_invest_grpc_sandbox_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandboxPayInRequest); i {
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
		file_tinkoff_invest_grpc_sandbox_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandboxPayInResponse); i {
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
			RawDescriptor: file_tinkoff_invest_grpc_sandbox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tinkoff_invest_grpc_sandbox_proto_goTypes,
		DependencyIndexes: file_tinkoff_invest_grpc_sandbox_proto_depIdxs,
		MessageInfos:      file_tinkoff_invest_grpc_sandbox_proto_msgTypes,
	}.Build()
	File_tinkoff_invest_grpc_sandbox_proto = out.File
	file_tinkoff_invest_grpc_sandbox_proto_rawDesc = nil
	file_tinkoff_invest_grpc_sandbox_proto_goTypes = nil
	file_tinkoff_invest_grpc_sandbox_proto_depIdxs = nil
}
