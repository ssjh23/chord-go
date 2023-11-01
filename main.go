package main

import (
	"log"
	"net"
	"time"

	"github.com/ssjh23/chord-go/gapi"
	"github.com/ssjh23/chord-go/pb"
	"github.com/ssjh23/chord-go/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	runGrpcServer(config)

	time.Sleep(10 * time.Second)

}

func runGrpcServer(config util.Config) {
	server, err := gapi.NewServer(config)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterChordServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Start gRPC server on %s", listener.Addr().String())
	// go func() {
	// 	for {
	// 		time.Sleep(5 * time.Second) // Adjust the interval based on your requirements
	// 		updateSuccessorAddress(config)
	// 	}
	// }()
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

// func joinChordRing(config util.Config, existingNodeAddress string){
// 	// Make an RPC call to the existing node to get its successor
// 	// Set the successor of the new node to the successor of the existing node

// 	log.Printf("Node address ___ joined the Chord ring with successor at %s\n", existingNodeAddress)
// 	return updatedConfig
// }

func updateSuccessorAddress(config util.Config) {
	// Implement the logic to dynamically obtain the successor address
	// and update the config.SuccessorAddress
	// ...

	log.Printf("Updated successor address to %s", config.SuccessorAddress)
}
