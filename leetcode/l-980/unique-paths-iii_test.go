package l_980

import "testing"

func Test_uniquePathsIII(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}}}, want: 4},
		{name: "test2", args: args{[][]int{{0, 1}, {2, 0}}}, want: 0},
		{name: "test3", args: args{[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsIII(tt.args.grid); got != tt.want {
				t.Errorf("uniquePathsIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
