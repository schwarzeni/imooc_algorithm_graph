package cpt8

import (
	"fmt"
	"imooc_graph_go/graph"
)

// 默认是连通图
func bridges(g graph.Graph, startVertex int) (result []string) {
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

		for _, v := range vs {
			if !visited[v] {
				order += 1
				dfs(currentNode, v, order)
			}
		}

		for _, v := range vs {
			if v != parentNode {
				if low[v] < low[currentNode] {
					// 相邻节点所能到达的最小序号比当前节点小，则更新当前节点的 low 值
					low[currentNode] = low[v]
				} else if low[v] > ord[currentNode] {
					// 相邻节点的所能到达的最小序号比当前节点序号要大，为桥
					result = append(result, fmt.Sprintf("%d-%d", v, currentNode))
				}
			}
		}
	}

	dfs(startVertex, startVertex, 0)

	return
}
