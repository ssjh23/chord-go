package gapi

import (
	"github.com/ssjh23/chord-go/pb"
	"github.com/ssjh23/chord-go/util"
)

// Server serves gRPC requests as Node in chord
type Server struct {
	pb.UnimplementedChordServer
	config util.Config
}

// Newserver creates a new gRPC server
func NewServer(config util.Config)(*Server, error){
	server := &Server{
		config: config,
	}
	return server, nil
}