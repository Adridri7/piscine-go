package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return root
	}
	x := root
	if x.Data == elem {
		return x
	}

	if x.Data < elem {
		found := BTreeSearchItem(root.Left, elem)
		if found != nil {
			return found
		}
	}

	if x.Data > elem {
		found := BTreeSearchItem(root.Right, elem)
		if found != nil {
			return found
		}
	}
	return nil
}
