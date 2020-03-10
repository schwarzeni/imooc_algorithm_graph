package adjmatrix

import (
	"bufio"
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
		adjMatrix.adj[x][y] = MatrixEdge
		adjMatrix.adj[y][x] = MatrixEdge
	}

	return adjMatrix
}

func NewAdjMatrix(filename string) *AdjMatrix {
	return newAdjMatrix(filename)
}
