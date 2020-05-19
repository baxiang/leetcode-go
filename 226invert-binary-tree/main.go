package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}


func invertTree3(root *TreeNode) *TreeNode {
	if root==nil {// 递归结束条件
		return nil
	}
	right:= invertTree3(root.Right)// 返回已经翻转的右节点
	left := invertTree3(root.Left) // 返回已经翻转的左节点
	// 当前需要一级需要做的是把 左节点指向已经翻转的右节点 右节点指向已经翻转的左节点
	root.Left = right
	root.Right = left
	return root
}
func invertTree1(root *TreeNode) *TreeNode {
	if root==nil {// 递归结束条件
		return nil
	}

	// 当前需要一级需要做的是把 左节点指向右节点 右节点指向已经翻转的左节点
	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	invertTree1(root.Left) // 翻转的左节点
	invertTree1(root.Right)// 返翻转的右节点

	return root
}


func invertTree(root *TreeNode) *TreeNode {
       if root==nil{
       	  return root
	   }
	   queue :=[]*TreeNode{root}
	   for len(queue)>0{
	   	  size  :=len(queue)
	   	  for i:=0;i<size;i++{
	   	  	  node := queue[i]
	   	  	  node.Left,node.Right = node.Right,node.Left
	   	  	  if node.Left!=nil{
				  queue = append(queue,node.Left)
			  }
			  if node.Right!=nil{
				  queue = append(queue,node.Right)
			  }
		  }
		  queue =queue[size:]
	   }
	   return root
}


func main() {
		a := &TreeNode{
			Val: 4,
		}
		b := &TreeNode{
			Val: 2,
		}
		a.Left = b
		c := &TreeNode{
			Val: 7,
		}
		a.Right = c
		d := &TreeNode{
			Val: 1,
		}
		b.Left = d
		e := &TreeNode{
			Val: 3,
		}
		b.Right = e
		//f := &TreeNode{
		//	Val: 6,
		//}
	    //c.Left = f
	    g := &TreeNode{
		Val: 9,
	    }
		c.Right = g
		// 9 7 64 3 2 1
	    r :=invertTree(a)
		printOrder(r)
}

func printOrder(root *TreeNode){
	if root== nil {
		return
	}
	fmt.Println(root.Val)
	printOrder(root.Left)
	printOrder(root.Right)
	return
}