package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}


func invertTree(root *TreeNode) *TreeNode {
	if root==nil {// 递归结束条件
		return nil
	}
	right:= invertTree(root.Right)// 返回已经翻转的右节点
	left := invertTree(root.Left) // 返回已经翻转的左节点
	// 当前需要一级需要做的是把 左节点指向已经翻转的右节点 右节点指向已经翻转的左节点
	root.Left = right
	root.Right = left
	return root
}
func invertTree1(root *TreeNode) *TreeNode {
	if root==nil {// 递归结束条件
		return nil
	}

	// 当前需要一级需要做的是把 左节点指向右节点 右节点指向已经翻转的左节点
	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	invertTree(root.Left) // 翻转的左节点
	invertTree(root.Right)// 返翻转的右节点

	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root==nil {
		return nil
	}
	l := make([]*TreeNode,0)
	l = append(l ,root)
	for len(l)>0{
		curr := l[0]
		l = l[1:]
		left :=curr.Left
		right :=curr.Right
		curr.Right = left
		curr.Left = right
		if curr.Left!=nil {
			l = append(l,curr.Left)
		}
		if curr.Right!=nil {
			l = append(l,curr.Right)
		}
	}
	return root
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
	    invertTree2(a)
		printOrder(a)
}

func printOrder(root *TreeNode){
	if root== nil {
		return
	}
	printOrder(root.Left)
	fmt.Println(root.Val)
	printOrder(root.Right)
	return
}