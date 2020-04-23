package graph

type WeightedGraph interface {
	Graph
	GetWeight(v, w int) int
}
