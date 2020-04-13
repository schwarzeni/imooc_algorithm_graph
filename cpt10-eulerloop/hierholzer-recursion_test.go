package cpt10_eulerloop

import (
	"fmt"
	"imooc_graph_go/graph"
	"testing"
)

func Test_hierholzerRecursion(t *testing.T) {
	g := graph.NewAdjList("g2.txt")
	result := hierholzerRecursion(g)
	if len(result)-1 != g.E() {
		t.Errorf("length not equal, expect %d, but got %d", g.E(), len(result)-1)
	}
	fmt.Println(result) // [0 3 4 6 7 9 10 8 7 5 2 1 5 4 1 0]
}
