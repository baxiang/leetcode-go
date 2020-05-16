package main

 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

func levelOrder(root *TreeNode) []int {
	var res []int
	if root==nil{
		return res
	}
	queue := []*TreeNode{root}
	for len(queue)>0{
		size :=len(queue)
		for i:=0;i<size;i++{
			node :=queue[i]
			if node.Left!=nil{
              queue = append(queue,node.Left)
			}
			if node.Right!=nil{
				queue = append(queue,node.Right)
			}
			res = append(res,node.Val)
		}
		queue = queue[size:]
	}
	return res
}

//面试题32 - I. 从上到下打印二叉树
func main() {
	
}
