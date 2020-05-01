package tree


func buildTree(nums []int, l int, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	node := &TreeNode{Val:nums[mid]}
	node.Left = buildTree(nums, l, mid-1)
	node.Right = buildTree(nums, mid+1, r)
	return node
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildTree(nums, 0, len(nums)-1)
}