package graphs

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	graph := NewGraph(8)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)
	graph.AddEdge(4, 6)
	graph.AddEdge(5, 7)
	graph.AddEdge(6, 7)
	t.Run("Graph BSF", func(t *testing.T) {
		graph.BFS(0, 7)
		fmt.Println()
		graph.BFS(1, 3)
		fmt.Println()
	})

	t.Run("Graph DSF", func(t *testing.T) {
		graph.DFS(0, 7)
		fmt.Println()
		graph.DFS(1, 3)
		fmt.Println()
	})
}