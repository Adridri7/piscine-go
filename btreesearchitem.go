package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return root
	}

	if found := BTreeSearchItem(root.Left, elem); found != nil {
		return found
	}

	if found := BTreeSearchItem(root.Right, elem); found != nil {
		return found
	}

	return nil
}
