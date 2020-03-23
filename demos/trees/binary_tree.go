package trees

type BTreeNode struct {
	Value        int
	Left         *BTreeNode
	Right        *BTreeNode
}

func PreOrderTraversal(root *BTreeNode) []int {
	if root == nil {
		return nil
	}
	// 没有子节点
    if root.Left == nil && root.Right == nil {
    	return []int{root.Value}
	}
	var res []int
	var stack []*BTreeNode
	stack = append(stack, root)

	for len(stack) != 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, e.Value)
		if e.Right != nil {
			stack = append(stack, e.Right)
		}
		if e.Left != nil {
			stack = append(stack, e.Left)
		}
	}
	return []int{}
}

func InOrderTraversal(root *BTreeNode) []int {
	if root == nil {
		return nil
	}
	// 没有子节点
	if root.Left == nil && root.Right == nil {
		return []int{root.Value}
	}
	res := InOrderTraversal(root.Left)
	res = append(res, root.Value)
	res = append(res, InOrderTraversal(root.Right)...)
    return res
}

func PostOrderTraversal(root *BTreeNode) []int {
	if root == nil {
		return nil
	}
	// 没有子节点
	if root.Left == nil && root.Right == nil {
		return []int{root.Value}
	}
	var res []int
	if root.Left != nil {
		lres := PostOrderTraversal(root.Left)
		if len(lres) > 0 {
			res = append(res, lres...)
		}
	}
	if root.Right != nil {
		rres := PostOrderTraversal(root.Right)
		if len(rres) > 0 {
			res = append(res, rres...)
		}
	}
	res = append(res, root.Value)
	return res
}