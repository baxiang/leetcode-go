package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


// 中序遍历
func kthSmallest(root *TreeNode, k int) int {
     if root==nil{
     	return 0
	 }
	 stack := make([]*TreeNode,0)
	 i :=0
	 for root!=nil||len(stack)>0{
		 if root!= nil {
			stack = append(stack,root)
			root = root.Left
		 }else {
		 	curr :=stack[len(stack)-1]
		 	stack = stack[:len(stack)-1]
		 	// 数据出站
			 if i==k-1  {
				 return curr.Val
			 }
			 i++
		 	root = curr.Right
		 }
	 }
	 return 0
}


func kthSmallest1(root *TreeNode, k int) int{
	//index = k
	dfs(root,k)
	return res

}
var res int
var index int
func dfs(root *TreeNode,k int){
	if root==nil {
		return
	}
	dfs(root.Left,k)
	index++
	if index==k {
		res = root.Val
		return
	}
	dfs(root.Right,k)
}

func main() {
	a := &TreeNode{
		Val: 3,
	}
	b := &TreeNode{
		Val: 1,
	}
	a.Left = b
	c := &TreeNode{
		Val: 4,
	}
	a.Right = c
	d := &TreeNode{
		Val: 2,
	}
	b.Right = d
	fmt.Println(kthSmallest1(a,1))
}
