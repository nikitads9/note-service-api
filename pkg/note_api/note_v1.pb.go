// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: note_v1.proto

package note_service_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Notes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Notes) Reset() {
	*x = Notes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notes) ProtoMessage() {}

func (x *Notes) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Notes.ProtoReflect.Descriptor instead.
func (*Notes) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{0}
}

func (x *Notes) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Notes) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type NoteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Note      *Notes                 `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *NoteInfo) Reset() {
	*x = NoteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteInfo) ProtoMessage() {}

func (x *NoteInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use NoteInfo.ProtoReflect.Descriptor instead.
func (*NoteInfo) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{1}
}

func (x *NoteInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NoteInfo) GetNote() *Notes {
	if x != nil {
		return x.Note
	}
	return nil
}

func (x *NoteInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *NoteInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type AddNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note *Notes `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *AddNoteRequest) Reset() {
	*x = AddNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNoteRequest) ProtoMessage() {}

func (x *AddNoteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AddNoteRequest.ProtoReflect.Descriptor instead.
func (*AddNoteRequest) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{2}
}

func (x *AddNoteRequest) GetNote() *Notes {
	if x != nil {
		return x.Note
	}
	return nil
}

type AddNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddNoteResponse) Reset() {
	*x = AddNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNoteResponse) ProtoMessage() {}

func (x *AddNoteResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AddNoteResponse.ProtoReflect.Descriptor instead.
func (*AddNoteResponse) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{3}
}

func (x *AddNoteResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
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
		mi := &file_note_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveNoteRequest) ProtoMessage() {}

func (x *RemoveNoteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RemoveNoteRequest.ProtoReflect.Descriptor instead.
func (*RemoveNoteRequest) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{4}
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
		mi := &file_note_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveNoteResponse) ProtoMessage() {}

