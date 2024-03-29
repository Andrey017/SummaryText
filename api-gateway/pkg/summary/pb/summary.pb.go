// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: pkg/summary/pb/summary.proto

package pb

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

type Summary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NameFile         string `protobuf:"bytes,2,opt,name=nameFile,proto3" json:"nameFile,omitempty"`
	Type             string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Note             string `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
	ArticleId        int64  `protobuf:"varint,5,opt,name=articleId,proto3" json:"articleId,omitempty"`
	CountWord        int64  `protobuf:"varint,6,opt,name=countWord,proto3" json:"countWord,omitempty"`
	CountSentense    int64  `protobuf:"varint,7,opt,name=countSentense,proto3" json:"countSentense,omitempty"`
	MinCountSentense int64  `protobuf:"varint,8,opt,name=minCountSentense,proto3" json:"minCountSentense,omitempty"`
	MaxCountSentense int64  `protobuf:"varint,9,opt,name=maxCountSentense,proto3" json:"maxCountSentense,omitempty"`
}

func (x *Summary) Reset() {
	*x = Summary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_summary_pb_summary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary) ProtoMessage() {}

func (x *Summary) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_summary_pb_summary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary.ProtoReflect.Descriptor instead.
func (*Summary) Descriptor() ([]byte, []int) {
	return file_pkg_summary_pb_summary_proto_rawDescGZIP(), []int{0}
}

func (x *Summary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Summary) GetNameFile() string {
	if x != nil {
		return x.NameFile
	}
	return ""
}

func (x *Summary) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Summary) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Summary) GetArticleId() int64 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

func (x *Summary) GetCountWord() int64 {
	if x != nil {
		return x.CountWord
	}
	return 0
}

func (x *Summary) GetCountSentense() int64 {
	if x != nil {
		return x.CountSentense
	}
	return 0
}

func (x *Summary) GetMinCountSentense() int64 {
	if x != nil {
		return x.MinCountSentense
	}
	return 0
}

func (x *Summary) GetMaxCountSentense() int64 {
	if x != nil {
		return x.MaxCountSentense
	}
	return 0
}

type SummaryTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Summary `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SummaryTextRequest) Reset() {
	*x = SummaryTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_summary_pb_summary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryTextRequest) ProtoMessage() {}

func (x *SummaryTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_summary_pb_summary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryTextRequest.ProtoReflect.Descriptor instead.
func (*SummaryTextRequest) Descriptor() ([]byte, []int) {
	return file_pkg_summary_pb_summary_proto_rawDescGZIP(), []int{1}
}

func (x *SummaryTextRequest) GetInfo() *Summary {
	if x != nil {
		return x.Info
	}
	return nil
}

type SummaryTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SummaryTextResponse) Reset() {
	*x = SummaryTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_summary_pb_summary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryTextResponse) ProtoMessage() {}

