package traverse_graph

import "imooc_graph_go/graph"

type gBipartiteType int

const (
	g_UNVISITED gBipartiteType = iota
	g_TYPE_1
	g_TYPE_2
)

const startV = 0

func IsBipartite(g graph.Graph) bool {
	if g.V() == 0 {
		return false
	}
	vmarked := make([]gBipartiteType, g.V())
	for i := 0; i < g.V(); i++ {
		vmarked[i] = g_UNVISITED
	}

	var dfs func(cv int) bool
	dfs = func(cv int) bool {
		currentType := vmarked[cv]
		oppositeType := g_TYPE_1
		if currentType == g_TYPE_1 {
			oppositeType = g_TYPE_2
		}
		vs, _ := g.Adj(cv)
		for _, v := range vs {
			if vmarked[v] == g_UNVISITED {
				vmarked[v] = oppositeType
				if validated := dfs(v); !validated {
					return false
				}
			} else if vmarked[v] == currentType {
				return false
			}
		}
		return true
	}

	vmarked[startV] = g_TYPE_1
	return dfs(startV)
}
