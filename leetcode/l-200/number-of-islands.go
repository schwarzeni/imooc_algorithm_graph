package l_200

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	w := len(grid[0])
	h := len(grid)
	cango := func(x, y int) bool { return x >= 0 && x < w && y >= 0 && y < h }
	isVisited := make([][]bool, h)
	for i := 0; i < h; i++ {
		isVisited[i] = make([]bool, w)
	}

	var (
		count     int
		search    func(x, y int)
		direction = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	)

	search = func(x, y int) {
		isVisited[y][x] = true
		for _, d := range direction {
			nx := x + d[0]
			ny := y + d[1]
			if cango(nx, ny) && !isVisited[ny][nx] && grid[ny][nx] == '1' {
				search(nx, ny)
			}
		}
		return
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == '1' && !isVisited[y][x] {
				search(x, y)
				count += 1
			}
		}
	}

	return count
}
