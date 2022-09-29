// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v1.0.0
// source: github.com/rancher/opni/plugins/model_training/pkg/model_training/model_training.proto

package model_training

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClusterID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ClusterID) Reset() {
	*x = ClusterID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterID) ProtoMessage() {}

func (x *ClusterID) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterID.ProtoReflect.Descriptor instead.
func (*ClusterID) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type WorkloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *WorkloadResponse) Reset() {
	*x = WorkloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadResponse) ProtoMessage() {}

func (x *WorkloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadResponse.ProtoReflect.Descriptor instead.
func (*WorkloadResponse) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_rawDescGZIP(), []int{1}
}

func (x *WorkloadResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_rawDesc = []byte{
	0x0a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x2c, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x75,
	0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x64, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73,
	0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x20, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x77, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x73, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_rawDescData = file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_rawDesc
)

func file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_rawDescData)
	})
	return file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_rawDescData
}

var file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_goTypes = []interface{}{
	(*ClusterID)(nil),        // 0: model_training.ClusterID
	(*WorkloadResponse)(nil), // 1: model_training.WorkloadResponse
}
var file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_depIdxs = []int32{
	0, // 0: model_training.ModelTraining.ListWorkloads:input_type -> model_training.ClusterID
	1, // 1: model_training.ModelTraining.ListWorkloads:output_type -> model_training.WorkloadResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_init()
}
func file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_init() {
	if File_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterID); i {
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
		file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadResponse); i {
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
			RawDescriptor: file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_depIdxs,
		MessageInfos:      file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto = out.File
	file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_rawDesc = nil
	file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_goTypes = nil
	file_github_com_rancher_opni_plugins_model_training_pkg_model_training_model_training_proto_depIdxs = nil
}