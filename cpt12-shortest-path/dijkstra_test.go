package cpt12_shortest_path

import (
	"imooc_graph_go/graph"
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	type args struct {
		g graph.WeightedGraph
		s int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{
			g: graph.NewWeightedAdjList("g.txt"),
			s: 0,
		}, want: []int{0, 3, 2, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dijkstra(tt.args.g, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dijkstra() = %v, want %v", got, tt.want)
			}
		})
	}
}
