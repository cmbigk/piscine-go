package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	h := BTreeLevelCount(root)
	for i := 0; i < h; i++ {
		applyGivenOrder(root, i, f)
	}
}

func applyGivenOrder(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	if level == 0 {
		arg := interface{}(root.Data)
		f(arg)
	} else if level > 0 {
		applyGivenOrder(root.Left, level-1, f)
		applyGivenOrder(root.Right, level-1, f)
	}
}
