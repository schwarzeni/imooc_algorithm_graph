package matrix

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MatrixEdge = 1

type AdjMatrix struct {
	v   int
	e   int
	adj [][]int
}

func (a *AdjMatrix) String() (output string) {
	output += fmt.Sprintf("v = %d, e = %d\n", a.v, a.e)
	for x := 0; x < a.v; x++ {
		for y := 0; y < a.v; y++ {
			output += fmt.Sprintf("%d ", a.adj[x][y])
		}
		output += "\n"
	}
	return
}

func (a *AdjMatrix) IsValid() error {
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

func (a *AdjMatrix) IsValidVertex(v int) error {
	if v < 0 || v >= a.v {
		return fmt.Errorf("vertex %d is invalid", v)
	}
	return nil
}

func (a *AdjMatrix) AddEdge(x int, y int) error {
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
	a.adj[y][x] = MatrixEdge
	return nil
}

func (a *AdjMatrix) V() int {
	return a.v
}

func (a *AdjMatrix) E() int {
	return a.e
}

func (a *AdjMatrix) HasEdge(x int, y int) (err error, hasEdge bool) {
	if err = a.IsValidVertex(x); err != nil {
		return
	}
	if err = a.IsValidVertex(y); err != nil {
		return
	}
	return nil, a.adj[x][y] == MatrixEdge
}

// 与顶点v相邻的顶点
func (a *AdjMatrix) Adj(v int) (vs []int, err error) {
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

// 顶点v的度
func (a *AdjMatrix) Degree(v int) (degree int, err error) {
	var vs []int
	if vs, err = a.Adj(v); err != nil {
		return -1, err
	}
	return len(vs), nil
}

func newAdjMatrix(filename string) *AdjMatrix {
	var (
		err       error
		file      *os.File
		scanner   *bufio.Scanner
		adjMatrix = &AdjMatrix{}
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

	for i := 0; i < adjMatrix.v; i++ {
		adjMatrix.adj = append(adjMatrix.adj, make([]int, adjMatrix.v))
	}

	for scanner.Scan() {
		lineData = strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(lineData[0])
		y, _ := strconv.Atoi(lineData[1])

		// 无向图
		if err := adjMatrix.AddEdge(x, y); err != nil {
			panic(err)
		}
	}

	if err := adjMatrix.IsValid(); err != nil {
		panic(err)
	}

	return adjMatrix
}

func NewAdjMatrix(filename string) *AdjMatrix {
	return newAdjMatrix(filename)
}
