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
	myHashedIp := Sha1Modulo(n.Node.myIpAddress, m)

	// connect to the node that is already in the ring
	conn, err := grpc.Dial(req.GetJoinAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	// Ask node to find successor using ip address
	findSuccessorMessage := &pb.FindSuccessorRequest{RequestedKey: n.Node.myIpAddress}
	successorResponse, _ := c.FindSuccessor(ctx, findSuccessorMessage)

	// connect to new successor
	n.Node.successorAddress = successorResponse.SuccessorAddress
	conn2, err := grpc.Dial(n.Node.successorAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn2.Close()
	successor := pb.NewChordClient(conn2)

	// Ask node to run notify and update its predecessor
	// log.Printf("My sucessor IP Updated: %s", n.Node.successorAddress)
	successorResp, err := successor.Notify(ctx, &pb.NotifyRequest{IpAddress: n.Node.myIpAddress})
	if err != nil {
		log.Fatalf("%v Fail to notify: %v", successorResp, err)
	}
	successorPredecessor := successorResp.PredecessorAddress
	log.Printf("Successor Predecessor Updated: %s", successorPredecessor)

	n.Node.predecessorAddress = n.Node.myIpAddress

	// update my data based on my successor's data
	new_data_to_store := successorResp.DataToBeAbsorbed
	for key, value := range new_data_to_store {
		n.Node.data[key] = value
	}

	// JUST ADDED TO POPULATE MY PREDECESSOR
	resp := &pb.JoinRingResponse{
		HashedID:           myHashedIp,
		Address:            n.Node.myIpAddress,
		SuccessorAddress:   n.Node.successorAddress,
		PredecessorAddress: n.Node.predecessorAddress,
	}
	return resp, nil
}
