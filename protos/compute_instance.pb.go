// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.2
// source: compute_instance.proto

package protos

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

type ComputeInstanceProfileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	SliceCount            uint32 `protobuf:"varint,2,opt,name=SliceCount,proto3" json:"SliceCount,omitempty"`
	InstanceCount         uint32 `protobuf:"varint,3,opt,name=InstanceCount,proto3" json:"InstanceCount,omitempty"`
	MultiprocessorCount   uint32 `protobuf:"varint,4,opt,name=MultiprocessorCount,proto3" json:"MultiprocessorCount,omitempty"`
	SharedCopyEngineCount uint32 `protobuf:"varint,5,opt,name=SharedCopyEngineCount,proto3" json:"SharedCopyEngineCount,omitempty"`
	SharedDecoderCount    uint32 `protobuf:"varint,6,opt,name=SharedDecoderCount,proto3" json:"SharedDecoderCount,omitempty"`
	SharedEncoderCount    uint32 `protobuf:"varint,7,opt,name=SharedEncoderCount,proto3" json:"SharedEncoderCount,omitempty"`
	SharedJpegCount       uint32 `protobuf:"varint,8,opt,name=SharedJpegCount,proto3" json:"SharedJpegCount,omitempty"`
	SharedOfaCount        uint32 `protobuf:"varint,9,opt,name=SharedOfaCount,proto3" json:"SharedOfaCount,omitempty"`
}

func (x *ComputeInstanceProfileInfo) Reset() {
	*x = ComputeInstanceProfileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compute_instance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeInstanceProfileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeInstanceProfileInfo) ProtoMessage() {}

func (x *ComputeInstanceProfileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_compute_instance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeInstanceProfileInfo.ProtoReflect.Descriptor instead.
func (*ComputeInstanceProfileInfo) Descriptor() ([]byte, []int) {
	return file_compute_instance_proto_rawDescGZIP(), []int{0}
}

func (x *ComputeInstanceProfileInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ComputeInstanceProfileInfo) GetSliceCount() uint32 {
	if x != nil {
		return x.SliceCount
	}
	return 0
}

func (x *ComputeInstanceProfileInfo) GetInstanceCount() uint32 {
	if x != nil {
		return x.InstanceCount
	}
	return 0
}

func (x *ComputeInstanceProfileInfo) GetMultiprocessorCount() uint32 {
	if x != nil {
		return x.MultiprocessorCount
	}
	return 0
}

func (x *ComputeInstanceProfileInfo) GetSharedCopyEngineCount() uint32 {
	if x != nil {
		return x.SharedCopyEngineCount
	}
	return 0
}

func (x *ComputeInstanceProfileInfo) GetSharedDecoderCount() uint32 {
	if x != nil {
		return x.SharedDecoderCount
	}
	return 0
}

func (x *ComputeInstanceProfileInfo) GetSharedEncoderCount() uint32 {
	if x != nil {
		return x.SharedEncoderCount
	}
	return 0
}

func (x *ComputeInstanceProfileInfo) GetSharedJpegCount() uint32 {
	if x != nil {
		return x.SharedJpegCount
	}
	return 0
}

func (x *ComputeInstanceProfileInfo) GetSharedOfaCount() uint32 {
	if x != nil {
		return x.SharedOfaCount
	}
	return 0
}

type ComputeInstanceProfileInfos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ComputeInstanceProfileInfo `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *ComputeInstanceProfileInfos) Reset() {
	*x = ComputeInstanceProfileInfos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compute_instance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeInstanceProfileInfos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeInstanceProfileInfos) ProtoMessage() {}

func (x *ComputeInstanceProfileInfos) ProtoReflect() protoreflect.Message {
	mi := &file_compute_instance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeInstanceProfileInfos.ProtoReflect.Descriptor instead.
func (*ComputeInstanceProfileInfos) Descriptor() ([]byte, []int) {
	return file_compute_instance_proto_rawDescGZIP(), []int{1}
}

func (x *ComputeInstanceProfileInfos) GetData() []*ComputeInstanceProfileInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type ComputeInstanceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GpuInstanceInfo *GpuInstanceInfo `protobuf:"bytes,1,opt,name=GpuInstanceInfo,proto3" json:"GpuInstanceInfo,omitempty"`
	Id              uint32           `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	ProfileId       uint32           `protobuf:"varint,3,opt,name=ProfileId,proto3" json:"ProfileId,omitempty"`
}

