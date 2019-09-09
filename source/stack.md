---
title: 栈
---

# 栈
关于“栈”的一个例子，就是一摞叠在一起的盘子。我们平时放盘子的时候，都是从下往上一个一个放；取的时候，我们也是从上往下一个一个地依次取，
不能从中间任意抽出。**先进后出**，这就是典型的“栈”结构。

栈是一种“操作受限”的线性表，只允许在一端插入和删除数据。

某个数据集合只涉及在一端插入和删除数据，并且满足后进先出、先进后出的特性，就应该首选“栈”这种数据结构。

## 实现一个栈
```go
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
```

对于出栈操作来说，我们不会涉及内存的重新申请和数据的搬移，所以出栈的时间复杂度仍然是 `O(1)`。但是，对于入栈操作来说，情况就不一样了。
当栈中有空闲空间时，入栈操作的时间复杂度为 `O(1)`。但当空间不够时，就需要重新申请内存和数据搬移，所以时间复杂度就变成了 `O(n)`。

## 栈在表达式求值中的应用
编译器如何利用栈来实现表达式求值，比如：`34+13*9+44-12/3`。

使用两个栈，其中一个保存操作数的栈，另一个是保存运算符的栈。我们从左向右遍历表达式，当遇到数字，就直接压入操作数栈；当遇到运算符，
就与运算符栈的栈顶元素进行比较。

如果比运算符栈顶元素的优先级高，就将当前运算符压入栈；如果比运算符栈顶元素的优先级低或者相同，从运算符栈中取栈顶运算符，从操作数栈
的栈顶取 2 个操作数，然后进行计算，再把计算完的结果压入操作数栈，继续比较。


## 如何实现浏览器的前进、后退功能
使用两个栈，X 和 Y，把首次浏览的页面依次压入栈 X，当点击后退按钮时，再依次从栈 X 中出栈，并将出栈的数据依次放入栈 Y。当点击前进按
钮时，依次从栈Y中取出数据，放入栈 X 中。当栈 X 中没有数据时，那就说明没有页面可以继续后退浏览了。当栈 Y 中没有数据，那就说明没有页面
可以点击前进按钮浏览了。