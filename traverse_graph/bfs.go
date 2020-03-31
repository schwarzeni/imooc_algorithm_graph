package traverse_graph

import "imooc_graph_go/graph"

// BFS 图的深度优先遍历
func BFS(g graph.Graph, startVertex int) (result []int) {
	var (
		isVisited = make([]bool, g.V())
		bfs       func(currV int)
	)

	bfs = func(currV int) {
		var (
			queue []int
			qIdx  int
		)
		queue = append(queue, currV)
		for len(queue) > qIdx {
			currV = queue[qIdx]
			result = append(result, currV)
			qIdx += 1
			vs, _ := g.Adj(currV)
			for _, v := range vs {
				if !isVisited[v] {
					queue = append(queue, v)
					isVisited[v] = true
				}
			}
		}
	}

	isVisited[startVertex] = true
	bfs(startVertex)

	for v := 0; v < g.V(); v++ {
		if !isVisited[v] {
			bfs(v)
		}
	}

	return
}
