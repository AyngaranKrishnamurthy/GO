// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: proto/services.proto

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

type CreateEmpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`       //Blog = Emp && blog = emp
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`   //author_id = name
	Level  string `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"` //title = level
	Stream string `protobuf:"bytes,4,opt,name=stream,proto3" json:"stream,omitempty"`
}

func (x *CreateEmpRequest) Reset() {
	*x = CreateEmpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmpRequest) ProtoMessage() {}

func (x *CreateEmpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmpRequest.ProtoReflect.Descriptor instead.
func (*CreateEmpRequest) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEmpRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateEmpRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateEmpRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *CreateEmpRequest) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

type CreateEmpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"` // will have a emp id
}

func (x *CreateEmpResponse) Reset() {
	*x = CreateEmpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmpResponse) ProtoMessage() {}

func (x *CreateEmpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmpResponse.ProtoReflect.Descriptor instead.
func (*CreateEmpResponse) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEmpResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ReadEmpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadEmpRequest) Reset() {
	*x = ReadEmpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadEmpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadEmpRequest) ProtoMessage() {}

func (x *ReadEmpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadEmpRequest.ProtoReflect.Descriptor instead.
func (*ReadEmpRequest) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{2}
}

func (x *ReadEmpRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadEmpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *ReadEmpResponse) Reset() {
	*x = ReadEmpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadEmpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadEmpResponse) ProtoMessage() {}

func (x *ReadEmpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadEmpResponse.ProtoReflect.Descriptor instead.
func (*ReadEmpResponse) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{3}
}

func (x *ReadEmpResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UpdateEmpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nname string `protobuf:"bytes,2,opt,name=nname,proto3" json:"nname,omitempty"`
}

func (x *UpdateEmpRequest) Reset() {
	*x = UpdateEmpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmpRequest) ProtoMessage() {}

func (x *UpdateEmpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmpRequest.ProtoReflect.Descriptor instead.
func (*UpdateEmpRequest) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEmpRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateEmpRequest) GetNname() string {
	if x != nil {
		return x.Nname
	}
	return ""
}

type UpdateEmpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *UpdateEmpResponse) Reset() {
	*x = UpdateEmpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmpResponse) ProtoMessage() {}

func (x *UpdateEmpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmpResponse.ProtoReflect.Descriptor instead.
func (*UpdateEmpResponse) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateEmpResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DeleteEmpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEmpRequest) Reset() {
	*x = DeleteEmpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEmpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmpRequest) ProtoMessage() {}

func (x *DeleteEmpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmpRequest.ProtoReflect.Descriptor instead.
func (*DeleteEmpRequest) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteEmpRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteEmpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *DeleteEmpResponse) Reset() {
	*x = DeleteEmpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEmpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmpResponse) ProtoMessage() {}

func (x *DeleteEmpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmpResponse.ProtoReflect.Descriptor instead.
func (*DeleteEmpResponse) Descriptor() ([]byte, []int) {
	return file_proto_services_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteEmpResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
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
		mi := &file_proto_services_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserRequest) ProtoMessage() {}

func (x *AuthUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[8]
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
	return file_proto_services_proto_rawDescGZIP(), []int{8}
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

	Eid     string `protobuf:"bytes,1,opt,name=eid,proto3" json:"eid,omitempty"`
	Empname string `protobuf:"bytes,2,opt,name=Empname,proto3" json:"Empname,omitempty"`
	Level   string `protobuf:"bytes,3,opt,name=Level,proto3" json:"Level,omitempty"`
	Stream  string `protobuf:"bytes,4,opt,name=Stream,proto3" json:"Stream,omitempty"`
}

func (x *AuthUserResponse) Reset() {
	*x = AuthUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserResponse) ProtoMessage() {}

func (x *AuthUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_proto_msgTypes[9]
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
	return file_proto_services_proto_rawDescGZIP(), []int{9}
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

func (x *AuthUserResponse) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *AuthUserResponse) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

var File_proto_services_proto protoreflect.FileDescriptor

var file_proto_services_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x22, 0x29, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x20,
	0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x27, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x38, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x22,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x29, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x27, 0x0a,
	0x0f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6c, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x45, 0x6d, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45,
	0x6d, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x32, 0xc3, 0x02, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70,
	0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6d, 0x70, 0x12, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6d, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x45, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a,
	0x08, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_services_proto_rawDescOnce sync.Once
	file_proto_services_proto_rawDescData = file_proto_services_proto_rawDesc
)

