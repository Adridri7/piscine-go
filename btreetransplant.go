package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	rplc.Parent = node.Parent

	if node.Parent.Left.Data == rplc.Data {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	return root
}
