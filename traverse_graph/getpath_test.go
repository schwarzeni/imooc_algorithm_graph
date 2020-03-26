package traverse_graph

import (
	"fmt"
	"imooc_graph_go/graph"
	"testing"
)

func TestGetPath(t *testing.T) {
	g := graph.NewAdjList("g.txt")
	path, _ := GetPath(g, 0, 6)
	fmt.Println(path)
}
