package cpt13_directed_graph

import (
	"imooc_graph_go/graph"
)

// 有向图的拓扑排序
// 若存在环则返回空数组
// 深度优先遍历实现
// 缺点：不能进行换检测
func TopologicalSort2(g graph.DirectedGraph) (order []int) {
	visited := make([]bool, g.V())
	var dfs func(v int)
	dfs = func(v int) {
		adjs, _ := g.Adj(v)
		for _, w := range adjs {
			if !visited[w] {
				visited[w] = true
				dfs(w)
			}
		}
		order = append(order, v)
	}
	for v := 0; v < g.V(); v++ {
		// 不用选取入度为 0 的点，没有影响
		//if !visited[v] {
		//	visited[v] = true
		//	dfs(v)
		//}
		if g.InDegree(v) == 0 {
			dfs(v)
		}
	}
	if len(order) != g.V() {
		return nil
	}
	i, j := 0, g.V()-1
	for i < j {
		order[i], order[j] = order[j], order[i]
		i, j = i+1, j-1
	}
	return
}
