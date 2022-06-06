// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: note_v1.proto

package note_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *AddNoteRequest) Reset() {
	*x = AddNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNoteRequest) ProtoMessage() {}

func (x *AddNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNoteRequest.ProtoReflect.Descriptor instead.
func (*AddNoteRequest) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{0}
}

func (x *AddNoteRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddNoteRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type AddNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *AddNoteResponse_Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AddNoteResponse) Reset() {
	*x = AddNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNoteResponse) ProtoMessage() {}

func (x *AddNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNoteResponse.ProtoReflect.Descriptor instead.
func (*AddNoteResponse) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{1}
}

func (x *AddNoteResponse) GetResult() *AddNoteResponse_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type RemoveNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveNoteRequest) Reset() {
	*x = RemoveNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveNoteRequest) ProtoMessage() {}

func (x *RemoveNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveNoteRequest.ProtoReflect.Descriptor instead.
func (*RemoveNoteRequest) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveNoteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Removed int64 `protobuf:"varint,1,opt,name=removed,proto3" json:"removed,omitempty"`
}

func (x *RemoveNoteResponse) Reset() {
	*x = RemoveNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveNoteResponse) ProtoMessage() {}

func (x *RemoveNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveNoteResponse.ProtoReflect.Descriptor instead.
func (*RemoveNoteResponse) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveNoteResponse) GetRemoved() int64 {
	if x != nil {
		return x.Removed
	}
	return 0
}

type MultiAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notes []*MultiAddRequest_Notes `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
}

func (x *MultiAddRequest) Reset() {
	*x = MultiAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiAddRequest) ProtoMessage() {}

func (x *MultiAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiAddRequest.ProtoReflect.Descriptor instead.
func (*MultiAddRequest) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{4}
}

func (x *MultiAddRequest) GetNotes() []*MultiAddRequest_Notes {
	if x != nil {
		return x.Notes
	}
	return nil
}

type MultiAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *MultiAddResponse_Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MultiAddResponse) Reset() {
	*x = MultiAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiAddResponse) ProtoMessage() {}

func (x *MultiAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiAddResponse.ProtoReflect.Descriptor instead.
func (*MultiAddResponse) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{5}
}

func (x *MultiAddResponse) GetResult() *MultiAddResponse_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type GetNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNoteRequest) Reset() {
	*x = GetNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteRequest) ProtoMessage() {}

func (x *GetNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoteRequest.ProtoReflect.Descriptor instead.
func (*GetNoteRequest) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{6}
}

func (x *GetNoteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GetNoteResponse) Reset() {
	*x = GetNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteResponse) ProtoMessage() {}

func (x *GetNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoteResponse.ProtoReflect.Descriptor instead.
func (*GetNoteResponse) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{7}
}

func (x *GetNoteResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetNoteResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetNoteResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type GetListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*GetListResponse_Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *GetListResponse) Reset() {
	*x = GetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResponse.ProtoReflect.Descriptor instead.
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{8}
}

func (x *GetListResponse) GetResults() []*GetListResponse_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type UpdateNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UpdateNoteRequest) Reset() {
	*x = UpdateNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoteRequest) ProtoMessage() {}

func (x *UpdateNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateNoteRequest) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateNoteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateNoteRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateNoteRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type AddNoteResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddNoteResponse_Result) Reset() {
	*x = AddNoteResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNoteResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNoteResponse_Result) ProtoMessage() {}

func (x *AddNoteResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNoteResponse_Result.ProtoReflect.Descriptor instead.
func (*AddNoteResponse_Result) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AddNoteResponse_Result) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MultiAddRequest_Notes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *MultiAddRequest_Notes) Reset() {
	*x = MultiAddRequest_Notes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiAddRequest_Notes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiAddRequest_Notes) ProtoMessage() {}

func (x *MultiAddRequest_Notes) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiAddRequest_Notes.ProtoReflect.Descriptor instead.
func (*MultiAddRequest_Notes) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{4, 0}
}

func (x *MultiAddRequest_Notes) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MultiAddRequest_Notes) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type MultiAddResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *MultiAddResponse_Result) Reset() {
	*x = MultiAddResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiAddResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiAddResponse_Result) ProtoMessage() {}

func (x *MultiAddResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiAddResponse_Result.ProtoReflect.Descriptor instead.
func (*MultiAddResponse_Result) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{5, 0}
}

func (x *MultiAddResponse_Result) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetListResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GetListResponse_Result) Reset() {
	*x = GetListResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse_Result) ProtoMessage() {}

func (x *GetListResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_note_v1_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResponse_Result.ProtoReflect.Descriptor instead.
func (*GetListResponse_Result) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{8, 0}
}

func (x *GetListResponse_Result) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetListResponse_Result) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_note_v1_proto protoreflect.FileDescriptor

var file_note_v1_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x14, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01,
	0x18, 0xe8, 0x07, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x5c, 0x0a, 0x0f,
	0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x1a, 0x18, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x11, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x6f,
	0x74, 0x65, 0x73, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x1a, 0x4e, 0x0a, 0x05, 0x4e, 0x6f,
	0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x14, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xe8,
	0x07, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x64, 0x0a, 0x10, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x1a, 0x1e, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x7e,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x1a, 0x38, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x73,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x01, 0x18, 0x14, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x24, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xe8, 0x07, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x32, 0xe8, 0x03, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31, 0x12, 0x45,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x41, 0x64, 0x64, 0x4e,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x41, 0x64, 0x64,
	0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x64, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x53, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4e, 0x0a, 0x08, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x41, 0x64, 0x64, 0x12, 0x10, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x2d, 0x61, 0x64, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x47, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x12, 0x11, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x53, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x12, 0x16, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x61,
	0x6c, 0x6c, 0x2d, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x1a, 0x0f, 0x2f, 0x6e, 0x6f, 0x74,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x34,
	0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6b,
	0x69, 0x74, 0x61, 0x64, 0x73, 0x39, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x6f, 0x74, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_note_v1_proto_rawDescOnce sync.Once
	file_note_v1_proto_rawDescData = file_note_v1_proto_rawDesc
)

func file_note_v1_proto_rawDescGZIP() []byte {
	file_note_v1_proto_rawDescOnce.Do(func() {
		file_note_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_note_v1_proto_rawDescData)
	})
	return file_note_v1_proto_rawDescData
}

var file_note_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_note_v1_proto_goTypes = []interface{}{
	(*AddNoteRequest)(nil),          // 0: AddNoteRequest
	(*AddNoteResponse)(nil),         // 1: AddNoteResponse
	(*RemoveNoteRequest)(nil),       // 2: RemoveNoteRequest
	(*RemoveNoteResponse)(nil),      // 3: RemoveNoteResponse
	(*MultiAddRequest)(nil),         // 4: MultiAddRequest
	(*MultiAddResponse)(nil),        // 5: MultiAddResponse
	(*GetNoteRequest)(nil),          // 6: GetNoteRequest
	(*GetNoteResponse)(nil),         // 7: GetNoteResponse
	(*GetListResponse)(nil),         // 8: GetListResponse
	(*UpdateNoteRequest)(nil),       // 9: UpdateNoteRequest
	(*AddNoteResponse_Result)(nil),  // 10: AddNoteResponse.Result
	(*MultiAddRequest_Notes)(nil),   // 11: MultiAddRequest.Notes
	(*MultiAddResponse_Result)(nil), // 12: MultiAddResponse.Result
	(*GetListResponse_Result)(nil),  // 13: GetListResponse.Result
	(*emptypb.Empty)(nil),           // 14: google.protobuf.Empty
}
var file_note_v1_proto_depIdxs = []int32{
	10, // 0: AddNoteResponse.result:type_name -> AddNoteResponse.Result
	11, // 1: MultiAddRequest.notes:type_name -> MultiAddRequest.Notes
	12, // 2: MultiAddResponse.result:type_name -> MultiAddResponse.Result
	13, // 3: GetListResponse.results:type_name -> GetListResponse.Result
	0,  // 4: NoteV1.AddNote:input_type -> AddNoteRequest
	2,  // 5: NoteV1.RemoveNote:input_type -> RemoveNoteRequest
	4,  // 6: NoteV1.MultiAdd:input_type -> MultiAddRequest
	6,  // 7: NoteV1.GetNote:input_type -> GetNoteRequest
	14, // 8: NoteV1.GetList:input_type -> google.protobuf.Empty
	9,  // 9: NoteV1.UpdateNote:input_type -> UpdateNoteRequest
	1,  // 10: NoteV1.AddNote:output_type -> AddNoteResponse
	3,  // 11: NoteV1.RemoveNote:output_type -> RemoveNoteResponse
	5,  // 12: NoteV1.MultiAdd:output_type -> MultiAddResponse
	7,  // 13: NoteV1.GetNote:output_type -> GetNoteResponse
	8,  // 14: NoteV1.GetList:output_type -> GetListResponse
	14, // 15: NoteV1.UpdateNote:output_type -> google.protobuf.Empty
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_note_v1_proto_init() }
func file_note_v1_proto_init() {
	if File_note_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_note_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNoteRequest); i {
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
		file_note_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNoteResponse); i {
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
		file_note_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveNoteRequest); i {
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
		file_note_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveNoteResponse); i {
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
		file_note_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiAddRequest); i {
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
		file_note_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiAddResponse); i {
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
		file_note_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoteRequest); i {
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
		file_note_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoteResponse); i {
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
		file_note_v1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListResponse); i {
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
		file_note_v1_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNoteRequest); i {
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
		file_note_v1_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNoteResponse_Result); i {
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
		file_note_v1_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiAddRequest_Notes); i {
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
		file_note_v1_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiAddResponse_Result); i {
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
		file_note_v1_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListResponse_Result); i {
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
			RawDescriptor: file_note_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_note_v1_proto_goTypes,
		DependencyIndexes: file_note_v1_proto_depIdxs,
		MessageInfos:      file_note_v1_proto_msgTypes,
	}.Build()
	File_note_v1_proto = out.File
	file_note_v1_proto_rawDesc = nil
	file_note_v1_proto_goTypes = nil
	file_note_v1_proto_depIdxs = nil
}
