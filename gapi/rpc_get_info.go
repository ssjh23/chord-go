package gapi

// quality of life function call to get information about the node and its data and replicated
// does not require any input
// return: successor, predecessor, own IP address, successor list, data and replicated data

import (
	"context"
	"flag"

	"github.com/ssjh23/chord-go/pb"
)

func (n *Server) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	flag.Parse()

	mapping := make(map[string]*pb.GetInfoResponse_Data)

	for key, value := range n.Node.replicaData {
		data := &pb.GetInfoResponse_Data{
			DataMap: value, // Assuming your DataMap field is defined in your protobuf
		}
		mapping[key] = data
	}

	resp := &pb.GetInfoResponse{
		MyIpAddress:        n.Node.myIpAddress,
		SuccessorAddress:   n.Node.successorAddress,
		PrecedessorAddress: n.Node.predecessorAddress,
		SuccessorList:      n.Node.successorList,
		Data:               n.Node.data,
		Replicated:         mapping,
	}
	return resp, nil
}
