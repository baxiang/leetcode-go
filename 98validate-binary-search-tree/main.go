package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
 }

func isValidBST1(root *TreeNode) bool {
	if root==nil{
		return true
	}
	stack := make([]*TreeNode,0)
	preVal :=math.MinInt64
	for root!=nil||len(stack)>0{
		for root!=nil{
			stack = append(stack,root)
			root =root.Left
		}
		node := stack[len(stack)-1]
		stack =stack[:len(stack)-1]
		if node.Val<=preVal {
			return false
		}
		preVal = node.Val
		root =node.Right
	}
	return true
}

func isValidBST(root *TreeNode) bool {
	if root==nil{
		return true
	}
	return helper(root,math.MinInt64,math.MaxInt64)
}

func helper(root *TreeNode,left,right int)bool{
	if root==nil{
		return true
	}
	if root.Val<=left||root.Val>=right{
		return false
	}
	return helper(root.Left,left,root.Val)&&helper(root.Right,root.Val,right)
}



func main() {
	t :=&TreeNode{Val: 1}
	one :=&TreeNode{Val: 1}
	//three :=&TreeNode{Val: 3}
	t.Left = one
	//t.Right = three
	fmt.Println(isValidBST(t))
}