func (x *SummaryTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_summary_pb_summary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryTextResponse.ProtoReflect.Descriptor instead.
func (*SummaryTextResponse) Descriptor() ([]byte, []int) {
	return file_pkg_summary_pb_summary_proto_rawDescGZIP(), []int{2}
}

func (x *SummaryTextResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SummaryLSARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Summary `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SummaryLSARequest) Reset() {
	*x = SummaryLSARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_summary_pb_summary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryLSARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryLSARequest) ProtoMessage() {}

func (x *SummaryLSARequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_summary_pb_summary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryLSARequest.ProtoReflect.Descriptor instead.
func (*SummaryLSARequest) Descriptor() ([]byte, []int) {
	return file_pkg_summary_pb_summary_proto_rawDescGZIP(), []int{3}
}

func (x *SummaryLSARequest) GetInfo() *Summary {
	if x != nil {
		return x.Info
	}
	return nil
}

type SummaryLSAResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SummaryLSAResponse) Reset() {
	*x = SummaryLSAResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_summary_pb_summary_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryLSAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryLSAResponse) ProtoMessage() {}

func (x *SummaryLSAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_summary_pb_summary_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryLSAResponse.ProtoReflect.Descriptor instead.
func (*SummaryLSAResponse) Descriptor() ([]byte, []int) {
	return file_pkg_summary_pb_summary_proto_rawDescGZIP(), []int{4}
}

func (x *SummaryLSAResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SummaryT5Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Summary `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SummaryT5Request) Reset() {
	*x = SummaryT5Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_summary_pb_summary_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryT5Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryT5Request) ProtoMessage() {}

func (x *SummaryT5Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_summary_pb_summary_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryT5Request.ProtoReflect.Descriptor instead.
func (*SummaryT5Request) Descriptor() ([]byte, []int) {
	return file_pkg_summary_pb_summary_proto_rawDescGZIP(), []int{5}
}

func (x *SummaryT5Request) GetInfo() *Summary {
	if x != nil {
		return x.Info
	}
	return nil
}

type SummaryT5Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SummaryT5Response) Reset() {
	*x = SummaryT5Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_summary_pb_summary_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryT5Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryT5Response) ProtoMessage() {}

func (x *SummaryT5Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_summary_pb_summary_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryT5Response.ProtoReflect.Descriptor instead.
func (*SummaryT5Response) Descriptor() ([]byte, []int) {
	return file_pkg_summary_pb_summary_proto_rawDescGZIP(), []int{6}
}

func (x *SummaryT5Response) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_pkg_summary_pb_summary_proto protoreflect.FileDescriptor

var file_pkg_summary_pb_summary_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2f, 0x70, 0x62,
	0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x9b, 0x02, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a,
	0x10, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x73,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6d, 0x61, 0x78,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x6e,
	0x74, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x0a, 0x12, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x25, 0x0a, 0x13, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x54, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x11, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x4c, 0x53, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x22, 0x24, 0x0a, 0x12, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4c, 0x53,
	0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x10, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x54, 0x35, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x22, 0x23, 0x0a, 0x11, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x54, 0x35,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xf3, 0x01, 0x0a, 0x12, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4e, 0x0a, 0x0f, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x54, 0x65, 0x78, 0x74, 0x52, 0x61,
	0x6e, 0x6b, 0x12, 0x1b, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x47, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4c, 0x53, 0x41, 0x12, 0x1a, 0x2e,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4c,
	0x53, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4c, 0x53, 0x41, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x09, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x54, 0x35, 0x12, 0x19, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x54, 0x35, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x54, 0x35, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12,
	0x5a, 0x10, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_summary_pb_summary_proto_rawDescOnce sync.Once
	file_pkg_summary_pb_summary_proto_rawDescData = file_pkg_summary_pb_summary_proto_rawDesc
)

func file_pkg_summary_pb_summary_proto_rawDescGZIP() []byte {
	file_pkg_summary_pb_summary_proto_rawDescOnce.Do(func() {
		file_pkg_summary_pb_summary_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_summary_pb_summary_proto_rawDescData)
	})
	return file_pkg_summary_pb_summary_proto_rawDescData
}

var file_pkg_summary_pb_summary_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_summary_pb_summary_proto_goTypes = []interface{}{
	(*Summary)(nil),             // 0: summary.Summary
	(*SummaryTextRequest)(nil),  // 1: summary.SummaryTextRequest
	(*SummaryTextResponse)(nil), // 2: summary.SummaryTextResponse
	(*SummaryLSARequest)(nil),   // 3: summary.SummaryLSARequest
	(*SummaryLSAResponse)(nil),  // 4: summary.SummaryLSAResponse
	(*SummaryT5Request)(nil),    // 5: summary.SummaryT5Request
	(*SummaryT5Response)(nil),   // 6: summary.SummaryT5Response
}
var file_pkg_summary_pb_summary_proto_depIdxs = []int32{
	0, // 0: summary.SummaryTextRequest.info:type_name -> summary.Summary
	0, // 1: summary.SummaryLSARequest.info:type_name -> summary.Summary
	0, // 2: summary.SummaryT5Request.info:type_name -> summary.Summary
	1, // 3: summary.SummaryTextService.SummaryTextRank:input_type -> summary.SummaryTextRequest
	3, // 4: summary.SummaryTextService.SummaryLSA:input_type -> summary.SummaryLSARequest
	5, // 5: summary.SummaryTextService.SummaryT5:input_type -> summary.SummaryT5Request
	2, // 6: summary.SummaryTextService.SummaryTextRank:output_type -> summary.SummaryTextResponse
	4, // 7: summary.SummaryTextService.SummaryLSA:output_type -> summary.SummaryLSAResponse
	6, // 8: summary.SummaryTextService.SummaryT5:output_type -> summary.SummaryT5Response
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_summary_pb_summary_proto_init() }
func file_pkg_summary_pb_summary_proto_init() {
	if File_pkg_summary_pb_summary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_summary_pb_summary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Summary); i {
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
		file_pkg_summary_pb_summary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryTextRequest); i {
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
		file_pkg_summary_pb_summary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryTextResponse); i {
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
		file_pkg_summary_pb_summary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryLSARequest); i {
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
		file_pkg_summary_pb_summary_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryLSAResponse); i {
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
		file_pkg_summary_pb_summary_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryT5Request); i {
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
		file_pkg_summary_pb_summary_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryT5Response); i {
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
			RawDescriptor: file_pkg_summary_pb_summary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_summary_pb_summary_proto_goTypes,
		DependencyIndexes: file_pkg_summary_pb_summary_proto_depIdxs,
		MessageInfos:      file_pkg_summary_pb_summary_proto_msgTypes,
	}.Build()
	File_pkg_summary_pb_summary_proto = out.File
	file_pkg_summary_pb_summary_proto_rawDesc = nil
	file_pkg_summary_pb_summary_proto_goTypes = nil
	file_pkg_summary_pb_summary_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SummaryTextServiceClient is the client API for SummaryTextService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SummaryTextServiceClient interface {
	SummaryTextRank(ctx context.Context, in *SummaryTextRequest, opts ...grpc.CallOption) (*SummaryTextResponse, error)
	SummaryLSA(ctx context.Context, in *SummaryLSARequest, opts ...grpc.CallOption) (*SummaryLSAResponse, error)
	SummaryT5(ctx context.Context, in *SummaryT5Request, opts ...grpc.CallOption) (*SummaryT5Response, error)
}

type summaryTextServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSummaryTextServiceClient(cc grpc.ClientConnInterface) SummaryTextServiceClient {
	return &summaryTextServiceClient{cc}
}

func (c *summaryTextServiceClient) SummaryTextRank(ctx context.Context, in *SummaryTextRequest, opts ...grpc.CallOption) (*SummaryTextResponse, error) {
	out := new(SummaryTextResponse)
	err := c.cc.Invoke(ctx, "/summary.SummaryTextService/SummaryTextRank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *summaryTextServiceClient) SummaryLSA(ctx context.Context, in *SummaryLSARequest, opts ...grpc.CallOption) (*SummaryLSAResponse, error) {
	out := new(SummaryLSAResponse)
	err := c.cc.Invoke(ctx, "/summary.SummaryTextService/SummaryLSA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *summaryTextServiceClient) SummaryT5(ctx context.Context, in *SummaryT5Request, opts ...grpc.CallOption) (*SummaryT5Response, error) {
	out := new(SummaryT5Response)
	err := c.cc.Invoke(ctx, "/summary.SummaryTextService/SummaryT5", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SummaryTextServiceServer is the server API for SummaryTextService service.
type SummaryTextServiceServer interface {
	SummaryTextRank(context.Context, *SummaryTextRequest) (*SummaryTextResponse, error)
	SummaryLSA(context.Context, *SummaryLSARequest) (*SummaryLSAResponse, error)
	SummaryT5(context.Context, *SummaryT5Request) (*SummaryT5Response, error)
}

// UnimplementedSummaryTextServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSummaryTextServiceServer struct {
}

func (*UnimplementedSummaryTextServiceServer) SummaryTextRank(context.Context, *SummaryTextRequest) (*SummaryTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SummaryTextRank not implemented")
}
func (*UnimplementedSummaryTextServiceServer) SummaryLSA(context.Context, *SummaryLSARequest) (*SummaryLSAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SummaryLSA not implemented")
}
func (*UnimplementedSummaryTextServiceServer) SummaryT5(context.Context, *SummaryT5Request) (*SummaryT5Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SummaryT5 not implemented")
}

func RegisterSummaryTextServiceServer(s *grpc.Server, srv SummaryTextServiceServer) {
	s.RegisterService(&_SummaryTextService_serviceDesc, srv)
}

func _SummaryTextService_SummaryTextRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummaryTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummaryTextServiceServer).SummaryTextRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/summary.SummaryTextService/SummaryTextRank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummaryTextServiceServer).SummaryTextRank(ctx, req.(*SummaryTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SummaryTextService_SummaryLSA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummaryLSARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummaryTextServiceServer).SummaryLSA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/summary.SummaryTextService/SummaryLSA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummaryTextServiceServer).SummaryLSA(ctx, req.(*SummaryLSARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SummaryTextService_SummaryT5_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummaryT5Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummaryTextServiceServer).SummaryT5(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/summary.SummaryTextService/SummaryT5",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummaryTextServiceServer).SummaryT5(ctx, req.(*SummaryT5Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _SummaryTextService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "summary.SummaryTextService",
	HandlerType: (*SummaryTextServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SummaryTextRank",
			Handler:    _SummaryTextService_SummaryTextRank_Handler,
		},
		{
			MethodName: "SummaryLSA",
			Handler:    _SummaryTextService_SummaryLSA_Handler,
		},
		{
			MethodName: "SummaryT5",
			Handler:    _SummaryTextService_SummaryT5_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/summary/pb/summary.proto",
}
