package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var res []int
	for len(queue) > 0 {
		l := len(queue)
		node := queue[l-1]
		res = append(res, node.Val)
		for i := range queue {
			currNode := queue[i]
			if currNode.Left != nil {
				queue = append(queue, currNode.Left)
			}
			if currNode.Right != nil {
				queue = append(queue, currNode.Right)

			}
		}
		queue = queue[l:]
	}
	return res
}

func main() {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	one.Left = two
	three := &TreeNode{Val: 3}
	one.Right = three
	//five :=&TreeNode{Val: 5}
	//two.Right = five
	four := &TreeNode{Val: 4}
	two.Left = four
	fmt.Println(rightSideView(one))

	//fmt.Println(rightSideView(one))
}
