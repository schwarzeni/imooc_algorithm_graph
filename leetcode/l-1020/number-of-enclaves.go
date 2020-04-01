package l_1020

func numEnclaves_v1(A [][]int) int {
	if len(A) == 0 {
		return 0
	}
	w := len(A[0])
	h := len(A)
	cango := func(x, y int) bool { return x >= 0 && x < w && y >= 0 && y < h }
	atSide := func(x, y int) bool { return x == 0 || x == w-1 || y == 0 || y == h-1 }
	isVisited := make([][]bool, h)
	for i := 0; i < h; i++ {
		isVisited[i] = make([]bool, w)
	}

	var (
		count     int
		search    func(x, y int) (int, bool)
		direction = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	)

	search = func(x, y int) (gcount int, isAtSide bool) {
		isVisited[y][x] = true
		isAtSide = atSide(x, y)
		gcount = 1
		for _, d := range direction {
			nx := x + d[0]
			ny := y + d[1]
			if cango(nx, ny) && !isVisited[ny][nx] && A[ny][nx] == 1 {
				c, a := search(nx, ny)
				if a {
					isAtSide = a
				}
				gcount += c
			}
		}
		return
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if A[y][x] == 1 && !isVisited[y][x] {
				if gcount, a := search(x, y); !a {
					count += gcount
				}
			}
		}
	}

	return count
}

func numEnclaves_v2(A [][]int) int {
	if len(A) == 0 {
		return 0
	}
	w := len(A[0])
	h := len(A)
	cango := func(x, y int) bool { return x >= 0 && x < w && y >= 0 && y < h }
	isVisited := make([][]bool, h)
	for i := 0; i < h; i++ {
		isVisited[i] = make([]bool, w)
	}

	var (
		count        int
		searchBorder func(x, y int)
		direction    = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if A[i][j] == 1 {
				count += 1
			}
		}
	}

	searchBorder = func(x, y int) {
		isVisited[y][x] = true
		count -= 1
		for _, d := range direction {
			nx := x + d[0]
			ny := y + d[1]
			if cango(nx, ny) && !isVisited[ny][nx] && A[ny][nx] == 1 {
				searchBorder(nx, ny)
			}
		}
		return
	}

	for x := 0; x < w; x++ {
		if A[0][x] == 1 && !isVisited[0][x] {
			searchBorder(x, 0)
		}
		if A[h-1][x] == 1 && !isVisited[h-1][x] {
			searchBorder(x, h-1)
		}
	}

	for y := 0; y < h; y++ {
		if A[y][0] == 1 && !isVisited[y][0] {
			searchBorder(0, y)
		}
		if A[y][w-1] == 1 && !isVisited[y][w-1] {
			searchBorder(w-1, y)
		}
	}

	return count
}
