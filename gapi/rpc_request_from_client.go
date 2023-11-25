package gapi

import (
	"context"
	"log"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func (server *Server) RequestFromClient(ctx context.Context, req *pb.GetRequestFromClient) (*pb.GetResponseToClient, error) {
	if req.GetInput() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty")
	}
	log.Printf("Server Address: %s", server.Node.myIpAddress)
	log.Printf("Successor Address: %s", server.Node.successorAddress)
	conn, err := grpc.Dial(server.config.SuccessorAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	chordResp, err := c.GetChordNode(ctx, &pb.GetChordNodeRequest{Id: req.GetInput()})
	if err != nil {
		log.Fatalf("could not get chord response: %v", err)
	}
	resp := &pb.GetResponseToClient{
		InputFromClient: req.GetInput(),
		Id:              chordResp.GetId(),
		Ip:              chordResp.GetIp(),
		Port:            chordResp.GetPort(),
	}
	return resp, nil
}
