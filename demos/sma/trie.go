package sma

type TrieNode struct {
	data            string
	isEndingChar    bool
	children        []*TrieNode
}

func NewTrieNode(data string) *TrieNode {
	return &TrieNode{data, false, make([]*TrieNode, 26)}
}

type TrieTree struct {
	root    *TrieNode
}

func NewTrieTree() *TrieTree {
	return &TrieTree{NewTrieNode("/")}
}

func (t *TrieTree) Inert(data string) {
	var basicChar rune = 'a'
	node := t.root
    for _, s := range data {
    	index := int(s) - int(basicChar)
    	if node.children[index] == nil {
			node.children[index] = NewTrieNode(string(s))
		}
		node = node.children[index]
	}
    node.isEndingChar = true
}

func (t *TrieTree) Find(pattern string) bool {
	var basicChar rune = 'a'
	node := t.root
	for _, s := range pattern {
		index := int(s) - int(basicChar)
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	if !node.isEndingChar { // 只匹配到部分前缀
		return false
	}
	return true
}