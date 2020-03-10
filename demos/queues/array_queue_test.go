package queues

import "testing"

func TestArrayQueue_Enqueue(t *testing.T) {
    q := NewArrayQueue(5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    q.Enqueue(4)
    q.Enqueue(5)
    q.Enqueue(6)
    t.Log(q)
}

func TestArrayQueue_Dequeue(t *testing.T) {
    q := NewArrayQueue(5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    q.Enqueue(4)
    q.Enqueue(5)
    q.Enqueue(6)
    t.Log(q)
    q.Dequeue()
    t.Log(q)
    q.Dequeue()
    t.Log(q)
    q.Dequeue()
    t.Log(q)
    q.Dequeue()
    t.Log(q)
    q.Dequeue()
    t.Log(q)
}