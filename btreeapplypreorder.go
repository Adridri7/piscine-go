package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyPostorder(root.Right, f)
	_, err := f(root.Data)
	if err != nil {
		return
	}
	BTreeApplyPostorder(root.Left, f)
}
