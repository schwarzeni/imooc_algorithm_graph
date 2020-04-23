package cpt11_minimum_spanning_tree

import "testing"

func TestUF(t *testing.T) {
	u := NewUF(5)
	u.unionElements(0, 1)
	u.unionElements(1, 2)
	u.unionElements(3, 4)

	if !u.isConnected(0, 2) {
		t.Errorf("expect 0 and 2 connected")
	}
	if u.isConnected(0, 4) {
		t.Errorf("expect 0 and 4 unconnected")
	}
}
