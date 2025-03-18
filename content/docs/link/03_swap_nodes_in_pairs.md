---
title: 两两交换链表中的节点
weight: 3
---

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例：

- 输入：`head = [1,2,3,4]`
- 输出：`[2,1,4,3]`

## 实现

```go
func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{ // 定义一个虚拟头结点
        Val: -1,
        Next: head,
    }
    cur := dummy
    for cur.Next != nil &&  cur.Next.Next != nil {
        tmpNode1 := cur.Next // 记录临时节点
        tmpNode3 := cur.Next.Next.Next // 记录临时节点
        
        cur.Next = cur.Next.Next // 步骤一， 将 cur 的 next 指向第二个节点
        cur.Next.Next = tmpNode1 // 步骤二，将第二个节点的 next 指向第一个节点
        cur.Next.Next.Next = tmpNode3 // 步骤三，将第一个节点的 next 指向第三个节点
    
        cur = cur.Next.Next // cur 移动两位，准备下一轮交换
    }

    return dummy.Next
}
```

### 递归法

```go
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    next := head.Next // 记录下一个节点
    head.Next = swapPairs(next.Next) // 递归处理下一个节点
    next.Next = head // 将下一个节点的 next 指向当前节点
    return next // 返回新的头结点
}
```