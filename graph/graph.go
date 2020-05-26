package graph

const MatrixEdge = 1

type Graph interface {
	AddEdge(x int, y int, w int) error
	V() int
	E() int
	HasEdge(x int, y int) (hasEdge bool, err error)
	Adj(v int) (vs []int, err error)
	Degree(v int) (degree int, err error)
}

type WeightedGraph interface {
	Graph
	GetWeight(v, w int) int
}

type DirectedGraph interface {
	Graph
	InDegree(v int) int
	OutDegree(v int) int
}

type DirectedWeightedGraph interface {
	WeightedGraph
	InDegree(v int) int
	OutDegree(v int) int
}
