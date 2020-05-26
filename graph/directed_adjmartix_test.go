package graph

import (
	"fmt"
	"testing"
)

func TestNewDirecteddjMatrix(t *testing.T) {
	g := newDirectedAdjMatrix("dug.txt")
	fmt.Println(g)
}

func TestDirectedAdjMatrix_Degree(t *testing.T) {
	g := newDirectedAdjMatrix("dug.txt")
	id, od := g.InDegree(1), g.OutDegree(1)
	if id != 1 || od != 2 {
		t.Errorf("expected id=1, od=2, but got id=%v, od=%v", id, od)
	}
}
