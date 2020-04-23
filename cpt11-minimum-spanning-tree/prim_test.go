package cpt11_minimum_spanning_tree

import (
	"imooc_graph_go/graph"
	"reflect"
	"sort"
	"testing"
)

func TestPrim(t *testing.T) {
	g := graph.NewWeightedAdjList("g.txt")
	r := Prim(g)
	expected := []*graph.WeightedEdge{
		{V: 1, W: 2, Weight: 1},
		{V: 4, W: 3, Weight: 1},
		{V: 0, W: 1, Weight: 2},
		{V: 0, W: 5, Weight: 2},
		{V: 1, W: 4, Weight: 3},
		{V: 3, W: 6, Weight: 5},
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i].Weight < r[j].Weight || (r[i].Weight == r[j].Weight && r[i].V < r[j].V)
	})
	if !reflect.DeepEqual(expected, r) {
		t.Errorf("expect %v, but got %v", expected, r)
	}
}
