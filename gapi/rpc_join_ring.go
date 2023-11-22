package gapi

import (
	"context"
	"flag"
	"log"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func (n *Server) JoinRing(ctx context.Context, req *pb.JoinRingRequest) (*pb.JoinRingResponse, error) {
	flag.Parse()
	if req.GetJoinAddress() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Address cannot be empty")
	}

	m := 6
	node := n.Node
	myHashedIp := Sha1Modulo(node.myIpAddress, m)

	// connect to the node that is already in the ring
	conn, err := grpc.Dial(req.GetJoinAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	// Ask node to find successor using ip address
	findSuccessorMessage := &pb.FindSuccessorRequest{RequestedKey: node.myIpAddress}
	successorResponse, _ := c.FindSuccessor(ctx, findSuccessorMessage)

	node.successorAddress = successorResponse.SuccessorAddress
	node.predecessorAddress = "nil"

	resp := &pb.JoinRingResponse{
		HashedID:           myHashedIp,
		Address:            node.myIpAddress,
		SuccessorAddress:   node.successorAddress,
		PredecessorAddress: node.predecessorAddress,
	}
	return resp, nil
}
