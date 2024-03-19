package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	_, err := f(root.Data)
	if err != nil {
		return
	}
	BTreeApplyByLevel(root.Left, f)
	BTreeApplyByLevel(root.Right, f)
}
