package queues

import "fmt"

type ArrayQueue struct {
    head     int
    tail     int
    capacity int
    q        []interface{}
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{0, 0, n, make([]interface{}, n)}
}

func (a *ArrayQueue) Enqueue(v interface{}) bool {
	if a.tail == a.capacity { // 队列已满
		if a.head == 0 { // 没有空闲空间
			return false
		}
		// 队列已满 但是有空闲空间 进行数据搬移
		for i := a.head; i < a.tail; i ++ { // 从 head 指向的头部数据开始搬移
			a.q[i - a.head] = a.q[i]
		}
		// 数据搬移后 更新 head tail
		a.tail = a.tail - a.head
		a.head = 0
	}
	a.q[a.tail] = v
	a.tail ++
	return true
}

func (a *ArrayQueue) Dequeue() interface{} {
	if a.head == a.tail { // 队列为空
		return nil
	}
	v := a.q[a.head]
	a.head ++
	return v
}

func (a *ArrayQueue) String() string {
	if a.head == a.tail {
		return "empty queue"
	}
	result := "head"
	for i := a.head; i <= a.tail - 1; i ++ {
		result += fmt.Sprintf("<-%+v", a.q[i])
	}
	result += "<-tail"
	return result
}