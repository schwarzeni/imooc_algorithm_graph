package graph

import (
	"fmt"
	"testing"
)

func TestNewDirectedWeightedAdjList(t *testing.T) {
	g := NewDirectedWeightedAdjList("dwg.txt")
	fmt.Println(g)
}

func TestDirectedWeightedAdjList_Degree(t *testing.T) {
	g := NewDirectedWeightedAdjList("dwg.txt")
	id, od := g.InDegree(1), g.OutDegree(1)
	if id != 1 || od != 3 {
		t.Errorf("expected id=1, od=2, but got id=%v, od=%v", id, od)
	}
}
