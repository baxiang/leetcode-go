package main

import "fmt"

type TreeNode struct {
	v int
	left *TreeNode
	right *TreeNode
}

func lowestCommonAncestor(root,left,right *TreeNode)int  {
	for root!=nil{
		if left.v>root.v&&right.v>root.v {
			root=root.right
		}else if left.v<root.v&&right.v<root.v {
			root = right.left
		}else {
			break
		}
	}
	return root.v
}

func main() {
	a :=TreeNode{v:6}
	b :=TreeNode{v:2}
	c :=TreeNode{v:8}
	a.left= &b
	a.right =&c
	d :=TreeNode{v:0}
	f :=TreeNode{v:4}
	b.left = &d
	b.right = &f
	g :=TreeNode{v:3}
	h :=TreeNode{v:5}
	f.left = &g
	f.right = &h
	i :=TreeNode{v:7}
	j :=TreeNode{v:9}
	c.left = &i
	c.right = &j
	fmt.Println(lowestCommonAncestor(&a,&b,&c))
	fmt.Println(lowestCommonAncestor(&a,&b,&f))
}
