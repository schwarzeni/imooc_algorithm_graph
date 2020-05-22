package cpt12_shortest_path

import (
	"imooc_graph_go/graph"
	"reflect"
	"testing"
)

func TestFloyed(t *testing.T) {
	type args struct {
		g graph.WeightedGraph
	}
	tests := []struct {
		name                 string
		args                 args
		wantDis              [][]int
		wantHasNegativeCycle bool
	}{
		{
			name: "test1",
			args: args{
				g: graph.NewWeightedAdjList("g.txt"),
			},
			wantHasNegativeCycle: false,
			wantDis: [][]int{
				{0, 3, 2, 5, 6},
				{3, 0, 1, 2, 3},
				{2, 1, 0, 3, 4},
				{5, 2, 3, 0, 1},
				{6, 3, 4, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDis, gotHasNegativeCycle := Floyed(tt.args.g)
			if !reflect.DeepEqual(gotDis, tt.wantDis) {
				t.Errorf("Floyed() gotDis = %v, want %v", gotDis, tt.wantDis)
			}
			if gotHasNegativeCycle != tt.wantHasNegativeCycle {
				t.Errorf("Floyed() gotHasNegativeCycle = %v, want %v", gotHasNegativeCycle, tt.wantHasNegativeCycle)
			}
		})
	}
}
