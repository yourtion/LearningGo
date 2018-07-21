package tree

import "fmt"

func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	fmt.Print(node.Value, " ")
	node.Right.Traverse()
}
