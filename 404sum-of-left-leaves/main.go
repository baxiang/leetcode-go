package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
     stack := make([]*TreeNode,0)
     var res int
     for root!= nil||len(stack)>0{
     	if root!= nil{
     		stack = append(stack,root)
     		root = root.Left
			if root!=nil&&root.Left==nil&&root.Right==nil{
				res +=root.Val
			}
		}else {
			node :=stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = node.Right
		}
	 }
	 return res
}

func main() {
	three :=&TreeNode{Val: 3}
	nine :=&TreeNode{Val: 9}
	three.Left = nine
	twenty :=&TreeNode{Val: 20}
	three.Right = twenty
	fifteen :=&TreeNode{Val: 15}
	twenty.Left = fifteen
	seven :=&TreeNode{Val: 7}
	twenty.Right = seven

    fmt.Println(sumOfLeftLeaves(three))

	t1 :=&TreeNode{Val: 1}
	t2 :=&TreeNode{Val: 2}
	t1.Left = t2
	t3 :=&TreeNode{Val: 3}
	t1.Right = t3
	t4 :=&TreeNode{Val: 4}
	t2.Left = t4
	t5 :=&TreeNode{Val: 5}
	t2.Right = t5

	fmt.Println(sumOfLeftLeaves(t1))



}
