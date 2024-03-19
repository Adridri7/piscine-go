package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	tab := []*TreeNode{root}
	for len(tab) > 0 {
		node := tab[0]
		tab = tab[1:]
		_, _ = f(node.Data)
		if node.Left != nil {
			tab = append(tab, node.Left)
		}

		if node.Right != nil {
			tab = append(tab, node.Right)
		}
	}
}
