package l_1091

import "testing"

func Test_shortestPathBinaryMatrix(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{[][]int{{0, 1}, {1, 0}}}, want: 2},
		{name: "test2", args: args{[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}}, want: 4},
		{name: "test3", args: args{[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 1}}}, want: -1},
		{name: "test_1", args: args{[][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}}, want: -1},
		{name: "test_2", args: args{[][]int{{0}}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPathBinaryMatrix(tt.args.grid); got != tt.want {
				t.Errorf("shortestPathBinaryMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
