package __4_IQ

import "fmt"

func solveProblem2() {
	const (
		S_A_SIDE    = '0' // 在河岸一边
		S_B_SIDE    = '1' // 在河岸另一边
		I_WOLF      = 0   // 数组中狼的 idx
		I_SHEEP     = 1   // 数组中羊的 idx
		I_VEGETABLE = 2   // 数组中菜的 idx
		I_FARMER    = 3   // 数组中农夫的 idx
	)
	var (
		//   				   狼        羊        菜       农夫
		initStatus = []byte{S_A_SIDE, S_A_SIDE, S_A_SIDE, S_A_SIDE}

		// 如果狼和羊、羊和菜在同一河岸的情况下，必须农夫和他们在一边才行，否则不行
		// 其他情况都是可以的
		cango = func(status []byte) bool {
			if status[I_WOLF] == status[I_SHEEP] {
				return status[I_FARMER] == status[I_SHEEP]
			}
			if status[I_SHEEP] == status[I_VEGETABLE] {
				return status[I_FARMER] == status[I_SHEEP]
			}
			return true
		}
		finish = func(status []byte) bool {
			return status[I_WOLF] == S_B_SIDE &&
				status[I_SHEEP] == S_B_SIDE &&
				status[I_VEGETABLE] == S_B_SIDE &&
				status[I_FARMER] == S_B_SIDE
		}
		toggleStatus = func(status byte) byte { return S_A_SIDE + S_B_SIDE - status }

		// 下一个状态：农夫带一个和在在同一河岸的物件过河，当然，只带自己是可以的
		nextStatus = func(status []byte) (result [][]byte) {
			farmerStatus := status[I_FARMER]
			for idx, v := range status {
				if v == farmerStatus {
					ns := make([]byte, len(initStatus))
					copy(ns, status)
					ns[idx] = toggleStatus(v)
					ns[I_FARMER] = toggleStatus(v)
					result = append(result, ns)
				}
			}
			return
		}

		preStatusMap = make(map[string]string)
		queue        []string
	)

	preStatusMap[string(initStatus)] = string(initStatus)
	queue = append(queue, string(initStatus))

LOOP:
	for len(queue) > 0 {
		s := []byte(queue[0])
		queue = queue[1:]
		nextStatuses := nextStatus(s)
		for _, ns := range nextStatuses {
			if cango(ns) {
				if _, ok := preStatusMap[string(ns)]; !ok {
					preStatusMap[string(ns)] = string(s)
					if finish(ns) {
						break LOOP
					}
					queue = append(queue, string(ns))
				}
			}
		}
	}

	fs := string([]byte{S_B_SIDE, S_B_SIDE, S_B_SIDE, S_B_SIDE})

	for fs != string(initStatus) {
		fmt.Println(fs)
		if _, ok := preStatusMap[fs]; !ok {
			panic("failed to get result")
		}
		fs = preStatusMap[fs]
	}
	fmt.Println(fs)
}
