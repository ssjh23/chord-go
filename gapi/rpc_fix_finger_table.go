package gapi

// updates the finger table entries of the node
// called periodically along with Stabilize

import (
	"context"
	"strconv"

	"github.com/ssjh23/chord-go/pb"
)

func (n *Server) FixFingerTable(ctx context.Context, req *pb.FixFingerTableRequest) (*pb.FixFingerTableResponse, error) {
	// fmt.Println("fixing finger table.....")
	for i, finger := range n.fTable {
		findSuccessorMessage := &pb.FindSuccessorRequest{RequestedKey: strconv.FormatInt(finger.key, 10)}
		successorResponse, _ := n.FindSuccessor(ctx, findSuccessorMessage)
		n.fTable[i].NodeAddress = successorResponse.SuccessorAddress
	}
	resp := &pb.FixFingerTableResponse{
		IpAddress: "fix fingers successful",
	}
	return resp, nil
}
