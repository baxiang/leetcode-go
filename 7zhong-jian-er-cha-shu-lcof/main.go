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
	rootIdx := 0
	for i := range inorder {
		if inorder[i] == preorder[0] {
			rootIdx = i
			break
		}
	}
	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:rootIdx+1], inorder[:rootIdx])
	root.Right = buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:])
	return root
}
func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	r := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	printTree(r)
}
