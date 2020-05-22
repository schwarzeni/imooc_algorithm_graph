package cpt12_shortest_path

import (
	"imooc_graph_go/graph"
	"math"
)

func Floyed(g graph.WeightedGraph) (dis [][]int, hasNegativeCycle bool) {
	// 初始化
	dis = make([][]int, g.V())
	for idx, _ := range dis {
		dis[idx] = make([]int, g.V())
		for i, _ := range dis[idx] {
			dis[idx][i] = math.MaxInt64
		}
	}
	for v := 0; v < g.V(); v++ {
		dis[v][v] = 0
		adjs, _ := g.Adj(v)
		for _, w := range adjs {
			dis[v][w] = g.GetWeight(v, w)
		}
	}

	for t := 0; t < g.V(); t++ {
		for v := 0; v < g.V(); v++ {
			for w := 0; w < g.V(); w++ {
				if dis[v][t] == math.MaxInt64 || dis[t][w] == math.MaxInt64 {
					continue
				}
				if l := dis[v][t] + dis[t][w]; l < dis[v][w] {
					dis[v][w] = l
				}
			}
		}
	}

	// 检测负权环
	for v := 0; v < g.V(); v++ {
		if dis[v][v] < 0 {
			return nil, true
		}
	}
	return
}
