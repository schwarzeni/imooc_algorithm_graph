package cpt15_match

import (
	"imooc_graph_go/graph"
	"imooc_graph_go/traverse_graph"
)

// HungarianBFS 匈牙利算法，BFS实现
type HungarianBFS struct {
	g           graph.Graph
	matching    []int
	maxMatching int
}

func (h *HungarianBFS) bfs(v int) bool {
	pre := make([]int, h.g.V())
	for idx, _ := range pre {
		pre[idx] = -1
	}
	q := []int{v}
	pre[v] = v

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		adj, _ := h.g.Adj(v)
		for _, next := range adj {
			if pre[next] == -1 {
				if h.matching[next] != -1 { // 位于同侧的节点还未匹配
					q = append(q, h.matching[next])
					pre[next] = cur
					pre[h.matching[next]] = next
				} else {
					pre[next] = cur
					augPath := h.getAugPath(pre, v, next)
					for i := 0; i < len(augPath); i += 2 { // 长度为奇数
						h.matching[augPath[i]] = augPath[i+1]
						h.matching[augPath[i+1]] = augPath[i]
					}
					return true
				}
			}
		}
	}
	return false
}

func (h *HungarianBFS) getAugPath(pre []int, start, end int) (p []int) {
	curr := end
	for curr != start {
		p = append(p, curr)
		curr = pre[curr]
	}
	p = append(p, start)
	return p
}

func (h *HungarianBFS) MaxMatching() int {
	return h.maxMatching
}

func NewHungarianBFS(g graph.Graph) *HungarianBFS {
	b := traverse_graph.NewBipartite(g)
	if !b.IsBitpartite() {
		panic("必须为二分图")
	}
	h := &HungarianBFS{g: g}
	h.g = g
	h.matching = make([]int, g.V())
	for idx, _ := range h.matching {
		h.matching[idx] = -1
	}
	for _, v := range b.GetBitpartite()[0] {
		if h.matching[v] == -1 {
			if h.bfs(v) {
				h.maxMatching++
			}
		}
	}
	return h
}
