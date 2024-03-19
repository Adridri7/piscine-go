package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root != nil {
		return 1 + BTreeLevelCount(root.Left) + BTreeLevelCount(root.Right)
	} else {
		return 0
	}
}
