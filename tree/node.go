package tree

import "fmt"

// Node ...
type Node struct {
	Value       int
	Left, Right *Node
}

// Print ...
func (node Node) Print() {
	fmt.Println(node.Value)
}

// SetValue ...
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("node is nil")
		return
	}
	node.Value = value
}

// CreateNode ...
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
