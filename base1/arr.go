package main

import "fmt"

type treeNode struct {
	data int64
	leftChild * treeNode
	rightChild * treeNode
}

func main() {
	arr1 := [5]int{0,0,0,0,0}
	fmt.Print(arr1)
	node :=treeNode{}
	fmt.Print(node)
}


