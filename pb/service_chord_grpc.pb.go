// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: service_chord.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChordClient is the client API for Chord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChordClient interface {
	GetChordNode(ctx context.Context, in *GetChordNodeRequest, opts ...grpc.CallOption) (*GetChordNodeResponse, error)
	RequestFromClient(ctx context.Context, in *GetRequestFromClient, opts ...grpc.CallOption) (*GetResponseToClient, error)
	FindSuccessor(ctx context.Context, in *FindSuccessorRequest, opts ...grpc.CallOption) (*FindSuccessorResponse, error)
	// rpc for external client to insert or retrieve a key-value pair into the DHT
	ClientRequestHandler(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ClientResponse, error)
	InsertKeyValuePair(ctx context.Context, in *InsertKeyValuePairRequest, opts ...grpc.CallOption) (*InsertKeyValuePairResponse, error)
	GetValueFromKey(ctx context.Context, in *GetValueFromKeyRequest, opts ...grpc.CallOption) (*GetValueFromKeyResponse, error)
	CreateRing(ctx context.Context, in *CreateRingRequest, opts ...grpc.CallOption) (*CreateRingResponse, error)
	JoinRing(ctx context.Context, in *JoinRingRequest, opts ...grpc.CallOption) (*JoinRingResponse, error)
}

type chordClient struct {
	cc grpc.ClientConnInterface
}

func NewChordClient(cc grpc.ClientConnInterface) ChordClient {
	return &chordClient{cc}
}

