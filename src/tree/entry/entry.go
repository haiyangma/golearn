package main

import (
		"fmt"
		"tree"

	"golang.org/x/tools/container/intsets"
)


func main() {
	var root tree.Node

	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5,nil,nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetVaue(4)

	root.Traverse()

	var count int
	root.TraverseFunc(func(node *tree.Node) {
		count++;
	})
	fmt.Println("node count :",count)

}


func testSparse(){
	s:= intsets.Sparse{}
	s.Insert(1)
	s.Insert(10)
	s.Insert(1000)
	fmt.Println(s.Has(1))
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(100))

}
