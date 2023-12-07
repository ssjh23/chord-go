package gapi

import (
	"fmt"
	"math"

	"github.com/ssjh23/chord-go/pb"
	"github.com/ssjh23/chord-go/util"
)

// Server serves gRPC requests as Node in chord
type Server struct {
	pb.UnimplementedChordServer
	config util.Config
	Node
}

type Node struct {
	myIpAddress        string
	fTable             []*Finger
	successorAddress   string
	predecessorAddress string
	successorList      []string
	data               map[string]string
	replicaData        map[string]map[string]string
}

type Finger struct {
	key         int64
	NodeAddress string
}

// Newserver creates a new gRPC server
func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}
	return server, nil
}

func InitNode(config util.Config) (Node, error) {
	// create a new node
	fmt.Println("Initialising a new node...\n")
	node := Node{
		myIpAddress:        config.ServerAddress,
		fTable:             []*Finger{},
		data:               map[string]string{},
		successorAddress:   config.ServerAddress,
		predecessorAddress: config.ServerAddress,
		successorList:      []string{},
		replicaData:        map[string]map[string]string{},
	}

	populateFingerTables(&node)

	return node, nil
}

func populateFingerTables(node *Node) {
	// populate finger tables
	fmt.Println("Populating finger tables...\n")
	for i := 0; i < m; i++ {
		myHashedIp := Sha1Modulo(node.myIpAddress, m)
		k := (myHashedIp + int64(math.Pow(2, float64(i)))) % int64(math.Pow(2, float64(m)))
		node.fTable = append(node.fTable, &Finger{key: k, NodeAddress: node.myIpAddress})
		fmt.Printf("Finger table %d: %v \n", i, node.fTable[i])
	}
}
