package main
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(root *TreeNode, s int) int {

	// If root is nil, that means we can't find anything,
	// we return -1
	if root == nil {
		return -1
	}

	// If the current tree value is bigger than the smallest
	// element we hat, that means we have found the second smallest
	// one in the tree, since all subtrees is greater or equal to this
	// value.
	if root.Val > s {
		return root.Val
	}

	sLeft := dfs(root.Left, s)
	sRight := dfs(root.Right, s)

	if sLeft == -1 {
		return sRight
	}

	if sRight == -1 {
		return sLeft
	}

	return min(sLeft, sRight)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	a := &TreeNode{
		Val: 4,
	}
	b := &TreeNode{
		Val: 2,
	}
	a.Left = b
	c := &TreeNode{
		Val: 7,
	}
	a.Right = c
	d := &TreeNode{
		Val: 1,
	}
	b.Left = d
	e := &TreeNode{
		Val: 3,
	}
	b.Right = e
	f := &TreeNode{
		Val: 6,
	}
	c.Left = f
	g := &TreeNode{
		Val: 9,
	}
	c.Right = g
	// 9 7 64 3 2 1
	findSecondMinimumValue(a)
}
