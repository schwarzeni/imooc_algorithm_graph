// https://leetcode-cn.com/problems/is-graph-bipartite/
package l_785

func isBipartite(graph [][]int) bool {
	var (
		status = make([]int, len(graph))
		dfs    func(startV int) bool
	)

	dfs = func(startV int) bool {
		for _, v := range graph[startV] {
			if status[v] == 0 { // adj vertex unvisited
				status[v] = -status[startV]
				if !dfs(v) {
					return false
				}
			} else {
				if status[v]+status[startV] != 0 {
					return false
				}
			}
		}
		return true
	}

	for idx, adj := range graph {
		if len(adj) != 0 && status[idx] == 0 {
			status[idx] = 1
			if !dfs(idx) {
				return false
			}
		}
	}
	return true
}
