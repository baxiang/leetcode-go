package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees1(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val = t1.Val + t2.Val
	t1.Left = mergeTrees1(t1.Left, t2.Left)
	t1.Right = mergeTrees1(t1.Right, t2.Right)
	return t1
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	queue := []*TreeNode{t1, t2}
	for len(queue) > 0 {
		n1 := queue[0]
		n2 := queue[1]
		n1.Val = n1.Val + n2.Val
		if n1.Left != nil && n2.Left != nil {
			queue = append(queue, n1.Left, n2.Left)
		} else if n1.Left == nil && n2.Left != nil {
			n1.Left = n2.Left
		}
		if n1.Right != nil && n2.Right != nil {
			queue = append(queue, n1.Right, n2.Right)
		} else if n1.Right == nil && n2.Right != nil {
			n1.Right = n2.Right
		}
		queue = queue[2:]
	}
	return t1
}

func main() {

	t1 := &TreeNode{Val: 1}
	t11 := &TreeNode{Val: 3}
	t1.Left = t11
	t12 := &TreeNode{Val: 2}
	t1.Right = t12
	t13 := &TreeNode{Val: 5}
	t11.Left = t13

	t2 := &TreeNode{Val: 2}
	t21 := &TreeNode{Val: 1}
	t2.Left = t21
	t22 := &TreeNode{Val: 3}
	t2.Right = t22
	t23 := &TreeNode{Val: 4}
	t21.Right = t23
	t24 := &TreeNode{Val: 7}
	t22.Right = t24

	printNode(mergeTrees(t1, t2)) //3 4 5 4 5 7
	fmt.Println()

	t3 := &TreeNode{Val: 1}
	t31 := &TreeNode{Val: 2}
	t3.Left = t31
	t32 := &TreeNode{Val: 3}
	t31.Left = t32

	t4 := &TreeNode{Val: 1}
	t41 := &TreeNode{Val: 2}
	t4.Right = t41
	t42 := &TreeNode{Val: 3}
	t41.Right = t42
	printNode(mergeTrees(t3, t4)) //2 2 3 2 3
}

func printNode(t1 *TreeNode) {
	if t1 == nil {
		return
	}
	fmt.Println(t1.Val)
	printNode(t1.Left)
	printNode(t1.Right)
}
