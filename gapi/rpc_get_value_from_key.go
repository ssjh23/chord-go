package gapi

import (
	"context"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetValueFromKey(ctx context.Context, req *pb.GetValueFromKeyRequest) (*pb.GetValueFromKeyResponse, error) {
	if req.GetKey() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "key cannot be empty")
	}
	value := s.Node.data[req.GetKey()]
	resp := &pb.GetValueFromKeyResponse{
		Value: value,
	}
	return resp, nil
}
