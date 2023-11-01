package gapi

import (
	"context"
	"fmt"
	"log"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// type Node struct {
// 	pb.Node
// }

func (n *Server) FindSuccessor(ctx context.Context, req *pb.FindSuccessorRequest) (*pb.FindSuccessorResponse, error) {

	log.Printf("STARTING FIND SUCCESSOR: %d", n.Port)

	for _, key := range n.Node.Keys {
		if key == req.RequestedKey {
			resp := &pb.FindSuccessorResponse{
				Response: &n.Node.Port,
			}
			log.Printf("YOUR KEY IS LOCATED AT PORT: %d", resp.Response)
			return resp, nil
		}
	}

	if req.RequestedKey > n.Node.Id && req.RequestedKey <= n.Node.Successor.Id {
		resp := &pb.FindSuccessorResponse{
			Response: &n.Node.Successor.Port,
		}
		log.Printf("YOUR KEY IS LOCATED AT PORT: %d", resp.Response)
		return resp, nil
	}
	next_node_port := n.ClosestPrecedingNode(req.RequestedKey)

	serverAddr := fmt.Sprintf("localhost:%d", next_node_port)

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// return next_node_port.FindSuccessor(ctx, recursiveReq)
	resp := &pb.FindSuccessorResponse{
		Response: nil,
	}
	log.Printf("YOUR QUERY IS BEING ROUTED TO PORT: %d", next_node_port)
	return resp, err
}

func (n *Server) ClosestPrecedingNode(k int64) int64 {
	fmt.Printf("n is: %v \n", n)
	fmt.Printf("len(n.fTable): %d\n", len(n.Node.FTable))

	for i := 6 - 1; i >= 0; i-- {
		if n.Node.Id < n.Node.FTable[i].Key && n.Node.FTable[i].Key < k {
			return n.Node.FTable[i].NodeAddr
		}
	}
	return n.Node.Port
}
