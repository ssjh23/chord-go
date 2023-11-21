package gapi

import (
	"context"
	"flag"
	"os"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) JoinNode(ctx context.Context, req *pb.JoinNodeRequest) (*pb.JoinNodeResponse, error) {
	flag.Parse()
	os.Setenv("SUCCESSOR_ADDRESS", "chord-go-chord1-1:9090")
	// to implement inform successor to update predecessor
	os.Setenv("PREDECESSOR_ADDRESS", "chord-go-chord4-1:9093")
	// to implement inform predecessor to update successor

	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty")
	}

	resp := &pb.JoinNodeResponse{
		Id:          os.Getenv("CHORD_ID"),
		Successor:   os.Getenv("SUCCESSOR_ADDRESS"),
		Address:     os.Getenv("SERVER_ADDRESS"),
		Predecessor: os.Getenv("PREDECESSOR_ADDRESS"),
	}
	return resp, nil
}
