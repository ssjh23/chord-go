package gapi

import (
	"context"
	"flag"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (n *Server) NewPreSuccesor(ctx context.Context, req *pb.NewPreSuccessorRequest) (*pb.NewPreSuccessorResponse, error) {
	flag.Parse()
	if req.GetIpAddress() == "" || req.GetAddressType() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "IP Address cannot be empty and addressType needs to be predecessor or successor")
	} else if req.GetAddressType() == "PREDECESSOR" || req.GetAddressType() == "predecessor" {
		n.Node.predecessorAddress = req.GetIpAddress()
	} else if req.GetAddressType() == "SUCCESSOR" || req.GetAddressType() == "successor" {
		n.Node.successorAddress = req.GetIpAddress()
	}

	resp := &pb.NewPreSuccessorResponse{
		PredecessorAddress: n.Node.predecessorAddress,
		SucessorAddress:    n.Node.successorAddress,
	}
	return resp, nil
}
