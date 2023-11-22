package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ssjh23/chord-go/gapi"
	"github.com/ssjh23/chord-go/pb"
	"github.com/ssjh23/chord-go/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("starting Main fn!")
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	runGrpcServer(config)
}

func runGrpcServer(config util.Config) {
	server, err := gapi.NewServer(config)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	grpcServer := grpc.NewServer()

	node, _ := gapi.InitNode(config)
	server.Node = node

	pb.RegisterChordServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Start gRPC server on %s\n", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
