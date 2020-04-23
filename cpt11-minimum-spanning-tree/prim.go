package cpt11_minimum_spanning_tree

import (
	"imooc_graph_go/graph"
	"math"
)

func Prim(g graph.WeightedGraph) []*graph.WeightedEdge {
	var (
		mst     []*graph.WeightedEdge
		visited = make([]bool, g.V())
	)

	if graph.CC(g, 0) > 1 {
		return nil
	}

	visited[0] = true

	for i := 1; i < g.V(); i++ {
		me := graph.NewWeightedEdge(-1, -1, math.MaxInt64)
		for v := 0; v < g.V(); v++ {
			if visited[v] {
				adjs, _ := g.Adj(v)
				for _, adj := range adjs {
					if !visited[adj] && g.GetWeight(v, adj) < me.Weight {
						me = graph.NewWeightedEdge(v, adj, g.GetWeight(v, adj))
					}
				}
			}
		}
		mst = append(mst, me)
		visited[me.V], visited[me.W] = true, true
	}

	return mst
}
