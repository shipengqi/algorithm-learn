---
title: 链表
weight: 3
---

链表是一种通过指针串联在一起的数据结构。每一个节点由两部分组成，一个是数据域一个是指针（指向下一个节点的指针），最后一个节点的指针指向 `null`。

常见的链表结构有：单链表，双向链表，循环链表。

链表节点的定义：

```go
type ListNode struct {
    Val int
    Next *ListNode
}
```

## 单链表

![single_link_list](https://raw.gitcode.com/shipengqi/illustrations/files/main/algo/single-link-list.png)

- **头结点用来记录链表的基地址**。有了它，就可以遍历得到整条链表。
- **尾结点**的指针不是指向下一个结点，而是指向一个空地址 NULL，表示这是链表上最后一个结点。

## 循环链表

循环链表只是一种特殊的单链表，唯一的不同就是循环链表的尾节点的指针不指向空地址 NULL，而是头节点。

![ring](https://raw.gitcode.com/shipengqi/illustrations/files/main/algo/ring.png)
 
循环链表的优点是从链尾到链头比较方便。当要处理的数据具有环型结构特点时，就特别适合采用循环链表。比如约瑟夫问题。

## 双向链表

双向链表是一种更加复杂的链表。它中的每个节点除了有一个 `next` 指针指向下一个节点，还有一个 `prev` 指针指向上一个节点。

![doubly-link-list](https://raw.gitcode.com/shipengqi/illustrations/files/main/algo/doubly-link-list.png)

双向链表比单链表需要额外的空间来存储前驱指针，因此同样的数据要比单链表占用更多的内存空间。

## 链表的操作

在链表中插入或者删除一个数据，不需要为了保持内存的连续性而搬移结点，因为链表的存储空间本身就不是连续的。所以，在链表中插入和删除一个数据是非常快速的。链表的插入和删除操作，只需要考虑相邻结点的指针改变，时间复杂度是 `O(1)`。

但是，有利就有弊。链表要想随机访问第 k 个元素，就没有数组那么高效了。因为链表中的数据并非连续存储的，所以无法像数组那样直接通过寻址公式计算出下标对应的内存地址，必须
根据后继指针来遍历每一个节点。所以链表的随机减访问时间复杂度为 `O(n)`。


## 链表实现

单链表实现：

```go
type MyListNode struct {
    Val int
    Next *MyListNode
}

type MyLinkedList struct {
    dummy *MyListNode
    Size int
}


func Constructor() MyLinkedList {
    return MyLinkedList{
        Size: 0,
        dummy: &MyListNode{
            Val: -1,
            Next: nil,
        },
    }
}


func (this *MyLinkedList) Get(index int) int {
    if this == nil || index < 0 || index >= this.Size { // 索引无效则返回 -1，注意这里的索引是从 0 开始的，所以这里的判断条件是 index >= this.Size
        return -1
    }

    cur := this.dummy.Next
    for i := 0; i < index; i++ {
        cur = cur.Next
    }
    return cur.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
    newNode := &MyListNode{Val: val}
    newNode.Next = this.dummy.Next
    this.dummy.Next = newNode
    this.Size++
}


func (this *MyLinkedList) AddAtTail(val int)  {
    cur := this.dummy
    if cur != nil {
        cur = cur.Next
    }
    newNode := &MyListNode{Val: val, Next: nil}
    cur.Next = newNode
    this.Size++
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index < 0 {
        index = 0
    } else if index > this.Size {
        return
    }

    cur := this.dummy
    for i := 0; i < index; i++ {
        cur = cur.Next
    }
    newNode := &MyListNode{Val: val}
    newNode.Next = cur.Next
    cur.Next = newNode
    this.Size++
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= this.Size { // 索引无效则直接返回
        return
    }
    cur := this.dummy
    for i := 0; i < index; i++ {
        cur = cur.Next
    }
    if cur.Next != nil {
        cur.Next = cur.Next.Next
    }
    
    this.Size--
}
```