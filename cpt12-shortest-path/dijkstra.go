package cpt12_shortest_path

import (
	"imooc_graph_go/graph"
	"math"
)

func Dijkstra(g graph.WeightedGraph, s int) []int {
	visited := make([]bool, g.V())
	dis := make([]int, g.V())
	for idx, _ := range dis {
		dis[idx] = math.MaxInt64
	}
	dis[s] = 0

	for {
		cur, curdis := -1, math.MaxInt64
		for v := 0; v < g.V(); v++ {
			if !visited[v] && dis[v] < curdis {
				curdis = dis[v]
				cur = v
			}
		}
		if cur == -1 {
			break
		}
		visited[cur] = true
		adj, _ := g.Adj(cur)
		for _, v := range adj {
			if !visited[v] {
				if dis[cur]+g.GetWeight(cur, v) < dis[v] {
					dis[v] = dis[cur] + g.GetWeight(cur, v)
				}
			}
		}
	}
	return dis
}
