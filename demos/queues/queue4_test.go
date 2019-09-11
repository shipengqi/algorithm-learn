package queues

import "testing"

func TestCircularQueue_Dequeue(t *testing.T) {
	q := NewCircularQueue(4)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	var want int
	t.Run("CircularQueue is full", func(t *testing.T) {
		if q.IsFull() == false {
			t.Fatalf("Wanted true, got false")
		}
	})

	t.Run("Get CircularQueue first item", func(t *testing.T) {
		want := 1
		got := q.Dequeue()
		if want != got {
			t.Fatalf("Wanted %v, got %v", want, got)
		}
	})


	t.Run("Get CircularQueue item 3", func(t *testing.T) {
		q.Dequeue()
		want = 3
		got := q.Dequeue()
		if want != got {
			t.Fatalf("Wanted %v, got %v", want, got)
		}
	})

	t.Run("Get CircularQueue item nil", func(t *testing.T) {
		got := q.Dequeue()
		if  got != nil {
			t.Fatalf("Wanted %v, got %v", nil, got)
		}
	})

	t.Run("CircularQueue is empty", func(t *testing.T) {
		if q.IsEmpty() == false {
			t.Fatalf("Wanted true, got false")
		}
	})

	t.Run("Get CircularQueue item 5", func(t *testing.T) {
		q.Enqueue(5)
		want = 5
		got := q.Dequeue()
		if want != got {
			t.Fatalf("Wanted %v, got %v", want, got)
		}
	})
}