package gapi

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (n *Server) Notify(ctx context.Context, req *pb.NotifyRequest) (*pb.NotifyResponse, error) {
	flag.Parse()
	fmt.Printf("%s NOTIFYING %s", req.IpAddress, n.myIpAddress)
	if req.GetIpAddress() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty")
	} else {
		m := 6
		myHashedIp := Sha1Modulo(n.Node.myIpAddress, m)
		myPredecessorHashedIp := Sha1Modulo(n.Node.predecessorAddress, m)
		new_predecessorHashedIp := Sha1Modulo(req.GetIpAddress(), m)

		if myPredecessorHashedIp < new_predecessorHashedIp && new_predecessorHashedIp < myHashedIp {
			n.Node.predecessorAddress = req.GetIpAddress()
			log.Printf("Predecessor Updated: %s\n", n.Node.predecessorAddress)
		}

		// handling edge case: in between "start" & "end" of ring
		if myPredecessorHashedIp > myHashedIp {
			if (new_predecessorHashedIp > myPredecessorHashedIp && new_predecessorHashedIp <= int64(math.Pow(float64(2), float64(m)))) || (new_predecessorHashedIp >= 0 && new_predecessorHashedIp < myHashedIp) {
				n.Node.predecessorAddress = req.GetIpAddress()
				log.Printf("Predecessor Updated: %s\n", n.Node.predecessorAddress)
			}
		}

		// handling case for new nodes
		if myPredecessorHashedIp == myHashedIp {
			n.Node.predecessorAddress = req.GetIpAddress()
			log.Printf("Predecessor Updated: %s\n", n.Node.predecessorAddress)
		}
	}
	resp := &pb.NotifyResponse{
		PredecessorAddress: n.Node.predecessorAddress,
	}
	return resp, nil
}
