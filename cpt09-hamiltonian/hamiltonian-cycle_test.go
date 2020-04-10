package cpt09_hamiltonian

import (
	"fmt"
	"imooc_graph_go/graph"
	"log"
	"testing"
)

func Test_hamiltonianCycle(t *testing.T) {
	g1 := graph.NewAdjList("hamiltonian-g-1.txt")
	result1 := hamiltonianCycle(g1)
	if len(result1) != g1.V()+1 {
		log.Fatal("length not equal, ", len(result1), g1.V()+1)
	}
	// 0 <- 3 <- 1 <- 2 <- 0
	printPath(result1)

	g2 := graph.NewAdjList("hamiltonian-g-2.txt")
	result2 := hamiltonianCycle(g2)
	if len(result2) != g2.V()+1 {
		log.Fatal("length not equal, ", len(result2), g2.V()+1)
	}
	// 0 <- 13 <- 14 <- 15 <- 16 <- 17 <- 18 <- 19 <- 12 <- 11 <- 10 <- 9 <- 8 <- 7 <- 6 <- 5 <- 4 <- 3 <- 2 <- 1 <- 0
	printPath(result2)

}

func printPath(result []int) {
	for idx, v := range result {
		if idx == 0 {
			fmt.Print(v)
		} else {
			fmt.Print(" <- ", v)
		}
	}
	fmt.Println()
}
