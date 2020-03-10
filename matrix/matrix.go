package matrix

type Matrix interface {
	AddEdge(x int, y int) error
	V() int
	E() int
	HasEdge(x int, y int) (err error, hasEdge bool)
	Adj(v int) (vs []int, err error)
	Degree(v int) (degree int, err error)
}
