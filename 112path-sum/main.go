package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//递归解决思路
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil{
		return false
	}
	// 当前是叶子节点 判断当前是否是等于目标值
	if root.Left==nil&&root.Right==nil{
		return root.Val==sum
	}
	return hasPathSum(root.Left,sum-root.Val)||hasPathSum(root.Right,sum-root.Val)
}


func hasPathSum1(root *TreeNode, sum int) bool {
	if root == nil{
		return false
	}
	stack := make([]*TreeNode,0)
	vals := make([]int,0)
	stack = append(stack,root)
	vals = append(vals,sum)
	for len(stack)>0{
		curr :=stack[len(stack)-1]
		stack = stack[:len(stack)-1]//出栈
		val :=vals[len(vals)-1]
		vals = vals[:len(vals)-1] // 出栈
		if curr.Left==nil&&curr.Right==nil&&val ==curr.Val {// 如果当前节点已经为叶子节点并且节点上的数字等于当前和s,则返回true
			return true
		}
		if curr.Left!=nil {
			stack = append(stack,curr.Left)
			vals = append(vals,val-curr.Val)
		}
		if curr.Right!=nil {
			stack = append(stack,curr.Right)
			vals = append(vals,val-curr.Val)
		}
	}
	return false
}

//https://leetcode-cn.com/problems/path-sum/
func main() {
	a := &TreeNode{
		Val: 5,
	}
	b := &TreeNode{
		Val: 4,
	}
	a.Left = b
	c := &TreeNode{
		Val: 8,
	}
	a.Right = c
	d := &TreeNode{
		Val: 11,
	}
	b.Left = d
	e := &TreeNode{
		Val: 7,
	}
	d.Left = e
	f := &TreeNode{
		Val: 2,
	}
	d.Right = f
	// 4,2,5,1,6,3
	fmt.Println(hasPathSum1(a,21))
	fmt.Println(hasPathSum1(a,22))
}
