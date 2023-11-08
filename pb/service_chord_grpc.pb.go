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
	JoinNode(ctx context.Context, in *JoinNodeRequest, opts ...grpc.CallOption) (*JoinNodeResponse, error)
	LeaveNode(ctx context.Context, in *LeaveNodeRequest, opts ...grpc.CallOption) (*LeaveNodeResponse, error)
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

func (c *chordClient) JoinNode(ctx context.Context, in *JoinNodeRequest, opts ...grpc.CallOption) (*JoinNodeResponse, error) {
	out := new(JoinNodeResponse)
	err := c.cc.Invoke(ctx, "/pb.Chord/JoinNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) LeaveNode(ctx context.Context, in *LeaveNodeRequest, opts ...grpc.CallOption) (*LeaveNodeResponse, error) {
	out := new(LeaveNodeResponse)
	err := c.cc.Invoke(ctx, "/pb.Chord/LeaveNode", in, out, opts...)
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
	JoinNode(context.Context, *JoinNodeRequest) (*JoinNodeResponse, error)
	LeaveNode(context.Context, *LeaveNodeRequest) (*LeaveNodeResponse, error)
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
func (UnimplementedChordServer) JoinNode(context.Context, *JoinNodeRequest) (*JoinNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinNode not implemented")
}
func (UnimplementedChordServer) LeaveNode(context.Context, *LeaveNodeRequest) (*LeaveNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveNode not implemented")
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

func _Chord_JoinNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).JoinNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/JoinNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).JoinNode(ctx, req.(*JoinNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_LeaveNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).LeaveNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/LeaveNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).LeaveNode(ctx, req.(*LeaveNodeRequest))
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
			MethodName: "JoinNode",
			Handler:    _Chord_JoinNode_Handler,
		},
		{
			MethodName: "LeaveNode",
			Handler:    _Chord_LeaveNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_chord.proto",
}
