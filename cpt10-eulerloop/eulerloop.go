package cpt10_eulerloop

import (
	cpt8 "imooc_graph_go/cpt08-cut-edge-points"
	"imooc_graph_go/graph"
)

func isEulerloop(g graph.Graph) bool {
	if len(cpt8.CutPoints(g, 0)) != 1 {
		return false
	}
	for v := 0; v < g.V(); v++ {
		if d, _ := g.Degree(v); d%2 == 1 {
			return false
		}
	}
	return true
}
