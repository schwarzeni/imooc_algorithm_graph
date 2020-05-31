package cpt15_match

import (
	"imooc_graph_go/graph"
	"testing"
)

func TestHungarianBFS_MaxMatching(t *testing.T) {

	tests := []struct {
		name string
		g    graph.Graph
		want int
	}{
		{name: "testing1", g: graph.NewAdjList("g1.txt"), want: 3},
		{name: "testing2", g: graph.NewAdjList("g2.txt"), want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHungarianBFS(tt.g)
			if got := h.MaxMatching(); got != tt.want {
				t.Errorf("MaxMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
