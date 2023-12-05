package gapi

import (
	"context"
	"flag"
	"fmt"
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

	// Get the successor list from its successor 
	successorListResp, _ := new_successor.GetSuccessorList(ctx, &pb.GetSuccessorListRequest{})
	successorList := successorListResp.SuccessorList
	fmt.Printf("successor ip: %s\n", n.Node.successorAddress)
	log.Printf("\nResponse from Successor List: %s", successorList)
	if (n.Node.successorAddress != n.Node.myIpAddress) {
		n.updateSuccessorList(successorList, n.Node.successorAddress)
	}
	// Print successor list
	log.Printf("\nSuccessor List Final: %s", n.Node.successorList)
	//
	resp := &pb.StabilizeResponse{
		SuccessorAddress: n.Node.successorAddress,
	}
	n.updateReplicasInSuccessors(ctx)
	// Print out all replica data
	log.Printf("\nReplica Data: %s", n.Node.replicaData)
	return resp, nil
}

/* 
Update Successor List Helper function
Successor list = [Last successor, ..., 2nd successor, 1st successor]
*/
func (n *Server)updateSuccessorList(successorList []string, successorAddress string) {
	successorListLength := 5
	finalSuccessorList := successorList
	if len(successorList) == 0 {
		finalSuccessorList = append(finalSuccessorList, successorAddress) 
	} else {
		if len(finalSuccessorList) < successorListLength {
			if finalSuccessorList[0] == n.Node.myIpAddress {
				finalSuccessorList = finalSuccessorList[1:]
			}
			finalSuccessorList = append(finalSuccessorList, successorAddress)
		} else {
			finalSuccessorList = finalSuccessorList[1:]
			finalSuccessorList = append(finalSuccessorList, successorAddress)
		}
	}
	n.Node.successorList = finalSuccessorList
}

// Helper function to update replicas in all successors in list
func (n *Server) updateReplicasInSuccessors(ctx context.Context) {
	// Print successor list in function
	log.Printf("\nSuccessor List in updateReplicasInSuccessors: %s", n.Node.successorList)
	// Iterate through the successor list
	for _, successor := range n.Node.successorList {
		// If the successor is not the current node, check the replicated data for that successor
		if successor != n.Node.myIpAddress {
			// Set up a connection to the successor
			conn, err := grpc.Dial(successor, grpc.WithTransportCredentials(insecure.NewCredentials()))
			// Invoke CheckReplicateData RPC
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			successorClient := pb.NewChordClient(conn)
			// Form the request
			checkReplicateDataRequest := &pb.CheckReplicateDataRequest{
				NodeAddress: n.Node.myIpAddress,
				Data: n.Node.data,
			}
			// Invoke CheckReplicateData RPC
			_, err = successorClient.CheckReplicateData(ctx, checkReplicateDataRequest)
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
		}
	}
}