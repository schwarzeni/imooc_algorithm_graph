package cpt11_minimum_spanning_tree

type UF struct {
	parent []int
}

func NewUF(size int) *UF {
	u := &UF{parent: make([]int, size)}
	for idx, _ := range u.parent {
		u.parent[idx] = idx
	}
	return u
}

func (u *UF) find(p int) int {
	if p != u.parent[p] {
		u.parent[p] = u.find(u.parent[p])
	}
	return u.parent[p]
}

func (u *UF) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *UF) unionElements(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)

	if pRoot == qRoot {
		return
	}
	u.parent[pRoot] = qRoot
}
