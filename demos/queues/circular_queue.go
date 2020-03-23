package queues

import "fmt"

type CircularQueue struct {
    q        []interface{}
    capacity int
    head     int
    tail     int
}

func NewCircularQueue(cap int) *CircularQueue {
    if cap == 0 {
        return nil
    }
    return &CircularQueue{
        q:        make([]interface{}, cap),
        capacity: cap,
        head:     0,
        tail:     0,
    }
}

func (c *CircularQueue) Enqueue(v interface{}) bool {
    if (c.tail + 1) % c.capacity == c.head { // 队列已满
        return false
    }
    c.q[c.tail] = v
    c.tail = (c.tail + 1) % c.capacity
    return true
}

func (c *CircularQueue) Dequeue() interface{} {
    if c.head == c.tail { // 队列为空
        return nil
    }
    v := c.q[c.head]
    c.head = (c.head + 1) % c.capacity
    return v
}

func (c *CircularQueue) String() string {
    if c.head == c.tail {
        return "empty queue"
    }
    result := "head"
    var i = c.head
    for true {
        result += fmt.Sprintf("<-%+v", c.q[i])
        i = (i + 1) % c.capacity
        if i == c.tail {
            break
        }
    }
    result += "<-tail"
    return result
}
