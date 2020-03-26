package traverse_graph

import "imooc_graph_go/graph"

// 求单源路径问题
func GetPath(g graph.Graph, startVertex int, endVertex int) (path []int, isFound bool) {
	var (
		isVisited = make([]bool, g.V())
		getPath   func(sv int, ev int) bool
	)
	getPath = func(sv int, ev int) bool {
		isVisited[sv] = true
		if sv == ev {
			return true
		}
		vs, _ := g.Adj(sv)
		for _, v := range vs {
			if !isVisited[v] && getPath(v, ev) {
				path = append(path, v)
				return true
			}
		}
		return false
	}

	isFound = getPath(startVertex, endVertex)
	path = append(path, startVertex)
	return
}
