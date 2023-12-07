package main

import (
	"crypto/sha1"
	"fmt"
	"math"
	"math/big"

	"github.com/ssjh23/chord-go/constant"
)

type Node struct {
	node_identifier int
	ip_address      string
	fTable          []*Finger
	keys            []int
	successor       *Node
	predecessor     *Node
}

type Finger struct {
	key  int
	node *Node
}

var m = constant.VALUE_OF_M

// math fn: given a string, hash it using SHA-1 and return the modulo of the hash by 2^m
// we use this function to hash both the keys and the ip addresses of the nodes
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

// math fn: given a node, populate its finger table
func populateFingerTable(node *Node) {
	// m := len(node.fTable.fingers)
	// fmt.Printf("%d node \n", node.id)
	for i := 0; i < m; i++ {
		k := (node.node_identifier + int(math.Pow(2, float64(i)))) % int(math.Pow(2, float64(m)))
		// fmt.Printf("k is: %v \n", k)
		node.fTable = append(node.fTable, &Finger{key: k})
	}
}

func (n *Node) findSuccessor(k int) *Node {
	// fmt.Println("starting findSuccessor of %v", n)
	if n.node_identifier == k {
		return n
	}

	if k > n.node_identifier && k <= n.successor.node_identifier {
		fmt.Println("inside if of findSuccessor")
		return n.successor
	}
	//  else
	fmt.Println("else of findSuccessor")
	next_node := n.ClosestPrecedingNode(k)

	for _, key := range next_node.keys {
		if key == k {
			return next_node
		}
	}

	// time.Sleep(1 * time.Second)
	fmt.Printf("next_node is: %v \n", next_node)
	return next_node.findSuccessor(k)
}

func (n *Node) ClosestPrecedingNode(k int) *Node {
	// fmt.Printf("n is: %v \n", n)
	// fmt.Printf("len(n.fTable): %d\n", len(n.fTable))

	for i := m - 1; i >= 0; i-- {
		if n.node_identifier < n.fTable[i].key && n.fTable[i].key < k {
			return n.fTable[i].node
		}
	}
	return n
}

// fix fingers function to be called periodically
func (n *Node) fixFingers() {
	for i := 0; i < m; i++ {
		k := (n.node_identifier + int(math.Pow(2, float64(i)))) % int(math.Pow(2, float64(m)))
		n.fTable[i].node = n.findSuccessor(k)
	}
	fmt.Println("Fingers fixed!")
}
