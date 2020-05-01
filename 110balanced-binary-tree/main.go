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

func isBalanced(root *TreeNode) bool {
	if root ==nil {
		return true
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	return math.Abs(float64(l)-float64(r)) <=1&&isBalanced(root.Right)&&isBalanced(root.Left)
}
func dfs(root *TreeNode)int{
	if root ==nil {
		return 0
	}
	return max(dfs(root.Left),dfs(root.Right))+1
}
func max(a ,b int)int{
	if a>b {
		return a
	}
	return b
}

func main() {
	a := &TreeNode{
		Val: 3,
	}
	b := &TreeNode{
		Val: 9,
	}
	a.Left = b
	c := &TreeNode{
		Val: 20,
	}
	a.Right = c
	d := &TreeNode{
		Val: 15,
	}
	c.Left = d
	e := &TreeNode{
		Val: 7,
	}
	c.Right = e

	fmt.Println(isBalanced(a))
	// 4,2,5,1,6,3


	a1 := &TreeNode{
		Val: 1,
	}
	b1 := &TreeNode{
		Val: 2,
	}
	a1.Left = b1
	c1 := &TreeNode{
		Val: 2,
	}
	a1.Right = c1
	d1 := &TreeNode{
		Val: 3,
	}
	c1.Left = d1
	e1 := &TreeNode{
		Val: 3,
	}
	c1.Right = e1
	f1 := &TreeNode{
		Val: 4,
	}
	e1.Left = f1
	g1 := &TreeNode{
		Val: 4,
	}
	e1.Right = g1
	fmt.Println(isBalanced(a1))
}
