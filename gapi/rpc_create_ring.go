package gapi

import (
	"context"
	"flag"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (n *Server) CreateRing(ctx context.Context, req *pb.CreateRingRequest) (*pb.CreateRingResponse, error) {
	flag.Parse()
	if req.GetAddress() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Address cannot be empty")
	}

	m := 6
	node := n.Node
	node.successorAddress = node.myIpAddress
	node.predecessorAddress = "nil"
	myHashedIp := Sha1Modulo(node.myIpAddress, m)

	resp := &pb.CreateRingResponse{
		HashedID:           myHashedIp,
		Address:            node.myIpAddress,
		SuccessorAddress:   node.successorAddress,
		PredecessorAddress: node.predecessorAddress,
	}
	return resp, nil
}
