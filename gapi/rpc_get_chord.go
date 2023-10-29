package gapi

import (
	"context"
	"flag"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetChordNode(ctx context.Context, req *pb.GetChordNodeRequest) (*pb.GetChordNodeResponse, error) {
	flag.Parse()
	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty")
	}

	resp := &pb.GetChordNodeResponse{
		Id: server.config.ChordId,
		Ip: "test",
		Port: 9090,
	}
	return resp, nil
}
