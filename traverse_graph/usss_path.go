package traverse_graph

import "imooc_graph_go/graph"

//Unweighted Single Source Shortest Path

func GetUSSSPath(g graph.Graph, startVertex int) []int {
	var (
		isVisited = make([]bool, g.V())
		bfs       func(currV int)
		paths     = make([]int, g.V())
	)

	paths[startVertex] = 0

	bfs = func(currV int) {
		var (
			queue []int
			qIdx  int
		)
		queue = append(queue, currV)
		for len(queue) > qIdx {
			currV = queue[qIdx]
			qIdx += 1
			vs, _ := g.Adj(currV)
			for _, v := range vs {
				if !isVisited[v] {
					queue = append(queue, v)
					isVisited[v] = true
					paths[v] = 1 + paths[currV]
				}
			}
		}
	}

	isVisited[startVertex] = true
	bfs(startVertex)

	return paths
}
