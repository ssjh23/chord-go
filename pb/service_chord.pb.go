// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.0
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
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x66, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbc, 0x04,
	0x0a, 0x05, 0x43, 0x68, 0x6f, 0x72, 0x64, 0x12, 0x43, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x6f, 0x72, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x6f, 0x72, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x6f, 0x72, 0x64, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x11,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x55, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69,
	0x6e, 0x67, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x69, 0x6e, 0x67, 0x12,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x73, 0x6a, 0x68, 0x32,
	0x33, 0x2f, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_chord_proto_goTypes = []interface{}{
	(*GetChordNodeRequest)(nil),        // 0: pb.GetChordNodeRequest
	(*GetRequestFromClient)(nil),       // 1: pb.GetRequestFromClient
	(*FindSuccessorRequest)(nil),       // 2: pb.FindSuccessorRequest
	(*ClientRequest)(nil),              // 3: pb.ClientRequest
	(*InsertKeyValuePairRequest)(nil),  // 4: pb.InsertKeyValuePairRequest
	(*GetValueFromKeyRequest)(nil),     // 5: pb.GetValueFromKeyRequest
	(*CreateRingRequest)(nil),          // 6: pb.CreateRingRequest
	(*JoinRingRequest)(nil),            // 7: pb.JoinRingRequest
	(*GetChordNodeResponse)(nil),       // 8: pb.GetChordNodeResponse
	(*GetResponseToClient)(nil),        // 9: pb.GetResponseToClient
	(*FindSuccessorResponse)(nil),      // 10: pb.FindSuccessorResponse
	(*ClientResponse)(nil),             // 11: pb.ClientResponse
	(*InsertKeyValuePairResponse)(nil), // 12: pb.InsertKeyValuePairResponse
	(*GetValueFromKeyResponse)(nil),    // 13: pb.GetValueFromKeyResponse
	(*CreateRingResponse)(nil),         // 14: pb.CreateRingResponse
	(*JoinRingResponse)(nil),           // 15: pb.JoinRingResponse
}
var file_service_chord_proto_depIdxs = []int32{
	0,  // 0: pb.Chord.GetChordNode:input_type -> pb.GetChordNodeRequest
	1,  // 1: pb.Chord.RequestFromClient:input_type -> pb.GetRequestFromClient
	2,  // 2: pb.Chord.FindSuccessor:input_type -> pb.FindSuccessorRequest
	3,  // 3: pb.Chord.ClientRequestHandler:input_type -> pb.ClientRequest
	4,  // 4: pb.Chord.InsertKeyValuePair:input_type -> pb.InsertKeyValuePairRequest
	5,  // 5: pb.Chord.GetValueFromKey:input_type -> pb.GetValueFromKeyRequest
	6,  // 6: pb.Chord.CreateRing:input_type -> pb.CreateRingRequest
	7,  // 7: pb.Chord.JoinRing:input_type -> pb.JoinRingRequest
	8,  // 8: pb.Chord.GetChordNode:output_type -> pb.GetChordNodeResponse
	9,  // 9: pb.Chord.RequestFromClient:output_type -> pb.GetResponseToClient
	10, // 10: pb.Chord.FindSuccessor:output_type -> pb.FindSuccessorResponse
	11, // 11: pb.Chord.ClientRequestHandler:output_type -> pb.ClientResponse
	12, // 12: pb.Chord.InsertKeyValuePair:output_type -> pb.InsertKeyValuePairResponse
	13, // 13: pb.Chord.GetValueFromKey:output_type -> pb.GetValueFromKeyResponse
	14, // 14: pb.Chord.CreateRing:output_type -> pb.CreateRingResponse
	15, // 15: pb.Chord.JoinRing:output_type -> pb.JoinRingResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_chord_proto_init() }
func file_service_chord_proto_init() {
	if File_service_chord_proto != nil {
		return
	}
	file_rpc_test_proto_init()
	file_finger_table_proto_init()
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
