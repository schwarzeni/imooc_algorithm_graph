package adjmatrix

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_newAdjMatrix(t *testing.T) {
	expectedMatrixData := [][]int{
		[]int{0, 1, 0, 1, 0, 0, 0},
		[]int{1, 0, 1, 0, 0, 0, 1},
		[]int{0, 1, 0, 1, 0, 1, 0},
		[]int{1, 0, 1, 0, 1, 0, 0},
		[]int{0, 0, 0, 1, 0, 1, 0},
		[]int{0, 0, 1, 0, 1, 0, 1},
		[]int{0, 1, 0, 0, 0, 1, 0},
	}

	newMatrix := newAdjMatrix("g.txt")
	if !reflect.DeepEqual(newMatrix.adj, expectedMatrixData) {
		t.Errorf("expect %v, \n but got %v", expectedMatrixData, newMatrix.adj)
	}
	fmt.Println(newMatrix)
}
