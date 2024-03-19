package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return root
	}
	x := root
	if x.Data == elem {
		return x
	}

	if found := BTreeSearchItem(root.Left, elem); found != nil {
		return found
	}

	if found := BTreeSearchItem(root.Right, elem); found != nil {
		return found
	}

	return nil
}
