package sma

import "testing"

func TestTrieTree_Find(t *testing.T) {
	tree := NewTrieTree()
	tree.Inert("abc")
	tree.Inert("hello")
	tree.Inert("hi")
	tree.Inert("her")
	tree.Inert("him")

	t.Run("Find he from trie tree", func(t *testing.T) {
		isFind := tree.Find("he")
		if isFind {
			t.Fatalf("Wanted false, got %v", isFind)
		}
	})

	t.Run("Find him from trie tree", func(t *testing.T) {
		isFind := tree.Find("him")
		if !isFind {
			t.Fatalf("Wanted true, got %v", isFind)
		}
	})
}
