package l_1034

// https://leetcode-cn.com/problems/coloring-a-border/
// 1 1
// 1 2
// r0 = 0, c0 = 0, color = 3
// 3 3
// 3 2

// 1 2 2
// 2 3 2
// r0 = 0, c0 = 1, color = 3
// 1 3 3
// 2 3 3

// 1 1 1
// 1 1 1
// 1 1 1
// r0 = 1, c0 = 1, color = 2
// 2 2 2
// 2 1 2
// 2 2 2
func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return grid
	}
	var (
		rlen       = len(grid)
		clen       = len(grid[0])
		visited    = make([][]bool, len(grid))
		dfs        func(cr, cc int)
		cango      = func(cr, cc int) bool { return !(cr < 0 || cr >= rlen || cc < 0 || cc >= clen) }
		directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		borders    [][]int
	)
	if !cango(r0, c0) {
		return grid
	}
	for i := 0; i < rlen; i++ {
		visited[i] = make([]bool, clen)
	}
	targetColor := grid[r0][c0]
	dfs = func(cr, cc int) {
		visited[cr][cc] = true
		atBorder := false
		inside := true
		for _, d := range directions {
			nr := cr + d[0]
			nc := cc + d[1]
			if cango(nr, nc) && !visited[nr][nc] && grid[nr][nc] == targetColor {
				dfs(nr, nc)
			}
			if !cango(nr, nc) { // at border
				atBorder = true
			} else if grid[nr][nc] != targetColor { // inner target graph
				inside = false
			}
		}
		if atBorder || !inside {
			borders = append(borders, []int{cr, cc})
		}
	}

	dfs(r0, c0)

	for _, b := range borders {
		grid[b[0]][b[1]] = color
	}

	return grid
}
