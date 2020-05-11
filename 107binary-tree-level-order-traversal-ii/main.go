package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
    if root==nil{
    	return res
	}
	queue := []*TreeNode{root}
	for len(queue)>0{
		l :=len(queue)
		list := make([]int,0)
		for i:=0;i<l;i++{
			node :=queue[i]
			list = append(list,node.Val)
			if node.Left!=nil{
				queue = append(queue,node.Left)
			}
			if node.Right!=nil{
				queue = append(queue,node.Right)
			}
		}
		res = append([][]int{list},res...)
		queue =queue[l:]
	}
	return res
}


func main() {
	three :=&TreeNode{Val: 3}
	nine :=&TreeNode{Val: 9}
	twenty :=&TreeNode{Val: 20}
	three.Left = nine
	three.Right = twenty
	fifteen :=&TreeNode{Val: 15}
	seven :=&TreeNode{Val: 7}
	twenty.Left = fifteen
	twenty.Right = seven
	fmt.Println(levelOrderBottom(three))
}
