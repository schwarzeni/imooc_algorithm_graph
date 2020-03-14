package matrix

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type AdjList struct {
	v   int
	e   int
	adj []map[int]int
}

func (a *AdjList) String() (output string) {
	output += fmt.Sprintf("v = %d, e = %d\n", a.v, a.e)
	for x, v := range a.adj {
		var xvs []int
		for e, _ := range v {
			xvs = append(xvs, e)
		}
		sort.Ints(xvs)
		output += fmt.Sprintf("%d: ", x)
		for _, y := range xvs {
			output += fmt.Sprintf("%d ", y)
		}
		output += "\n"
	}
	return
}

func (a *AdjList) IsValidVertex(v int) error {
	if v < 0 || v >= a.v {
		return fmt.Errorf("vertex %d is invalid", v)
	}
	return nil
}

func (a *AdjList) AddEdge(x int, y int) error {
	if err := a.IsValidVertex(x); err != nil {
		return err
	}
	if err := a.IsValidVertex(y); err != nil {
		return err
	}

	// TODO: 这里没有考虑排除平行边情况

	if x == y {
		return errors.New("self loop is detected")
	}
	a.adj[x][y] = MatrixEdge
	a.adj[y][x] = MatrixEdge
	return nil
}

func (a *AdjList) V() int {
	return a.v
}

func (a *AdjList) E() int {
	return a.e
}

func (a *AdjList) HasEdge(x int, y int) (hasEdge bool, err error) {
	if err := a.IsValidVertex(x); err != nil {
		return false, err
	}
	if err := a.IsValidVertex(y); err != nil {
		return false, err
	}
	if isEdge, ok := a.adj[x][y]; ok && isEdge == MatrixEdge {
		return true, nil
	}
	return false, nil
}

func (a *AdjList) Adj(v int) (vs []int, err error) {
	if err = a.IsValidVertex(v); err != nil {
		return nil, err
	}
	for k, _ := range a.adj[v] {
		vs = append(vs, k)
	}
	sort.Ints(vs)
	return
}

func (a *AdjList) Degree(v int) (degree int, err error) {
	if err = a.IsValidVertex(v); err != nil {
		return -1, err
	}
	return len(a.adj[v]), nil
}

func (a *AdjList) IsValid() error {
	return nil
}

func newAdjList(filename string) *AdjList {
	var (
		err      error
		file     *os.File
		scanner  *bufio.Scanner
		adjList  = &AdjList{}
		lineData []string
	)
	if file, err = os.Open(filename); err != nil {
		panic(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	scanner.Scan()

	lineData = strings.Split(scanner.Text(), " ")
	adjList.v, _ = strconv.Atoi(lineData[0])
	adjList.e, _ = strconv.Atoi(lineData[1])

	for i := 0; i < adjList.v; i++ {
		adjList.adj = append(adjList.adj, make(map[int]int))
	}
	//adjList.adj = make([]map[int]int, adjList.v)

	for scanner.Scan() {
		lineData = strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(lineData[0])
		y, _ := strconv.Atoi(lineData[1])

		// 无向图
		if err := adjList.AddEdge(x, y); err != nil {
			panic(err)
		}
	}

	if err := adjList.IsValid(); err != nil {
		panic(err)
	}

	return adjList
}

func NewAdjList(filename string) *AdjList {
	return newAdjList(filename)
}
