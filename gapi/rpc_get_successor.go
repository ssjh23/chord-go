package gapi

import (
	"context"
	"flag"
	"os"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetSuccessor(ctx context.Context, req *pb.GetSuccessorRequest) (*pb.GetSuccessorResponse, error) {
	flag.Parse()
	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty")
	}

	resp := &pb.GetSuccessorResponse{
		Successor: os.Getenv("SUCCESSOR_ADDRESS"),
	}
	return resp, nil
}
