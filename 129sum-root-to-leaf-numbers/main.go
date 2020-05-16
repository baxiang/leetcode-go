package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var res int
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, val int, res *int) {
	if root == nil {
		return
	}
	val = val*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res = *res + val
	}
	dfs(root.Left, val, res)
	dfs(root.Right, val, res)
}

func main() {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	one.Left = two
	three := &TreeNode{Val: 3}
	one.Right = three
	fmt.Println(sumNumbers(one))

	four := &TreeNode{Val: 4}
	nine := &TreeNode{Val: 9}
	four.Left = nine
	zero := &TreeNode{Val: 0}
	four.Right = zero
	five := &TreeNode{Val: 5}
	nine.Left = five
	one = &TreeNode{Val: 1}
	nine.Right = one
	fmt.Println(sumNumbers(four))
}
