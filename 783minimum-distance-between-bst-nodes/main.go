package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
 }
 // 中序遍历
func minDiffInBST(root *TreeNode) int {
	var res []int
	stack := make([]*TreeNode,0)
	min := 1<<63-1
    for root!=nil||len(stack)>0{
    	for root!=nil{
    		stack = append(stack,root)
    		root =root.Left
		}
		node :=stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res,node.Val)
		if len(res)>1 {
			r := node.Val-res[len(res)-2]
			if min>r{
				min = r
			}
		}
		root = node.Right
	}
	if len(res)<2 {
		return res[0]
	}
	return min
}

func BuildBinaryTree(src string)*TreeNode{
	src =src[1:len(src)-1]
	strList := strings.Split(src, ",")
	var currNode *TreeNode
	var root *TreeNode
	var queue []*TreeNode
	for i :=0;i<len(strList);i+=2{
		if i==0 {
			v,_:=strconv.Atoi(strList[i])
			currNode =&TreeNode{Val:v}
			root =currNode
			queue = append(queue,currNode)
		}
		if len(queue)>0 {
			currNode =queue[0]
			queue =queue[1:]
		}else {
			break
		}
		if i+1<len(strList)&&strList[i+1]!="null"{
			v,_:=strconv.Atoi(strList[i+1])
			currNode.Left = &TreeNode{Val:v}
			queue = append(queue,currNode.Left)
		}
		if i+2<len(strList)&&strList[i+2]!="null"{
			v,_:=strconv.Atoi(strList[i+2])
			currNode.Right = &TreeNode{Val:v}
			queue = append(queue,currNode.Right)
		}
	}
	return root
}

func main() {
	root := BuildBinaryTree("[4,2,6,1,3,null,null]")
	fmt.Println(minDiffInBST(root))
	root1 :=BuildBinaryTree("[27,null,34,null,58,50,null,44,null,null,null]")
	fmt.Println(minDiffInBST(root1))
}
