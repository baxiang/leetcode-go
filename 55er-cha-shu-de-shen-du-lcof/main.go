package main

type TreeNode struct {
	     Val int
	    Left *TreeNode
	     Right *TreeNode
}

func Max(a,b int)int  {
	if a>b{
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	lHeight := maxDepth(root.Left)
	rHeight := maxDepth(root.Right)
	return Max(lHeight,rHeight)+1
}
func maxDepth1(root *TreeNode) int {
	if root==nil{
		return 0
	}
	ques := make([]*TreeNode,0)
	ques = append(ques,root)
	height :=0
	for len(ques)!=0{
		l :=len(ques)
		for i:=0;i<l;i++{
			node :=ques[i]
			if node.Left!=nil {
				ques = append(ques,node.Left)
			}
			if node.Right!=nil {
				ques = append(ques,node.Right)
			}
		}
		ques =ques[l:]
		height++
	}
	return height
}

func main() {
	
}
