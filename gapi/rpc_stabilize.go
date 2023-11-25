package gapi

import (
	"context"
	"flag"
	"log"
	"math"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (n *Server) Stabilize(ctx context.Context, req *pb.StabilizeRequest) (*pb.StabilizeResponse, error) {
	flag.Parse()
	// connect to the successor Node
	conn, err := grpc.Dial(n.Node.successorAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	successor := pb.NewChordClient(conn)

	// Ask node to for its predecessor
	successorResp, err := successor.GetPredecessor(ctx, &pb.GetPredecessorRequest{IpAddress: ""})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	successorPredecessor := successorResp.PredecessorAddress
	defer conn.Close()
	m := 6
	myHashedIp := Sha1Modulo(n.Node.myIpAddress, m)
	successorHashedIP := Sha1Modulo(n.Node.successorAddress, m)
	successorpredecessorHashedIp := Sha1Modulo(successorPredecessor, m)
	// log.Printf("My hashed IP: %+v", myHashedIp)
	// log.Printf("successor haship: %+v", successorHashedIP)
	// log.Printf("sucessorpredecessor haship: %+v", successorpredecessorHashedIp)

	// handling edge case: i am my own successor
	if myHashedIp == successorHashedIP {
		n.Node.successorAddress = successorPredecessor
	}

	// handling edge case: in between "start" & "end" of ring
	if myHashedIp > successorHashedIP {
		if (successorpredecessorHashedIp > myHashedIp && successorpredecessorHashedIp <= int64(math.Pow(float64(2), float64(m)))) || (successorpredecessorHashedIp >= 0 && successorpredecessorHashedIp < successorHashedIP) {
			n.Node.successorAddress = successorPredecessor
		}
	}

	if myHashedIp < successorpredecessorHashedIp && successorpredecessorHashedIp < successorHashedIP {
		// update new successors
		log.Printf("i entered\n")

		n.Node.successorAddress = successorPredecessor
		// } else if myHashedIp > successorpredecessorHashedIp && successorpredecessorHashedIp > successorHashedIP {
		// 	// update new successor
		// 	log.Printf("i entered there\n")

		// 	n.Node.successorAddress = successorPredecessor
	}

	// connect to new successor Node to run notify
	// connect to the successor Node
	conn2, err := grpc.Dial(n.Node.successorAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn2.Close()
	new_successor := pb.NewChordClient(conn2)

	// Ask node to run notify and update its predecessor
	log.Printf("My sucessor IP Updated: %s", n.Node.successorAddress)
	new_successorResp, err := new_successor.Notify(ctx, &pb.NotifyRequest{IpAddress: n.Node.myIpAddress})
	if err != nil {
		log.Fatalf("%v Fail to notify: %v", new_successorResp, err)
	}
	// new_successorPredecessor := new_successorResp.PredecessorAddress
	// log.Printf("Successor Predecessor Updated: %s", new_successorPredecessor)

	resp := &pb.StabilizeResponse{
		SuccessorAddress: n.Node.successorAddress,
	}
	return resp, nil
}
