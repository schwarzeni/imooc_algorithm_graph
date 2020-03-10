package adjmatrix

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
	V   int
	E   int
	adj [][]int
}

func (a *AdjMatrix) String() (output string) {
	output += fmt.Sprintf("V = %d, E = %d\n", a.V, a.E)
	for x := 0; x < a.V; x++ {
		for y := 0; y < a.V; y++ {
			output += fmt.Sprintf("%d ", a.adj[x][y])
		}
		output += "\n"
	}
	return
}

func (a *AdjMatrix) IsValid() error {
	if a.V < 0 {
		return errors.New("vertex must be non-negative")
	}
	if a.E < 0 {
		return errors.New("edge must be non-negative")
	}
	for i := 0; i < a.V; i++ {
		if a.adj[i][i] == MatrixEdge {
			return errors.New("self loop is detected") // 排除自环边情况
		}
	}
	return nil
}

func (a *AdjMatrix) AddEdge(x int, y int) error {
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
	adjMatrix.V, _ = strconv.Atoi(lineData[0])
	adjMatrix.E, _ = strconv.Atoi(lineData[1])

	for i := 0; i < adjMatrix.V; i++ {
		adjMatrix.adj = append(adjMatrix.adj, make([]int, adjMatrix.V))
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