func (x *ComputeInstanceInfo) Reset() {
	*x = ComputeInstanceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compute_instance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeInstanceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeInstanceInfo) ProtoMessage() {}

func (x *ComputeInstanceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_compute_instance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeInstanceInfo.ProtoReflect.Descriptor instead.
func (*ComputeInstanceInfo) Descriptor() ([]byte, []int) {
	return file_compute_instance_proto_rawDescGZIP(), []int{2}
}

func (x *ComputeInstanceInfo) GetGpuInstanceInfo() *GpuInstanceInfo {
	if x != nil {
		return x.GpuInstanceInfo
	}
	return nil
}

func (x *ComputeInstanceInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ComputeInstanceInfo) GetProfileId() uint32 {
	if x != nil {
		return x.ProfileId
	}
	return 0
}

type ComputeInstanceInfos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ComputeInstanceInfo `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *ComputeInstanceInfos) Reset() {
	*x = ComputeInstanceInfos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compute_instance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeInstanceInfos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeInstanceInfos) ProtoMessage() {}

func (x *ComputeInstanceInfos) ProtoReflect() protoreflect.Message {
	mi := &file_compute_instance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeInstanceInfos.ProtoReflect.Descriptor instead.
func (*ComputeInstanceInfos) Descriptor() ([]byte, []int) {
	return file_compute_instance_proto_rawDescGZIP(), []int{3}
}

func (x *ComputeInstanceInfos) GetData() []*ComputeInstanceInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_compute_instance_proto protoreflect.FileDescriptor

var file_compute_instance_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x1a, 0x12, 0x67, 0x70, 0x75, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x03, 0x0a, 0x1a, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x70, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x43, 0x6f, 0x70, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2e, 0x0a, 0x12, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x44, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2e, 0x0a, 0x12, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4a, 0x70, 0x65, 0x67, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x4a, 0x70, 0x65, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x4f, 0x66, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4f, 0x66, 0x61, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x55, 0x0a, 0x1b, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x12, 0x36, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x86, 0x01, 0x0a, 0x13, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x41, 0x0a, 0x0f, 0x47, 0x70, 0x75, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x70, 0x75, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x47, 0x70, 0x75, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_compute_instance_proto_rawDescOnce sync.Once
	file_compute_instance_proto_rawDescData = file_compute_instance_proto_rawDesc
)

func file_compute_instance_proto_rawDescGZIP() []byte {
	file_compute_instance_proto_rawDescOnce.Do(func() {
		file_compute_instance_proto_rawDescData = protoimpl.X.CompressGZIP(file_compute_instance_proto_rawDescData)
	})
	return file_compute_instance_proto_rawDescData
}

var file_compute_instance_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_compute_instance_proto_goTypes = []interface{}{
	(*ComputeInstanceProfileInfo)(nil),  // 0: protos.ComputeInstanceProfileInfo
	(*ComputeInstanceProfileInfos)(nil), // 1: protos.ComputeInstanceProfileInfos
	(*ComputeInstanceInfo)(nil),         // 2: protos.ComputeInstanceInfo
	(*ComputeInstanceInfos)(nil),        // 3: protos.ComputeInstanceInfos
	(*GpuInstanceInfo)(nil),             // 4: protos.GpuInstanceInfo
}
var file_compute_instance_proto_depIdxs = []int32{
	0, // 0: protos.ComputeInstanceProfileInfos.Data:type_name -> protos.ComputeInstanceProfileInfo
	4, // 1: protos.ComputeInstanceInfo.GpuInstanceInfo:type_name -> protos.GpuInstanceInfo
	2, // 2: protos.ComputeInstanceInfos.Data:type_name -> protos.ComputeInstanceInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_compute_instance_proto_init() }
func file_compute_instance_proto_init() {
	if File_compute_instance_proto != nil {
		return
	}
	file_gpu_instance_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_compute_instance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeInstanceProfileInfo); i {
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
		file_compute_instance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeInstanceProfileInfos); i {
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
		file_compute_instance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeInstanceInfo); i {
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
		file_compute_instance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeInstanceInfos); i {
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
			RawDescriptor: file_compute_instance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_compute_instance_proto_goTypes,
		DependencyIndexes: file_compute_instance_proto_depIdxs,
		MessageInfos:      file_compute_instance_proto_msgTypes,
	}.Build()
	File_compute_instance_proto = out.File
	file_compute_instance_proto_rawDesc = nil
	file_compute_instance_proto_goTypes = nil
	file_compute_instance_proto_depIdxs = nil
}
