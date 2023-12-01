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
	return resp, nil
}

/* 
Update Successor List Helper function
Successor list = [Last successor, ..., 2nd successor, 1st successor]
*/
func (n *Server)updateSuccessorList(successorList []string, successorAddress string) {
	successorListLength := 5
	finalSuccessorList := successorList
	// Edge case: if the successor list is empty, do not add my own IP address
	if len(successorList) == 0 {
		fmt.Printf("Successor list, Empty: %s\n", successorList)
		successorList = append(successorList, successorAddress)
		finalSuccessorList = successorList
	} else {
		finalSuccessorList = successorList
		// Remove from list if last address is the same address as the current node's address
		if successorList[len(successorList)-1] == n.Node.myIpAddress {
			fmt.Printf("Successor list, Remove last : %s\n", successorList)
			finalSuccessorList = finalSuccessorList[:len(finalSuccessorList)-1]
		} 
		// Add successor to the list
		if len(finalSuccessorList) < successorListLength {
			// If after removing the last address, the list is empty or the last address is not the successor's address, add the successor's address to the list
			if len(finalSuccessorList) == 0 || (finalSuccessorList[len(finalSuccessorList)-1] != successorAddress && finalSuccessorList[0] != n.Node.myIpAddress) {
				fmt.Printf("Successor list, not full: %s\n", finalSuccessorList)
				finalSuccessorList = append(finalSuccessorList, successorAddress)
			} else {

			}
		} else {
			finalSuccessorList = finalSuccessorList[1:]
			fmt.Printf("Successor list, full: %s\n", finalSuccessorList)
			fmt.Printf("Successor address: %s\n", successorAddress)
			// Add successor to the list
			finalSuccessorList = append(finalSuccessorList, successorAddress)
		}
		// // Edge case where the last address in the list is the same as the node's ip address. Only happens when the successor list size is same as the number of nodes in the network
		// if finalSuccessorList[0] == n.Node.myIpAddress {
		// 	fmt.Printf("Successor list, same first and last: %s\n", finalSuccessorList)
		// 	finalSuccessorList = finalSuccessorList[1:]
		// }
	}
	n.Node.successorList = finalSuccessorList
}