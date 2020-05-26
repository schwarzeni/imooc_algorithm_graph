package cpt13_directed_graph

import "imooc_graph_go/graph"

// 有向图的拓扑排序
// 若存在环则返回空数组
// 队列实现
func TopologicalSort(g graph.DirectedGraph) (order []int) {
	indegree := make(map[int]int)
	queue := make([]int, 0)

	for v := 0; v < g.V(); v++ {
		if indegree[v] = g.InDegree(v); indegree[v] == 0 {
			queue = append(queue, v)
		}
	}

	for len(queue) > 0 {
		v := queue[0]
		order = append(order, v)
		queue = queue[1:]
		adjs, _ := g.Adj(v)
		for _, w := range adjs {
			if indegree[w] -= 1; indegree[w] == 0 {
				queue = append(queue, w)
			}
		}
	}

	if len(order) != g.V() {
		return nil
	}
	return
}
