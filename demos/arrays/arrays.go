package arrays

import (
    "errors"
    "fmt"
)

/*
* int 数组，数组插入，删除
 */

type Array struct {
    length uint
    data   []int
}

func NewArray(capacity uint) *Array {
    if capacity == 0 {
        return nil
    }
    return &Array{
        length: 0,
        data:   make([]int, capacity, capacity),
    }
}

func (a *Array) Len() uint {
    return a.length
}

func (a *Array) Find(index uint) (int, error) {
    if a.isOutOffRange(index) {
        return 0, errors.New("index out of range")
    }
    return a.data[index], nil
}

func (a *Array) Insert(index uint, value int) error {
    if a.isFull() {
        return errors.New("full array")
    }
    if a.isOutOffRange(index) {
        return errors.New("index out of range")
    }

    for i := a.Len(); i > index; i -- {
        a.data[i] = a.data[i - 1]
    }
    a.data[index] = value
    a.length ++
    return nil
}

func (a *Array) Delete(index uint) (int, error) {
    if a.isOutOffRange(index) {
        return 0, errors.New("index out of range")
    }
    temp := a.data[index]
    for i := index; i < a.Len() - 1; i ++ {
        a.data[i] = a.data[i + 1]
    }
    a.length --
    return temp, nil
}

func (a *Array) Append(v int) error {
    return a.Insert(a.Len(), v)
}

func (a *Array) Print() {
    var format string
    for i := uint(0); i < a.Len(); i++ {
        format += fmt.Sprintf("|%+v", a.data[i])
    }
    fmt.Println(format)
}

func (a *Array) isOutOffRange(index uint) bool {
    if index >= uint(cap(a.data)) {
        return true
    }
    return false
}

func (a *Array) isFull() bool {
    if a.Len() == uint(cap(a.data)) {
        return true
    }
    return false
}