package mypkg

import (
	"fmt"
)

func CreateNode(value int) *NodeTree {
	return &NodeTree{Value: value}
}

func (node NodeTree) Print(){
	fmt.Println(node.Value)
}

func (node *NodeTree) Setvalue(value int){
	if node==nil{
		fmt.Println("setting value to nil node,Ignored")
		return
	}
	node.Value =value
}

