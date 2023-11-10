package gapi

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) JoinNode(ctx context.Context, req *pb.JoinNodeRequest) (*pb.JoinNodeResponse, error) {
	flag.Parse()
	// Print environment variables before changes
	fmt.Println("Before changes:")
	fmt.Println("CHORD_ID:", os.Getenv("CHORD_ID"))
	fmt.Println("SUCCESSOR_ADDRESS:", os.Getenv("SUCCESSOR_ADDRESS"))
	fmt.Println("SERVER_ADDRESS:", os.Getenv("SERVER_ADDRESS"))
	fmt.Println("PREDECESSOR_ADDRESS:", os.Getenv("PREDECESSOR_ADDRESS"))

	os.Setenv("SUCCESSOR_ADDRESS", "chord-go-chord1-1:9090")
	os.Setenv("PREDECESSOR_ADDRESS", "chord-go-chord4-1:9093")

	// Print environment variables after changes
	fmt.Println("After changes:")
	fmt.Println("CHORD_ID:", os.Getenv("CHORD_ID"))
	fmt.Println("SUCCESSOR_ADDRESS:", os.Getenv("SUCCESSOR_ADDRESS"))
	fmt.Println("SERVER_ADDRESS:", os.Getenv("SERVER_ADDRESS"))
	fmt.Println("PREDECESSOR_ADDRESS:", os.Getenv("PREDECESSOR_ADDRESS"))
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
