package bstree

import "fmt"

type TreeNode struct {
	data  interface{}
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(data interface{}) *TreeNode {
	return &TreeNode{data, nil, nil}
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", t.data, t.left, t.right)
}