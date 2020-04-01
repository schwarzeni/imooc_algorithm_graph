package l_529

// https://leetcode-cn.com/problems/minesweeper/
func updateBoard(board [][]byte, click []int) [][]byte {
	const (
		T_UNSEARCHED_BOMB     = 'M' // 未找到的炸弹
		T_UNSEARCHED_BLOCK    = 'E' // 未搜寻的方块
		T_SEARCHED_SAFE_BLOCK = 'B' // 已经搜寻过的并且周围没有炸弹的方块
		T_SEARCHED_BOMB       = 'X' // 已经找到的炸弹
	)
	var (
		rlen                    = len(board)
		clen                    = len(board[0])
		isNearBombSearchedBlock = func(cr, cc int) bool { return board[cr][cc] >= '1' && board[cr][cc] <= '8' }
		dfs                     func(cr, cc int)
		cango                   = func(cr, cc int) bool { return !(cr < 0 || cr >= rlen || cc < 0 || cc >= clen) }
		fulldirections          = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, 1}, {1, -1}, {-1, 1}}
		visited                 = make([][]bool, len(board))
		clickr                  = click[0]
		clickc                  = click[1]
		nearBlockBombCount      = func(cr, cc int) (count int) {
			for _, d := range fulldirections {
				nr := cr + d[0]
				nc := cc + d[1]
				// 不清楚是否要加上 T_SEARCHED_BOMB
				if cango(nr, nc) && (board[nr][nc] == T_UNSEARCHED_BOMB || board[nr][nc] == T_SEARCHED_BOMB) {
					count += 1
				}
			}
			return
		}
	)
	if !cango(clickr, clickc) || // 错误点击的情况
		board[clickr][clickc] == T_SEARCHED_BOMB || // X 点击到已经搜寻过 bomb 的情况
		board[clickr][clickc] == T_SEARCHED_SAFE_BLOCK || // B 点击到已经搜寻过的并且周围没有炸弹的方块
		isNearBombSearchedBlock(clickr, clickc) { // 点击到周围有炸弹的方块的情况
		return board
	}
	if board[clickr][clickc] == T_UNSEARCHED_BOMB { // M 点击到未找到炸弹的情况
		board[clickr][clickc] = T_SEARCHED_BOMB
		return board
	}

	// 以下为点击到 'E' 未搜寻的方块的情况
	for i := 0; i < rlen; i++ {
		visited[i] = make([]bool, clen)
	}
	dfs = func(cr, cc int) {
		visited[cr][cc] = true
		count := nearBlockBombCount(cr, cc)
		if count > 0 {
			board[cr][cc] = byte(48 + count)
			return
		}
		board[cr][cc] = T_SEARCHED_SAFE_BLOCK
		for _, d := range fulldirections {
			nr := cr + d[0]
			nc := cc + d[1]
			if cango(nr, nc) && board[nr][nc] == T_UNSEARCHED_BLOCK && !visited[nr][nc] {
				dfs(nr, nc)
			}
		}
	}

	dfs(clickr, clickc)
	return board
}
