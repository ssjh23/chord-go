// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: service_chord.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_chord_proto protoreflect.FileDescriptor

var file_service_chord_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0e, 0x72, 0x70, 0x63, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x96, 0x01, 0x0a, 0x05, 0x43, 0x68,
	0x6f, 0x72, 0x64, 0x12, 0x43, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x68, 0x6f, 0x72, 0x64, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x6f, 0x72,
	0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x6f, 0x72, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x72, 0x6f,
	0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x73, 0x6a, 0x68, 0x32, 0x33, 0x2f, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2d, 0x67, 0x6f,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_chord_proto_goTypes = []interface{}{
	(*GetChordNodeRequest)(nil),  // 0: pb.GetChordNodeRequest
	(*GetRequestFromClient)(nil), // 1: pb.GetRequestFromClient
	(*GetChordNodeResponse)(nil), // 2: pb.GetChordNodeResponse
	(*GetResponseToClient)(nil),  // 3: pb.GetResponseToClient
}
var file_service_chord_proto_depIdxs = []int32{
	0, // 0: pb.Chord.GetChordNode:input_type -> pb.GetChordNodeRequest
	1, // 1: pb.Chord.RequestFromClient:input_type -> pb.GetRequestFromClient
	2, // 2: pb.Chord.GetChordNode:output_type -> pb.GetChordNodeResponse
	3, // 3: pb.Chord.RequestFromClient:output_type -> pb.GetResponseToClient
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_chord_proto_init() }
func file_service_chord_proto_init() {
	if File_service_chord_proto != nil {
		return
	}
	file_rpc_test_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_chord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_chord_proto_goTypes,
		DependencyIndexes: file_service_chord_proto_depIdxs,
	}.Build()
	File_service_chord_proto = out.File
	file_service_chord_proto_rawDesc = nil
	file_service_chord_proto_goTypes = nil
	file_service_chord_proto_depIdxs = nil
}
