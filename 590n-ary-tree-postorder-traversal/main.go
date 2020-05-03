package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root==nil{
		return nil
	}
	res :=make([]int,0)
	helper(root,&res)
	return res
}

func helper(root *Node,res *[]int){
	if root==nil{
		return
	}
	for _,v:=range root.Children{
		helper(v,res)
	}
	*res = append(*res,root.Val)
}


func postorder1(root *Node) []int {
	if root==nil{
		return nil
	}
	stack := make([]*Node,0)
	stack = append(stack,root)
	res := make([]int,0)
	for len(stack)>0{
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{node.Val},res...)
		for i:=0;i<len(node.Children);i++{
			stack = append(stack,node.Children[i])
		}
	}
	return res
}

func main() {
	
}
