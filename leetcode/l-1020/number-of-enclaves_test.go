package l_1020

import "testing"

func Test_numEnclaves(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{[][]int{
			{0, 0, 0, 0},
			{1, 0, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0}}}, want: 3},
		{name: "test2", args: args{[][]int{
			{0, 1, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 0}}}, want: 0},
		{name: "test3", args: args{
			[][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}}, want: 15},
	}
	for _, tt := range tests {
		//t.Run(tt.name, func(t *testing.T) {
		//	if got := numEnclaves_v1(tt.args.A); got != tt.want {
		//		t.Errorf("numEnclaves_v1() = %v, want %v", got, tt.want)
		//	}
		//})
		t.Run(tt.name, func(t *testing.T) {
			if got := numEnclaves_v2(tt.args.A); got != tt.want {
				t.Errorf("numEnclaves_v2() = %v, want %v", got, tt.want)
			}
		})
	}
}
