// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: grpc/proto/aiges_inner.proto

package proto

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

type Ret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret int32 `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
}

func (x *Ret) Reset() {
	*x = Ret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_aiges_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ret) ProtoMessage() {}

func (x *Ret) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_aiges_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ret.ProtoReflect.Descriptor instead.
func (*Ret) Descriptor() ([]byte, []int) {
	return file_grpc_proto_aiges_inner_proto_rawDescGZIP(), []int{0}
}

func (x *Ret) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

type RequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Len    uint64            `protobuf:"varint,2,opt,name=len,proto3" json:"len,omitempty"`
	Type   uint32            `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Status uint32            `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Desc   map[string]string `protobuf:"bytes,5,rep,name=desc,proto3" json:"desc,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data   []byte            `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RequestData) Reset() {
	*x = RequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_aiges_inner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestData) ProtoMessage() {}

func (x *RequestData) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_aiges_inner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestData.ProtoReflect.Descriptor instead.
func (*RequestData) Descriptor() ([]byte, []int) {
	return file_grpc_proto_aiges_inner_proto_rawDescGZIP(), []int{1}
}

func (x *RequestData) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RequestData) GetLen() uint64 {
	if x != nil {
		return x.Len
	}
	return 0
}

func (x *RequestData) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *RequestData) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RequestData) GetDesc() map[string]string {
	if x != nil {
		return x.Desc
	}
	return nil
}

func (x *RequestData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Len    uint64            `protobuf:"varint,2,opt,name=len,proto3" json:"len,omitempty"`
	Type   uint32            `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Status uint32            `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Desc   map[string]string `protobuf:"bytes,5,rep,name=desc,proto3" json:"desc,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data   []byte            `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseData) Reset() {
	*x = ResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_aiges_inner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseData) ProtoMessage() {}

func (x *ResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_aiges_inner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseData.ProtoReflect.Descriptor instead.
func (*ResponseData) Descriptor() ([]byte, []int) {
	return file_grpc_proto_aiges_inner_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseData) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ResponseData) GetLen() uint64 {
	if x != nil {
		return x.Len
	}
	return 0
}

func (x *ResponseData) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ResponseData) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ResponseData) GetDesc() map[string]string {
	if x != nil {
		return x.Desc
	}
	return nil
}

