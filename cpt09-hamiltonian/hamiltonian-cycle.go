package cpt09_hamiltonian

import "imooc_graph_go/graph"

func hamiltonianCycle(g graph.Graph) (result []int) {
	var (
		visited      = make([]bool, g.V())
		pre          = make([]int, g.V())
		endVertex    int
		startVertex  = 0
		visitedCount int
		dfs          func(currVertex int) bool
	)

	dfs = func(currVertex int) bool {
		adjvs, _ := g.Adj(currVertex)
		for _, v := range adjvs {
			if v == startVertex && visitedCount == g.V() {
				endVertex = currVertex
				return true
			}
			if !visited[v] {
				visited[v] = true
				visitedCount += 1

				if dfs(v) {
					pre[v] = currVertex
					return true
				}

				visitedCount -= 1
				visited[v] = false
			}
		}
		return false
	}

	visited[startVertex] = true
	visitedCount = 1

	if dfs(startVertex) {
		result = append(result, startVertex)
		v := endVertex
		for v != startVertex {
			result = append(result, v)
			v = pre[v]
		}
		result = append(result, v)
	}
	return
}
