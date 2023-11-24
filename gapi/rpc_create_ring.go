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

	n.Node.successorAddress = n.Node.myIpAddress
	n.Node.predecessorAddress = n.Node.myIpAddress
	myHashedIp := Sha1Modulo(n.Node.myIpAddress, m)

	// JUST ADDED TO POPULATE MY PREDECESSOR
	n.Stabilize(ctx, &pb.StabilizeRequest{IpAddress: n.Node.myIpAddress})

	resp := &pb.CreateRingResponse{
		HashedID:           myHashedIp,
		Address:            n.Node.myIpAddress,
		SuccessorAddress:   n.Node.successorAddress,
		PredecessorAddress: n.Node.predecessorAddress,
	}

	return resp, nil
}
