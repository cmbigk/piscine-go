package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(bt *TreeNode, elem string) *TreeNode {
	if bt == nil {
		return &TreeNode{Data: elem}
	}

	if elem < bt.Data {
		bt.Left = BTreeInsertData(bt.Left, elem)
		bt.Left.Parent = bt
	} else if elem >= bt.Data {
		bt.Right = BTreeInsertData(bt.Right, elem)
		bt.Right.Parent = bt
	}
	return bt
}
