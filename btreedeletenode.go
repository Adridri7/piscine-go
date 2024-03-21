package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if node.Data < root.Data {
        root.Left = BTreeDeleteNode(root.Left, node)
    } else if node.Data > root.Data {
        root.Right = BTreeDeleteNode(root.Right, node)
    } else {
        if root.Left == nil {
            return root.Right
        } else if root.Right == nil {
            return root.Left
        }

        minRight := BTreeMin(root.Right)
        root.Data = minRight.Data
        root.Right = BTreeDeleteNode(root.Right, minRight)
    }

    return root
}

func BTreeMin(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    for root.Left != nil {
        root = root.Left
    }
    return root
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    if node == root {
        return rplc
    }

    isLeftChild := false
    parent := root
    for parent != nil {
        if parent.Left == node {
            isLeftChild = true
            break
        } else if parent.Right == node {
            isLeftChild = false
            break
        }
        if node.Data < parent.Data {
            parent = parent.Left
        } else {
            parent = parent.Right
        }
    }
    if isLeftChild {
        parent.Left = rplc
    } else {
        parent.Right = rplc
    }

    return root
}

type TreeNode struct {
    Left, Right, Parent *TreeNode
    Data                 string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
    if root == nil {
        return &TreeNode{Data: data}
    }

    if data < root.Data {
        root.Left = BTreeInsertData(root.Left, data)
        root.Left.Parent = root
    } else {
        root.Right = BTreeInsertData(root.Right, data)
        root.Right.Parent = root
    }
    return root
}

func BTreeIsBinary(root *TreeNode) bool {
    if root == nil || root.Left == nil || root.Right == nil {
        return true
    }
    if root.Data > root.Right.Data || root.Data < root.Left.Data {
        return false
    }

    if !BTreeIsBinary(root.Left) {
        return false
    }
    if !BTreeIsBinary(root.Right) {
        return false
    }
    return true
}


func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
    if root == nil {
        return root
    }
    if root.Data == elem {
        return root
    } else if root.Data < elem {
        return BTreeSearchItem(root.Right, elem)
    } else if root.Data > elem {
        return BTreeSearchItem(root.Left, elem)
    } else {
        return nil
    }
}