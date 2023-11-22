package gapi

import (
	"context"
	"flag"
	"log"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (n *Server) Notify(ctx context.Context, req *pb.NotifyRequest) (*pb.NotifyResponse, error) {
	flag.Parse()

	if req.GetIpAddress() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty")
	} else {
		m := 6
		myHashedIp := Sha1Modulo(n.Node.myIpAddress, m)
		predecessoressorHashedIP := Sha1Modulo(n.Node.predecessorAddress, m)
		new_predecessorHashedIp := Sha1Modulo(req.GetIpAddress(), m)

		if n.Node.predecessorAddress == "nil" || (predecessoressorHashedIP < new_predecessorHashedIp && new_predecessorHashedIp < myHashedIp) {
			n.Node.predecessorAddress = req.GetIpAddress()
			log.Printf("Predecessor Updated: %s\n", n.Node.predecessorAddress)
		}
	}
	resp := &pb.NotifyResponse{
		PredecessorAddress: n.Node.predecessorAddress,
	}
	return resp, nil
}
