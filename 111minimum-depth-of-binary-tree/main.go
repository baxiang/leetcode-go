package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var min = 1<<63 - 1

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dfs(root, 1) //深度递归
	return min
}

func dfs(node *TreeNode, depth int) {
	if node.Right == nil && node.Left == nil { //递归结束条件
		if depth < min {
			min = depth
		}
		return
	}
	if node.Left != nil { // 遍历左节点
		dfs(node.Left, depth+1)
	}

	if node.Right != nil { // 遍历右节点
		dfs(node.Right, depth+1)
	}
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
	fmt.Println(minDepth(a))
}
