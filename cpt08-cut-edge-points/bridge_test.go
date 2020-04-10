package cpt8

import (
	"fmt"
	"imooc_graph_go/graph"
	"testing"
)

func Test_bridge(t *testing.T) {
	g1 := graph.NewAdjList("bridge-1-g.txt")
	result1 := bridges(g1, 0)
	fmt.Println(result1) // expected [5-3]

	g2 := graph.NewAdjList("bridge-2-g.txt")
	result2 := bridges(g2, 0)
	fmt.Println(result2) // expected [3-1] [4-1] [7-6] [5-2] [6-2] [1-0] [2-0]

	g3 := graph.NewAdjList("bridge-3-g.txt")
	result3 := bridges(g3, 0)
	fmt.Println(result3) // expected [8-6] [7-4] [5-3]

	g4 := graph.NewAdjList("bridge-4-g.txt")
	result4 := bridges(g4, 0)
	fmt.Println(result4) // expected []
}
