package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	BTreeApplyByLevel(root.Left, f)
	_, err := f(root.Data)
	if err != nil {
		return
	}

	BTreeApplyByLevel(root.Right, f)
}
