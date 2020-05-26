package cpt13_directed_graph

import (
	"imooc_graph_go/graph"
	"testing"
)

func TestCycleDetection(t *testing.T) {
	type args struct {
		g graph.Graph
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{graph.NewDirecteddjMatrix("dug.txt")}, want: false},
		{name: "test2", args: args{graph.NewDirecteddjMatrix("dug-cycle.txt")}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CycleDetection(tt.args.g); got != tt.want {
				t.Errorf("CycleDetection() = %v, want %v", got, tt.want)
			}
		})
	}
}
