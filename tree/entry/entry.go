package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	*tree.Node // Embedding
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}
	left.postOrder()
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("This mehtod is shadowed.")
}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{
		Value: 5,
		Left:  nil,
		Right: nil,
	}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Print()

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	pRoot := &root
	pRoot.SetValue(100)
	pRoot.Print()
	root.Print()

	root.Traverse()
	root.Node.Traverse()

	var nRoot *tree.Node
	nRoot.SetValue(10)

	root.postOrder()
}
