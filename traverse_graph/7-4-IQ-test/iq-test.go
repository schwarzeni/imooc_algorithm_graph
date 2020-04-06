package __4_IQ

import "fmt"

// 倒水问题：5升和3升的水桶，得到2升
func solveProblem() {
	var (
		status        = make(map[int]int) // 10 * x +y
		queue         []int
		maxX          = 5
		maxY          = 3
		aim           = 4
		aimX          = -1
		aimY          = -1
		compressOpt   = func(x, y int) int { return x*10 + y }
		uncompressOpt = func(x_y int) (int, int) { return x_y / 10, x_y % 10 }
		nextOpts      = func(x, y int) (opts [][]int) {
			opts = append(opts, []int{5, y}) // opt1 5, y
			opts = append(opts, []int{x, 3}) // opt2 x,3
			opts = append(opts, []int{0, y}) // opt3 0, y
			opts = append(opts, []int{x, 0}) // opt3 0, y

			// opt5 y --> x
			if newx := x + y; newx <= maxX {
				opts = append(opts, []int{newx, 0})
			} else {
				opts = append(opts, []int{maxX, y - (maxX - x)})
			}

			// opt6 x --> y
			if newy := x + y; newy <= maxY {
				opts = append(opts, []int{0, newy})
			} else {
				opts = append(opts, []int{x - (maxY - y), maxY})
			}
			return
		}
	)

	queue = append(queue, 0)
	status[0] = 0

LOOP:
	for len(queue) > 0 {
		x_y := queue[0]
		queue = queue[1:]
		x, y := uncompressOpt(x_y)

		for _, n_x_y_arr := range nextOpts(x, y) {
			nx, ny := n_x_y_arr[0], n_x_y_arr[1]
			n_x_y := compressOpt(nx, ny)
			if _, ok := status[n_x_y]; !ok {
				status[n_x_y] = x_y
				queue = append(queue, n_x_y)
				if nx == aim {
					aimX = nx
					aimY = ny
					break LOOP
				}
			}
		}
	}

	if aimX == -1 || aimY == -1 {
		fmt.Println("not found")
		return
	}
	x_y := compressOpt(aimX, aimY)
	for x_y != 0 {
		fmt.Println(uncompressOpt(x_y))
		x_y = status[x_y]
	}
}
