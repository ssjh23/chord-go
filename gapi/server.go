package gapi

import (
	"fmt"

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
	data               map[string]string
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
		successorAddress:   config.SuccessorAddress,
		predecessorAddress: config.PredecessorAddress,
	}

	populateFingerTables(&node)

	return node, nil
}

func populateFingerTables(node *Node) {
	// populate finger tables
	fmt.Println("Populating finger tables...\n")
	// for i := 0; i < m; i++ {
	// k := (nodeIdentifier + int64(math.Pow(2, float64(i)))) % int64(math.Pow(2, float64(m)))
	// 	nodeIdentifier := Sha1Modulo(node.myIpAddress, m)
	// 	node.fTable = append(node.fTable, &Finger{key: k, NodeAddress: node.myIpAddress})
	// 	fmt.Printf("Finger table %d: %v \n", i, node.fTable[i])
	// }

	// hardcode finger tables
	// for i := 0; i < m; i++ {
	ip := node.myIpAddress
	if ip == "chord-go-chord1-1:9091" {
		node.fTable = append(node.fTable, &Finger{key: 7, NodeAddress: "chord-go-chord3-1:9093"})
		node.fTable = append(node.fTable, &Finger{key: 8, NodeAddress: "chord-go-chord3-1:9093"})
		node.fTable = append(node.fTable, &Finger{key: 10, NodeAddress: "chord-go-chord3-1:9093"})
		node.fTable = append(node.fTable, &Finger{key: 14, NodeAddress: "chord-go-chord2-1:9092"})
		node.fTable = append(node.fTable, &Finger{key: 22, NodeAddress: "chord-go-chord2-1:9092"})
		node.fTable = append(node.fTable, &Finger{key: 38, NodeAddress: "chord-go-chord4-1:9094"})
	} else if ip == "chord-go-chord2-1:9092" {
		node.fTable = append(node.fTable, &Finger{key: 30, NodeAddress: "chord-go-chord4-1:9094"})
		node.fTable = append(node.fTable, &Finger{key: 31, NodeAddress: "chord-go-chord4-1:9094"})
		node.fTable = append(node.fTable, &Finger{key: 33, NodeAddress: "chord-go-chord4-1:9094"})
		node.fTable = append(node.fTable, &Finger{key: 37, NodeAddress: "chord-go-chord4-1:9094"})
		node.fTable = append(node.fTable, &Finger{key: 45, NodeAddress: "chord-go-chord4-1:9094"})
		node.fTable = append(node.fTable, &Finger{key: 61, NodeAddress: "chord-go-chord1-1:9091"})
	} else if ip == "chord-go-chord3-1:9093" {
		node.fTable = append(node.fTable, &Finger{key: 12, NodeAddress: "chord-go-chord2-1:9092"})
		node.fTable = append(node.fTable, &Finger{key: 13, NodeAddress: "chord-go-chord2-1:9092"})
		node.fTable = append(node.fTable, &Finger{key: 15, NodeAddress: "chord-go-chord2-1:9092"})
		node.fTable = append(node.fTable, &Finger{key: 19, NodeAddress: "chord-go-chord2-1:9092"})
		node.fTable = append(node.fTable, &Finger{key: 27, NodeAddress: "chord-go-chord2-1:9092"})
		node.fTable = append(node.fTable, &Finger{key: 43, NodeAddress: "chord-go-chord4-1:9094"})
	} else if ip == "chord-go-chord4-1:9094" {
		node.fTable = append(node.fTable, &Finger{key: 50, NodeAddress: "chord-go-chord1-1:9091"})
		node.fTable = append(node.fTable, &Finger{key: 51, NodeAddress: "chord-go-chord1-1:9091"})
		node.fTable = append(node.fTable, &Finger{key: 53, NodeAddress: "chord-go-chord1-1:9091"})
		node.fTable = append(node.fTable, &Finger{key: 57, NodeAddress: "chord-go-chord1-1:9091"})
		node.fTable = append(node.fTable, &Finger{key: 1, NodeAddress: "chord-go-chord1-1:9091"})
		node.fTable = append(node.fTable, &Finger{key: 17, NodeAddress: "chord-go-chord3-1:9093"})
	} else {
		fmt.Println("ERROR: node.myIpAddress is not chord-go-chord1-1:9091, chord-go-chord2-1:9092, chord-go-chord3-1:9093, chord-go-chord4-1:9094")
	}

	fmt.Printf("Finger table: %v \n", node.fTable)
	// }

}
