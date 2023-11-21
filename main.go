package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ssjh23/chord-go/gapi"
	"github.com/ssjh23/chord-go/pb"
	"github.com/ssjh23/chord-go/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// runGrpcServer(config)
	joiningNetwork(config)
}

func initNode(config util.Config, s *grpc.Server) {
	chordID := config.ChordId
	var num int64
	_, err := fmt.Sscanf(chordID, "%d", &num)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	node := &pb.Node{Port: num}
	log.Printf("My ID is %s and my node is: %v", chordID, node)
}

func runGrpcServer(config util.Config) {

	server, err := gapi.NewServer(config)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	grpcServer := grpc.NewServer()

	// call initNode fn
	initNode(config, grpcServer)

	pb.RegisterChordServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func joiningNetwork(config util.Config) {
	fmt.Println("i got here")
	bootstrap_address := config.BootstrapNode
	fmt.Println(bootstrap_address)
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChordClient(conn)

	response, err := client.FindSuccessor(context.Background(), &pb.FindSuccessorRequest{ClientAddr: "127.0.0.1:9092", RequestedKey: 2})
	if err != nil {
		log.Fatalf("Error calling FindSuccessor: %v", err)
	}

	fmt.Printf("Received response from successor: %+v\n", response)

	// log.Printf("Connected to %+v\n", conn)

	// client := pb.NewChordServiceClient(conn)

	// // Implement logic to get information about the Chord ring
	// nodeInfo := &pb.NodeInfo{} // Replace with actual information

	// response, err := client.FindSuccessor(context.Background(), nodeInfo)
	// if err != nil {
	// 	log.Fatalf("Error calling FindSuccessor: %v", err)
	// }

	// // Process the response and continue joining the Chord ring
	// fmt.Printf("Received response from successor: %+v\n", response)
}
