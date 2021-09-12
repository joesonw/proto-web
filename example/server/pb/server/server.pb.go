// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: server.proto

package pbgo

import (
	_ "github.com/joesonw/proto-web/pbgo/openapi"
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

type Unary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Unary) Reset() {
	*x = Unary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unary) ProtoMessage() {}

func (x *Unary) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unary.ProtoReflect.Descriptor instead.
func (*Unary) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

type Stream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Stream) Reset() {
	*x = Stream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream) ProtoMessage() {}

func (x *Stream) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream.ProtoReflect.Descriptor instead.
func (*Stream) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

type Unary_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Extra   string `protobuf:"bytes,3,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *Unary_Request) Reset() {
	*x = Unary_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unary_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unary_Request) ProtoMessage() {}

func (x *Unary_Request) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unary_Request.ProtoReflect.Descriptor instead.
func (*Unary_Request) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Unary_Request) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Unary_Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Unary_Request) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

type Unary_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	TestHeader int64  `protobuf:"varint,2,opt,name=test_header,json=testHeader,proto3" json:"test_header,omitempty"`
}

func (x *Unary_Response) Reset() {
	*x = Unary_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unary_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unary_Response) ProtoMessage() {}

func (x *Unary_Response) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unary_Response.ProtoReflect.Descriptor instead.
func (*Unary_Response) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Unary_Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Unary_Response) GetTestHeader() int64 {
	if x != nil {
		return x.TestHeader
	}
	return 0
}

type Stream_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Stream_Request) Reset() {
	*x = Stream_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream_Request) ProtoMessage() {}

func (x *Stream_Request) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream_Request.ProtoReflect.Descriptor instead.
func (*Stream_Request) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Stream_Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Stream_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Stream_Response) Reset() {
	*x = Stream_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream_Response) ProtoMessage() {}

func (x *Stream_Response) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream_Response.ProtoReflect.Descriptor instead.
func (*Stream_Response) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Stream_Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x05, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x1a,
	0x66, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0e, 0xba, 0xd5, 0xaa, 0xb1, 0x02, 0x02, 0x69, 0x64,
	0xc8, 0xd5, 0xaa, 0xb1, 0x02, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0b, 0xaa, 0xd5, 0xaa, 0xb1, 0x02, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x51, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a,
	0x0b, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x0a, 0xb2, 0xd5, 0xaa, 0xb1, 0x02, 0x04, 0x74, 0x65, 0x73, 0x74, 0x52, 0x0a,
	0x74, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x53, 0x0a, 0x06, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x1a, 0x23, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0xa3, 0x03, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x6c, 0x0a, 0x05, 0x55,
	0x6e, 0x61, 0x72, 0x79, 0x12, 0x15, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x55, 0x6e,
	0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x34, 0x8a, 0xcf, 0xaa, 0xb1, 0x02, 0x2e, 0x0a, 0x0d, 0x75, 0x6e, 0x61,
	0x72, 0x79, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5a, 0x0f, 0x2f, 0x75, 0x6e, 0x61, 0x72, 0x79,
	0x5f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x3a, 0x69, 0x64, 0x12, 0x5e, 0x0a, 0x0e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x8a, 0xcf,
	0xaa, 0xb1, 0x02, 0x13, 0x8a, 0x01, 0x10, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x5c, 0x0a, 0x0d, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x8a, 0xcf, 0xaa,
	0xb1, 0x02, 0x12, 0x8a, 0x01, 0x0f, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x28, 0x01, 0x12, 0x5c, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x44, 0x75, 0x70, 0x6c, 0x65, 0x78, 0x12, 0x16, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x8a, 0xcf, 0xaa, 0xb1, 0x02, 0x11,
	0x8a, 0x01, 0x0e, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x64, 0x75, 0x70, 0x6c, 0x65,
	0x78, 0x28, 0x01, 0x30, 0x01, 0x1a, 0x0e, 0xea, 0xc8, 0xaa, 0xb1, 0x02, 0x08, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0xa9, 0x02, 0x5a, 0x0e, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x3b, 0x70, 0x62, 0x67, 0x6f, 0xd2, 0xc2, 0xaa, 0xb1, 0x02, 0x6c, 0x0a, 0x20,
	0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x7b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x7d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x12, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x75, 0x72, 0x6c, 0x1a, 0x3c, 0x0a, 0x0b,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x03, 0x64,
	0x65, 0x76, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x0a, 0x04, 0x70, 0x72, 0x6f,
	0x64, 0x12, 0x03, 0x64, 0x65, 0x76, 0x1a, 0x12, 0x63, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x20, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0xda, 0xc2, 0xaa, 0xb1, 0x02, 0x16,
	0x0a, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0xe2, 0xc2, 0xaa, 0xb1, 0x02, 0x0e, 0x0a, 0x0c, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0xca, 0xc2, 0xaa, 0xb1, 0x02, 0x71,
	0x0a, 0x0f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1f, 0x74, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x6a, 0x75, 0x73, 0x74, 0x20,
	0x61, 0x6e, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x13, 0x74, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x1a, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x1a, 0x11, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x40, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x32, 0x05, 0x0a, 0x03, 0x4d, 0x49, 0x54, 0x3a, 0x05, 0x30, 0x2e, 0x30, 0x2e,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_server_proto_goTypes = []interface{}{
	(*Unary)(nil),           // 0: errors.Unary
	(*Stream)(nil),          // 1: errors.Stream
	(*Unary_Request)(nil),   // 2: errors.Unary.Request
	(*Unary_Response)(nil),  // 3: errors.Unary.Response
	(*Stream_Request)(nil),  // 4: errors.Stream.Request
	(*Stream_Response)(nil), // 5: errors.Stream.Response
}
var file_server_proto_depIdxs = []int32{
	2, // 0: errors.Example.Unary:input_type -> errors.Unary.Request
	4, // 1: errors.Example.StreamResponse:input_type -> errors.Stream.Request
	4, // 2: errors.Example.StreamRequest:input_type -> errors.Stream.Request
	4, // 3: errors.Example.StreamDuplex:input_type -> errors.Stream.Request
	3, // 4: errors.Example.Unary:output_type -> errors.Unary.Response
	5, // 5: errors.Example.StreamResponse:output_type -> errors.Stream.Response
	5, // 6: errors.Example.StreamRequest:output_type -> errors.Stream.Response
	5, // 7: errors.Example.StreamDuplex:output_type -> errors.Stream.Response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unary); i {
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
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stream); i {
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
		file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unary_Request); i {
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
		file_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unary_Response); i {
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
		file_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stream_Request); i {
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
		file_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stream_Response); i {
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
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
