package graph

const MatrixEdge = 1

type Graph interface {
	AddEdge(x int, y int) error
	V() int
	E() int
	HasEdge(x int, y int) (hasEdge bool, err error)
	Adj(v int) (vs []int, err error)
	Degree(v int) (degree int, err error)
}
