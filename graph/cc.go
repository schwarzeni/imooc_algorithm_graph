package graph

// 求联通分量
func CC(g Graph, startVertex int) (count int) {
	var (
		isVisited = make([]bool, g.V())
		dfs       func(currV int)
	)
	dfs = func(currV int) {
		isVisited[currV] = true
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
			count++
		}
	}
	return
}
