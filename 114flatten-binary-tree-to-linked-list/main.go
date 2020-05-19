package main

type TreeNode struct {
	v int
	left *TreeNode
	right *TreeNode
}

func flatten(root *TreeNode)  {
	if root== nil{
		return
	}
	stack :=[]*TreeNode{root}
	var preNode *TreeNode
	for len(stack)>0{
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		preNode.left = nil
		preNode.right = node
		if node.right !=nil{
			stack = append(stack,node.right)
		}
		if node.left !=nil{
			stack = append(stack,node.left)
		}
		preNode = preNode.right
	}
}

func main() {
	one :=&TreeNode{v:1}
	two :=&TreeNode{v:2}
	one.left = two
	three :=&TreeNode{v:3}
	two.left = three
	four :=&TreeNode{v:4}
	two.right = four
	five :=&TreeNode{v:5}
	one.right = five
	six :=&TreeNode{v:6}
	five.right= six

	flatten(one)


}
