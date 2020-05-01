package tree

import (
	"strconv"
	"strings"
)
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
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