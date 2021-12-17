// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: proto/service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eid     string `protobuf:"bytes,1,opt,name=Eid,proto3" json:"Eid,omitempty"`
	Empname string `protobuf:"bytes,2,opt,name=Empname,proto3" json:"Empname,omitempty"`
	Elevel  int32  `protobuf:"varint,3,opt,name=Elevel,proto3" json:"Elevel,omitempty"`
	Estream string `protobuf:"bytes,4,opt,name=Estream,proto3" json:"Estream,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetEid() string {
	if x != nil {
		return x.Eid
	}
	return ""
}

func (x *CreateRequest) GetEmpname() string {
	if x != nil {
		return x.Empname
	}
	return ""
}

func (x *CreateRequest) GetElevel() int32 {
	if x != nil {
		return x.Elevel
	}
	return 0
}

func (x *CreateRequest) GetEstream() string {
	if x != nil {
		return x.Estream
	}
	return ""
}

type RetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eid string `protobuf:"bytes,1,opt,name=Eid,proto3" json:"Eid,omitempty"`
}

func (x *RetrieveRequest) Reset() {
	*x = RetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveRequest) ProtoMessage() {}

func (x *RetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveRequest.ProtoReflect.Descriptor instead.
func (*RetrieveRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *RetrieveRequest) GetEid() string {
	if x != nil {
		return x.Eid
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eid      string `protobuf:"bytes,1,opt,name=Eid,proto3" json:"Eid,omitempty"`
	Empname1 string `protobuf:"bytes,2,opt,name=Empname1,proto3" json:"Empname1,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetEid() string {
	if x != nil {
		return x.Eid
	}
	return ""
}

