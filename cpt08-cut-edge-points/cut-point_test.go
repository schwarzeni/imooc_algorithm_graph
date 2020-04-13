package cpt8

import (
	"fmt"
	"imooc_graph_go/graph"
	"testing"
)

func Test_cutPoints(t *testing.T) {
	g1 := graph.NewAdjList("cutpoints-1-g.txt")
	result1 := CutPoints(g1, 0)
	fmt.Println(result1) // expected [3 5]

	g2 := graph.NewAdjList("cutpoints-2-g.txt")
	result2 := CutPoints(g2, 0)
	fmt.Println(result2) // expected [0 1]

	g3 := graph.NewAdjList("cutpoints-3-g.txt")
	result3 := CutPoints(g3, 0)
	fmt.Println(result3) // expected [0 1 2 6]

	g4 := graph.NewAdjList("cutpoints-4-g.txt")
	result4 := CutPoints(g4, 0)
	fmt.Println(result4) // expected 2
}
