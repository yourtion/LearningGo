package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

// Go 语言都是值传递（只有值传递）
func (node treeNode) setValueWrong(value int) {
	node.value = value
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}

	node.left.traverse()
	fmt.Print(node.value, " ")
	node.right.traverse()
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)

	root.print()

	root.right.left.setValueWrong(4)
	root.right.left.print() // still 0

	root.right.left.setValue(4)
	root.right.left.print() // 4 now

	root.traverse()
	fmt.Println()

	// nil 也可以调用方法
	var pRoot *treeNode
	// error
	pRoot.setValue(200)
	pRoot = &root
	pRoot.setValue(300)
	pRoot.print()
	pRoot.setValue(3)

}
