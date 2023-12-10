// function call on a node to initialise a node ring
// does not require any input
// return: successor, predecessor, own IP address and hashed id

package gapi

import (
	"context"
	"flag"

	"github.com/ssjh23/chord-go/constant"
	"github.com/ssjh23/chord-go/pb"
)

func (n *Server) CreateRing(ctx context.Context, req *pb.CreateRingRequest) (*pb.CreateRingResponse, error) {
	flag.Parse()

	m = constant.VALUE_OF_M

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
