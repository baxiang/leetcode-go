package main

import "fmt"

// 前序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// 递归方法
func preorderTraversal(root *TreeNode) []int{
	if root==nil {
		return nil
	}
	fmt.Println(root.Val)
	preorderTraversal1(root.Left)
	preorderTraversal1(root.Right)
	return nil
}

// 中左右
func preorderTraversal1(root *TreeNode) []int {
	var l []int
	//如果是空直接返回
	if root == nil {
		return nil
	}
	// 创建栈 递归从根源就是栈结构
	stack := make([]*TreeNode, 0)
	 curr := root
	// 只要当前节点不为空，或者栈不为空，就一直循环查找左节点
	for curr != nil || len(stack) > 0 {
		//查找左节点
		for (curr != nil){
			stack = append(stack,curr)
			fmt.Println(curr.Val)
			l = append(l,curr.Val)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr = curr.Right
	}
	return l
}

// 中左右
func preorderTraversal2(root *TreeNode) []int {
	var l []int
	//如果是空直接返回
	if root == nil {
		return nil
	}
	// 创建栈 递归从根源就是栈结构
	stack := make([]*TreeNode, 0)
	stack = append(stack,root)
	// 只要栈不为空
	for  len(stack) > 0 {
		//直接出栈
		 curr := stack[len(stack)-1]
		 stack = stack[:len(stack)-1]
		 l = append(l,curr.Val)
		 //先压入右节点
		if curr.Right != nil {
			stack = append(stack,curr.Right)
		}
		 //先压入左节点
		if curr.Left != nil {
			stack = append(stack,curr.Left)
		}
	}
	return l
}


//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
func main() {
	a := &TreeNode{
		Val: 1,
	}
	b := &TreeNode{
		Val: 2,
	}
	a.Left = b
	c := &TreeNode{
		Val: 3,
	}
	a.Right = c
	d := &TreeNode{
		Val: 4,
	}
	b.Left = d
	e := &TreeNode{
		Val: 5,
	}
	b.Right = e
	f := &TreeNode{
		Val: 6,
	}
	c.Left = f
	// 1 2 4 5 3 6
	preorderTraversal1(a)
}
