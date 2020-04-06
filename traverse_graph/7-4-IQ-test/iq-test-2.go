package __4_IQ

import "fmt"

// 000 -->  100 010 001
// 010 -->  110 000 011   100 001
//
// 0000
// 0101
// 0100
// 1101
// 1000
// 1011
// 1010
// 1111

func solveProblem2() {
	const (
		S_A_SIDE = '0'
		S_B_SIDE = '1'
	)
	var (
		initStatus = []byte{S_A_SIDE, S_A_SIDE, S_A_SIDE, S_A_SIDE} // 狼 羊 菜 农夫
		cango      = func(status []byte) bool {
			if status[0] == status[1] && status[3] == status[1] {
				return false
			}
			if status[1] == status[2] && status[1] == status[3] {
				return false
			}
			return true
		}
		finish = func(status []byte) bool {
			return status[0] == S_B_SIDE && status[1] == S_B_SIDE && status[2] == S_B_SIDE && status[3] == S_B_SIDE
		}
		toggleStatus = func(status byte) byte { return 97 - status }
		nextStatus   = func(status []byte) [][]byte {
			farmerStatus := status[3]
			if status[0] == status[1] && status[1] == status[2] {
				if status[0] != farmerStatus {
					return [][]byte{}
				} else {
					return [][]byte{{status[0], toggleStatus(status[1]), status[2], toggleStatus(status[3])}}
				}
			}
			if status[0] == status[1] && status[1] == farmerStatus {
				return [][]byte{
					{status[0], toggleStatus(status[1]), status[2], toggleStatus(status[3])},
					{toggleStatus(status[0]), status[1], status[2], toggleStatus(status[3])},
				}
			}
			if status[1] == status[2] && status[1] == farmerStatus {
				return [][]byte{
					{status[0], toggleStatus(status[1]), status[2], toggleStatus(status[3])},
					{status[0], status[1], toggleStatus(status[2]), toggleStatus(status[3])},
				}
			}
			var result [][]byte
			for idx, v := range status {
				newStatus := make([]byte, len(status))
				copy(newStatus, status)
				if v == farmerStatus {
					newStatus[idx] = toggleStatus(status[idx])
					result = append(result, newStatus)
				}
			}
			return result
		}
		preStatusMap = make(map[string]string)
		queue        []string
	)
	_ = cango

	preStatusMap[string(initStatus)] = string(initStatus)
	queue = append(queue, string(initStatus))

LOOP:
	for len(queue) > 0 {
		s := []byte(queue[0])
		queue = queue[1:]
		nextStatuses := nextStatus(s)
		for _, ns := range nextStatuses {
			if _, ok := preStatusMap[string(ns)]; !ok {
				preStatusMap[string(ns)] = string(s)
				if finish(ns) {
					break LOOP
				}
				queue = append(queue, string(ns))
			}
		}
	}

	fs := string([]byte{S_B_SIDE, S_B_SIDE, S_B_SIDE, S_B_SIDE})

	for fs != string(initStatus) {
		fmt.Println(fs)
		if _, ok := preStatusMap[fs]; !ok {
			fmt.Println("failed to get result")
			return
		}
		fs = preStatusMap[fs]
	}
}
