package traverse_graph

import "imooc_graph_go/graph"

type gBipartiteType int

const (
	g_UNVISITED gBipartiteType = iota
	g_TYPE_1
	g_TYPE_2
)

type Bipartite struct {
	g            graph.Graph
	part         [2][]int
	isBitpartite bool
}

func (b *Bipartite) GetBitpartite() [2][]int {
	return b.part
}

func (b *Bipartite) IsBitpartite() bool {
	return b.isBitpartite
}

func (b *Bipartite) testIsBipartite() {
	if b.g.V() == 0 {
		return
	}
	vmarked := make([]gBipartiteType, b.g.V())
	for i := 0; i < b.g.V(); i++ {
		vmarked[i] = g_UNVISITED
	}

	var dfs func(cv int) bool
	dfs = func(cv int) bool {
		currentType := vmarked[cv]
		oppositeType := g_TYPE_1
		if currentType == g_TYPE_1 {
			oppositeType = g_TYPE_2
		}
		vs, _ := b.g.Adj(cv)
		for _, v := range vs {
			if vmarked[v] == g_UNVISITED {
				vmarked[v] = oppositeType
				if validated := dfs(v); !validated {
					return false
				}
			} else if vmarked[v] == currentType {
				return false
			}
		}
		return true
	}

	// fix bug loop
	vmarked[0] = g_TYPE_1
	for v := 0; v < b.g.V(); v++ {
		if v == 0 || vmarked[v] == g_UNVISITED {
			if !dfs(v) {
				b.isBitpartite = false
				return
			}
		}
	}
	b.isBitpartite = true
	for idx, t := range vmarked {
		if t == g_TYPE_1 {
			b.part[0] = append(b.part[0], idx)
		} else if t == g_TYPE_2 {
			b.part[1] = append(b.part[1], idx)
		}
	}
	return
}

func NewBipartite(g graph.Graph) *Bipartite {
	b := &Bipartite{g: g}
	b.testIsBipartite()
	return b
}
