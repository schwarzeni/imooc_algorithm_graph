// https://leetcode-cn.com/problems/max-area-of-island/
// 不清楚是否能改变 grid 的值
package l_695

func maxAreaOfIsland(grid [][]int) int {
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
		maxSize int
		search  func(x, y int) int
	)

	search = func(x, y int) int {
		if !cango(x, y) || isVisited[y][x] || grid[y][x] == 0 {
			return 0
		}
		isVisited[y][x] = true
		size := search(x-1, y) + search(x+1, y) + search(x, y-1) + search(x, y+1) + 1
		return size
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == 1 && !isVisited[y][x] {
				size := search(x, y)
				if size > maxSize {
					maxSize = size
				}
			}
		}
	}

	return maxSize
}
