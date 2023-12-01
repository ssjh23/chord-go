package gapi

import (
	"context"
	"log"
	"reflect"

	"github.com/ssjh23/chord-go/pb"
)

func (n *Server) CheckReplicateData(ctx context.Context, req *pb.CheckReplicateDataRequest) (*pb.CheckReplicateDataResponse, error) {
	// Get the id of the node in the request
	requestingNodeAddress := req.GetNodeAddress()
	// Check that the requesting node address is in the replica data list
	if _, ok := n.Node.replicaData[requestingNodeAddress]; !ok {
		// log that the replica does not exist yet
		log.Printf("Replica data for node %s does not exist yet", requestingNodeAddress)
		// Put the replica data in the replica data list
		n.Node.replicaData[requestingNodeAddress] = req.GetData()	
	}
	// If the data already exist, check that the maps are exactly the same. use reflect.DeepEqual to compare the maps
	eq := reflect.DeepEqual(n.Node.replicaData[requestingNodeAddress], req.GetData())
	if !eq {
		// log that the replica data does not match
		log.Printf("Replica data for node %s does not match", requestingNodeAddress)
		// Put the replica data in the replica data list
		n.Node.replicaData[requestingNodeAddress] = req.GetData()
	} else {
		// log that the replica data matches
		log.Printf("Replica data for node %s matches", requestingNodeAddress)
	}
	// Return the response
	resp := &pb.CheckReplicateDataResponse{
		Replicated: true,
	}
	return resp, nil
}