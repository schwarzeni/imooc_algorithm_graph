package cpt15_match

import (
	"bufio"
	"fmt"
	cpt14_network_flow "imooc_graph_go/cpt14-network-flow"
	"imooc_graph_go/graph"
	"imooc_graph_go/traverse_graph"
	"os"
)

// MostMatchWithNetworkFlow 使用最大流解决二分图最大匹配问题
// 二分图求解参考 ../traverse_graph/isbipartite.go
// 网络流求解 ../cpt14-network-flow/edmonds-karp.go
func MostMatchWithNetworkFlow(g graph.Graph) int {
	const TMP_FILE_NAME = "tmp.txt"
	bp := traverse_graph.NewBipartite(g)
	if !bp.IsBitpartite() {
		panic("not bitpartite")
	}

	twoParts := bp.GetBitpartite()
	l := g.V()
	s, t := l, l+1
	f, _ := os.Create(TMP_FILE_NAME)
	writer := bufio.NewWriter(f)
	_, _ = writer.WriteString(fmt.Sprintf("%d %d\n", g.V()+2, g.E()+l))
	for _, v := range twoParts[0] {
		adjs, _ := g.Adj(v)
		for _, w := range adjs {
			_, _ = writer.WriteString(fmt.Sprintf("%d %d 1\n", v, w))
			_, _ = writer.WriteString(fmt.Sprintf("%d %d 1\n", s, v))
		}
	}
	for _, v := range twoParts[1] {
		_, _ = writer.WriteString(fmt.Sprintf("%d %d 1\n", v, t))
	}
	_ = writer.Flush()
	_ = f.Close()
	defer os.Remove(TMP_FILE_NAME)

	return cpt14_network_flow.NewEdmondsKarp(TMP_FILE_NAME, s, t).Result()
}
