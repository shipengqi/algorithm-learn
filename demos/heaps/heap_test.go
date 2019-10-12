package heaps

import "testing"

func TestHeapInsert(t *testing.T) {
	h := NewHeap(10)
	h.insert(10)
	h.insert(9)
	h.insert(2)
	h.insert(5)
	h.insert(4)
	h.insert(7)
	h.insert(3)
	h.insert(1)
	h.insert(6)
	t.Run("Get Heap length", func(t *testing.T) {
        t.Log(h)
	})
}
