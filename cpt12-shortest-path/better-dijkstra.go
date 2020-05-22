package cpt12_shortest_path

import (
	"container/heap"
	"imooc_graph_go/graph"
	"math"
)

func BetterDijkstra(g graph.WeightedGraph, s int) []int {
	visited := make([]bool, g.V())
	dis := make([]int, g.V())
	for idx, _ := range dis {
		dis[idx] = math.MaxInt64
	}
	dis[s] = 0
	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &Node{v: s, w: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Node).v
		if visited[cur] {
			continue
		}
		visited[cur] = true
		adj, _ := g.Adj(cur)
		for _, v := range adj {
			if !visited[v] {
				if dis[cur]+g.GetWeight(cur, v) < dis[v] {
					dis[v] = dis[cur] + g.GetWeight(cur, v)
					heap.Push(&pq, &Node{v: v, w: dis[v]})
				}
			}
		}
	}
	return dis
}

type Node struct {
	v int
	w int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].w < pq[j].w }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(n interface{}) { *pq = append(*pq, n.(*Node)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
