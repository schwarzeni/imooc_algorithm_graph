package traverse_graph

import (
	"fmt"
	"imooc_graph_go/graph"
	"testing"
)

func TestDFS(t *testing.T) {
	g := graph.NewAdjList("g.txt")
	vertexList := DFS(g, 0)
	fmt.Println(vertexList) // [0 1 3 2 6 5 4]
}
