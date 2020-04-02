package a_star

import (
	"container/heap"
	"fmt"
)

func AStar(g [][]int, start []int, goal []int) (steps [][]int) {
	const (
		T_BLOCK = 0
		T_WALL  = 1
	)
	pq := NewPriorityQueue()
	heap.Push(pq, &Item{Value: start, Priority: 0})

	var (
		rLen       = len(g)
		cLen       = len(g[0])
		cameFrom   = make([][]int, rLen)
		costAsFar  = make([][]int, rLen)
		canGo      = func(l []int) bool { return l[0] >= 0 && l[0] < rLen && l[1] >= 0 && l[1] < cLen }
		reached    = func(l []int) bool { return l[0] == goal[0] && l[1] == goal[1] }
		directions = [][]int{{0, 1}, {-1, 0}, {1, 0}, {0, -1}}
		heuristic  = func(l []int) int {
			r1 := l[0] - goal[0]
			r2 := l[1] - goal[1]
			if r1 < 0 {
				r1 = -r1
			}
			if r2 < 0 {
				r2 = -r2
			}
			return r1 + r2
		}
	)

	markCost := make([][]int, cLen)
	for i := 0; i < rLen; i++ {
		cameFrom[i] = make([]int, cLen)
		for j := 0; j < cLen; j++ {
			cameFrom[i][j] = -1
		}
		costAsFar[i] = make([]int, cLen)
		markCost[i] = make([]int, cLen)
	}

	cameFrom[start[0]][start[1]] = start[0]*rLen + start[1]
	costAsFar[start[0]][start[1]] = 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)

		if reached(current.Value) {
			break
		}

		cr := current.Value[0]
		cc := current.Value[1]

		for _, d := range directions {
			nr := cr + d[0]
			nc := cc + d[1]
			if canGo([]int{nr, nc}) && g[nr][nc] == T_BLOCK {
				newCost := costAsFar[cr][cc] + 1
				if costAsFar[nr][nc] <= 0 || newCost < costAsFar[nr][nc] {
					priority := newCost + heuristic([]int{nr, nc})
					heap.Push(pq, &Item{Value: []int{nr, nc}, Priority: priority})
					markCost[nr][nc] = priority
					costAsFar[nr][nc] = newCost
					cameFrom[nr][nc] = cr*rLen + cc
				}
			}
		}
	}

	if cameFrom[goal[0]][goal[1]] == -1 {
		fmt.Println("not found path")
		return
	}

	currL := goal
	steps = append(steps, goal)
	for currL[0] != start[0] && currL[1] != start[0] {
		nn := cameFrom[currL[0]][currL[1]]
		if nn == -1 {
			fmt.Println("not found path")
			return
		}
		nr := nn / rLen
		nc := nn % rLen
		steps = append(steps, []int{nr, nc})
		currL = []int{nr, nc}
	}

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			// this is search area
			if g[i][j] == T_WALL {
				fmt.Print("X")
			} else {
				fmt.Print(markCost[i][j])
			}
			if markCost[i][j] > 9 {
				fmt.Print(" ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	return
}
