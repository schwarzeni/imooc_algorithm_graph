package cpt8

import (
	"imooc_graph_go/graph"
)

// 默认是连通图
func cutPoints(g graph.Graph, startVertex int) (result []int) {
	var (
		ord     = make([]int, g.V())
		low     = make([]int, g.V())
		visited = make([]bool, g.V())
		dfs     func(parentNode int, currentNode int, order int)
	)

	dfs = func(parentNode int, currentNode int, order int) {
		ord[currentNode] = order
		low[currentNode] = order
		visited[currentNode] = true

		vs, _ := g.Adj(currentNode)

		childCount := 0
		for _, v := range vs {
			if !visited[v] {
				order += 1
				dfs(currentNode, v, order)
				childCount += 1
			}
		}

		// 处理根节点的情况
		if startVertex == currentNode {
			if childCount > 1 {
				result = append(result, currentNode)
			}
			return
		}

		for _, v := range vs {
			if v != parentNode {
				if low[v] < low[currentNode] {
					// 相邻节点所能到达的最小序号比当前节点小，则更新当前节点的 low 值
					low[currentNode] = low[v]
				} else if low[v] >= ord[currentNode] {
					// 相邻节点的所能到达的最小序号比当前节点序号大，或者等于
					result = append(result, currentNode)
					break
				}
			}
		}
	}

	dfs(startVertex, startVertex, 0)

	return
}
