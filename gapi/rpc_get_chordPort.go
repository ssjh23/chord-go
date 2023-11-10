package gapi

import (
	"context"
	"flag"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetChordPort(ctx context.Context, req *pb.GetChordPortRequest) (*pb.GetChordPortResponse, error) {
	flag.Parse()
	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty")
	}

	resp := &pb.GetChordPortResponse{
		Port: 3,
	}
	return resp, nil
}
