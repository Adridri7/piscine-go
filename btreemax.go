package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	curr := root
	for curr.Right != nil {
		curr = curr.Right
	}
	return curr
}
