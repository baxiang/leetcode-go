package main

type Node struct {
	Val      int
	Children []*Node
}


func preorder(root *Node) []int {
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
	*res = append(*res,root.Val)
	for i:=0;i<len(root.Children);i++{
		helper(root.Children[i],res)
	}
}

func preorder1(root *Node) []int {
	if root==nil{
		return nil
	}
	stack := make([]*Node,0)
	stack = append(stack,root)
	res :=make([]int,0)
	for len(stack)>0{
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res,node.Val)
		for i:=len(node.Children)-1;i>=0;i--{
			stack = append(stack,node.Children[i])
		}
	}
	return res
}



func main() {
	
}
