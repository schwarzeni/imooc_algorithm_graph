package l_130

func solve(board [][]byte) {
	if len(board) < 2 || (len(board) > 0 && len(board[0]) < 2) {
		return
	}
	w := len(board[0])
	h := len(board)
	cango := func(x, y int) bool { return x >= 0 && x < w && y >= 0 && y < h }
	isVisited := make([][]bool, h)
	for i := 0; i < h; i++ {
		isVisited[i] = make([]bool, w)
	}

	var (
		search    func(x, y int)
		direction = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	)

	search = func(x, y int) {
		isVisited[y][x] = true

		for _, d := range direction {
			nx := x + d[0]
			ny := y + d[1]
			if cango(nx, ny) && !isVisited[ny][nx] && board[ny][nx] == 'O' {
				search(nx, ny)
			}
		}
		return
	}

	for x := 0; x < w; x++ {
		if board[0][x] == 'O' && !isVisited[0][x] {
			search(x, 0)
		}
		if board[h-1][x] == 'O' && !isVisited[h-1][x] {
			search(x, h-1)
		}
	}

	for y := 0; y < h; y++ {
		if board[y][0] == 'O' && !isVisited[y][0] {
			search(0, y)
		}
		if board[y][w-1] == 'O' && !isVisited[y][w-1] {
			search(w-1, y)
		}
	}

	for i := 1; i < h-1; i++ {
		for j := 1; j < w-1; j++ {
			if board[i][j] == 'O' && !isVisited[i][j] {
				board[i][j] = 'X'
			}
		}
	}

}
