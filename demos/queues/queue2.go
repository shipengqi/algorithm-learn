package queues

type ArrayQueue struct {
	items      []string
	n          int       // 数组 size
	head       int       // 队列头下标
	tail       int       // 队列尾下标
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{make([]string, capacity), capacity, 0, 0}
}

// 入队
func (a *ArrayQueue) Enqueue(item string) bool {
	if a.tail == a.n { // 队列已满
		if a.head == 0 { // 没有空闲空间
			return false
		}
		// 数据搬移
		for i := a.head; i < a.tail; i ++ { // 从 head 指向的头部数据开始搬移
			a.items[i - a.head] = a.items[i]
		}
		// 数据搬移后 更新 head tail
		a.tail -= a.head
		a.head = 0
	}
	a.items[a.tail] = item
	a.tail ++
	return true
}

// 出队
func (a *ArrayQueue) Dequeue() string {
	if a.head == a.tail { // 队列为空
		return ""
	}
	item := a.items[a.head]
	a.head ++
	return item
}
