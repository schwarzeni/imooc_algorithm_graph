package cpt12_shortest_path

import (
	"imooc_graph_go/graph"
	"reflect"
	"testing"
)

func TestBellmanFord(t *testing.T) {
	type args struct {
		s int
		g graph.WeightedGraph
	}
	tests := []struct {
		name                 string
		args                 args
		wantDis              []int
		wantHasNegativeCycle bool
	}{
		{
			name: "test1",
			args: args{
				g: graph.NewWeightedAdjList("g.txt"),
				s: 0,
			},
			wantHasNegativeCycle: false,
			wantDis:              []int{0, 3, 2, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDis, gotHasNegativeCycle := BellmanFord(tt.args.s, tt.args.g)
			if !reflect.DeepEqual(gotDis, tt.wantDis) {
				t.Errorf("BellmanFord() gotDis = %v, want %v", gotDis, tt.wantDis)
			}
			if gotHasNegativeCycle != tt.wantHasNegativeCycle {
				t.Errorf("BellmanFord() gotHasNegativeCycle = %v, want %v", gotHasNegativeCycle, tt.wantHasNegativeCycle)
			}
		})
	}
}
