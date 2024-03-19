package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	_, err := f(root.Data)
	if err != nil {
		return
	}
	BTreeApplyPreorder(root.Left, f)
	BTreeApplyPreorder(root.Right, f)
}
