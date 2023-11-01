package main

import (
	"fmt"
	"math"
)

type Node struct {
	id          int
	fTable      []*Finger
	keys        []int
	successor   *Node
	predecessor *Node
}

type Finger struct {
	key  int
	node *Node
}

var m = 6

// math fn: given a node, populate its finger table
func populateFingerTable(node *Node) {
	// m := len(node.fTable.fingers)
	// fmt.Printf("%d node \n", node.id)
	for i := 0; i < m; i++ {
		k := (node.id + int(math.Pow(2, float64(i)))) % int(math.Pow(2, float64(m)))
		// fmt.Printf("k is: %v \n", k)
		node.fTable = append(node.fTable, &Finger{key: k})
	}
}

func (n *Node) findSuccessor(k int) *Node {
	// fmt.Println("starting findSuccessor of %v", n)
	if n.id == k {
		return n
	}

	if k > n.id && k <= n.successor.id {
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
		if n.id < n.fTable[i].key && n.fTable[i].key < k {
			return n.fTable[i].node
		}
	}
	return n
}
