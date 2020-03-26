package traverse_graph

import (
	"fmt"
	"imooc_graph_go/graph"
	"testing"
)

func TestIsBipartite(t *testing.T) {
	g := graph.NewAdjList("g_bipartite.txt")
	fmt.Println(IsBipartite(g))
}
