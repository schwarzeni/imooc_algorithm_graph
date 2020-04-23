package graph

import (
	"fmt"
	"testing"
)

func TestNewWeightedAdjListt(t *testing.T) {
	g := NewWeightedAdjList("gw.txt")
	fmt.Println(g)
	// v = 7, e = 12
	// 0: 1(2) 3(7) 5(2)
	// 1: 0(2) 2(1) 3(4) 4(3) 5(5)
	// 2: 1(1) 4(4) 5(4)
	// 3: 0(7) 1(4) 4(1) 6(5)
	// 4: 1(3) 2(4) 3(1) 6(7)
	// 5: 0(2) 1(5) 2(4)
	// 6: 3(5) 4(7)
}
