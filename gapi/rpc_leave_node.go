package gapi

import (
	"context"
	"flag"
	"os"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) LeaveNode(ctx context.Context, req *pb.LeaveNodeRequest) (*pb.LeaveNodeResponse, error) {
	flag.Parse()
	os.Setenv("SUCCESSOR_ADDRESS", "")
	os.Setenv("PREDECESSOR_ADDRESS", "")
	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty")
	}

	resp := &pb.LeaveNodeResponse{
		Id:          os.Getenv("CHORD_ID"),
		Successor:   os.Getenv("SUCCESSOR_ADDRESS"),
		Address:     os.Getenv("SERVER_ADDRESS"),
		Predecessor: os.Getenv("PREDECESSOR_ADDRESS"),
	}
	return resp, nil
}
