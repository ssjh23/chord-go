// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.0
// source: rpc_test.proto

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

type GetChordNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetChordNodeRequest) Reset() {
	*x = GetChordNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChordNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChordNodeRequest) ProtoMessage() {}

func (x *GetChordNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChordNodeRequest.ProtoReflect.Descriptor instead.
func (*GetChordNodeRequest) Descriptor() ([]byte, []int) {
	return file_rpc_test_proto_rawDescGZIP(), []int{0}
}

func (x *GetChordNodeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetChordNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ip   string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *GetChordNodeResponse) Reset() {
	*x = GetChordNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChordNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChordNodeResponse) ProtoMessage() {}

func (x *GetChordNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChordNodeResponse.ProtoReflect.Descriptor instead.
func (*GetChordNodeResponse) Descriptor() ([]byte, []int) {
	return file_rpc_test_proto_rawDescGZIP(), []int{1}
}

func (x *GetChordNodeResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetChordNodeResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *GetChordNodeResponse) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type GetRequestFromClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *GetRequestFromClient) Reset() {
	*x = GetRequestFromClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequestFromClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequestFromClient) ProtoMessage() {}

func (x *GetRequestFromClient) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequestFromClient.ProtoReflect.Descriptor instead.
func (*GetRequestFromClient) Descriptor() ([]byte, []int) {
	return file_rpc_test_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequestFromClient) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type GetResponseToClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputFromClient string `protobuf:"bytes,1,opt,name=inputFromClient,proto3" json:"inputFromClient,omitempty"`
	Id              string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Ip              string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Port            int32  `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *GetResponseToClient) Reset() {
	*x = GetResponseToClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponseToClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponseToClient) ProtoMessage() {}

func (x *GetResponseToClient) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponseToClient.ProtoReflect.Descriptor instead.
func (*GetResponseToClient) Descriptor() ([]byte, []int) {
	return file_rpc_test_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponseToClient) GetInputFromClient() string {
	if x != nil {
		return x.InputFromClient
	}
	return ""
}

func (x *GetResponseToClient) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetResponseToClient) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *GetResponseToClient) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_rpc_test_proto protoreflect.FileDescriptor

var file_rpc_test_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x68, 0x6f, 0x72, 0x64,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x6f, 0x72, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x2c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x73, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x72, 0x6f, 0x6d,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x73, 0x6a, 0x68, 0x32, 0x33, 0x2f,
	0x63, 0x68, 0x6f, 0x72, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rpc_test_proto_rawDescOnce sync.Once
	file_rpc_test_proto_rawDescData = file_rpc_test_proto_rawDesc
)

func file_rpc_test_proto_rawDescGZIP() []byte {
	file_rpc_test_proto_rawDescOnce.Do(func() {
		file_rpc_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_test_proto_rawDescData)
	})
	return file_rpc_test_proto_rawDescData
}

var file_rpc_test_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_test_proto_goTypes = []interface{}{
	(*GetChordNodeRequest)(nil),  // 0: pb.GetChordNodeRequest
	(*GetChordNodeResponse)(nil), // 1: pb.GetChordNodeResponse
	(*GetRequestFromClient)(nil), // 2: pb.GetRequestFromClient
	(*GetResponseToClient)(nil),  // 3: pb.GetResponseToClient
}
var file_rpc_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_test_proto_init() }
func file_rpc_test_proto_init() {
	if File_rpc_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChordNodeRequest); i {
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
		file_rpc_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChordNodeResponse); i {
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
		file_rpc_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequestFromClient); i {
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
		file_rpc_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponseToClient); i {
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
			RawDescriptor: file_rpc_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_test_proto_goTypes,
		DependencyIndexes: file_rpc_test_proto_depIdxs,
		MessageInfos:      file_rpc_test_proto_msgTypes,
	}.Build()
	File_rpc_test_proto = out.File
	file_rpc_test_proto_rawDesc = nil
	file_rpc_test_proto_goTypes = nil
	file_rpc_test_proto_depIdxs = nil
}
