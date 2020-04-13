package cpt10_eulerloop

import (
	"fmt"
	"imooc_graph_go/graph"
)

// 默认合法！
func HierholzerRecursion(g graph.Graph) []int {
	return hierholzerRecursion(g)
}

func hierholzerRecursion(g graph.Graph) []int {
	var (
		visited   = make(map[string]struct{})
		visitEdge func(v1, v2 int) bool
		dfs       func(currVertex, preVertex int)
		next      = make(map[int][]int)
	)
	visitEdge = func(v1, v2 int) bool {
		if v1 > v2 {
			v1, v2 = v2, v1
		}
		key := fmt.Sprintf("%d_%d", v1, v2)
		if _, ok := visited[key]; ok {
			return false
		}
		visited[key] = struct{}{}
		return true
	}
	dfs = func(currVertex, preVertex int) {
		adjs, _ := g.Adj(currVertex)

		for _, adj := range adjs {
			if visitEdge(adj, currVertex) {
				dfs(adj, currVertex)
			}
		}
		if preVertex != -1 {
			next[preVertex] = append(next[preVertex], currVertex)
		}
	}

	dfs(0, -1)

	var (
		stack      []int
		result     []int
		currVertex = 0
	)
	stack = append(stack, currVertex)
	for {
		nvs, ok := next[currVertex]
		if !ok {
			result = append(result, currVertex)
			currVertex = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			nv := nvs[0]
			nvs = nvs[1:]
			stack = append(stack, currVertex)
			if len(nvs) == 0 {
				delete(next, currVertex)
			} else {
				next[currVertex] = nvs
			}
			currVertex = nv
		}
		if len(stack) <= 0 {
			break
		}
	}

	return result
}
