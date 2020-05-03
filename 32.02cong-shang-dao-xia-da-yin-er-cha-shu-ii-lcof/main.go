package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)
	var res [][]int
	for len(nodes) > 0 {
		var nodeList []*TreeNode
		var values []int
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
			}
			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
			}
			values = append(values, node.Val)
		}
		if len(values) > 0 {
			res = append(res, values, )
		}
		nodes = append(nodes[len(nodes):], nodeList...)
	}
	return res
}
func main() {
	three := &TreeNode{Val: 3}
	nine := &TreeNode{Val: 9}
	three.Left = nine
	twenty := &TreeNode{Val: 20}
	three.Right = twenty
	fifteen := &TreeNode{Val: 15}
	twenty.Left = fifteen
	severn := &TreeNode{Val: 7}
	twenty.Right = severn
	fmt.Println(levelOrder(three))
}