func file_proto_services_proto_rawDescGZIP() []byte {
	file_proto_services_proto_rawDescOnce.Do(func() {
		file_proto_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_services_proto_rawDescData)
	})
	return file_proto_services_proto_rawDescData
}

var file_proto_services_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_services_proto_goTypes = []interface{}{
	(*CreateEmpRequest)(nil),  // 0: proto.CreateEmpRequest
	(*CreateEmpResponse)(nil), // 1: proto.CreateEmpResponse
	(*ReadEmpRequest)(nil),    // 2: proto.ReadEmpRequest
	(*ReadEmpResponse)(nil),   // 3: proto.ReadEmpResponse
	(*UpdateEmpRequest)(nil),  // 4: proto.UpdateEmpRequest
	(*UpdateEmpResponse)(nil), // 5: proto.UpdateEmpResponse
	(*DeleteEmpRequest)(nil),  // 6: proto.DeleteEmpRequest
	(*DeleteEmpResponse)(nil), // 7: proto.DeleteEmpResponse
	(*AuthUserRequest)(nil),   // 8: proto.AuthUserRequest
	(*AuthUserResponse)(nil),  // 9: proto.AuthUserResponse
}
var file_proto_services_proto_depIdxs = []int32{
	0, // 0: proto.EmpService.CreateEmp:input_type -> proto.CreateEmpRequest
	2, // 1: proto.EmpService.ReadEmp:input_type -> proto.ReadEmpRequest
	4, // 2: proto.EmpService.UpdateEmp:input_type -> proto.UpdateEmpRequest
	6, // 3: proto.EmpService.DeleteEmp:input_type -> proto.DeleteEmpRequest
	8, // 4: proto.EmpService.AuthUser:input_type -> proto.AuthUserRequest
	1, // 5: proto.EmpService.CreateEmp:output_type -> proto.CreateEmpResponse
	3, // 6: proto.EmpService.ReadEmp:output_type -> proto.ReadEmpResponse
	5, // 7: proto.EmpService.UpdateEmp:output_type -> proto.UpdateEmpResponse
	7, // 8: proto.EmpService.DeleteEmp:output_type -> proto.DeleteEmpResponse
	9, // 9: proto.EmpService.AuthUser:output_type -> proto.AuthUserResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_services_proto_init() }
func file_proto_services_proto_init() {
	if File_proto_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmpRequest); i {
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
		file_proto_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmpResponse); i {
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
		file_proto_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadEmpRequest); i {
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
		file_proto_services_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadEmpResponse); i {
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
		file_proto_services_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmpRequest); i {
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
		file_proto_services_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmpResponse); i {
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
		file_proto_services_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEmpRequest); i {
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
		file_proto_services_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEmpResponse); i {
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
		file_proto_services_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_services_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_proto_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_services_proto_goTypes,
		DependencyIndexes: file_proto_services_proto_depIdxs,
		MessageInfos:      file_proto_services_proto_msgTypes,
	}.Build()
	File_proto_services_proto = out.File
	file_proto_services_proto_rawDesc = nil
	file_proto_services_proto_goTypes = nil
	file_proto_services_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EmpServiceClient is the client API for EmpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmpServiceClient interface {
	CreateEmp(ctx context.Context, in *CreateEmpRequest, opts ...grpc.CallOption) (*CreateEmpResponse, error)
	ReadEmp(ctx context.Context, in *ReadEmpRequest, opts ...grpc.CallOption) (*ReadEmpResponse, error)
	UpdateEmp(ctx context.Context, in *UpdateEmpRequest, opts ...grpc.CallOption) (*UpdateEmpResponse, error)
	DeleteEmp(ctx context.Context, in *DeleteEmpRequest, opts ...grpc.CallOption) (*DeleteEmpResponse, error)
	AuthUser(ctx context.Context, in *AuthUserRequest, opts ...grpc.CallOption) (*AuthUserResponse, error)
}

type empServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmpServiceClient(cc grpc.ClientConnInterface) EmpServiceClient {
	return &empServiceClient{cc}
}

func (c *empServiceClient) CreateEmp(ctx context.Context, in *CreateEmpRequest, opts ...grpc.CallOption) (*CreateEmpResponse, error) {
	out := new(CreateEmpResponse)
	err := c.cc.Invoke(ctx, "/proto.EmpService/CreateEmp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *empServiceClient) ReadEmp(ctx context.Context, in *ReadEmpRequest, opts ...grpc.CallOption) (*ReadEmpResponse, error) {
	out := new(ReadEmpResponse)
	err := c.cc.Invoke(ctx, "/proto.EmpService/ReadEmp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *empServiceClient) UpdateEmp(ctx context.Context, in *UpdateEmpRequest, opts ...grpc.CallOption) (*UpdateEmpResponse, error) {
	out := new(UpdateEmpResponse)
	err := c.cc.Invoke(ctx, "/proto.EmpService/UpdateEmp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *empServiceClient) DeleteEmp(ctx context.Context, in *DeleteEmpRequest, opts ...grpc.CallOption) (*DeleteEmpResponse, error) {
	out := new(DeleteEmpResponse)
	err := c.cc.Invoke(ctx, "/proto.EmpService/DeleteEmp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *empServiceClient) AuthUser(ctx context.Context, in *AuthUserRequest, opts ...grpc.CallOption) (*AuthUserResponse, error) {
	out := new(AuthUserResponse)
	err := c.cc.Invoke(ctx, "/proto.EmpService/AuthUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmpServiceServer is the server API for EmpService service.
type EmpServiceServer interface {
	CreateEmp(context.Context, *CreateEmpRequest) (*CreateEmpResponse, error)
	ReadEmp(context.Context, *ReadEmpRequest) (*ReadEmpResponse, error)
	UpdateEmp(context.Context, *UpdateEmpRequest) (*UpdateEmpResponse, error)
	DeleteEmp(context.Context, *DeleteEmpRequest) (*DeleteEmpResponse, error)
	AuthUser(context.Context, *AuthUserRequest) (*AuthUserResponse, error)
}

// UnimplementedEmpServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmpServiceServer struct {
}

func (*UnimplementedEmpServiceServer) CreateEmp(context.Context, *CreateEmpRequest) (*CreateEmpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmp not implemented")
}
func (*UnimplementedEmpServiceServer) ReadEmp(context.Context, *ReadEmpRequest) (*ReadEmpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEmp not implemented")
}
func (*UnimplementedEmpServiceServer) UpdateEmp(context.Context, *UpdateEmpRequest) (*UpdateEmpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmp not implemented")
}
func (*UnimplementedEmpServiceServer) DeleteEmp(context.Context, *DeleteEmpRequest) (*DeleteEmpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmp not implemented")
}
func (*UnimplementedEmpServiceServer) AuthUser(context.Context, *AuthUserRequest) (*AuthUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthUser not implemented")
}

func RegisterEmpServiceServer(s *grpc.Server, srv EmpServiceServer) {
	s.RegisterService(&_EmpService_serviceDesc, srv)
}

func _EmpService_CreateEmp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmpServiceServer).CreateEmp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmpService/CreateEmp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmpServiceServer).CreateEmp(ctx, req.(*CreateEmpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmpService_ReadEmp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadEmpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmpServiceServer).ReadEmp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmpService/ReadEmp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmpServiceServer).ReadEmp(ctx, req.(*ReadEmpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmpService_UpdateEmp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmpServiceServer).UpdateEmp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmpService/UpdateEmp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmpServiceServer).UpdateEmp(ctx, req.(*UpdateEmpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmpService_DeleteEmp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEmpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmpServiceServer).DeleteEmp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmpService/DeleteEmp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmpServiceServer).DeleteEmp(ctx, req.(*DeleteEmpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmpService_AuthUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmpServiceServer).AuthUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmpService/AuthUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmpServiceServer).AuthUser(ctx, req.(*AuthUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmpService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EmpService",
	HandlerType: (*EmpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmp",
			Handler:    _EmpService_CreateEmp_Handler,
		},
		{
			MethodName: "ReadEmp",
			Handler:    _EmpService_ReadEmp_Handler,
		},
		{
			MethodName: "UpdateEmp",
			Handler:    _EmpService_UpdateEmp_Handler,
		},
		{
			MethodName: "DeleteEmp",
			Handler:    _EmpService_DeleteEmp_Handler,
		},
		{
			MethodName: "AuthUser",
			Handler:    _EmpService_AuthUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/services.proto",
}
