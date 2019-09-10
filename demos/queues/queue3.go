package queues

type LinkedListQueue struct {
	head       *ListNode
	tail       *ListNode
	Length     int
}

type ListNode struct {
	data        interface{}
	next        *ListNode
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

// 入队
func (l *LinkedListQueue) Enqueue(item interface{}) {
	node := &ListNode{item, nil}
	if nil == l.tail { // 空队列
		l.tail = node
		l.head = node
	} else {
		l.tail.next = node
		l.tail = node
	}
    l.Length ++
}

// 出队
func (l *LinkedListQueue) Dequeue() interface{} {
	if l.head == nil { // 队列为空
	    l.tail = nil // clean
		return nil
	}
    item := l.head.data
    l.head = l.head.next
    l.Length --
    return  item
}
