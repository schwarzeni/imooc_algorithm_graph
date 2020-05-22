package cpt12_shortest_path

import (
	"imooc_graph_go/graph"
	"math"
)

func BellmanFord(s int, g graph.WeightedGraph) (dis []int, hasNegativeCycle bool) {
	dis = make([]int, g.V())
	for idx, _ := range dis {
		dis[idx] = math.MaxInt64
	}
	dis[s] = 0

	for pass := 0; pass < g.V(); pass++ {
		for v := 0; v < g.V(); v++ {
			adjs, _ := g.Adj(v)
			for _, w := range adjs {
				if dis[v] != math.MaxInt64 && dis[v]+g.GetWeight(v, w) < dis[w] {
					dis[w] = dis[v] + g.GetWeight(v, w)
				}
			}
		}
	}
	// check if there is a negative cycle
	for v := 0; v < g.V(); v++ {
		adjs, _ := g.Adj(v)
		for _, w := range adjs {
			if dis[v] != math.MaxInt64 && dis[v]+g.GetWeight(v, w) < dis[w] {
				return nil, true
			}
		}
	}
	return
}
