package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	    Right *TreeNode
}

func findMode(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	var max int
	m := make(map[int]int)
	var res []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		m[node.Val]++
		if max < m[node.Val] {
			max = m[node.Val]
			res = nil
			res = append(res,node.Val)
		}else if max == m[node.Val]{
			res = append(res,node.Val)
		}
		root = node.Right
	}
	return res
}


func main() {
	one := &TreeNode{Val:1}
	two := &TreeNode{Val:2}
	one.Right = two
	three :=&TreeNode{Val:2}
	two.Left = three
	fmt.Println(findMode(one))
}
