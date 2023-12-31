package gapi

// main function that user will call via evans
// input: requestType (GET or INSERT), key, value
// return: response

import (
	"context"
	"log"
	"strconv"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func (server *Server) ClientRequestHandler(ctx context.Context, req *pb.ClientRequest) (*pb.ClientResponse, error) {
	if (req.GetRequestType() != "GET" && req.GetRequestType() != "INSERT") || req.GetRequestedKey() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty and requestType needs to be GET or INSERT")
	}
	hashedKey := Sha1Modulo(req.GetRequestedKey(), m)
	log.Printf("%+v hashedKey in ClientRequestHandler : %d", req.GetRequestedKey(), hashedKey)
	findSuccessorMessage := &pb.FindSuccessorRequest{RequestedKey: strconv.FormatInt(hashedKey, 10)}
	successorResponse, _ := server.FindSuccessor(ctx, findSuccessorMessage)

	if req.GetRequestType() == "GET" {
		conn, err := grpc.Dial(successorResponse.SuccessorAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewChordClient(conn)
		chordResp, err := c.GetValueFromKey(ctx, &pb.GetValueFromKeyRequest{Key: req.GetRequestedKey()})
		if err != nil {
			log.Printf("could not get chord response while GetValueFromKey in ClientRequestHandler: %v", err)
		}
		resp := &pb.ClientResponse{
			RequestType:   "Successful GET",
			RequestedKey:  req.GetRequestedKey(),
			Value:         chordResp.GetValue(),
			RetrievedFrom: chordResp.GetRetrievedFrom(),
		}
		return resp, nil
	} else {
		conn, err := grpc.Dial(successorResponse.SuccessorAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewChordClient(conn)
		chordResp, err := c.InsertKeyValuePair(ctx, &pb.InsertKeyValuePairRequest{Key: req.GetRequestedKey(), Value: req.GetValue()})
		if err != nil {
			log.Printf("could not get chord response while InsertKeyValuePair in ClientRequestHandler: %v", err)
		}
		resp := &pb.ClientResponse{
			RequestType:  chordResp.GetMessage(),
			RequestedKey: req.GetRequestedKey(),
			Value:        req.GetValue(),
		}
		return resp, nil
	}
}
