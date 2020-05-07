package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool{
    if s==nil{
    	return false
	}
	return helper(s,t)||isSubtree(s.Left,t)||isSubtree(s.Right,t)
}

func helper(a,b *TreeNode)bool{
	if a==nil&&b==nil{
		return true
	}
	if a==nil||b==nil{
		return false
	}
	if a.Val==b.Val{
		return helper(a.Left,b.Left)&&helper(a.Right,b.Right)
	}
	return false
}

func isSubtree1(s *TreeNode, t *TreeNode) bool{
	var s1 strings.Builder
	preOrder(s,&s1,"")
	var s2 strings.Builder
	preOrder(t,&s2,"")
	fmt.Println(s1.String())
	fmt.Println(s2.String())
	return strings.Contains(s1.String(),s2.String())
}
func preOrder(root *TreeNode, res *strings.Builder,tag string){
     if root ==nil{
		 res.WriteString(tag)
		 return
	 }
	res.WriteString(fmt.Sprintf(" %d ",root.Val))
	preOrder(root.Left,res,"L")
	preOrder(root.Right,res,"R")
}

func main() {
	three := &TreeNode{Val: 3}
	four := &TreeNode{Val: 4}
	three.Left = four
	five := &TreeNode{Val: 5}
	three.Right = five
	one := &TreeNode{Val: 1}
	four.Left = one
	two := &TreeNode{Val: 2}
	four.Right = two
	//zero := &TreeNode{Val: 0}
	//two.Left = zero

	four1 := &TreeNode{Val: 4}
	one1 := &TreeNode{Val: 1}
	four1.Left = one1
	two2 := &TreeNode{Val: 2}
	four1.Right = two2


	fmt.Println(isSubtree1(&TreeNode{Val:12},&TreeNode{Val:2}))
    fmt.Println(strings.Contains("(12)","(2)"))


}
