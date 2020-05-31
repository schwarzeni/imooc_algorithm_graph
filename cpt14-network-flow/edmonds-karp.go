package cpt14_network_flow

import (
	"bytes"
	"fmt"
	"imooc_graph_go/graph"
	"math"
)

// EdmondsKarp edmonds karp 算法
type EdmondsKarp struct {
	network, residualNetwork graph.DirectedWeightedGraph
	s, t, maxFlow            int
}

func (ek *EdmondsKarp) Result() int {
	return ek.maxFlow
}

func (ek *EdmondsKarp) Flow(v, w int) int {
	return ek.residualNetwork.GetWeight(w, v)
}

func (ek *EdmondsKarp) getAugmentingPath() (res []int) {
	q := []int{}
	pre := make([]int, ek.network.V())
	for idx, _ := range pre {
		pre[idx] = -1
	}
	q = append(q, ek.s)
	pre[ek.s] = ek.s
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == ek.t {
			break
		}
		adjs, _ := ek.residualNetwork.Adj(cur)
		for _, next := range adjs {
			if pre[next] == -1 && ek.residualNetwork.GetWeight(cur, next) > 0 {
				pre[next] = cur
				q = append(q, next)
			}
		}
	}
	if pre[ek.t] == -1 {
		return nil
	}
	cur := ek.t
	for cur != ek.s {
		res = append(res, cur)
		cur = pre[cur]
	}
	res = append(res, ek.s)
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return
}

func (ek *EdmondsKarp) String() string {
	sbuilder := bytes.Buffer{}
	for v := 0; v < ek.network.V(); v++ {
		adjs, _ := ek.network.Adj(v)
		for _, w := range adjs {
			sbuilder.WriteString(fmt.Sprintf("%d-%d: %d / %d\n", v, w, ek.Flow(v, w), ek.network.GetWeight(v, w)))
		}
	}
	return sbuilder.String()
}

// NewEdmondsKarp 新建，默认输入数据合法
func NewEdmondsKarp(src string, s, t int) *EdmondsKarp {
	ek := &EdmondsKarp{network: graph.NewDirectedWeightedAdjList(src), s: s, t: t}
	rg := graph.NewDirectedWeightedAdjList(src)
	for v := 0; v < ek.network.V(); v++ {
		adjs, _ := ek.network.Adj(v)
		for _, w := range adjs {
			c := ek.network.GetWeight(v, w)
			_ = rg.AddEdge(v, w, c)
			_ = rg.AddEdge(w, v, 0)
		}
	}
	ek.residualNetwork = rg

	for {
		augPath := ek.getAugmentingPath()
		if len(augPath) == 0 {
			break
		}
		f := math.MaxInt64
		for i := 1; i < len(augPath); i++ {
			v, w := augPath[i-1], augPath[i]
			if f > ek.residualNetwork.GetWeight(v, w) {
				f = ek.residualNetwork.GetWeight(v, w)
			}
		}
		ek.maxFlow += f

		for i := 1; i < len(augPath); i++ {
			v, w := augPath[i-1], augPath[i]
			ek.residualNetwork.SetWeight(v, w, ek.residualNetwork.GetWeight(v, w)-f)
			ek.residualNetwork.SetWeight(w, v, ek.residualNetwork.GetWeight(w, v)+f)
		}
	}

	return ek
}
