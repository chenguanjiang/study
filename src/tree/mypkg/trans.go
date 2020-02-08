package mypkg

func (node *NodeTree) Trans(){
	if node==nil{
		return
	}
	node.Left.Trans()
	node.Print()
	node.Right.Trans()
}
