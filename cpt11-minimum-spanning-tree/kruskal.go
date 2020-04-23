package cpt11_minimum_spanning_tree

import (
	"imooc_graph_go/graph"
	"sort"
)

func Kruskal(g graph.WeightedGraph) []*graph.WeightedEdge {
	var (
		edges []*graph.WeightedEdge
		mst   []*graph.WeightedEdge
		uf    = NewUF(g.E())
	)

	if graph.CC(g, 0) > 1 {
		return nil
	}

	// 添加边到数组中，并从小到大排序
	for v := 0; v < g.V(); v++ {
		adjs, _ := g.Adj(v)
		for _, adj := range adjs {
			if v < adj {
				edges = append(edges, graph.NewWeightedEdge(v, adj, g.GetWeight(v, adj)))
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].Weight < edges[j].Weight })

	// 使用并查集检测环是否存在
	// 环的意思就是在一个集合中
	for _, e := range edges {
		if !uf.isConnected(e.V, e.W) {
			mst = append(mst, e)
			uf.unionElements(e.V, e.W)
		}
	}

	return mst
}
