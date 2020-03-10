package linklist

import "fmt"

type ListNode struct {
    data interface{}
    next *ListNode
}

type SingleLinkList struct {
    head   *ListNode
    length uint
}

func NewListNode(v interface{}) *ListNode {
    return &ListNode{ v, nil}
}

func NewSingleLinkList() *SingleLinkList {
    return &SingleLinkList{
        head:   NewListNode(0),
        length: 0,
    }
}

func (l *ListNode) GetNext() *ListNode {
    return l.next
}

func (l *ListNode) GetValue() interface{} {
    return l.data
}

// 在某个节点后面插入节点
func (s *SingleLinkList) InsertAfter(p *ListNode, v interface{}) bool {
    if p == nil {
        return false
    }
    newNode := NewListNode(v)
    oldNode := p.next
    p.next = newNode
    newNode.next = oldNode
    s.length ++
    return true
}

// 在某个节点前面插入节点
func (s *SingleLinkList) InsertBefore(p *ListNode, v interface{}) bool {
    if nil == p || p == s.head {
        return false
    }
    cur := s.head.next
    pre := s.head
    for nil != cur {
        if cur == p {
            break
        }
        pre = cur
        cur = cur.next
    }
    if nil == cur {
        return false
    }
    newNode := NewListNode(v)
    pre.next = newNode
    newNode.next = cur
    s.length ++
    return true
}

// 在链表头部插入节点
func (s *SingleLinkList) InsertToHead(v interface{}) bool {
    return s.InsertAfter(s.head, v)
}

// 在链表尾部插入节点
func (s *SingleLinkList) InsertToTail(v interface{}) bool {
    cur := s.head
    for nil != cur.next {
        cur = cur.next
    }
    return s.InsertAfter(cur, v)
}

// 通过索引查找节点
func (s *SingleLinkList) FindByIndex(index uint) *ListNode {
    if index >= s.length {
        return nil
    }
    cur := s.head
    var i uint = 0
    for ; i < index; i ++ {
        cur = cur.next
    }
    return cur
}

// 删除传入的节点
func (s *SingleLinkList) DeleteNode(p *ListNode) bool {
    if p == nil {
        return false
    }
    cur := s.head.next
    pre := s.head
    for nil != cur {
        if cur == p {
            break
        }
        pre = cur
        cur = cur.next
    }
    if nil == cur {
        return false
    }
    pre.next = cur.next
    p = nil
    s.length--
    return true
}

// 打印链表
func (s *SingleLinkList) Print() {
    cur := s.head.next
    format := ""
    for nil != cur {
        format += fmt.Sprintf("%+v", cur.GetValue())
        cur = cur.next
        if nil != cur {
            format += "->"
        }
    }
    fmt.Println(format)
}

