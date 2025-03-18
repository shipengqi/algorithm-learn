---
title: 移除元素
weight: 1
---

给定一个链表的头节点 head 和一个整数 val ，删除链表中所有满足 `Node.val == val` 的节点，并返回新的头节点 。

示例：

- 输入：`head = [1,2,6,3,4,5,6], val = 6`
- 输出：`[1,2,3,4,5]`

## 实现

移除链表中元素的两种方式：

### 直接使用原来的链表来进行删除操作

移除头结点和移除其他节点的操作是不一样的，因为链表的其他节点都是**通过前一个节点来移除当前节点**，而头结点没有前一个节点。

移除头结点只需要向后移动一位，这样就从链表中移除了一个头结点。

```go
func removeElements(head *ListNode, val int) *ListNode {
	// 如果原链表的头节点为 val 的话，head=head.next，且为持续过程，防止头节点后面的节点也为 Val
	//这里前置循环 并且要判定 head 是否为 nil，防止出错
	for head != nil && head.Val == val {
		head = head.Next
	}
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head

}
```

这种方式下，单链表中**移除头结点**和**移除其他节点**的操作方式是不一样的。可以通过设置一个**虚拟头结点**，这样原链表的所有节点就都可以按照统一的方式进行移除了。

### 设置一个虚拟头结点在进行删除操作

```go
func removeElements(head *ListNode, val int) *ListNode {
    fake := &ListNode{}
    fake.Next = head
    cur := fake
    for cur != nil && cur.Next != nil {
        if cur.Next.Val == val {
            cur.Next = cur.Next.Next
        } else {
            cur = cur.Next
        }
    }
    return fake.Next
}
```