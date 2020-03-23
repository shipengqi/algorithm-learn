package stacks

import "fmt"

type StackNode struct {
    data interface{}
    next *StackNode
}

type LinkedListStack struct {
    top    *StackNode
}

func NewLinkedListStack() *LinkedListStack {
    return &LinkedListStack{nil}
}

func (l *LinkedListStack) Push(v interface{}) {
    l.top = &StackNode{v, l.top}
}

func (l *LinkedListStack) Pop() interface{} {
    if l.isEmpty() {
        return nil
    }

    v := l.top.data
    l.top = l.top.next
    return v
}

func (l *LinkedListStack) isEmpty() bool {
    if l.top == nil {
        return true
    }
    return false
}

func (l *LinkedListStack) Print() {
    if l.isEmpty() {
        fmt.Println("empty stack")
    } else {
        cur := l.top
        for nil != cur {
            fmt.Println(cur.data)
            cur = cur.next
        }
    }
}