package gapi

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ssjh23/chord-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var m = 6

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

	log.Printf("STARTING FIND SUCCESSOR: %s", n.config.ServerAddress)

	// hashing the key
	hashedKey := Sha1Modulo(req.GetRequestedKey(), m)
	myHashedIp := Sha1Modulo(n.Node.myIpAddress, m)
	mySuccessorHashedIp := Sha1Modulo(n.Node.successorAddress, m)
	myPredecessorHashedIp := Sha1Modulo(n.Node.predecessorAddress, m)

	// if my predecessor is higher than me, check if hashed key is between my predecessor to 2**m, or 0 and me, then i'm the successor
	if myPredecessorHashedIp > myHashedIp {
		if (hashedKey > myPredecessorHashedIp && hashedKey <= int64(math.Pow(float64(2), float64(m)))) || (hashedKey >= 0 && hashedKey <= myHashedIp) {
			resp := &pb.FindSuccessorResponse{
				SuccessorAddress: n.Node.myIpAddress,
			}
			log.Printf("YOUR KEY IS LOCATED AT PORT: %s", resp.SuccessorAddress)
			return resp, nil
		}
	}

	// if the key is between my predecessor and me, then I am the successor
	if hashedKey > myPredecessorHashedIp && hashedKey <= myHashedIp {
		// if myPredecessorHashedIp < hashedKey && hashedKey <= myHashedIp {
		resp := &pb.FindSuccessorResponse{
			SuccessorAddress: n.Node.myIpAddress,
		}
		log.Printf("YOUR KEY IS LOCATED AT PORT: %s", resp.SuccessorAddress)
		return resp, nil
	}
	// if the key is between my sucessor and me, then my successor the successor
	if hashedKey >= myHashedIp && hashedKey < mySuccessorHashedIp {
		// if myPredecessorHashedIp < hashedKey && hashedKey <= myHashedIp {
		resp := &pb.FindSuccessorResponse{
			SuccessorAddress: n.Node.successorAddress,
		}
		log.Printf("YOUR KEY IS LOCATED AT PORT: %s", resp.SuccessorAddress)
		return resp, nil
	}
	// else, find the closest preceding node
	nextNodeAddress := n.ClosestPrecedingNode(hashedKey)

	conn, err := grpc.Dial(nextNodeAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	nextNode := pb.NewChordClient(conn)
	nextNodeResp, err := nextNode.FindSuccessor(ctx, &pb.FindSuccessorRequest{RequestedKey: req.GetRequestedKey()})

	if err != nil {
		log.Fatalf("could not get chord response: %v", err)
	}
	resp := &pb.FindSuccessorResponse{
		SuccessorAddress: nextNodeResp.SuccessorAddress,
	}
	log.Printf("The successor is: %s", nextNodeResp.SuccessorAddress)
	return resp, err
}

func (n *Server) ClosestPrecedingNode(hashedKey int64) string {
	fmt.Printf("n is: %v \n", n)
	fmt.Printf("len(n.fTable): %d\n", len(n.Node.fTable))

	// myHashedIp := Sha1Modulo(n.Node.myIpAddress, m)

	for i := 6 - 1; i >= 0; i-- {
		if n.Node.fTable[i].key < hashedKey {
			// if myHashedIp < n.Node.fTable[i].key && n.Node.fTable[i].key < hashedKey {
			fmt.Println("INSIDE IF OF CLOSEST PRECEDING NODE")
			return n.Node.fTable[i].NodeAddress
		}
	}
	return n.Node.fTable[m-1].NodeAddress
}
