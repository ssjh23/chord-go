package main

import (
	"log"
	"net"
	"os"
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
	go func() {
		for {
			time.Sleep(5 * time.Second) // Adjust the interval based on your requirements
			changeSuccessorAddress(config, "new-successor-address")
		}
	}()
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func changeSuccessorAddress(config util.Config, newSuccessorAddress string) {
	// Implement the logic to dynamically obtain the successor address
	// and update the config.SuccessorAddress
	if config.ChordId == "5" || config.ChordId == "6" { //cords that do not have a successor_address
		// Update the environment variable within the code
		os.Setenv("SUCCESSOR_ADDRESS", newSuccessorAddress)

		// Also, update the config if needed
		config.SuccessorAddress = os.Getenv("SUCCESSOR_ADDRESS")
	}

	log.Printf("I am chord %s and my address is %s, my successor is %s", config.ChordId, config.ServerAddress, config.SuccessorAddress)
}