func (x *UpdateRequest) GetEmpname1() string {
	if x != nil {
		return x.Empname1
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eid string `protobuf:"bytes,1,opt,name=Eid,proto3" json:"Eid,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRequest) GetEid() string {
	if x != nil {
		return x.Eid
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *AuthResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type EidUsedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eid string `protobuf:"bytes,1,opt,name=Eid,proto3" json:"Eid,omitempty"` //username
}

func (x *EidUsedRequest) Reset() {
	*x = EidUsedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EidUsedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EidUsedRequest) ProtoMessage() {}

func (x *EidUsedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EidUsedRequest.ProtoReflect.Descriptor instead.
func (*EidUsedRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{5}
}

func (x *EidUsedRequest) GetEid() string {
	if x != nil {
		return x.Eid
	}
	return ""
}

type UsedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Used bool `protobuf:"varint,1,opt,name=used,proto3" json:"used,omitempty"`
}

func (x *UsedResponse) Reset() {
	*x = UsedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsedResponse) ProtoMessage() {}

func (x *UsedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsedResponse.ProtoReflect.Descriptor instead.
func (*UsedResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{6}
}

func (x *UsedResponse) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

type AuthUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *AuthUserRequest) Reset() {
	*x = AuthUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserRequest) ProtoMessage() {}

func (x *AuthUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserRequest.ProtoReflect.Descriptor instead.
func (*AuthUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{7}
}

func (x *AuthUserRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eid     string `protobuf:"bytes,1,opt,name=Eid,proto3" json:"Eid,omitempty"`
	Empname string `protobuf:"bytes,2,opt,name=Empname,proto3" json:"Empname,omitempty"`
	Elevel  int32  `protobuf:"varint,3,opt,name=Elevel,proto3" json:"Elevel,omitempty"`
	Estream string `protobuf:"bytes,4,opt,name=Estream,proto3" json:"Estream,omitempty"`
}

func (x *AuthUserResponse) Reset() {
	*x = AuthUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserResponse) ProtoMessage() {}

func (x *AuthUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserResponse.ProtoReflect.Descriptor instead.
func (*AuthUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{8}
}

func (x *AuthUserResponse) GetEid() string {
	if x != nil {
		return x.Eid
	}
	return ""
}

func (x *AuthUserResponse) GetEmpname() string {
	if x != nil {
		return x.Empname
	}
	return ""
}

func (x *AuthUserResponse) GetElevel() int32 {
	if x != nil {
		return x.Elevel
	}
	return 0
}

func (x *AuthUserResponse) GetEstream() string {
	if x != nil {
		return x.Estream
	}
	return ""
}

var File_proto_service_proto protoreflect.FileDescriptor

var file_proto_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x45, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x45, 0x6d, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x45, 0x6d, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x45, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x45, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x23, 0x0a, 0x0f, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x45, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x69, 0x64,
	0x22, 0x3d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x45, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x31, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x45, 0x6d, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x31, 0x22,
	0x21, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x45, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45,
	0x69, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x22, 0x0a, 0x0e, 0x45, 0x69, 0x64, 0x55,
	0x73, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x0c,
	0x55, 0x73, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64,
	0x22, 0x27, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x70, 0x0a, 0x10, 0x41, 0x75, 0x74,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x45, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x45, 0x6d, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x45, 0x6d, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x45, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x45, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x32, 0xd9, 0x02, 0x0a, 0x0b,
	0x63, 0x72, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x45, 0x69, 0x64, 0x55, 0x73, 0x65, 0x64, 0x12, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x69, 0x64, 0x55, 0x73, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x41, 0x75,
	0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_proto_rawDescOnce sync.Once
	file_proto_service_proto_rawDescData = file_proto_service_proto_rawDesc
)

func file_proto_service_proto_rawDescGZIP() []byte {
	file_proto_service_proto_rawDescOnce.Do(func() {
		file_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_proto_rawDescData)
	})
	return file_proto_service_proto_rawDescData
}

var file_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_service_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),    // 0: proto.CreateRequest
	(*RetrieveRequest)(nil),  // 1: proto.RetrieveRequest
	(*UpdateRequest)(nil),    // 2: proto.UpdateRequest
	(*DeleteRequest)(nil),    // 3: proto.DeleteRequest
	(*AuthResponse)(nil),     // 4: proto.AuthResponse
	(*EidUsedRequest)(nil),   // 5: proto.EidUsedRequest
	(*UsedResponse)(nil),     // 6: proto.UsedResponse
	(*AuthUserRequest)(nil),  // 7: proto.AuthUserRequest
	(*AuthUserResponse)(nil), // 8: proto.AuthUserResponse
}
var file_proto_service_proto_depIdxs = []int32{
	0, // 0: proto.crudService.Create:input_type -> proto.CreateRequest
	1, // 1: proto.crudService.Retrieve:input_type -> proto.RetrieveRequest
	2, // 2: proto.crudService.Update:input_type -> proto.UpdateRequest
	3, // 3: proto.crudService.Delete:input_type -> proto.DeleteRequest
	5, // 4: proto.crudService.EidUsed:input_type -> proto.EidUsedRequest
	7, // 5: proto.crudService.AuthUser:input_type -> proto.AuthUserRequest
	4, // 6: proto.crudService.Create:output_type -> proto.AuthResponse
	4, // 7: proto.crudService.Retrieve:output_type -> proto.AuthResponse
	4, // 8: proto.crudService.Update:output_type -> proto.AuthResponse
	4, // 9: proto.crudService.Delete:output_type -> proto.AuthResponse
	6, // 10: proto.crudService.EidUsed:output_type -> proto.UsedResponse
	8, // 11: proto.crudService.AuthUser:output_type -> proto.AuthUserResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_proto_init() }
func file_proto_service_proto_init() {
	if File_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveRequest); i {
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
		file_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_proto_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EidUsedRequest); i {
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
		file_proto_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsedResponse); i {
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
		file_proto_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUserRequest); i {
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
		file_proto_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUserResponse); i {
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
			RawDescriptor: file_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_proto_goTypes,
		DependencyIndexes: file_proto_service_proto_depIdxs,
		MessageInfos:      file_proto_service_proto_msgTypes,
	}.Build()
	File_proto_service_proto = out.File
	file_proto_service_proto_rawDesc = nil
	file_proto_service_proto_goTypes = nil
	file_proto_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CrudServiceClient is the client API for CrudService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CrudServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	EidUsed(ctx context.Context, in *EidUsedRequest, opts ...grpc.CallOption) (*UsedResponse, error)
	AuthUser(ctx context.Context, in *AuthUserRequest, opts ...grpc.CallOption) (*AuthUserResponse, error)
}

type crudServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrudServiceClient(cc grpc.ClientConnInterface) CrudServiceClient {
	return &crudServiceClient{cc}
}

func (c *crudServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.crudService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudServiceClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.crudService/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.crudService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.crudService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudServiceClient) EidUsed(ctx context.Context, in *EidUsedRequest, opts ...grpc.CallOption) (*UsedResponse, error) {
	out := new(UsedResponse)
	err := c.cc.Invoke(ctx, "/proto.crudService/EidUsed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudServiceClient) AuthUser(ctx context.Context, in *AuthUserRequest, opts ...grpc.CallOption) (*AuthUserResponse, error) {
	out := new(AuthUserResponse)
	err := c.cc.Invoke(ctx, "/proto.crudService/AuthUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrudServiceServer is the server API for CrudService service.
type CrudServiceServer interface {
	Create(context.Context, *CreateRequest) (*AuthResponse, error)
	Retrieve(context.Context, *RetrieveRequest) (*AuthResponse, error)
	Update(context.Context, *UpdateRequest) (*AuthResponse, error)
	Delete(context.Context, *DeleteRequest) (*AuthResponse, error)
	EidUsed(context.Context, *EidUsedRequest) (*UsedResponse, error)
	AuthUser(context.Context, *AuthUserRequest) (*AuthUserResponse, error)
}

// UnimplementedCrudServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCrudServiceServer struct {
}

func (*UnimplementedCrudServiceServer) Create(context.Context, *CreateRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCrudServiceServer) Retrieve(context.Context, *RetrieveRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (*UnimplementedCrudServiceServer) Update(context.Context, *UpdateRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCrudServiceServer) Delete(context.Context, *DeleteRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedCrudServiceServer) EidUsed(context.Context, *EidUsedRequest) (*UsedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EidUsed not implemented")
}
func (*UnimplementedCrudServiceServer) AuthUser(context.Context, *AuthUserRequest) (*AuthUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthUser not implemented")
}

func RegisterCrudServiceServer(s *grpc.Server, srv CrudServiceServer) {
	s.RegisterService(&_CrudService_serviceDesc, srv)
}

func _CrudService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.crudService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrudService_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServiceServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.crudService/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServiceServer).Retrieve(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrudService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.crudService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrudService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.crudService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrudService_EidUsed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EidUsedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServiceServer).EidUsed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.crudService/EidUsed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServiceServer).EidUsed(ctx, req.(*EidUsedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrudService_AuthUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServiceServer).AuthUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.crudService/AuthUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServiceServer).AuthUser(ctx, req.(*AuthUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CrudService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.crudService",
	HandlerType: (*CrudServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CrudService_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _CrudService_Retrieve_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CrudService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CrudService_Delete_Handler,
		},
		{
			MethodName: "EidUsed",
			Handler:    _CrudService_EidUsed_Handler,
		},
		{
			MethodName: "AuthUser",
			Handler:    _CrudService_AuthUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}