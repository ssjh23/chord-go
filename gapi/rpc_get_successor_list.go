package gapi

import (
	"context"
	"github.com/ssjh23/chord-go/pb"
)

func (n *Server) GetSuccessorList(ctx context.Context, req *pb.GetSuccessorListRequest) (*pb.GetSuccessorListResponse, error) {
	resp := &pb.GetSuccessorListResponse{
		SuccessorList: n.Node.successorList,
	}
	return resp, nil
}