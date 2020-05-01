package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	stack :=make([]*TreeNode,0)
	curr :=root
	var s []int
    for curr!=nil||len(stack)>0{
		if curr!= nil {
			stack = append(stack,curr)
			curr = curr.Left
		}else {
			v := stack[len(stack)-1] // 获取栈顶元素
			stack = stack[:len(stack)-1] // 出栈
			s = append(s,v.Val)
			curr = v.Right
		}
	}
	return s[len(s)-k]
}

func kthLargest1(root *TreeNode, k int) int {
	stack :=make([]int,0)
	inorder(root,&stack)
	return stack[k-1]
}
func inorder(root *TreeNode,a *[]int){
	if root==nil {
		return
	}
	inorder(root.Right,a)
	*a = append(*a,root.Val)
	inorder(root.Left,a)
}
// 面试题54. 二叉搜索树的第k大节点 https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
func main() {
	a := &TreeNode{Val:3}
	b := &TreeNode{Val:1}
	c := &TreeNode{Val:4}
	d := &TreeNode{Val:2}
	a.Left = b
	a.Right= c
	b.Right = d
	fmt.Println(kthLargest1(a,1))
}
