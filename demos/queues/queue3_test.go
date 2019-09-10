package queues

import (
	"testing"
)

func TestLinkedListQueue_Dequeue(t *testing.T) {
	q := NewLinkedListQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
    var want int
	t.Run("Get LinkedListQueue length", func(t *testing.T) {
		if q.Length != 4 {
			t.Fatalf("Wanted %v, got %v", 4, q.Length)
		}
	})

	t.Run("Get LinkedListQueue head node", func(t *testing.T) {
		want := 1
		got := q.Dequeue()
		if want != got {
			t.Fatalf("Wanted %v, got %v", want, got)
		}
	})


	t.Run("Get LinkedListQueue node 3", func(t *testing.T) {
		q.Dequeue()
		want = 3
		got := q.Dequeue()
		if want != got {
			t.Fatalf("Wanted %v, got %v", want, got)
		}
	})

	t.Run("Get LinkedListQueue last node", func(t *testing.T) {
		want = 4
		got := q.Dequeue()
		if want != got {
			t.Fatalf("Wanted %v, got %v", want, got)
		}
	})

	t.Run("Get LinkedListQueue node nil", func(t *testing.T) {
		got := q.Dequeue()

		if  got != nil {
			t.Fatalf("Wanted %v, got %v", nil, got)
		}
		t.Log(q)
	})

	t.Run("Get LinkedListQueue head node 5", func(t *testing.T) {
		q.Enqueue(5)
		t.Log(q)
		want = 5
		got := q.Dequeue()

		if want != got {
			t.Fatalf("Wanted %v, got %v", want, got)
		}
	})
}