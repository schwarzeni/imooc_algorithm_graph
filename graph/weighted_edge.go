package graph

import "fmt"

type WeightedEdge struct {
	V      int
	W      int
	Weight int
}

func NewWeightedEdge(v int, w int, weight int) *WeightedEdge {
	return &WeightedEdge{V: v, W: w, Weight: weight}
}

func (w *WeightedEdge) String() string {
	return fmt.Sprintf("(%d-%d: %d)", w.V, w.W, w.Weight)
}