func (x *RemoveNoteResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RemoveNoteResponse.ProtoReflect.Descriptor instead.
func (*RemoveNoteResponse) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{5}
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

	Notes []*Notes `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
}

func (x *MultiAddRequest) Reset() {
	*x = MultiAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiAddRequest) ProtoMessage() {}

func (x *MultiAddRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MultiAddRequest.ProtoReflect.Descriptor instead.
func (*MultiAddRequest) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{6}
}

func (x *MultiAddRequest) GetNotes() []*Notes {
	if x != nil {
		return x.Notes
	}
	return nil
}

type MultiAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *MultiAddResponse) Reset() {
	*x = MultiAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiAddResponse) ProtoMessage() {}

func (x *MultiAddResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MultiAddResponse.ProtoReflect.Descriptor instead.
func (*MultiAddResponse) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{7}
}

func (x *MultiAddResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
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
		mi := &file_note_v1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteRequest) ProtoMessage() {}

func (x *GetNoteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetNoteRequest.ProtoReflect.Descriptor instead.
func (*GetNoteRequest) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{8}
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

	NoteInfo *NoteInfo `protobuf:"bytes,1,opt,name=noteInfo,proto3" json:"noteInfo,omitempty"`
}

func (x *GetNoteResponse) Reset() {
	*x = GetNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteResponse) ProtoMessage() {}

func (x *GetNoteResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetNoteResponse.ProtoReflect.Descriptor instead.
func (*GetNoteResponse) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{9}
}

func (x *GetNoteResponse) GetNoteInfo() *NoteInfo {
	if x != nil {
		return x.NoteInfo
	}
	return nil
}

type GetListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteInfo []*NoteInfo `protobuf:"bytes,1,rep,name=noteInfo,proto3" json:"noteInfo,omitempty"`
}

func (x *GetListResponse) Reset() {
	*x = GetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetListResponse.ProtoReflect.Descriptor instead.
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{10}
}

func (x *GetListResponse) GetNoteInfo() []*NoteInfo {
	if x != nil {
		return x.NoteInfo
	}
	return nil
}

type UpdateNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UpdateNoteRequest) Reset() {
	*x = UpdateNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoteRequest) ProtoMessage() {}

func (x *UpdateNoteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateNoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateNoteRequest) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateNoteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateNoteRequest) GetTitle() *wrapperspb.StringValue {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *UpdateNoteRequest) GetContent() *wrapperspb.StringValue {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_note_v1_proto protoreflect.FileDescriptor

var file_note_v1_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x05, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12,
	0x1f, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x14, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x24, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xe8, 0x07, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xbb, 0x01, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x3d, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x64, 0x22, 0x40, 0x0a, 0x0f, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52,
	0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x10, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6e, 0x6f,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x49, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x6e, 0x6f, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x6f,
	0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e,
	0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x3d, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x00, 0x18, 0x14, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x42, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x00, 0x18, 0xe8, 0x07, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x32, 0x92, 0x05, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31, 0x12, 0x67,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64,
	0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x6f,
	0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x64, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x75, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x70,
	0x0a, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x64, 0x64, 0x12, 0x21, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x6e, 0x6f, 0x74, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2d, 0x61, 0x64, 0x64, 0x3a, 0x01, 0x2a,
	0x12, 0x69, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x6e, 0x6f,
	0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x64, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21,
	0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x6e, 0x6f, 0x74, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x61, 0x6c, 0x6c, 0x2d, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x12, 0x65, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12,
	0x23, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x1a, 0x0f, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6b, 0x69, 0x74, 0x61, 0x64, 0x73, 0x39,
	0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x3b, 0x6e,
	0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_note_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_note_v1_proto_goTypes = []interface{}{
	(*Notes)(nil),                  // 0: note.service.api.Notes
	(*NoteInfo)(nil),               // 1: note.service.api.NoteInfo
	(*AddNoteRequest)(nil),         // 2: note.service.api.AddNoteRequest
	(*AddNoteResponse)(nil),        // 3: note.service.api.AddNoteResponse
	(*RemoveNoteRequest)(nil),      // 4: note.service.api.RemoveNoteRequest
	(*RemoveNoteResponse)(nil),     // 5: note.service.api.RemoveNoteResponse
	(*MultiAddRequest)(nil),        // 6: note.service.api.MultiAddRequest
	(*MultiAddResponse)(nil),       // 7: note.service.api.MultiAddResponse
	(*GetNoteRequest)(nil),         // 8: note.service.api.GetNoteRequest
	(*GetNoteResponse)(nil),        // 9: note.service.api.GetNoteResponse
	(*GetListResponse)(nil),        // 10: note.service.api.GetListResponse
	(*UpdateNoteRequest)(nil),      // 11: note.service.api.UpdateNoteRequest
	(*timestamppb.Timestamp)(nil),  // 12: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 13: google.protobuf.StringValue
	(*emptypb.Empty)(nil),          // 14: google.protobuf.Empty
}
var file_note_v1_proto_depIdxs = []int32{
	0,  // 0: note.service.api.NoteInfo.note:type_name -> note.service.api.Notes
	12, // 1: note.service.api.NoteInfo.createdAt:type_name -> google.protobuf.Timestamp
	12, // 2: note.service.api.NoteInfo.updatedAt:type_name -> google.protobuf.Timestamp
	0,  // 3: note.service.api.AddNoteRequest.note:type_name -> note.service.api.Notes
	0,  // 4: note.service.api.MultiAddRequest.notes:type_name -> note.service.api.Notes
	1,  // 5: note.service.api.GetNoteResponse.noteInfo:type_name -> note.service.api.NoteInfo
	1,  // 6: note.service.api.GetListResponse.noteInfo:type_name -> note.service.api.NoteInfo
	13, // 7: note.service.api.UpdateNoteRequest.title:type_name -> google.protobuf.StringValue
	13, // 8: note.service.api.UpdateNoteRequest.content:type_name -> google.protobuf.StringValue
	2,  // 9: note.service.api.NoteV1.AddNote:input_type -> note.service.api.AddNoteRequest
	4,  // 10: note.service.api.NoteV1.RemoveNote:input_type -> note.service.api.RemoveNoteRequest
	6,  // 11: note.service.api.NoteV1.MultiAdd:input_type -> note.service.api.MultiAddRequest
	8,  // 12: note.service.api.NoteV1.GetNote:input_type -> note.service.api.GetNoteRequest
	14, // 13: note.service.api.NoteV1.GetList:input_type -> google.protobuf.Empty
	11, // 14: note.service.api.NoteV1.UpdateNote:input_type -> note.service.api.UpdateNoteRequest
	3,  // 15: note.service.api.NoteV1.AddNote:output_type -> note.service.api.AddNoteResponse
	5,  // 16: note.service.api.NoteV1.RemoveNote:output_type -> note.service.api.RemoveNoteResponse
	7,  // 17: note.service.api.NoteV1.MultiAdd:output_type -> note.service.api.MultiAddResponse
	9,  // 18: note.service.api.NoteV1.GetNote:output_type -> note.service.api.GetNoteResponse
	10, // 19: note.service.api.NoteV1.GetList:output_type -> note.service.api.GetListResponse
	14, // 20: note.service.api.NoteV1.UpdateNote:output_type -> google.protobuf.Empty
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_note_v1_proto_init() }
func file_note_v1_proto_init() {
	if File_note_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_note_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notes); i {
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
			switch v := v.(*NoteInfo); i {
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
		file_note_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_note_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_note_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_note_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_note_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_note_v1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_note_v1_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
		file_note_v1_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
		file_note_v1_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_note_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
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
