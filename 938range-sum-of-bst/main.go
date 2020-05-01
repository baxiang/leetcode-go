package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
 }
 //
func rangeSumBST(root *TreeNode, L int, R int) int {
    stack := make([]*TreeNode,0)
    var res =0
    for root!=nil||len(stack)>0{
    	for root!=nil{
    		stack = append(stack,root)
    		root = root.Left
		}
		node :=stack[len(stack)-1]
		stack =stack[:len(stack)-1]
		 if node.Val>=L&&node.Val<=R{
			res +=node.Val
			if node.Val==R{
				return res
			}
		}
		root = node.Right

	}
	return res
}

func main() {
	ten :=&TreeNode{Val:10}
	five :=&TreeNode{Val:5}
	ten.Left = five
	fifteen :=&TreeNode{Val:15}
	ten.Right = fifteen
	three :=&TreeNode{Val:3}
	five.Left = three
	seven :=&TreeNode{Val:7}
	five.Right = seven
	eighteen :=&TreeNode{Val:18}
	fifteen.Right = eighteen
	fmt.Println(rangeSumBST(ten,7,15))
}
