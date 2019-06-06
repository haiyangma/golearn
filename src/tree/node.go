package tree

import "fmt"

type Node struct {
	Value int
	Left,Right *Node
}

func (node Node) Print(){
	fmt.Println(node.Value)
}

func (node *Node) SetVaue(v int){
	if node == nil {
		fmt.Println("Setting value to nil nood, ignored !")
		return
	}
	node.Value = v
}


func CreateNode(value int) *Node{
	return &Node{Value:value}
}

type MyTreeNode struct{
	node *Node
}

func (myTreeNode *MyTreeNode) postOrder(){
	if myTreeNode == nil || myTreeNode.node == nil{
		return
	}
	left := MyTreeNode{myTreeNode.node.Left}
	right := MyTreeNode{myTreeNode.node.Right}
	left.postOrder()
	right.postOrder()
	self := myTreeNode.node;
	self.Print()

}

