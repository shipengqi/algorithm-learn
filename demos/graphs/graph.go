package graphs

import (
	"container/list"
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