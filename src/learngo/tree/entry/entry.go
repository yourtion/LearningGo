package main

import (
	"fmt"
	"learngo/tree"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)

	root.Print()

	root.Right.Left.SetValueWrong(4)
	root.Right.Left.Print() // still 0

	root.Right.Left.SetValue(4)
	root.Right.Left.Print() // 4 now

	root.Traverse()
	fmt.Println()

	// nil 也可以调用方法
	var pRoot *tree.Node
	// error
	pRoot.SetValue(200)
	pRoot = &root
	pRoot.SetValue(300)
	pRoot.Print()
	pRoot.SetValue(3)

}
