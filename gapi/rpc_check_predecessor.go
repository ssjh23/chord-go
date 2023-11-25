package gapi

import (
	"context"
	"flag"
	"log"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (n *Server) CheckPredecessor(ctx context.Context, req *pb.CheckPredecessorRequest) (*pb.CheckPredecessorResponse, error) {
	flag.Parse()
	conn, err := grpc.Dial(n.Node.predecessorAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to predecessor: %v", err)
		n.Node.predecessorAddress = "nil"
	}
	defer conn.Close()

	resp := &pb.CheckPredecessorResponse{}
	return resp, nil
}
