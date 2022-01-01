// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: protobuf/mock.proto

package protobuf

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

type IncrementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IncrementRequest) Reset() {
	*x = IncrementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementRequest) ProtoMessage() {}

func (x *IncrementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementRequest.ProtoReflect.Descriptor instead.
func (*IncrementRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_mock_proto_rawDescGZIP(), []int{0}
}

type IncrementReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewValue int32 `protobuf:"varint,1,opt,name=newValue,proto3" json:"newValue,omitempty"`
}

func (x *IncrementReply) Reset() {
	*x = IncrementReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrementReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementReply) ProtoMessage() {}

func (x *IncrementReply) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementReply.ProtoReflect.Descriptor instead.
func (*IncrementReply) Descriptor() ([]byte, []int) {
	return file_protobuf_mock_proto_rawDescGZIP(), []int{1}
}

func (x *IncrementReply) GetNewValue() int32 {
	if x != nil {
		return x.NewValue
	}
	return 0
}

var File_protobuf_mock_proto protoreflect.FileDescriptor

var file_protobuf_mock_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x6f, 0x63, 0x6b, 0x22, 0x12, 0x0a, 0x10, 0x49,
	0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x2c, 0x0a, 0x0e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x43, 0x0a,
	0x04, 0x4d, 0x6f, 0x63, 0x6b, 0x12, 0x3b, 0x0a, 0x09, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x16, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x6f, 0x63,
	0x6b, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x72, 0x65, 0x6a, 0x61, 0x62, 0x76,
	0x2f, 0x44, 0x69, 0x73, 0x79, 0x73, 0x4d, 0x6f, 0x63, 0x6b, 0x45, 0x78, 0x61, 0x6d, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_mock_proto_rawDescOnce sync.Once
	file_protobuf_mock_proto_rawDescData = file_protobuf_mock_proto_rawDesc
)

func file_protobuf_mock_proto_rawDescGZIP() []byte {
	file_protobuf_mock_proto_rawDescOnce.Do(func() {
		file_protobuf_mock_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_mock_proto_rawDescData)
	})
	return file_protobuf_mock_proto_rawDescData
}

var file_protobuf_mock_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_mock_proto_goTypes = []interface{}{
	(*IncrementRequest)(nil), // 0: mock.IncrementRequest
	(*IncrementReply)(nil),   // 1: mock.IncrementReply
}
var file_protobuf_mock_proto_depIdxs = []int32{
	0, // 0: mock.Mock.Increment:input_type -> mock.IncrementRequest
	1, // 1: mock.Mock.Increment:output_type -> mock.IncrementReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_mock_proto_init() }
func file_protobuf_mock_proto_init() {
	if File_protobuf_mock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_mock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrementRequest); i {
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
		file_protobuf_mock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrementReply); i {
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
			RawDescriptor: file_protobuf_mock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_mock_proto_goTypes,
		DependencyIndexes: file_protobuf_mock_proto_depIdxs,
		MessageInfos:      file_protobuf_mock_proto_msgTypes,
	}.Build()
	File_protobuf_mock_proto = out.File
	file_protobuf_mock_proto_rawDesc = nil
	file_protobuf_mock_proto_goTypes = nil
	file_protobuf_mock_proto_depIdxs = nil
}
