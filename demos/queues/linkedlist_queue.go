package queues

import "fmt"

type QueueNode struct {
    data interface{}
    next *QueueNode
}

type LinkedListQueue struct {
    head   *QueueNode
    tail   *QueueNode
    Length int
}

func NewLinkedListQueue() *LinkedListQueue {
    return &LinkedListQueue{
        head:   nil,
        tail:   nil,
        Length: 0,
    }
}

func NewQueueNode(v interface{}) *QueueNode {
    return &QueueNode{
        data: v,
        next: nil,
    }
}

func (l *LinkedListQueue) Enqueue(v interface{}) {
    newNode := NewQueueNode(v)
    if l.head == nil {
        l.head = newNode
        l.tail = newNode
        l.Length++
        return
    }
    l.tail.next = newNode
    l.tail = newNode
    l.Length++
}

func (l *LinkedListQueue) Dequeue() interface{} {
    if l.head == nil {
        return nil
    }
    v := l.head.data
    l.head = l.head.next
    l.Length--
    return v
}

func (l *LinkedListQueue) String() string {
    if l.head == nil {
        return "empty queue"
    }
    result := "head"
    for cur := l.head; cur != nil; cur = cur.next {
        result += fmt.Sprintf("<-%+v", cur.data)
    }
    result += "<-tail"
    return result
}
