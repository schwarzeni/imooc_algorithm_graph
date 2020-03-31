package traverse_graph

import (
	"fmt"
	"imooc_graph_go/graph"
	"testing"
)

func TestGetPathDFS(t *testing.T) {
	g := graph.NewAdjList("g.txt")
	path, _ := GetPathDFS(g, 0, 6)
	fmt.Println(path) // [6 2 3 1 0]
}

func TestGetPathBFS(t *testing.T) {
	g := graph.NewAdjList("g.txt")
	path, _ := GetPathBFS(g, 0, 6)
	fmt.Println(path) // [6 2 0]

	path, _ = GetPathBFS(g, 0, 5)
	fmt.Println(path) // [5 3 1 0]
}
