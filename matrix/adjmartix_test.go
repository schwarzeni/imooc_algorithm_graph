package matrix

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

func TestAdjMatrix_HasEdge(t *testing.T) {
	type fields struct {
		v   int
		e   int
		adj [][]int
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name        string
		adjMatrix   *AdjMatrix
		args        args
		wantErr     error
		wantHasEdge bool
	}{
		{name: "1", adjMatrix: newAdjMatrix("g.txt"), args: args{x: 1, y: 2}, wantErr: nil, wantHasEdge: true},
		{name: "2", adjMatrix: newAdjMatrix("g.txt"), args: args{x: 6, y: 5}, wantErr: nil, wantHasEdge: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.adjMatrix
			gotErr, gotHasEdge := a.HasEdge(tt.args.x, tt.args.y)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("HasEdge() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			if gotHasEdge != tt.wantHasEdge {
				t.Errorf("HasEdge() gotHasEdge = %v, want %v", gotHasEdge, tt.wantHasEdge)
			}
		})
	}
}

func TestAdjMatrix_Adj(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name      string
		adjMatrix *AdjMatrix
		args      args
		wantVs    []int
		wantErr   error
	}{
		{name: "1", adjMatrix: newAdjMatrix("g.txt"), args: args{v: 0}, wantErr: nil, wantVs: []int{1, 3}},
		{name: "2", adjMatrix: newAdjMatrix("g.txt"), args: args{v: 5}, wantErr: nil, wantVs: []int{2, 4, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.adjMatrix
			gotVs, err := a.Adj(tt.args.v)
			if err != tt.wantErr {
				t.Errorf("Adj() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVs, tt.wantVs) {
				t.Errorf("Adj() gotVs = %v, want %v", gotVs, tt.wantVs)
			}
		})
	}
}

func TestAdjMatrix_Degree(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name       string
		adjMatrix  *AdjMatrix
		args       args
		wantDegree int
		wantErr    error
	}{
		{name: "1", adjMatrix: newAdjMatrix("g.txt"), args: args{v: 0}, wantErr: nil, wantDegree: 2},
		{name: "2", adjMatrix: newAdjMatrix("g.txt"), args: args{v: 5}, wantErr: nil, wantDegree: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.adjMatrix
			gotDegree, err := a.Degree(tt.args.v)
			if err != tt.wantErr {
				t.Errorf("Degree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDegree != tt.wantDegree {
				t.Errorf("Degree() gotDegree = %v, want %v", gotDegree, tt.wantDegree)
			}
		})
	}
}
