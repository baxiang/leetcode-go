package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
 }


func sortedArrayToBST(nums []int) *TreeNode {
     if len(nums)==0{
     	return  nil
	 }
	 mid := len(nums)/2
	 root := &TreeNode{Val:nums[mid]}
	 root.Left = sortedArrayToBST(nums[:mid])
	 root.Right = sortedArrayToBST(nums[mid+1:])
	 return root
}

func printTreeNode(h *TreeNode){
	if h==nil{
		return
	}
	printTreeNode(h.Left)
	fmt.Println(h.Val)
	printTreeNode(h.Right)
}


func main() {
	r := sortedArrayToBST([]int{-10,-3,0,5,9})
	printTreeNode(r)
}
