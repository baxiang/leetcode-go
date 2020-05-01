package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root!= nil{
		if root.Val>p.Val&&root.Val>q.Val {
			root = root.Left
		} else if root.Val<p.Val&&root.Val<q.Val {
			root = root.Right
		}else {
			return root
		}
	}
  return nil
}
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		root = lowestCommonAncestor1(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		root = lowestCommonAncestor1(root.Right, p, q)
	}
	return root
}

func main() {
	six := &TreeNode{
		Val: 6,
	}
	two := &TreeNode{
		Val: 2,
	}
	six.Left = two
	eight := &TreeNode{
		Val: 8,
	}
	six.Right = eight
	zero := &TreeNode{
		Val: 0,
	}
	two.Left = zero
	four := &TreeNode{
		Val: 4,
	}
	two.Right = four
	three := &TreeNode{
		Val: 3,
	}
	four.Left = three
	five := &TreeNode{
		Val: 5,
	}
	four.Right = five

	seven := &TreeNode{
		Val: 7,
	}
	eight.Left = seven
	nine := &TreeNode{
		Val: 9,
	}
	eight.Right = nine
	fmt.Println(lowestCommonAncestor1(six,two,eight).Val)
	fmt.Println(lowestCommonAncestor1(six,two,four).Val)
}
