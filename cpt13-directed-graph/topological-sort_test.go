package cpt13_directed_graph

import (
	"imooc_graph_go/graph"
	"reflect"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	type args struct {
		g graph.DirectedGraph
	}
	tests := []struct {
		name      string
		args      args
		wantOrder []int
	}{
		{name: "test1", args: args{g: graph.NewDirecteddjMatrix("tsg.txt")}, wantOrder: []int{0, 2, 8, 1, 3, 4, 6, 5, 7}},
		{name: "test2", args: args{g: graph.NewDirecteddjMatrix("dug-cycle.txt")}, wantOrder: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOrder := TopologicalSort(tt.args.g); !reflect.DeepEqual(gotOrder, tt.wantOrder) {
				t.Errorf("TopologicalSort() = %v, want %v", gotOrder, tt.wantOrder)
			}
		})
	}

}
