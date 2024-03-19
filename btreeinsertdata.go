package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	newNode := &TreeNode{Data: data}

	if root == nil {
		return newNode
	}
	current := root
	var parent *TreeNode
	for current != nil {
		parent = current
		if data < current.Data {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	newNode.Parent = parent
	if data < parent.Data {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}
	return root
}
