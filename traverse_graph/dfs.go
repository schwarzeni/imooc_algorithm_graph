package traverse_graph

import "imooc_graph_go/graph"

// DFS 图的深度优先遍历
func DFS(g graph.Graph, startVertex int) (result []int) {
	var (
		isVisited = make([]bool, g.V())
		dfs       func(currV int)
	)
	dfs = func(currV int) {
		isVisited[currV] = true
		result = append(result, currV)
		adjVs, _ := g.Adj(currV)
		for _, v := range adjVs {
			if !isVisited[v] {
				dfs(v)
			}
		}
	}
	dfs(startVertex)

	// 联通分量大于1的情况
	for i := 0; i < g.V(); i++ {
		if !isVisited[i] {
			dfs(i)
		}
	}
	return
}
