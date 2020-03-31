package traverse_graph

import (
	"imooc_graph_go/graph"
)

// 求单源路径问题
func GetPathDFS(g graph.Graph, startVertex int, endVertex int) (path []int, isFound bool) {
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

func GetPathBFS(g graph.Graph, startVertex int, endVertex int) (path []int, isFound bool) {
	var (
		isVisited = make([]bool, g.V())
		pre       = make([]int, g.V())
		currV     = startVertex
		queue     []int
		qIdx      int
	)
	for idx, _ := range pre {
		pre[idx] = -1
	}
	pre[startVertex] = startVertex

	queue = append(queue, currV)
	for len(queue) > qIdx {
		currV = queue[qIdx]
		if currV == endVertex {
			break
		}
		qIdx += 1
		vs, _ := g.Adj(currV)
		for _, v := range vs {
			if !isVisited[v] {
				queue = append(queue, v)
				isVisited[v] = true
				pre[v] = currV
			}
		}
	}

	path = append(path, endVertex)
	currV = endVertex
	for currV != startVertex {
		if pre[currV] == -1 {
			return nil, false
		}
		path = append(path, pre[currV])
		currV = pre[currV]
	}

	return path, true
}
