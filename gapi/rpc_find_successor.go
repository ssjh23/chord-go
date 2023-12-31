package gapi

// function that is part of ClientRequestHandler
// returns the successor of the key
// it is also used by the FixFingerTable function to update the finger table entries

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"

	"github.com/ssjh23/chord-go/constant"
	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var m = constant.VALUE_OF_M

// function to hash the input string using SHA-1 and return the modulo
func Sha1Modulo(inputString string, m int) int64 {
	// Hash the input string using SHA-1
	hash := sha1.New()
	hash.Write([]byte(inputString))
	hashBytes := hash.Sum(nil)

	// Convert the hashed bytes to a big integer
	hashValue := new(big.Int).SetBytes(hashBytes)

	// Set divisor
	divisor := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(m)), nil)

	// Get the modulo
	modulo := new(big.Int).Mod(hashValue, divisor)
	return modulo.Int64()
}

func (n *Server) FindSuccessor(ctx context.Context, req *pb.FindSuccessorRequest) (*pb.FindSuccessorResponse, error) {

	// hashing the key
	hashedKey, _ := strconv.ParseInt(req.GetRequestedKey(), 10, 64)
	myHashedIp := Sha1Modulo(n.Node.myIpAddress, m)
	mySuccessorHashedIp := Sha1Modulo(n.Node.successorAddress, m)
	myPredecessorHashedIp := Sha1Modulo(n.Node.predecessorAddress, m)

	// if i am the only one in the network, return me
	if myPredecessorHashedIp == myHashedIp && mySuccessorHashedIp == myHashedIp {
		resp := &pb.FindSuccessorResponse{
			SuccessorAddress: n.Node.myIpAddress,
		}
		return resp, nil
	}

	// if my predecessor is higher than me, check if hashed key is between my predecessor to 2**m, or 0 and me, then i'm the successor
	if myPredecessorHashedIp > myHashedIp {
		if (hashedKey > myPredecessorHashedIp && hashedKey <= int64(math.Pow(float64(2), float64(m)))) || (hashedKey >= 0 && hashedKey <= myHashedIp) {
			resp := &pb.FindSuccessorResponse{
				SuccessorAddress: n.Node.myIpAddress,
			}
			return resp, nil
		}
	}

	// if my successor is lower than me, check if hashed key is between me to 2**m or 0 and my successor, then my successor is the successor
	if mySuccessorHashedIp < myHashedIp {
		if (hashedKey > myHashedIp && hashedKey <= int64(math.Pow(float64(2), float64(m)))) || (hashedKey >= 0 && hashedKey <= mySuccessorHashedIp) {
			resp := &pb.FindSuccessorResponse{
				SuccessorAddress: n.Node.successorAddress,
			}
			return resp, nil
		}
	}

	// if the key is between my predecessor and me, then I am the successor
	if hashedKey > myPredecessorHashedIp && hashedKey <= myHashedIp {
		resp := &pb.FindSuccessorResponse{
			SuccessorAddress: n.Node.myIpAddress,
		}
		return resp, nil
	}
	// if the key is between my sucessor and me, then my successor the successor
	if hashedKey >= myHashedIp && hashedKey < mySuccessorHashedIp {
		resp := &pb.FindSuccessorResponse{
			SuccessorAddress: n.Node.successorAddress,
		}
		return resp, nil
	}

	// else, find the closest preceding node
	nextNodeAddress := n.ClosestPrecedingNode(hashedKey)

	// if by now the closest preceding node is me, my fTable is wrong.
	if nextNodeAddress == n.Node.myIpAddress {
		conn, err := grpc.Dial(n.Node.successorAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("did not connect: %v", err)
		}
		defer conn.Close()
		nextNode := pb.NewChordClient(conn)
		nextNodeResp, err := nextNode.FindSuccessor(ctx, &pb.FindSuccessorRequest{RequestedKey: req.GetRequestedKey()})
		if err != nil {
			log.Printf("could not get chord response while calling FindSuccessor in FindSuccessor: %v", err)
		}
		resp := &pb.FindSuccessorResponse{
			SuccessorAddress: nextNodeResp.SuccessorAddress,
		}
		return resp, err
	}

	conn, err := grpc.Dial(nextNodeAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	nextNode := pb.NewChordClient(conn)
	nextNodeResp, err := nextNode.FindSuccessor(ctx, &pb.FindSuccessorRequest{RequestedKey: req.GetRequestedKey()})

	if err != nil {
		log.Printf("could not get chord response while nextNode.FindSuccessor in FindSuccessor: %v", err)
	}
	resp := &pb.FindSuccessorResponse{
		SuccessorAddress: nextNodeResp.SuccessorAddress,
	}
	return resp, err
}

func (n *Server) ClosestPrecedingNode(hashedKey int64) string {
	fmt.Printf("n is: %v \n", n)
	fmt.Printf("len(n.fTable): %d\n", len(n.Node.fTable))

	for i := 6 - 1; i >= 0; i-- {
		if n.Node.fTable[i].key < hashedKey {
			fmt.Println("INSIDE IF OF CLOSEST PRECEDING NODE")
			return n.Node.fTable[i].NodeAddress
		}
	}
	return n.Node.fTable[m-1].NodeAddress
}
