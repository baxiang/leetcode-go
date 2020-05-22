package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	i := 0
	// 查找根节点位置
	for ; i < len(inorder); i++ {
		if rootVal == inorder[i] {
			break
		}
	}
	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

func main() {
	r := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	printPreorder(r)
	fmt.Println()
	printInorder(r)

}
func printPreorder(r *TreeNode) {
	if r == nil {
		return
	}
	fmt.Print(r.Val)
	printPreorder(r.Left)
	printPreorder(r.Right)
}

func printInorder(r *TreeNode) {
	if r == nil {
		return
	}

	printInorder(r.Left)
	fmt.Print(r.Val)
	printInorder(r.Right)
}
