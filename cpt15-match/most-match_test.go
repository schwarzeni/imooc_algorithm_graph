package cpt15_match

import (
	"imooc_graph_go/graph"
	"testing"
)

func TestMostMatchWithNetworkFlow(t *testing.T) {
	type args struct {
		g graph.Graph
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{graph.NewAdjList("g1.txt")}, want: 3},
		{name: "test2", args: args{graph.NewAdjList("g2.txt")}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MostMatchWithNetworkFlow(tt.args.g); got != tt.want {
				t.Errorf("MostMatchWithNetworkFlow() = %v, want %v", got, tt.want)
			}
		})
	}
}
