package main

import "fmt"

type ArrayStack struct {
	items      []string // slice
	count      int      // 栈中元素的个数
	n          int      // 栈的 size
}

// 初始化栈
func (a *ArrayStack) InitArrayStack(n int) {
	a.items = []string{}
	a.count = 0
	a.n = n
}

// 入栈
func (a *ArrayStack) Push(item string) bool {
	if a.count == a.n { // 空间已满
		return false
	}
	a.items = append(a.items, item)
    a.count ++
	return true
}

// 出栈
func (a *ArrayStack) Pop() string {
	if a.count == 0 { // 栈为空
		return ""
	}
	fmt.Println(len(a.items), a.count)
    item := a.items[len(a.items) - 1]
    a.items = a.items[: len(a.items) - 1]
    a.count --
    return item
}

func main() {
   s := ArrayStack{}
   s.InitArrayStack(5)
   s.Push("Hello")
   fmt.Println(s)
   s.Push("World")
   fmt.Println(s)
   i := s.Pop()
   fmt.Println(s, i)
   i = s.Pop()
   fmt.Println(s, i)
}
