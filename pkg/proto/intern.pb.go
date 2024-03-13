// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: pkg/proto/intern.proto

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

type GetInternInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetInternInfoReq) Reset() {
	*x = GetInternInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_intern_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInternInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInternInfoReq) ProtoMessage() {}

func (x *GetInternInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_intern_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInternInfoReq.ProtoReflect.Descriptor instead.
func (*GetInternInfoReq) Descriptor() ([]byte, []int) {
	return file_pkg_proto_intern_proto_rawDescGZIP(), []int{0}
}

func (x *GetInternInfoReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetInternInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Hub    string `protobuf:"bytes,3,opt,name=hub,proto3" json:"hub,omitempty"`
}

func (x *GetInternInfoResponse) Reset() {
	*x = GetInternInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_intern_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInternInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInternInfoResponse) ProtoMessage() {}

func (x *GetInternInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_intern_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInternInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInternInfoResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_intern_proto_rawDescGZIP(), []int{1}
}

func (x *GetInternInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetInternInfoResponse) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *GetInternInfoResponse) GetHub() string {
	if x != nil {
		return x.Hub
	}
	return ""
}

var File_pkg_proto_intern_proto protoreflect.FileDescriptor

var file_pkg_proto_intern_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x68, 0x75, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x68, 0x75, 0x62, 0x32, 0x52,
	0x0a, 0x06, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_intern_proto_rawDescOnce sync.Once
	file_pkg_proto_intern_proto_rawDescData = file_pkg_proto_intern_proto_rawDesc
)

func file_pkg_proto_intern_proto_rawDescGZIP() []byte {
	file_pkg_proto_intern_proto_rawDescOnce.Do(func() {
		file_pkg_proto_intern_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_intern_proto_rawDescData)
	})
	return file_pkg_proto_intern_proto_rawDescData
}

var file_pkg_proto_intern_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_intern_proto_goTypes = []interface{}{
	(*GetInternInfoReq)(nil),      // 0: proto.GetInternInfoReq
	(*GetInternInfoResponse)(nil), // 1: proto.GetInternInfoResponse
}
var file_pkg_proto_intern_proto_depIdxs = []int32{
	0, // 0: proto.Intern.GetInternInfo:input_type -> proto.GetInternInfoReq
	1, // 1: proto.Intern.GetInternInfo:output_type -> proto.GetInternInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_intern_proto_init() }
func file_pkg_proto_intern_proto_init() {
	if File_pkg_proto_intern_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_intern_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInternInfoReq); i {
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
		file_pkg_proto_intern_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInternInfoResponse); i {
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
			RawDescriptor: file_pkg_proto_intern_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_intern_proto_goTypes,
		DependencyIndexes: file_pkg_proto_intern_proto_depIdxs,
		MessageInfos:      file_pkg_proto_intern_proto_msgTypes,
	}.Build()
	File_pkg_proto_intern_proto = out.File
	file_pkg_proto_intern_proto_rawDesc = nil
	file_pkg_proto_intern_proto_goTypes = nil
	file_pkg_proto_intern_proto_depIdxs = nil
}