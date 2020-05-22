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

func tree2str(t *TreeNode) string {
	var str strings.Builder
	if t == nil {
		return ""
	}
	stack := []interface{}{t}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if n, ok := node.(*TreeNode); ok {
			str.WriteString(fmt.Sprintf("%d", n.Val))
			if n.Right != nil {
				stack = append(stack, ")")
				stack = append(stack, n.Right)
				stack = append(stack, "(")
			}
			if n.Right != nil && n.Left == nil {
				stack = append(stack, "()")
			}
			if n.Left != nil {
				stack = append(stack, ")")
				stack = append(stack, n.Left)
				stack = append(stack, "(")
			}
		} else {
			s := node.(string)
			str.WriteString(s)
		}
	}
	return str.String()
}
func main() {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	one.Left = two
	three := &TreeNode{Val: 3}
	one.Right = three
	four := &TreeNode{Val: 4}
	two.Right = four
	fmt.Println(tree2str(one))

	//one1 :=&TreeNode{Val: 1}
	//two1 :=&TreeNode{Val: 2}
	//one1.Left = two1
	//three1 :=&TreeNode{Val: 3}
	//one1.Right = three1
	//four1 :=&TreeNode{Val: 4}
	//two1.Right = four1
	//fmt.Println(tree2str(one1))

}
