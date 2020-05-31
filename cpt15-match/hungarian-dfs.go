package cpt15_match

import (
	"imooc_graph_go/graph"
	"imooc_graph_go/traverse_graph"
)

// HungarianDFS 匈牙利算法，DFS实现
type HungarianDFS struct {
	g           graph.Graph
	matching    []int
	visited     []bool
	maxMatching int
}

func (h *HungarianDFS) MaxMatching() int {
	return h.maxMatching
}

func (h *HungarianDFS) dfs(v int) bool {
	h.visited[v] = true
	adj, _ := h.g.Adj(v)
	for _, w := range adj { // 二分图另一半的点
		if !h.visited[w] {
			h.visited[w] = true
			if h.matching[w] == -1 || h.dfs(h.matching[w]) {
				h.matching[v] = w
				h.matching[w] = v
				return true
			}
		}
	}
	return false
}

func NewHungarianDFS(g graph.Graph) *HungarianDFS {
	b := traverse_graph.NewBipartite(g)
	if !b.IsBitpartite() {
		panic("非二分图")
	}
	h := &HungarianDFS{g: g, matching: make([]int, g.V())}
	for idx, _ := range h.matching {
		h.matching[idx] = -1
	}
	for _, v := range b.GetBitpartite()[0] {
		if h.matching[v] == -1 {
			h.visited = make([]bool, g.V())
			if h.dfs(v) {
				h.maxMatching++
			}
		}
	}

	return h
}
