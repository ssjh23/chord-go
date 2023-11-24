package gapi

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ssjh23/chord-go/pb"
)

func (n *Server) FixFingerTable(ctx context.Context, req *pb.FixFingerTableRequest) (*pb.FixFingerTableResponse, error) {
	fmt.Println("fixing finger table.....")
	for i, finger := range n.fTable {
		fmt.Printf("KEY: %v, ADDRESS: %v\n", finger.key, finger.NodeAddress)
		findSuccessorMessage := &pb.FindSuccessorRequest{RequestedKey: strconv.FormatInt(finger.key, 10)}
		successorResponse, _ := n.FindSuccessor(ctx, findSuccessorMessage)
		fmt.Println("successorResponse: ", successorResponse.SuccessorAddress)
		n.fTable[i].NodeAddress = successorResponse.SuccessorAddress
	}
	resp := &pb.FixFingerTableResponse{
		IpAddress: "fix fingers successful",
	}
	return resp, nil
}
