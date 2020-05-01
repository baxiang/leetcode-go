package main

import (
	"fmt"
	"math"
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
var pre *TreeNode
var res int
func dfs(root *TreeNode){
	if root==nil{
		return
	}
	dfs(root.Left)
	if pre!=nil {
		res = int(math.Min(math.Abs(float64(root.Val-pre.Val)),float64(res)))
	}
	pre = root
	dfs(root.Right)
}
func getMinimumDifference(root *TreeNode) int {
	res = math.MaxInt32
    dfs(root)
    return res
}
func main() {
	root :=&TreeNode{Val:   1}
	two :=&TreeNode{Val:   2}
	three :=&TreeNode{Val:   3}
	root.Right = three
	three.Left = two
	fmt.Println(getMinimumDifference(root))
}
