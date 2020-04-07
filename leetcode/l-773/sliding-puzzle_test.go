package l_773

import "testing"

func Test_slidingPuzzle(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{[][]int{{4, 1, 2}, {5, 0, 3}}}, want: 5},
		{name: "test2", args: args{[][]int{{3, 2, 4}, {1, 5, 0}}}, want: 14},
		{name: "test3", args: args{[][]int{{1, 2, 3}, {5, 4, 0}}}, want: -1},
		{name: "test4", args: args{[][]int{{1, 2, 3}, {4, 0, 5}}}, want: 1},
		{name: "test5", args: args{[][]int{{1, 2, 3}, {4, 5, 0}}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slidingPuzzle(tt.args.board); got != tt.want {
				t.Errorf("slidingPuzzle() = %v, want %v", got, tt.want)
			}
		})
	}
}
