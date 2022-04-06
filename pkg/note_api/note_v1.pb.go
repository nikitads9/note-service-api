// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: note_v1.proto

package note_api

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_note_v1_proto_rawDescGZIP(), []int{3}
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
		mi := &file_note_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNoteResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNoteResponse_Result) ProtoMessage() {}

func (x *AddNoteResponse_Result) ProtoReflect() protoreflect.Message {
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

var File_note_v1_proto protoreflect.FileDescriptor

var file_note_v1_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x40, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x5c, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x18, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x23, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x60, 0x0a,
	0x06, 0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31, 0x12, 0x2c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4e, 0x6f,
	0x74, 0x65, 0x12, 0x0f, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69,
	0x6b, 0x69, 0x74, 0x61, 0x64, 0x73, 0x39, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x6f, 0x74,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_note_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_note_v1_proto_goTypes = []interface{}{
	(*AddNoteRequest)(nil),         // 0: AddNoteRequest
	(*AddNoteResponse)(nil),        // 1: AddNoteResponse
	(*RemoveNoteRequest)(nil),      // 2: RemoveNoteRequest
	(*Empty)(nil),                  // 3: Empty
	(*AddNoteResponse_Result)(nil), // 4: AddNoteResponse.Result
}
var file_note_v1_proto_depIdxs = []int32{
	4, // 0: AddNoteResponse.result:type_name -> AddNoteResponse.Result
	0, // 1: NoteV1.AddNote:input_type -> AddNoteRequest
	2, // 2: NoteV1.RemoveNote:input_type -> RemoveNoteRequest
	1, // 3: NoteV1.AddNote:output_type -> AddNoteResponse
	3, // 4: NoteV1.RemoveNote:output_type -> Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
			switch v := v.(*Empty); i {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_note_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
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
