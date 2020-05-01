package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
var res *int // leetcode bug 全局变量最好是指针
func diameterOfBinaryTree(root *TreeNode) int {
	t :=1
	res =&t
	if root==nil {
		return 0
	}
	dfs(root)
	return *res
}

func max(a,b int)int{
	if a>b {
		return a
	}
	return b
}

func dfs(root *TreeNode)int{
	if root==nil {
		return 0
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	*res = max(*res,l+r)
	return max(l,r)+1
}


func main() {
	a := &TreeNode{
		Val: 1,
	}
	b := &TreeNode{
		Val: 2,
	}
	a.Left = b
	c := &TreeNode{
		Val: 3,
	}
	a.Right = c
	d := &TreeNode{
		Val: 4,
	}
	b.Left = d
	e := &TreeNode{
		Val: 5,
	}
	b.Right = e
	fmt.Println(diameterOfBinaryTree(a))

	g := &TreeNode{
		Val: 1,
	}
	fmt.Println(diameterOfBinaryTree(g))
}
