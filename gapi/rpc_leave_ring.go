package gapi

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

	// connect to the successor to add its back up data into its data
	conn, err := grpc.Dial(successor, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	// Ask node to find successor using ip address
	for key, value := range n.Node.data {
		migrateDataMessage := &pb.InsertKeyValuePairRequest{Key: key, Value: value}
		migrateDataResponse, err := c.InsertKeyValuePair(ctx, migrateDataMessage)
		if err != nil {
			log.Fatalf("could not migrate data: %v", err)
		}
		log.Printf("Migrate Data Response: Key: %s, Value: %s - %s", key, value, migrateDataResponse.Message)
	}

	// // connect to new successor
	// n.Node.successorAddress = successorResponse.SuccessorAddress
	// conn2, err := grpc.Dial(n.Node.successorAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn2.Close()
	// successor := pb.NewChordClient(conn2)

	// // Ask node to run notify and update its predecessor
	// successorResp, err := successor.Notify(ctx, &pb.NotifyRequest{IpAddress: n.Node.myIpAddress})
	// if err != nil {
	// 	log.Fatalf("%v Fail to notify: %v", successorResp, err)
	// }
	// successorPredecessor := successorResp.PredecessorAddress
	// log.Printf("Successor Predecessor Updated: %s", successorPredecessor)

	// n.Node.predecessorAddress = n.Node.myIpAddress

	// util.ChangeEnvVariable(".", "SUCCESSOR_ADDRESS", n.Node.successorAddress)
	// JUST ADDED TO POPULATE MY PREDECESSOR
	// n.Stabilize(ctx, &pb.StabilizeRequest{IpAddress: n.Node.myIpAddress})

	resp := &pb.LeaveRingResponse{
		MyIpAddress: n.Node.myIpAddress + "has successfully left the ring",
	}
	return resp, nil
}
