package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
 }
func isSameTree(p *TreeNode, q *TreeNode) bool {
     if p==nil&&q==nil{
     	return  true
	 }
	 if p!=nil&&q==nil||p==nil&&q!=nil||p.Val!=q.Val{
	 	return false
	 }
	 return isSameTree(p.Left,q.Left)&&isSameTree(p.Right,q.Right)
}


func isSameTree1(p *TreeNode, q *TreeNode) bool {
	if p==nil&&q==nil{
		return  true
	}
	if p!=nil&&q==nil||p==nil&&q!=nil||p.Val!=q.Val{
		return false
	}
	queue :=make([]*TreeNode,0)
	queue = append(queue,p)
	queue = append(queue,q)
	for len(queue)>0{
		pNode :=queue[0]
		qNode :=queue[1]

		if pNode.Val!=qNode.Val{
			return false
		}
		if pNode.Left!=nil&&qNode.Left!=nil {
			queue = append(queue,pNode.Left)
			queue = append(queue,qNode.Left)
		} else if pNode.Left==nil&&qNode.Left!=nil {
			return false
		}else if pNode.Left!=nil&&qNode.Left==nil {
			return false
		}
		if pNode.Right!=nil&&qNode.Right!=nil {
			queue = append(queue,pNode.Right)
			queue = append(queue,qNode.Right)
		}else if pNode.Right==nil&&qNode.Right!=nil {
			return false
		}else if pNode.Right!=nil&&qNode.Right==nil {
			return false
		}
		queue =queue[2:]
	}
	return true
}



func main() {
	t11 := &TreeNode{Val: 1}
	t12 := &TreeNode{Val: 2}
	t11.Left = t12
	t13 := &TreeNode{Val: 1}
	t11.Right = t13

	t21 := &TreeNode{Val: 1}
	t22 := &TreeNode{Val: 1}
	t21.Left = t22
	t23 := &TreeNode{Val: 2}
	t21.Right = t23


	fmt.Println(isSameTree1(t11,t21))
}
