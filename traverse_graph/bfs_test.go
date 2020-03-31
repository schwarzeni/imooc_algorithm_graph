package traverse_graph

import (
	"fmt"
	"imooc_graph_go/graph"
	"testing"
)

func TestBFS(t *testing.T) {
	g := graph.NewAdjList("g.txt")
	vertexList := BFS(g, 0)
	fmt.Println(vertexList)
}
