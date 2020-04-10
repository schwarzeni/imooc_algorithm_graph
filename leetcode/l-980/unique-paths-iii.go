package l_980

// 2020.04.10
// https://leetcode-cn.com/problems/unique-paths-iii/
func uniquePathsIII(grid [][]int) int {
	const (
		T_START      = 1
		T_END        = 2
		T_CAN_GO     = 0
		T_CAN_NOT_GO = -1
	)
	var (
		dfs          func(cr, cc int)
		gridNum      = 2
		rLen         = len(grid)
		cLen         = len(grid[0])
		cango        = func(r, c int) bool { return r >= 0 && r < rLen && c >= 0 && c < cLen && grid[r][c] != T_CAN_NOT_GO }
		directions   = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		visitedCount = 0
		visited      = make([][]bool, rLen)
		startR       = -1
		startC       = -1
		endR         = -1
		endC         = -1
		solutionNum  int
	)
	for r, cs := range grid {
		for c, _ := range cs {
			if grid[r][c] == T_START {
				startR = r
				startC = c
			} else if grid[r][c] == T_END {
				endR = r
				endC = c
			} else if grid[r][c] == T_CAN_GO {
				gridNum += 1
			}
		}
	}
	if startR == -1 || startC == -1 || endR == -1 || endC == -1 {
		return 0
	}
	for idx, _ := range visited {
		visited[idx] = make([]bool, cLen)
	}

	dfs = func(cr, cc int) {
		if cr == endR && cc == endC && visitedCount == gridNum {
			solutionNum += 1
			return
		}
		for _, d := range directions {
			nr := cr + d[0]
			nc := cc + d[1]
			if cango(nr, nc) && !visited[nr][nc] {
				visited[nr][nc] = true
				visitedCount += 1

				dfs(nr, nc)

				visitedCount -= 1
				visited[nr][nc] = false
			}
		}
	}

	visited[startR][startC] = true
	visitedCount = 1
	dfs(startR, startC)

	return solutionNum
}