func (x *ResponseData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag  string          `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Ret  int32           `protobuf:"varint,2,opt,name=ret,proto3" json:"ret,omitempty"`
	List []*ResponseData `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_aiges_inner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_aiges_inner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_grpc_proto_aiges_inner_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Response) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *Response) GetList() []*ResponseData {
	if x != nil {
		return x.List
	}
	return nil
}

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config map[string]string `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_aiges_inner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_aiges_inner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_aiges_inner_proto_rawDescGZIP(), []int{4}
}

func (x *InitRequest) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag    string            `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Params map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	List   []*RequestData    `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_aiges_inner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_aiges_inner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_grpc_proto_aiges_inner_proto_rawDescGZIP(), []int{5}
}

func (x *Request) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Request) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *Request) GetList() []*RequestData {
	if x != nil {
		return x.List
	}
	return nil
}

type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_aiges_inner_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_aiges_inner_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_aiges_inner_proto_rawDescGZIP(), []int{6}
}

func (x *StreamRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_grpc_proto_aiges_inner_proto protoreflect.FileDescriptor

var file_grpc_proto_aiges_inner_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x69, 0x67,
	0x65, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x61, 0x69, 0x67, 0x65, 0x73, 0x22, 0x17, 0x0a, 0x03, 0x52, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x22, 0xdc,
	0x01, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6c,
	0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61,
	0x69, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x65, 0x73, 0x63, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xde, 0x01,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6c,
	0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61,
	0x69, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x65, 0x73, 0x63, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x57,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x27,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61,
	0x69, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x69, 0x67, 0x65, 0x73, 0x2e,
	0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a,
	0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb2, 0x01, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x32, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x69, 0x67, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x69, 0x67,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x23, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0xe8, 0x01, 0x0a, 0x0e, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x12, 0x2e, 0x61, 0x69, 0x67, 0x65, 0x73, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x69, 0x67,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x4f, 0x6e, 0x63, 0x65, 0x45, 0x78, 0x65, 0x63, 0x12, 0x0e, 0x2e, 0x61, 0x69,
	0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x69,
	0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x61,
	0x69, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x69, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x61, 0x69, 0x67, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x69, 0x67, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_grpc_proto_aiges_inner_proto_rawDescOnce sync.Once
	file_grpc_proto_aiges_inner_proto_rawDescData = file_grpc_proto_aiges_inner_proto_rawDesc
)

func file_grpc_proto_aiges_inner_proto_rawDescGZIP() []byte {
	file_grpc_proto_aiges_inner_proto_rawDescOnce.Do(func() {
		file_grpc_proto_aiges_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_aiges_inner_proto_rawDescData)
	})
	return file_grpc_proto_aiges_inner_proto_rawDescData
}

var file_grpc_proto_aiges_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_grpc_proto_aiges_inner_proto_goTypes = []interface{}{
	(*Ret)(nil),           // 0: aiges.Ret
	(*RequestData)(nil),   // 1: aiges.RequestData
	(*ResponseData)(nil),  // 2: aiges.ResponseData
	(*Response)(nil),      // 3: aiges.Response
	(*InitRequest)(nil),   // 4: aiges.InitRequest
	(*Request)(nil),       // 5: aiges.Request
	(*StreamRequest)(nil), // 6: aiges.StreamRequest
	nil,                   // 7: aiges.RequestData.DescEntry
	nil,                   // 8: aiges.ResponseData.DescEntry
	nil,                   // 9: aiges.InitRequest.ConfigEntry
	nil,                   // 10: aiges.Request.ParamsEntry
}
var file_grpc_proto_aiges_inner_proto_depIdxs = []int32{
	7,  // 0: aiges.RequestData.desc:type_name -> aiges.RequestData.DescEntry
	8,  // 1: aiges.ResponseData.desc:type_name -> aiges.ResponseData.DescEntry
	2,  // 2: aiges.Response.list:type_name -> aiges.ResponseData
	9,  // 3: aiges.InitRequest.config:type_name -> aiges.InitRequest.ConfigEntry
	10, // 4: aiges.Request.params:type_name -> aiges.Request.ParamsEntry
	1,  // 5: aiges.Request.list:type_name -> aiges.RequestData
	4,  // 6: aiges.WrapperService.wrapperInit:input_type -> aiges.InitRequest
	5,  // 7: aiges.WrapperService.wrapperOnceExec:input_type -> aiges.Request
	6,  // 8: aiges.WrapperService.testStream:input_type -> aiges.StreamRequest
	5,  // 9: aiges.WrapperService.communicate:input_type -> aiges.Request
	0,  // 10: aiges.WrapperService.wrapperInit:output_type -> aiges.Ret
	3,  // 11: aiges.WrapperService.wrapperOnceExec:output_type -> aiges.Response
	3,  // 12: aiges.WrapperService.testStream:output_type -> aiges.Response
	3,  // 13: aiges.WrapperService.communicate:output_type -> aiges.Response
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_grpc_proto_aiges_inner_proto_init() }
func file_grpc_proto_aiges_inner_proto_init() {
	if File_grpc_proto_aiges_inner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_aiges_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ret); i {
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
		file_grpc_proto_aiges_inner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestData); i {
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
		file_grpc_proto_aiges_inner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseData); i {
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
		file_grpc_proto_aiges_inner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_grpc_proto_aiges_inner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
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
		file_grpc_proto_aiges_inner_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_grpc_proto_aiges_inner_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequest); i {
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
			RawDescriptor: file_grpc_proto_aiges_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_aiges_inner_proto_goTypes,
		DependencyIndexes: file_grpc_proto_aiges_inner_proto_depIdxs,
		MessageInfos:      file_grpc_proto_aiges_inner_proto_msgTypes,
	}.Build()
	File_grpc_proto_aiges_inner_proto = out.File
	file_grpc_proto_aiges_inner_proto_rawDesc = nil
	file_grpc_proto_aiges_inner_proto_goTypes = nil
	file_grpc_proto_aiges_inner_proto_depIdxs = nil
}
