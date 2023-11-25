package gapi

import (
	"context"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) InsertKeyValuePair(ctx context.Context, req *pb.InsertKeyValuePairRequest) (*pb.InsertKeyValuePairResponse, error) {
	if req.GetKey() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty")
	}
	s.Node.data[req.GetKey()] = req.GetValue()

	resp := &pb.InsertKeyValuePairResponse{
		Message: "Successful INSERT into " + s.Node.myIpAddress,
	}
	return resp, nil
}
