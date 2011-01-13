package main

import(
	"fmt"
)

type Node struct {
	label string
	neighbors []*Node
}

// Just for fmt.printf :)
func (n *Node) String() string {
	return n.label
}

// If the Node is a Part Of Speech
type PartOfSpeech interface {
	Label() string
	Arity() int
}

func (n *Node) Label() string {
	return n.label
}

func (n *Node) Arity() int {
	return len(n.neighbors)
}

// If it's a term
type Term interface {
	Label() string
	Definition() string
}

func (n *Node) Definition() string {
	def := ""
	if len(n.neighbors) == 0 {
		def = "Null"
	} else {
		for _, neighbor := range n.neighbors {
			def += neighbor.Label()
			def += " "
		}
	}
	return def
}

// Noodle around with them
func main() {
	node1 := &Node{ label: "foo" }
	node2 := &Node{ label: "bar" }
	node3 := &Node{ label: "baz", neighbors: []*Node{node1,node2} }
	fmt.Printf("Node 1: %s\tDefinition: %s\t Arity: %d\n", node1, node1.Definition(), node1.Arity())
	fmt.Printf("Node 3: %s\tDefinition: %s\t Arity: %d\n", node3, node3.Definition(), node3.Arity())
}
