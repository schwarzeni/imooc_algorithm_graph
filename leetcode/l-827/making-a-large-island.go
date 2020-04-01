package l_827

// https://leetcode-cn.com/problems/making-a-large-island/

func largestIsland(grid [][]int) int {
	const (
		T_WATER = 0
		T_LAND  = 1
	)
	var (
		rlen        = len(grid)
		clen        = len(grid)
		areaList    []int
		blockToArea = make([][]int, rlen)
		cango       = func(r, c int) bool { return r >= 0 && r < rlen && c >= 0 && c < clen }
		dfs         func(r, c int)
		directions  = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		areaID      = 0
		maxArea     int
	)
	for idx, _ := range blockToArea {
		blockToArea[idx] = make([]int, rlen)
	}
	areaList = append(areaList, 0)

	dfs = func(r, c int) {
		areaList[areaID] += 1
		blockToArea[r][c] = areaID
		for _, d := range directions {
			nr := r + d[0]
			nc := c + d[1]
			if cango(nr, nc) && blockToArea[nr][nc] == 0 && grid[nr][nc] == T_LAND {
				dfs(nr, nc)
			}
		}
	}

	for r, ca := range grid {
		for c, _ := range ca {
			if blockToArea[r][c] == 0 && grid[r][c] == T_LAND {
				areaList = append(areaList, 0)
				areaID += 1
				dfs(r, c)
			}
		}
	}

	if areaID == 0 { // 无陆地
		return 1
	}

	if areaList[1] == rlen*clen { // 全是陆地
		return areaList[1]
	}

	if areaID == 1 { // 只有一块陆地，且有海
		return areaList[1] + 1
	}

	for r, ca := range grid {
		for c, _ := range ca {
			if grid[r][c] == T_WATER {
				set := make(map[int]struct{})
				area := 0
				for _, d := range directions {
					nr := r + d[0]
					nc := c + d[1]
					if cango(nr, nc) && grid[nr][nc] == T_LAND {
						if _, ok := set[blockToArea[nr][nc]]; !ok {
							area += areaList[blockToArea[nr][nc]]
							set[blockToArea[nr][nc]] = struct{}{}
						}
					}
				}
				area += 1
				if maxArea < area {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}
