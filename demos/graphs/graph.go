package graphs

import (
	"container/list"
	"fmt"
)

// adjacency table, 无向图
type Graph struct {
	v      int // 顶点的个数
	adj    []*list.List  // 邻接表
}

// init graph according to capacity
func NewGraph(v int) *Graph {
	graphh := &Graph{}
	graphh.v = v
	graphh.adj = make([]*list.List, v)
	for i := range graphh.adj {
		graphh.adj[i] = list.New()
	}
	return graphh
}

func (g *Graph) AddEdge(s, t int) {
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}

// s 起始顶点，t 终止顶点
func (g *Graph) BFS(s, t int) {
	if s == t {
		return
	}

	// init prev 记录搜索路径
	prev := make([]int, g.v)
	for index := range prev {
		prev[index] = -1
	}

	// search by queue
	var queue []int // 一个队列，存储已经被访问、但相连的顶点还没有被访问的顶点
	visited := make([]bool, g.v) // 记录已经被访问的顶点，避免顶点被重复访问
	queue = append(queue, s)
	visited[s] = true  // 顶点设置为 true 表示已经被访问
	isFound := false
	for len(queue) > 0 && !isFound {
		top := queue[0]
		queue = queue[1:]
		linkedList := g.adj[top]
		for e := linkedList.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visited[k] {
				prev[k] = top
				if k == t {
					isFound = true
					break
				}
				queue = append(queue, k)
				visited[k] = true
			}
		}
	}

	if isFound {
		printPrev(prev, s, t)
	} else {
		fmt.Printf("no path found from %d to %d\n", s, t)
	}
}

func (g *Graph) DFS(s, t int) {
	if s == t {
		return
	}
	// init prev 记录搜索路径
	prev := make([]int, g.v)
	for index := range prev {
		prev[index] = -1
	}

	visited := make([]bool, g.v) // 记录已经被访问的顶点，避免顶点被重复访问
	visited[s] = true  // 顶点设置为 true 表示已经被访问

	g.recurse(s, t, prev, visited, false)

	printPrev(prev, s, t)
}

func (g *Graph) recurse(s int, t int, prev []int, visited []bool, isFound bool) {

	if isFound {
		return
	}

	visited[s] = true

	if s == t {
		isFound = true
		return
	}

	linkedList := g.adj[s]
	for e := linkedList.Front(); e != nil; e = e.Next() {
		k := e.Value.(int)
		if !visited[k] {
			prev[k] = s
			g.recurse(k, t, prev, visited, false)
		}
	}

}
// print path recursively
func printPrev(prev []int, s int, t int) {

	if t == s || prev[t] == -1 {
		fmt.Printf("%d ", t)
	} else {
		printPrev(prev, s, prev[t])
		fmt.Printf("%d ", t)
	}

}