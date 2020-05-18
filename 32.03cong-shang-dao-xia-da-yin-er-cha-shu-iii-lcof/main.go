package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	var isReverse bool
	for len(queue) > 0 {
		size := len(queue)
		list := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[i]
			if !isReverse {
				list[i] = node.Val
			} else {
				list[size-i-1] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		isReverse = !isReverse
		res = append(res, list)
	}
	return res
}

func BuildBinaryTree(src string) *TreeNode {
	src = src[1 : len(src)-1]
	strList := strings.Split(src, ",")
	var currNode *TreeNode
	var root *TreeNode
	var queue []*TreeNode
	for i := 0; i < len(strList); i += 2 {
		if i == 0 {
			v, _ := strconv.Atoi(strList[i])
			currNode = &TreeNode{Val: v}
			root = currNode
			queue = append(queue, currNode)
		}
		if len(queue) > 0 {
			currNode = queue[0]
			queue = queue[1:]
		} else {
			break
		}
		if i+1 < len(strList) && strList[i+1] != "null" {
			v, _ := strconv.Atoi(strList[i+1])
			currNode.Left = &TreeNode{Val: v}
			queue = append(queue, currNode.Left)
		}
		if i+2 < len(strList) && strList[i+2] != "null" {
			v, _ := strconv.Atoi(strList[i+2])
			currNode.Right = &TreeNode{Val: v}
			queue = append(queue, currNode.Right)
		}
	}
	return root
}
func main() {
	root := BuildBinaryTree("[3,9,20,null,null,15,7]")
	fmt.Println(levelOrder(root))
}
