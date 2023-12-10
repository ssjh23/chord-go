package gapi

// function call to leave a node ring
// does not require any input
// retrun: confirmation of ip that has left the ring
import (
	"context"
	"flag"
	"log"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (n *Server) LeaveRing(ctx context.Context, req *pb.LeaveRingRequest) (*pb.LeaveRingResponse, error) {
	flag.Parse()
	successor := n.Node.successorAddress

	// connect to the successor to migrate data
	conn, err := grpc.Dial(successor, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	// Migrate data to successor
	for key, value := range n.Node.data {
		migrateDataMessage := &pb.InsertKeyValuePairRequest{Key: key, Value: value}
		migrateDataResponse, err := c.InsertKeyValuePair(ctx, migrateDataMessage)
		if err != nil {
			log.Printf("could not migrate data: %v", err)
		}
		log.Printf("Migrate Data Response: Key: %s, Value: %s - %s", key, value, migrateDataResponse.Message)
	}

	// change to this node's successor to change its predecessor to this node's predecessor
	succesorResp, err := c.NewPreSuccessor(ctx, &pb.NewPreSuccessorRequest{IpAddress: n.Node.predecessorAddress, AddressType: "predecessor"})
	if err != nil {
		log.Printf("%v Failed to update predecessor: %v", succesorResp.PredecessorAddress, err)
	}

	// connect to this node's predecessor to change its successor to this node's sucessor
	predecessor := n.Node.predecessorAddress
	conn2, err := grpc.Dial(predecessor, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn2.Close()
	c2 := pb.NewChordClient(conn2)

	// new predecessor's successor
	predecessorResp, err := c2.NewPreSuccessor(ctx, &pb.NewPreSuccessorRequest{IpAddress: n.Node.successorAddress, AddressType: "successor"})
	if err != nil {
		log.Printf("%v Failed to update successor: %v", predecessorResp.SucessorAddress, err)
	}

	// change this node's predecessor and successor to prevent stabalize from running
	n.Node.predecessorAddress = n.Node.myIpAddress
	n.Node.successorAddress = n.Node.myIpAddress

	resp := &pb.LeaveRingResponse{
		MyIpAddress: n.Node.myIpAddress + "has successfully left the ring",
	}
	return resp, nil
}
