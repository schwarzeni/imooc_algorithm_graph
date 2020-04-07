package l_773

import (
	"bytes"
)

// 2020.04.07
// https://leetcode-cn.com/problems/sliding-puzzle/
func slidingPuzzle(board [][]int) int {
	bboard := [][]byte{{'0', '0', '0'}, {'0', '0', '0'}}
	for r, cs := range board {
		for c, v := range cs {
			bboard[r][c] = 48 + byte(v)
		}
	}
	var (
		rLen = 2
		cLen = 3

		finish = func(status string) bool {
			return status == "123450"
		}
		compress = func(status [][]byte) string {
			var b bytes.Buffer
			for _, s := range status {
				b.Write(s)
			}
			return b.String()
		}
		uncompress = func(s string) (status [][]byte) {
			status = [][]byte{{'0', '0', '0'}, {'0', '0', '0'}}
			copy(status[0], s[0:3])
			copy(status[1], s[3:])
			return
		}
		findZero = func(status [][]byte) (zeroR, zeroC int) {
			for r, cs := range status {
				for c, v := range cs {
					if v == '0' {
						return r, c
					}
				}
			}
			return -1, -1
		}
		nextStatus = func(status [][]byte) (ns []string) {
			zeroR, zeroC := findZero(status)
			for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nr := zeroR + d[0]
				nc := zeroC + d[1]
				if nr >= 0 && nr < rLen && nc >= 0 && nc < cLen {
					status[nr][nc], status[zeroR][zeroC] = status[zeroR][zeroC], status[nr][nc]
					ns = append(ns, compress(status))
					status[nr][nc], status[zeroR][zeroC] = status[zeroR][zeroC], status[nr][nc]
				}
			}
			return
		}
		marked = make(map[string]int)
		queue  []string
	)

	init := compress(bboard)
	if finish(init) {
		return 0
	}
	queue = append(queue, init)
	marked[init] = 0

	for len(queue) > 0 {
		status := queue[0]
		queue = queue[1:]
		nss := nextStatus(uncompress(status))
		for _, ns := range nss {
			if _, ok := marked[ns]; !ok {
				marked[ns] = marked[status] + 1
				if finish(ns) {
					return marked[ns]
				}
				queue = append(queue, ns)
			}
		}
	}

	return -1
}
