package cpt13_directed_graph

import "imooc_graph_go/graph"

// 有向图的环检测
func CycleDetection(g graph.Graph) bool {
	var (
		visited = make([]bool, g.V())
		path    = make(map[int]struct{})
		dfs     func(prev, curr int) bool
	)

	dfs = func(prev, curr int) bool {
		visited[curr] = true
		path[curr] = struct{}{}
		adjvs, _ := g.Adj(curr)
		for _, w := range adjvs {
			if !visited[w] {
				if dfs(curr, w) {
					return true
				}
			} else if _, ok := path[w]; ok && prev != w {
				return true
			}
		}
		delete(path, curr)
		return false
	}

	return dfs(0, 0)
}
