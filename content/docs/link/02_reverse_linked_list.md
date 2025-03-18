---
title: 反转链表
weight: 2
---

反转一个单链表，返回反转后的链表。

- 输入：`1->2->3->4->5->NULL`
- 输出：`5->4->3->2->1->NULL`

## 实现

### 双指针法

```go
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var pre *ListNode // 定义一个 pre 指针，初始化为 null
 
    cur := head // 先定义一个 cur 指针，指向头结点
    for cur != nil {
        tmp := cur.Next // 保存一下 cur 的下一个节点，因为接下来要改变 cur->next
        cur.Next = pre // 翻转操作
        pre = cur
        cur = tmp
    }
    return pre
}
```

### 递归法

```go
func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    cur := head
    return reverse(pre, cur)
}

func reverse(pre, cur *ListNode) *ListNode {
    if cur != nil {
        tmp := cur.Next
        cur.Next = pre
        return reverse(cur, tmp)
    }
    return pre
}
```