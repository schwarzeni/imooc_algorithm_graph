package l_827

import "testing"

func Test_largestIsland(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{[][]int{
			{1, 0},
			{0, 1},
		}}, want: 3},
		{name: "test2", args: args{[][]int{
			{1, 0},
			{1, 1},
		}}, want: 4},
		{name: "test3", args: args{[][]int{
			{1, 1},
			{1, 1},
		}}, want: 4},
		{name: "test4", args: args{[][]int{
			{0, 0},
			{0, 0},
		}}, want: 1},
		{name: "test5", args: args{[][]int{
			{0, 0, 0, 0, 0, 0, 1},
			{0, 1, 1, 0, 1, 1, 1},
			{0, 1, 1, 0, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 1, 1, 0, 1, 1, 0},
			{1, 1, 1, 0, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 1},
		}}, want: 14},
		{name: "test_1", args: args{[][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 1, 0, 0},
			{0, 1, 0, 0, 1, 0, 0},
			{1, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 0, 0},
			{0, 1, 1, 1, 1, 0, 0},
		}}, want: 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestIsland(tt.args.grid); got != tt.want {
				t.Errorf("largestIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}
