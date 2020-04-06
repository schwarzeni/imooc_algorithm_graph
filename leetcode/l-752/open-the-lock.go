package l_752

// 2020.04.06
// https://leetcode-cn.com/problems/open-the-lock/
func openLock(deadends []string, target string) int {
	var (
		beginStr    = "0000"
		deadendsMap = make(map[string]struct{})
		visitedMap  = make(map[string]int)
		queue       []string
		move        = func(data byte) []byte {
			if data == '0' {
				return []byte{'9', '1'}
			} else if data == '9' {
				return []byte{'8', '0'}
			} else {
				return []byte{data - 1, data + 1}
			}
		}
		cango = func(curr []byte) bool {
			_, ok1 := deadendsMap[string(curr)]
			_, ok2 := visitedMap[string(curr)]
			return !ok1 && !ok2
		}
	)
	if target == beginStr {
		return 0
	}
	for _, v := range deadends {
		if v == beginStr {
			return -1
		}
		deadendsMap[v] = struct{}{}
	}

	visitedMap["0000"] = 0
	queue = append(queue, beginStr)
	for len(queue) > 0 {
		curr := []byte(queue[0])
		queue = queue[1:]

		curr2 := make([]byte, len(curr))
		copy(curr2, curr)
		for idx, b := range curr {
			ds := move(b)
			for _, d := range ds {
				curr2[idx] = d
				if cango(curr2) {
					visitedMap[string(curr2)] = visitedMap[string(curr)] + 1
					if string(curr2) == target {
						return visitedMap[string(curr2)]
					}
					queue = append(queue, string(curr2))
				}
				curr2[idx] = b
			}
		}
	}

	return -1
}
