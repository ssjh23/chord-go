package gapi

import (
	"context"
	"flag"

	"github.com/ssjh23/chord-go/pb"
)

func (n *Server) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	flag.Parse()
	resp := &pb.GetInfoResponse{
		MyIpAddress:        n.Node.myIpAddress,
		SuccessorAddress:   n.Node.successorAddress,
		PrecedessorAddress: n.Node.predecessorAddress,
		SuccessorList:      n.Node.successorList,
		Data:               n.Node.data,
	}
	return resp, nil
}
