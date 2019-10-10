package bstree

type BinaryTree struct {
    root      *TreeNode
}

func NewBinaryTree(rootV interface{}) *BinaryTree {
	return &BinaryTree{NewTreeNode(rootV)}
}

func (b *BinaryTree) PreOrderTraversal() {

}

func (b *BinaryTree) InOrderTraversal() {

}

func (b *BinaryTree) PostOrderTraversal() {

}