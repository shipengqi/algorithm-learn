package bstree

type BinaryTree struct {
    root      *TreeNode
}

func NewBinaryTree(rootV interface{}) *BinaryTree {
	return &BinaryTree{NewTreeNode(rootV)}
}