package lru_cache

const (
    hostbit = uint64(^uint(0)) == ^uint64(0)
    LENGTH  = 100
)

type lruNode struct {
    prev *lruNode
    next *lruNode

    key   int // lru key
    value int // lru value

    hnext *lruNode // 拉链
}

type LRUCache struct {
    node []lruNode // hash list

    head *lruNode // lru head node
    tail *lruNode // lru tail node

    capacity int //
    used     int //
}

func NewLRUCache(capacity int) LRUCache {
    return LRUCache{
        node:     make([]lruNode, LENGTH),
        head:     nil,
        tail:     nil,
        capacity: capacity,
        used:     0,
    }
}

func (l *LRUCache) Get(key int) int {
    if l.tail == nil {
        return -1
    }

    if tmp := l.searchNode(key); tmp != nil {
        l.moveToTail(tmp)
        return tmp.value
    }
    return -1
}

func (l *LRUCache) Put(key int, value int) {
    // 1. 首次插入数据
    // 2. 插入数据不在 LRU 中
    // 3. 插入数据在 LRU 中
    // 4. 插入数据不在 LRU 中, 并且 LRU 已满

    if tmp := l.searchNode(key); tmp != nil {
        tmp.value = value
        l.moveToTail(tmp)
        return
    }
    l.addNode(key, value)

    if l.used > l.capacity {
        l.delNode()
    }
}

func (l *LRUCache) searchNode(key int) *lruNode {
    if l.tail == nil {
        return nil
    }

    // 查找
    tmp := l.node[hash(key)].hnext
    for tmp != nil {
        if tmp.key == key {
            return tmp
        }
        tmp = tmp.hnext
    }
    return nil
}

func (l *LRUCache) addNode(key int, value int) {
    newNode := &lruNode{
        key:   key,
        value: value,
    }

    tmp := &l.node[hash(key)]
    newNode.hnext = tmp.hnext
    tmp.hnext = newNode
    l.used++

    if l.tail == nil {
        l.tail, l.head = newNode, newNode
        return
    }
    l.tail.next = newNode
    newNode.prev = l.tail
    l.tail = newNode
}

func (l *LRUCache) delNode() {
    if l.head == nil {
        return
    }
    prev := &l.node[hash(l.head.key)]
    tmp := prev.hnext

    for tmp != nil && tmp.key != l.head.key {
        prev = tmp
        tmp = tmp.hnext
    }
    if tmp == nil {
        return
    }
    prev.hnext = tmp.hnext
    l.head = l.head.next
    l.head.prev = nil
    l.used--
}

func (l *LRUCache) moveToTail(node *lruNode) {
    if l.tail == node {
        return
    }
    if l.head == node {
        l.head = node.next
        l.head.prev = nil
    } else {
        node.next.prev = node.prev
        node.prev.next = node.next
    }

    node.next = nil
    l.tail.next = node
    node.prev = l.tail

    l.tail = node
}

func hash(key int) int {
    if hostbit {
        return (key ^ (key >> 32)) & (LENGTH - 1)
    }
    return (key ^ (key >> 16)) & (LENGTH - 1)
}