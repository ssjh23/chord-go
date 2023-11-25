package gapi

import (
	"context"
	"flag"
	"log"

	"github.com/ssjh23/chord-go/pb"
)

func (n *Server) GetPredecessor(ctx context.Context, req *pb.GetPredecessorRequest) (*pb.GetPredecessorResponse, error) {
	flag.Parse()
	if req.GetIpAddress() != "" { // empty string is a get predecessor, non empty is a update predecessor
		n.Node.predecessorAddress = req.GetIpAddress()
		log.Printf("Predecessor Updated %s", req.GetIpAddress())
	}

	resp := &pb.GetPredecessorResponse{
		PredecessorAddress: n.Node.predecessorAddress,
	}
	return resp, nil
}
