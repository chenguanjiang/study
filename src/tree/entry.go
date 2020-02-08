package main

import (
	"fmt"
	"tree/mypkg"
)

type myNode struct {
	treenode *mypkg.NodeTree
}

func (node *myNode) postOder(){
	if node==nil || node.treenode==nil{
		fmt.Println("setting value to nil node,Ignored")
		return
	}
	left:=myNode{node.treenode.Left}
	left.postOder()
	right:=myNode{node.treenode.Right}
	right.postOder()
	node.treenode.Print()
}

func main() {
	var  root mypkg.NodeTree = mypkg.NodeTree{Value: 3}
	root.Print()
	fmt.Println(root)
	root.Left =&mypkg.NodeTree{}
	root.Right = mypkg.CreateNode(1)
	root.Left.Right = mypkg.CreateNode(2)
	root.Print()
	root.Setvalue(100)
	root.Print()
	proot:=&root
	proot.Setvalue(200)
	proot.Print()
	var qroot *mypkg.NodeTree
	qroot.Setvalue(1)
	fmt.Println("================================")
	root.Trans()
	fmt.Println("================================")
	mynode:=myNode{&root}
	mynode.postOder()
}