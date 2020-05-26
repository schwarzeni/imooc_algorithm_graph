package cpt13_directed_graph

import (
	"imooc_graph_go/graph"
	"reflect"
	"testing"
)

func TestTopologicalSort2(t *testing.T) {
	type args struct {
		g graph.DirectedGraph
	}
	tests := []struct {
		name      string
		args      args
		wantOrder []int
	}{
		{name: "test1", args: args{graph.NewDirecteddjMatrix("tsg.txt")}, wantOrder: []int{8, 2, 3, 6, 0, 1, 4, 5, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOrder := TopologicalSort2(tt.args.g); !reflect.DeepEqual(gotOrder, tt.wantOrder) {
				t.Errorf("TopologicalSort2() = %v, want %v", gotOrder, tt.wantOrder)
			}
		})
	}
}
