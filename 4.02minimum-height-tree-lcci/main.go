package main

 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
 }

func sortedArrayToBST(nums []int) *TreeNode {
       if len(nums)==0{
       	return nil
	   }
	   mid := len(nums)/2
	   root :=&TreeNode{Val:nums[mid]}
	   root.Left = sortedArrayToBST(nums[:mid])
	   root.Right = sortedArrayToBST(nums[mid+1:])
	   return root
}

//面试题 04.02. 最小高度树
func main() {
	
}
