package cpt15_match

import (
	"imooc_graph_go/graph"
	"reflect"
	"testing"
)

func TestHungarianDFS_MaxMatching(t *testing.T) {
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
			h := NewHungarianDFS(tt.g)
			if got := h.MaxMatching(); got != tt.want {
				t.Errorf("MaxMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHungarianDFS_dfs(t *testing.T) {
	type fields struct {
		g           graph.Graph
		matching    []int
		visited     []bool
		maxMatching int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HungarianDFS{
				g:           tt.fields.g,
				matching:    tt.fields.matching,
				visited:     tt.fields.visited,
				maxMatching: tt.fields.maxMatching,
			}
			if got := h.dfs(tt.args.v); got != tt.want {
				t.Errorf("dfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHungarianDFS(t *testing.T) {
	type args struct {
		g graph.Graph
	}
	tests := []struct {
		name string
		args args
		want *HungarianDFS
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHungarianDFS(tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHungarianDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
