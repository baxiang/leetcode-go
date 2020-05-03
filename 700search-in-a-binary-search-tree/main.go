package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST1(root *TreeNode, val int) *TreeNode {
	if root==nil{
		return nil
	}
	if root.Val==val{
		return root
	}else if root.Val>val{
	     return searchBST(root.Left,val)
	}else {
		return searchBST(root.Right,val)
	}
}


func searchBST(root *TreeNode, val int) *TreeNode {
	stack := make([]*TreeNode,0)
	if root.Val==val{
		return root
	}else if root.Val>val&&root.Left!=nil{
		stack = append(stack,root.Left)
	}else if root.Val<val&&root.Right!=nil{
		stack = append(stack,root.Right)
	}
	for len(stack)>0{
		node :=stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Val==val{
			return node
		}else if node.Val>val&&node.Left!=nil{
			stack = append(stack,node.Left)
		}else if node.Val<val&&node.Right!=nil{
			stack = append(stack,node.Right)
		}
	}
	return nil
}



func main() {
	four :=&TreeNode{Val:   4,}
	two :=&TreeNode{Val:   2,}
	seven :=&TreeNode{Val:   7,}
	one :=&TreeNode{Val:   1,}
	three :=&TreeNode{Val:   3,}
	four.Left = two
	four.Right = seven
	two.Left =one
	two.Right = three
	r :=searchBST1(four,7)
	if r!=nil{
		fmt.Println(r.Val)
	}
}
