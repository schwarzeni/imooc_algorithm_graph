package graph

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DirectedAdjMatrix struct {
	v          int
	e          int
	adj        [][]int
	indegrees  []int
	outdegrees []int
}

func (a *DirectedAdjMatrix) String() (output string) {
	output += fmt.Sprintf("v = %d, e = %d\n", a.v, a.e)
	for x := 0; x < a.v; x++ {
		for y := 0; y < a.v; y++ {
			output += fmt.Sprintf("%d ", a.adj[x][y])
		}
		output += "\n"
	}
	return
}

func (a *DirectedAdjMatrix) IsValid() error {
	if a.v < 0 {
		return errors.New("vertex must be non-negative")
	}
	if a.e < 0 {
		return errors.New("edge must be non-negative")
	}
	for i := 0; i < a.v; i++ {
		if a.adj[i][i] == MatrixEdge {
			return errors.New("self loop is detected") // 排除自环边情况
		}
	}
	return nil
}

func (a *DirectedAdjMatrix) IsValidVertex(v int) error {
	if v < 0 || v >= a.v {
		return fmt.Errorf("vertex %d is invalid", v)
	}
	return nil
}

func (a *DirectedAdjMatrix) AddEdge(x int, y int, w int) error {
	if err := a.IsValidVertex(x); err != nil {
		return err
	}
	if err := a.IsValidVertex(y); err != nil {
		return err
	}
	if a.adj[x][y] == MatrixEdge || a.adj[y][x] == MatrixEdge {
		return errors.New("parallel edges are detected") // 排除平行边情况
	}
	if x == y {
		return errors.New("self loop is detected")
	}
	a.adj[x][y] = MatrixEdge
	return nil
}

func (a *DirectedAdjMatrix) V() int {
	return a.v
}

func (a *DirectedAdjMatrix) E() int {
	return a.e
}

func (a *DirectedAdjMatrix) HasEdge(x int, y int) (hasEdge bool, err error) {
	if err = a.IsValidVertex(x); err != nil {
		return
	}
	if err = a.IsValidVertex(y); err != nil {
		return
	}
	return a.adj[x][y] == MatrixEdge, nil
}

// 与顶点v相邻的顶点
func (a *DirectedAdjMatrix) Adj(v int) (vs []int, err error) {
	if err = a.IsValidVertex(v); err != nil {
		return
	}
	for vertex, hasEdge := range a.adj[v] {
		if hasEdge == MatrixEdge {
			vs = append(vs, vertex)
		}
	}
	return
}

func (a *DirectedAdjMatrix) Degree(v int) (degree int, err error) {
	var vs []int
	if vs, err = a.Adj(v); err != nil {
		return -1, err
	}
	return len(vs), nil
}

func (a *DirectedAdjMatrix) InDegree(v int) (degree int) {
	return a.indegrees[v]
}

func (a *DirectedAdjMatrix) OutDegree(v int) (degree int) {
	return a.outdegrees[v]
}

func newDirectedAdjMatrix(filename string) *DirectedAdjMatrix {
	var (
		err       error
		file      *os.File
		scanner   *bufio.Scanner
		adjMatrix = &DirectedAdjMatrix{}
		lineData  []string
	)
	if file, err = os.Open(filename); err != nil {
		panic(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	scanner.Scan()

	lineData = strings.Split(scanner.Text(), " ")
	adjMatrix.v, _ = strconv.Atoi(lineData[0])
	adjMatrix.e, _ = strconv.Atoi(lineData[1])
	adjMatrix.indegrees = make([]int, adjMatrix.v)
	adjMatrix.outdegrees = make([]int, adjMatrix.v)

	for i := 0; i < adjMatrix.v; i++ {
		adjMatrix.adj = append(adjMatrix.adj, make([]int, adjMatrix.v))
	}

	for scanner.Scan() {
		lineData = strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(lineData[0])
		y, _ := strconv.Atoi(lineData[1])

		adjMatrix.indegrees[y]++
		adjMatrix.outdegrees[x]++
		if err := adjMatrix.AddEdge(x, y, 1); err != nil {
			panic(err)
		}
	}

	if err := adjMatrix.IsValid(); err != nil {
		panic(err)
	}

	return adjMatrix
}

func NewDirecteddjMatrix(filename string) *DirectedAdjMatrix {
	return newDirectedAdjMatrix(filename)
}
