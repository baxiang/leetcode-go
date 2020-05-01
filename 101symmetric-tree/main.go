package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root.Left,root.Right)
}

func isMirror(t1 *TreeNode,t2 *TreeNode) bool{
	if t1==nil&&t2==nil {
		return true
	}
	if t1==nil||t2==nil{
		return false
	}
	return (t1.Val == t2.Val)&&isMirror(t1.Left,t2.Right)&&isMirror(t1.Right,t2.Left)
}

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return  true // []认为是true
	}
	queue := make([]*TreeNode,0)
	queue = append(queue,root.Left)
	queue = append(queue,root.Right)
	for len(queue)>=2{
		 left := queue[0]
		 right := queue[1]
		queue = queue[2:]//出栈操作
		if left==nil&&right==nil {
			continue
		}
		if left==nil||right==nil{
			return false
		}
		if left.Val != right.Val {
			return false
		}

		queue = append(queue,left.Left)
		queue = append(queue,right.Right)
		queue = append(queue,left.Right)
		queue = append(queue,right.Left)
	}
	return true
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
		Val: 2,
	}
	a.Right = c
	d := &TreeNode{
		Val: 3,
	}
	b.Left = d
	e := &TreeNode{
		Val: 4,
	}
	b.Right = e
	f := &TreeNode{
		Val: 4,
	}
	c.Left = f
	g := &TreeNode{
		Val: 3,
	}
	c.Right = g
	// 4,2,5,1,6,3
	fmt.Println(isSymmetric1(a))
}
