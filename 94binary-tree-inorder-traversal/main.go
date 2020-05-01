package main

// 中序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// 左中右
func inorderTraversal(root *TreeNode) []int {
	var l []int
	//如果是空直接返回
	if root == nil {
		return nil
	}
	// 创建栈 递归从根源就是栈结构
	stack := make([]*TreeNode, 0)
	node := root
	// 只要当前节点不为空，或者栈不为空，就一直循环查找左节点
	for node != nil || len(stack) > 0 {
		//查找左节点
		if node != nil {
			// 将左节点放入栈中
			stack = append(stack, node)
			node = node.Left
		} else {
			// 获取站顶节点
			curr := stack[len(stack)-1]
			// 出栈操作
			stack = stack[:len(stack)-1]
			l = append(l, curr.Val)
			// 遍历右节点
			node = curr.Right
		}
	}
	return l
}





func main() {
	a := &TreeNode{
		Val: 1,
	}
	b := &TreeNode{
		Val: 2,
	}
	a.Left = b
	c := &TreeNode{
		Val: 3,
	}
	a.Right = c
	d := &TreeNode{
		Val: 4,
	}
	b.Left = d
	e := &TreeNode{
		Val: 5,
	}
	b.Right = e
	f := &TreeNode{
		Val: 6,
	}
	c.Left = f
	// 4,2,5,1,6,3
	inorderTraversal(a)
}
