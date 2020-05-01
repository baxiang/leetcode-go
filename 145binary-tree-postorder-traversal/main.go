package main

import "fmt"

// 后续遍历 https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var list []int
	if root == nil {
		return list
	}
	stack1 := make([]*TreeNode,0)
	stack2 := make([]*TreeNode,0)
	stack1 = append(stack1,root)
	for len(stack1)>0{
		curr := stack1[len(stack1)-1]
		fmt.Println(curr.Val)
		//list = append(list,curr.Val)
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2,curr)
		if curr.Left!= nil{
			stack1 = append(stack1,curr.Left)
		}
		if curr.Right!= nil{
			stack1 = append(stack1,curr.Right)
		}
	}
	for len(stack2)>0 {
		curr := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		list = append(list,curr.Val)
	}
	return list
}
func postorderTraversal1(root *TreeNode) []int {
	if root==nil{
		return nil
	}
	postorderTraversal1(root.Left)
	postorderTraversal1(root.Right)
	fmt.Println(root.Val)
	return nil
}

func postorderTraversal2(root *TreeNode) []int {
	var result []int
	if root==nil {
		return result
	}
	stack :=make([]*TreeNode,0)
	curr := root
	var lastNodeVistited *TreeNode //保存上一个遍历过的节点TreeNode  = null; //保存上一个遍历过的节点
	for curr!= nil||len(stack)>0{
		// 把左孩子押入栈 ，到叶子节点为止
		if curr != nil {
			stack = append(stack,curr)
			curr = curr.Left
		}else {
			top := stack[len(stack)-1]
			////如果栈顶节点没有右孩子，或者已经遍历过了它的右孩子，就遍历该节点
			if top.Right==nil||lastNodeVistited == top.Right {
				result = append(result,top.Val)
				lastNodeVistited = top
				stack =stack[:len(stack)-1]
			}else {
				curr = top.Right////否则，准备将栈顶节点的右孩子入栈
			}
		}
	}
	return result
}

func main() {
	a := &TreeNode{
		Val: 1,
	}
	b := &TreeNode{
		Val: 2,
	}
	a.Left = b
	c := &TreeNode{
		Val: 3,
	}
	a.Right = c
	d := &TreeNode{
		Val: 4,
	}
	b.Left = d
	e := &TreeNode{
		Val: 5,
	}
	b.Right = e
	f := &TreeNode{
		Val: 6,
	}
	c.Left = f
	// 4 5 2 6 3 1
	fmt.Println(postorderTraversal2(a))
}
