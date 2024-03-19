package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	curr := root
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr
}
