package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeVal struct {
	Node *TreeNode
	Str string
}
func binaryTreePaths1(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	queue := []*TreeNodeVal{&TreeNodeVal{
		Node: root,
		Str: fmt.Sprintf("%d->",root.Val),
	}}
	var res []string
	for len(queue)>0{
		node :=queue[0]
		queue= queue[1:]
		if node.Node.Left!= nil{
			queue = append(queue,&TreeNodeVal{
				Node: node.Node.Left,
				Str: fmt.Sprintf("%s%d->",node.Str,node.Node.Left.Val),
			})
		}
		if node.Node.Right!= nil{
			queue = append(queue,&TreeNodeVal{
				Node: node.Node.Right,
				Str: fmt.Sprintf("%s%d->",node.Str,node.Node.Right.Val),
			})
		}
		if node.Node.Left==nil&&node.Node.Right==nil{
			res = append(res,node.Str[:len(node.Str)-2])
		}
	}
	return res
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res := make([]string, 0)
	dfs(root, "", &res)
	return res
}

func dfs(root *TreeNode, str string, res *[]string){
	if root == nil{
		return
	}
	str = fmt.Sprintf("%s%d->",str, root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res, str[:len(str)-2])
		return
	}
	dfs(root.Right, str, res)
	dfs(root.Left, str, res)
}

func main() {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	one.Left = two
	three := &TreeNode{Val: 3}
	one.Right = three
	five := &TreeNode{Val: 5}
	two.Right = five
	fmt.Println(binaryTreePaths1(one))
}
