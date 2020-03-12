package skip_list

import (
    "errors"
    "fmt"
    "math"
    "math/rand"
)

const MAX_LEVEL = 16

type SkipListNode struct {
    v        interface{}
    level    int
    score    int
    forwards []*SkipListNode
}

func NewSkipListNode(v interface{}, score, level int) *SkipListNode {
    return &SkipListNode{
        v:        v,
        level:    level,
        score:    score,
        forwards: make([]*SkipListNode, level, level),
    }
}

type SkipList struct {
    head   *SkipListNode
    level  int
    length int
}

func NewSkipList() *SkipList {
    return &SkipList{
        head:   NewSkipListNode(0, math.MaxInt32, MAX_LEVEL),
        level:  0,
        length: 0,
    }
}

func (s *SkipList) Length(v interface{}) int {
    return s.length
}

func (s *SkipList) Level(v interface{}) int {
    return s.level
}


func (s *SkipList) Insert(v interface{}, score int) error {
    if nil == v {
        return errors.New("nil value")
    }

    // 查找插入的位置
    current := s.head
    trace := [MAX_LEVEL]*SkipListNode{}
    for i := MAX_LEVEL - 1; i >= 0; i -- {
        for nil != current.forwards[i] {
            if current.forwards[i].v == v {
                return errors.New("already exists")
            }
            if current.forwards[i].score > score {
                trace[i] = current
                break
            }
            current = current.forwards[i]
        }

        if nil == current.forwards[i] { // i 层最后一个节点
            trace[i] = current
        }
    }

    // 随机函数计算节点要插入的索引层数
    level := 1
    for i := 1; i < MAX_LEVEL; i++ {
        if rand.Int31() % 7 == 1 {
            level ++
        }
    }

    // 创建一个新的跳表节点
    newNode := NewSkipListNode(v, score, level)

    // 原有节点连接
    for i := 0; i <= level-1; i++ {
        next := trace[i].forwards[i]
        trace[i].forwards[i] = newNode
        newNode.forwards[i] = next
    }

    // 如果当前节点的层数大于之前跳表的层数
    // 更新当前跳表层数
    if level > s.level {
        s.level = level
    }

    // 更新跳表长度
    s.length++

    return nil
}

func (s *SkipList) Find(v interface{}, score int) *SkipListNode {
    if nil == v || s.length == 0 {
        return nil
    }

    current := s.head
    for i := s.level - 1; i >= 0; i -- {
        for nil != current.forwards[i] {
            if current.forwards[i].score == score && current.forwards[i].v == v {
                return current.forwards[i]
            } else if current.forwards[i].score > score {
                break
            }
            current = current.forwards[i]
        }
    }
    return nil
}

func (s *SkipList) Delete(v interface{}, score int) error {
    if nil == v || s.length == 0 {
        return nil
    }


    // 查找前驱节点
    current := s.head
    // 记录前驱路径
    trace := [MAX_LEVEL]*SkipListNode{}
    for i := s.level - 1; i >= 0; i -- {
        trace[i] = s.head
        for nil != current.forwards[i] {
            if current.forwards[i].score == score && current.forwards[i].v == v {
                trace[i] = current
                break
            }
            current = current.forwards[i]
        }
    }

    current = trace[0].forwards[0]
    for i := current.level - 1; i >= 0; i-- {
        if trace[i] == s.head && current.forwards[i] == nil {
            s.level = i
        }

        if nil != trace[i].forwards[i] {
            trace[i].forwards[i] = trace[i].forwards[i].forwards[i]
        }
    }

    s.length--

    return nil
}

func (s *SkipList) String() string {
    return fmt.Sprintf("level:%+v, length:%+v", s.level, s.length)
}