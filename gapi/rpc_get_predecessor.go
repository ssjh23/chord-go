package gapi

import (
	"context"
	"flag"
	"os"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetPredecessor(ctx context.Context, req *pb.GetPredecessorRequest) (*pb.GetPredecessorResponse, error) {
	flag.Parse()
	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty")
	}

	resp := &pb.GetPredecessorResponse{
		Predecessor: os.Getenv("PREDECESSOR_ADDRESS"),
	}
	return resp, nil
}
