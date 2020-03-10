package stacks

import (
    "errors"
    "fmt"
)

type ArrayStack struct {
    top  int
    cap  uint
    data []interface{}
}

func NewArrayStack(capacity uint) *ArrayStack {
    return &ArrayStack{
        top:  -1,
        cap:  capacity,
        data: make([]interface{}, 0, capacity),
    }
}

func (a *ArrayStack) Push(v interface{}) error {
    if a.isFull() {
        return errors.New("full stack")
    }
    a.top ++
    if a.top > len(a.data) - 1 {
        a.data = append(a.data, v)
    } else {
        a.data[a.top] = v
    }
    return nil
}

func (a *ArrayStack) Pop() interface{} {
    if a.isEmpty() {
        return nil
    }
    v := a.data[a.top]
    a.top--
    return v
}

func (a *ArrayStack) isFull() bool {
    if a.top < 0 {
        return false
    }
    if a.top >= int(a.cap)-1 {
        return true
    }
    return false
}

func (a *ArrayStack) isEmpty() bool {
    if a.top < 0 {
        return true
    }
    return false
}

func (a *ArrayStack) Print() {
    if a.isEmpty() {
        fmt.Println("empty stack")
    } else {
        for i := a.top; i >= 0; i-- {
            fmt.Println(a.data[i])
        }
    }
}
