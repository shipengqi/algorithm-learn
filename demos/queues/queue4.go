package queues

type CircularQueue struct {
	items      []interface{}
	capacity   int       // 数组 size
	head       int       // 队列头下标
	tail       int       // 队列尾下标
}

func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{make([]interface{}, n), n, 0, 0}
}

func (c *CircularQueue) IsEmpty() bool {
    if c.head == c.tail {
    	return true
	}
    return false
}

func (c *CircularQueue) IsFull() bool {
    if c.head == (c.tail + 1) % c.capacity {
    	return true
	}
    return false
}


func (c *CircularQueue) Enqueue(item interface{}) bool {
    if c.IsFull() {
    	return false
	}
    c.items[c.tail] = item
    c.tail = (c.tail + 1) % c.capacity
    return true
}

func (c *CircularQueue) Dequeue() interface{} {
	if c.IsEmpty() {
		return nil
	}
	v := c.items[c.head]
	c.head = (c.head + 1) % c.capacity
	return v
}
