package traverse_graph

import (
	"imooc_graph_go/graph"
	"reflect"
	"testing"
)

func TestIsBipartite(t *testing.T) {
	g := graph.NewAdjList("g_bipartite.txt")
	b := NewBipartite(g)
	if !b.IsBitpartite() {
		t.Errorf("expect is bitpartite")
	}
	expect := [2][]int{{0, 3, 4, 6}, {1, 2, 5}}
	if !reflect.DeepEqual(expect, b.GetBitpartite()) {
		t.Errorf("expect %v, but got %v", expect, b.GetBitpartite())
	}

	// fix bug loop
	g = graph.NewAdjList("g_bipartite_2.txt")
	b = NewBipartite(g)
	if !b.IsBitpartite() {
		t.Errorf("expect is bitpartite")
	}
	expect = [2][]int{{0, 1, 2, 5, 7}, {3, 4, 6}}
	if !reflect.DeepEqual(expect, b.GetBitpartite()) {
		t.Errorf("expect %v, but got %v", expect, b.GetBitpartite())
	}
}
