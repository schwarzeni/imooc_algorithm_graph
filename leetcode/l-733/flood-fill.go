package l_733

// https://leetcode-cn.com/problems/flood-fill/
// 1 1 1
// 1 1 0
// 1 0 1
//
// (1, 1) 2
//
// 2 2 2
// 2 2 0
// 2 0 1
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return image
	}
	var (
		rlen       = len(image)
		clen       = len(image[0])
		visited    = make([][]bool, len(image))
		dfs        func(cr, cc int)
		cango      = func(cr, cc int) bool { return !(cr < 0 || cr >= rlen || cc < 0 || cc >= clen) }
		directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	)
	if !cango(sr, sc) {
		return image
	}
	for i := 0; i < rlen; i++ {
		visited[i] = make([]bool, clen)
	}
	targetColor := image[sr][sc]
	dfs = func(cr, cc int) {
		image[cr][cc] = newColor
		visited[cr][cc] = true

		for _, d := range directions {
			nr := cr + d[0]
			nc := cc + d[1]
			if cango(nr, nc) && !visited[nr][nc] && image[nr][nc] == targetColor {
				dfs(nr, nc)
			}
		}
	}

	dfs(sr, sc)
	return image
}
