// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.2
// source: proto/reverse.proto

package pb

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

type ReverseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ReverseRequest) Reset() {
	*x = ReverseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reverse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseRequest) ProtoMessage() {}

func (x *ReverseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reverse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseRequest.ProtoReflect.Descriptor instead.
func (*ReverseRequest) Descriptor() ([]byte, []int) {
	return file_proto_reverse_proto_rawDescGZIP(), []int{0}
}

func (x *ReverseRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type ReverseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reversed string `protobuf:"bytes,2,opt,name=reversed,proto3" json:"reversed,omitempty"`
}

func (x *ReverseReply) Reset() {
	*x = ReverseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reverse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseReply) ProtoMessage() {}

func (x *ReverseReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reverse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseReply.ProtoReflect.Descriptor instead.
func (*ReverseReply) Descriptor() ([]byte, []int) {
	return file_proto_reverse_proto_rawDescGZIP(), []int{1}
}

func (x *ReverseReply) GetReversed() string {
	if x != nil {
		return x.Reversed
	}
	return ""
}

var File_proto_reverse_proto protoreflect.FileDescriptor

var file_proto_reverse_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x22, 0x24,
	0x0a, 0x0e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x2a, 0x0a, 0x0c, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x64,
	0x32, 0x53, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x52, 0x65,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_reverse_proto_rawDescOnce sync.Once
	file_proto_reverse_proto_rawDescData = file_proto_reverse_proto_rawDesc
)

func file_proto_reverse_proto_rawDescGZIP() []byte {
	file_proto_reverse_proto_rawDescOnce.Do(func() {
		file_proto_reverse_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_reverse_proto_rawDescData)
	})
	return file_proto_reverse_proto_rawDescData
}

var file_proto_reverse_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_reverse_proto_goTypes = []interface{}{
	(*ReverseRequest)(nil), // 0: reverse.ReverseRequest
	(*ReverseReply)(nil),   // 1: reverse.ReverseReply
}
var file_proto_reverse_proto_depIdxs = []int32{
	0, // 0: reverse.ReverseService.ReverseString:input_type -> reverse.ReverseRequest
	1, // 1: reverse.ReverseService.ReverseString:output_type -> reverse.ReverseReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_reverse_proto_init() }
func file_proto_reverse_proto_init() {
	if File_proto_reverse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_reverse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseRequest); i {
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
		file_proto_reverse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseReply); i {
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
			RawDescriptor: file_proto_reverse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_reverse_proto_goTypes,
		DependencyIndexes: file_proto_reverse_proto_depIdxs,
		MessageInfos:      file_proto_reverse_proto_msgTypes,
	}.Build()
	File_proto_reverse_proto = out.File
	file_proto_reverse_proto_rawDesc = nil
	file_proto_reverse_proto_goTypes = nil
	file_proto_reverse_proto_depIdxs = nil
}
