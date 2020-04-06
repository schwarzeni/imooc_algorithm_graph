package l_1091

// 2020.04.06
// https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/
func shortestPathBinaryMatrix(grid [][]int) int {
	var (
		directions    = [][]int{{-1, -1}, {1, 1}, {-1, 1}, {1, -1}, {0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		rLen          = len(grid)
		cLen          = len(grid[0])
		isEnd         = func(r, c int) bool { return r == rLen-1 && c == cLen-1 }
		idx           = 0
		queue         [][]int
		shortPathSize = make([][]int, rLen)
		cango         = func(r, c int) bool {
			return r >= 0 && r < rLen && c >= 0 && c < cLen && grid[r][c] == 0 && shortPathSize[r][c] == -1
		}
	)
	if grid[0][0] == 1 || grid[rLen-1][cLen-1] == 1 {
		return -1
	}
	for idx, _ := range grid {
		shortPathSize[idx] = make([]int, cLen)
		for i, _ := range shortPathSize {
			shortPathSize[idx][i] = -1
		}
	}

	queue = append(queue, []int{0, 0})
	shortPathSize[0][0] = 1

	for idx < len(queue) {
		r_c := queue[idx]
		cr := r_c[0]
		cc := r_c[1]

		for _, d := range directions {
			nr := cr + d[0]
			nc := cc + d[1]
			if cango(nr, nc) {
				queue = append(queue, []int{nr, nc})
				shortPathSize[nr][nc] = shortPathSize[cr][cc] + 1
				if isEnd(nr, nc) {
					return shortPathSize[nr][nc]
				}
			}
		}
		idx += 1
	}

	return shortPathSize[rLen-1][cLen-1]

	//var (
	//	directions = [][]int{{-1, -1}, {1, 1}, {-1, 1}, {1, -1}, {0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	//	rLen       = len(grid)
	//	cLen       = len(grid[0])
	//	cango      = func(r, c int) bool { return r >= 0 && r < rLen && c >= 0 && c < cLen && grid[r][c] == 0 }
	//	isEnd      = func(r, c int) bool { return r == rLen-1 && c == cLen-1 }
	//	idx        = 0
	//	queue      [][]int
	//	visited    = make([][]bool, rLen)
	//	pre        = make([]int, rLen*cLen)
	//	count      = 0
	//)
	//if grid[0][0] == 1 || grid[rLen-1][cLen-1] == 1 {
	//	return -1
	//}
	//for idx, _ := range visited {
	//	visited[idx] = make([]bool, cLen)
	//}
	//for idx, _ := range pre {
	//	pre[idx] = -1
	//}
	//
	//queue = append(queue, []int{0, 0})
	//visited[0][0] = true
	//pre[0] = 0
	//
	//for idx < len(queue) {
	//	r_c := queue[idx]
	//	cr := r_c[0]
	//	cc := r_c[1]
	//	if isEnd(cr, cc) {
	//		break
	//	}
	//	for _, d := range directions {
	//		nr := cr + d[0]
	//		nc := cc + d[1]
	//		if cango(nr, nc) && !visited[nr][nc] {
	//			visited[nr][nc] = true
	//			queue = append(queue, []int{nr, nc})
	//			pre[nr*rLen+nc] = cr*rLen + cc
	//		}
	//	}
	//	idx += 1
	//}
	//
	//node := (rLen-1)*rLen + cLen - 1
	//count = 1
	//for {
	//	if node == 0 {
	//		return count
	//	}
	//	node = pre[node]
	//	count += 1
	//	if node == -1 {
	//		return -1
	//	}
	//}
}
