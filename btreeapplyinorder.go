package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	_, err := f(root.Data)
	if err != nil {
		return
	}
	BTreeApplyInorder(root.Right, f)
}