func (c *chordClient) GetChordNode(ctx context.Context, in *GetChordNodeRequest, opts ...grpc.CallOption) (*GetChordNodeResponse, error) {
	out := new(GetChordNodeResponse)
	err := c.cc.Invoke(ctx, "/pb.Chord/GetChordNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) RequestFromClient(ctx context.Context, in *GetRequestFromClient, opts ...grpc.CallOption) (*GetResponseToClient, error) {
	out := new(GetResponseToClient)
	err := c.cc.Invoke(ctx, "/pb.Chord/RequestFromClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) FindSuccessor(ctx context.Context, in *FindSuccessorRequest, opts ...grpc.CallOption) (*FindSuccessorResponse, error) {
	out := new(FindSuccessorResponse)
	err := c.cc.Invoke(ctx, "/pb.Chord/FindSuccessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) ClientRequestHandler(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ClientResponse, error) {
	out := new(ClientResponse)
	err := c.cc.Invoke(ctx, "/pb.Chord/ClientRequestHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) InsertKeyValuePair(ctx context.Context, in *InsertKeyValuePairRequest, opts ...grpc.CallOption) (*InsertKeyValuePairResponse, error) {
	out := new(InsertKeyValuePairResponse)
	err := c.cc.Invoke(ctx, "/pb.Chord/InsertKeyValuePair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) GetValueFromKey(ctx context.Context, in *GetValueFromKeyRequest, opts ...grpc.CallOption) (*GetValueFromKeyResponse, error) {
	out := new(GetValueFromKeyResponse)
	err := c.cc.Invoke(ctx, "/pb.Chord/GetValueFromKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) CreateRing(ctx context.Context, in *CreateRingRequest, opts ...grpc.CallOption) (*CreateRingResponse, error) {
	out := new(CreateRingResponse)
	err := c.cc.Invoke(ctx, "/pb.Chord/CreateRing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) JoinRing(ctx context.Context, in *JoinRingRequest, opts ...grpc.CallOption) (*JoinRingResponse, error) {
	out := new(JoinRingResponse)
	err := c.cc.Invoke(ctx, "/pb.Chord/JoinRing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChordServer is the server API for Chord service.
// All implementations must embed UnimplementedChordServer
// for forward compatibility
type ChordServer interface {
	GetChordNode(context.Context, *GetChordNodeRequest) (*GetChordNodeResponse, error)
	RequestFromClient(context.Context, *GetRequestFromClient) (*GetResponseToClient, error)
	FindSuccessor(context.Context, *FindSuccessorRequest) (*FindSuccessorResponse, error)
	// rpc for external client to insert or retrieve a key-value pair into the DHT
	ClientRequestHandler(context.Context, *ClientRequest) (*ClientResponse, error)
	InsertKeyValuePair(context.Context, *InsertKeyValuePairRequest) (*InsertKeyValuePairResponse, error)
	GetValueFromKey(context.Context, *GetValueFromKeyRequest) (*GetValueFromKeyResponse, error)
	CreateRing(context.Context, *CreateRingRequest) (*CreateRingResponse, error)
	JoinRing(context.Context, *JoinRingRequest) (*JoinRingResponse, error)
	mustEmbedUnimplementedChordServer()
}

// UnimplementedChordServer must be embedded to have forward compatible implementations.
type UnimplementedChordServer struct {
}

func (UnimplementedChordServer) GetChordNode(context.Context, *GetChordNodeRequest) (*GetChordNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChordNode not implemented")
}
func (UnimplementedChordServer) RequestFromClient(context.Context, *GetRequestFromClient) (*GetResponseToClient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestFromClient not implemented")
}
func (UnimplementedChordServer) FindSuccessor(context.Context, *FindSuccessorRequest) (*FindSuccessorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessor not implemented")
}
func (UnimplementedChordServer) ClientRequestHandler(context.Context, *ClientRequest) (*ClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientRequestHandler not implemented")
}
func (UnimplementedChordServer) InsertKeyValuePair(context.Context, *InsertKeyValuePairRequest) (*InsertKeyValuePairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertKeyValuePair not implemented")
}
func (UnimplementedChordServer) GetValueFromKey(context.Context, *GetValueFromKeyRequest) (*GetValueFromKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValueFromKey not implemented")
}
func (UnimplementedChordServer) CreateRing(context.Context, *CreateRingRequest) (*CreateRingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRing not implemented")
}
func (UnimplementedChordServer) JoinRing(context.Context, *JoinRingRequest) (*JoinRingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRing not implemented")
}
func (UnimplementedChordServer) mustEmbedUnimplementedChordServer() {}

// UnsafeChordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChordServer will
// result in compilation errors.
type UnsafeChordServer interface {
	mustEmbedUnimplementedChordServer()
}

func RegisterChordServer(s grpc.ServiceRegistrar, srv ChordServer) {
	s.RegisterService(&Chord_ServiceDesc, srv)
}

func _Chord_GetChordNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChordNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).GetChordNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/GetChordNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).GetChordNode(ctx, req.(*GetChordNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_RequestFromClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestFromClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).RequestFromClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/RequestFromClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).RequestFromClient(ctx, req.(*GetRequestFromClient))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_FindSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindSuccessorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).FindSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/FindSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).FindSuccessor(ctx, req.(*FindSuccessorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_ClientRequestHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).ClientRequestHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/ClientRequestHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).ClientRequestHandler(ctx, req.(*ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_InsertKeyValuePair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertKeyValuePairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).InsertKeyValuePair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/InsertKeyValuePair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).InsertKeyValuePair(ctx, req.(*InsertKeyValuePairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_GetValueFromKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueFromKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).GetValueFromKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/GetValueFromKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).GetValueFromKey(ctx, req.(*GetValueFromKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_CreateRing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).CreateRing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/CreateRing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).CreateRing(ctx, req.(*CreateRingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_JoinRing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).JoinRing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/JoinRing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).JoinRing(ctx, req.(*JoinRingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chord_ServiceDesc is the grpc.ServiceDesc for Chord service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chord_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Chord",
	HandlerType: (*ChordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChordNode",
			Handler:    _Chord_GetChordNode_Handler,
		},
		{
			MethodName: "RequestFromClient",
			Handler:    _Chord_RequestFromClient_Handler,
		},
		{
			MethodName: "FindSuccessor",
			Handler:    _Chord_FindSuccessor_Handler,
		},
		{
			MethodName: "ClientRequestHandler",
			Handler:    _Chord_ClientRequestHandler_Handler,
		},
		{
			MethodName: "InsertKeyValuePair",
			Handler:    _Chord_InsertKeyValuePair_Handler,
		},
		{
			MethodName: "GetValueFromKey",
			Handler:    _Chord_GetValueFromKey_Handler,
		},
		{
			MethodName: "CreateRing",
			Handler:    _Chord_CreateRing_Handler,
		},
		{
			MethodName: "JoinRing",
			Handler:    _Chord_JoinRing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_chord.proto",
}
