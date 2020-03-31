package traverse_graph

import (
	"fmt"
	"imooc_graph_go/graph"
	"testing"
)

func Test_getShortedPath(t *testing.T) {
	g := graph.NewAdjList("g.txt")
	vertexList := GetUSSSPath(g, 0)
	fmt.Println(vertexList) // [0 1 1 2 2 3 2]
}
