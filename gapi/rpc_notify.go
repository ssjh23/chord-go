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
	data_to_be_sent_back := make(map[string]string)
	flag.Parse()
	fmt.Printf("%s NOTIFYING %s", req.IpAddress, n.myIpAddress)
	if req.GetIpAddress() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty")
	} else {
		m := 6
		myHashedIp := Sha1Modulo(n.Node.myIpAddress, m)
		myPredecessorHashedIp := Sha1Modulo(n.Node.predecessorAddress, m)
		new_predecessorHashedIp := Sha1Modulo(req.GetIpAddress(), m)

		// if my new predecessor is in between my current predecessor and me, update my predecessor and send some of my data to my new predecessor
		if myPredecessorHashedIp < new_predecessorHashedIp && new_predecessorHashedIp < myHashedIp {
			n.Node.predecessorAddress = req.GetIpAddress()
			// send some of my data to my new predecessor by checking which data is in between my new predecessor and me
			for key, value := range n.Node.data {
				hashedKey := Sha1Modulo(key, m)
				if hashedKey > myPredecessorHashedIp && hashedKey <= new_predecessorHashedIp {
					// send data back to my new predecessor as response
					data_to_be_sent_back[key] = value
					delete(n.Node.data, key)
				}
			}
			log.Printf("Predecessor Updated: %s\n", n.Node.predecessorAddress)
		}

		// handling edge case: in between "start" & "end" of ring
		if myPredecessorHashedIp > myHashedIp {
			if (new_predecessorHashedIp > myPredecessorHashedIp && new_predecessorHashedIp <= int64(math.Pow(float64(2), float64(m)))) || (new_predecessorHashedIp >= 0 && new_predecessorHashedIp < myHashedIp) {
				n.Node.predecessorAddress = req.GetIpAddress()
				for key, value := range n.Node.data {
					hashedKey := Sha1Modulo(key, m)
					if hashedKey > myPredecessorHashedIp && hashedKey <= new_predecessorHashedIp {
						// send data back to my new predecessor as response
						data_to_be_sent_back[key] = value
						delete(n.Node.data, key)
					}
				}
			}
			log.Printf("Predecessor Updated: %s\n", n.Node.predecessorAddress)
		}

		// handling case for new nodes
		if myPredecessorHashedIp == myHashedIp {
			n.Node.predecessorAddress = req.GetIpAddress()
			log.Printf("Predecessor Updated: %s\n", n.Node.predecessorAddress)
		}
	}
	resp := &pb.NotifyResponse{
		PredecessorAddress: n.Node.predecessorAddress,
		DataToBeAbsorbed:   data_to_be_sent_back,
	}
	return resp, nil
}
