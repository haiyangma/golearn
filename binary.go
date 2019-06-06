package main

import (
	"container/list"
)

// Binary Tree
type BinaryTree struct {
	Data  interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

// Constructor
func NewBinaryTree(data interface{}) *BinaryTree {
	return &BinaryTree{Data: data}
}

// 先序遍历-非递归
func (bt *BinaryTree) PreOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	for t != nil || stack.Len() != 0 {
		for t != nil {
			res = append(res, t.Data) //visit
			stack.PushBack(t)
			t = t.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*BinaryTree)
			t = t.Right
			stack.Remove(v)
		}
	}
	return res
}

// 中序遍历-非递归
func (bt *BinaryTree) InOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*BinaryTree)
			res = append(res, t.Data) //visit
			t = t.Right
			stack.Remove(v)
		}
	}
	return res
}

// 后序遍历-非递归
func (bt *BinaryTree) PostOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	var preVisited *BinaryTree

	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}

		v := stack.Back()
		top := v.Value.(*BinaryTree)

		if (top.Left == nil && top.Right == nil) || (top.Right == nil && preVisited == top.Left) || preVisited == top.Right {
			res = append(res, top.Data) //visit
			preVisited = top
			stack.Remove(v)
		} else {
			t = top.Right
		}
	}
	return res
}

func pinck(x int32) int64 {
	if x == 10 {
		return 1
	}
	return 2 * (pinck(x+1) + 1)
}

type Node struct {
	Index int32
	Value int32
	Next *Node
}

func makeChain(m int32) *Node {
	start := &Node{0,0,nil}
	current := start
	for i := int32(0); i < m; i++ {
		v := i%3
		current.Next = &Node{i,v + 1,nil}
		current = current.Next
	}
	current.Next = start.Next
	return start.Next
}

//func removeNum(n Node) Node {
//	for n.Next != nil {
//
//	}
//}

func main() {
	//t := NewBinaryTree(1)
	//t.Left  = NewBinaryTree(3)
	//t.Right = NewBinaryTree(6)
	//t.Left.Left = NewBinaryTree(4)
	//t.Left.Right = NewBinaryTree(5)
	//t.Left.Left.Left = NewBinaryTree(7)
	//t.Right.Left = NewBinaryTree(8)
	//t.Right.Right = NewBinaryTree(9)
	//t.Right.Right.Left = NewBinaryTree(10)
	//
	//fmt.Println(t.PreOrderNoRecursion())
	//fmt.Println(t.InOrderNoRecursion())
	//fmt.Println(t.PostOrderNoRecursion())

	//fmt.Println(pinck(1))
	n := makeChain(3)
	c := n
	i := 1
	for {
		if c.Next.Index != c.Index {
			c = c.Next
			i++
			if i==3 {
				c.Next = c.Next.Next
				i=1
			}
		} else {
			break
		}
	}
	println(c.Index)

}
