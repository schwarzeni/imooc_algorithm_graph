package l_1034

import (
	"reflect"
	"testing"
)

func Test_colorBorder(t *testing.T) {
	type args struct {
		grid  [][]int
		r0    int
		c0    int
		color int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test1", args: args{
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			r0:    1,
			c0:    1,
			color: 2,
		}, want: [][]int{
			{2, 2, 2},
			{2, 1, 2},
			{2, 2, 2},
		}},
		{name: "test2", args: args{
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			r0:    0,
			c0:    0,
			color: 2,
		}, want: [][]int{
			{2, 2, 2},
			{2, 1, 2},
			{2, 2, 2},
		}},
		{name: "test3", args: args{
			grid: [][]int{
				{1, 1},
				{1, 2},
			},
			r0:    0,
			c0:    0,
			color: 3,
		}, want: [][]int{
			{3, 3},
			{3, 2},
		}},
		{name: "test4", args: args{
			grid: [][]int{
				{1, 2, 2},
				{2, 3, 2},
			},
			r0:    0,
			c0:    1,
			color: 3,
		}, want: [][]int{
			{1, 3, 3},
			{2, 3, 3},
		}},
		{name: "test5", args: args{
			grid: [][]int{
				{1, 1, 1, 5},
				{1, 2, 1, 1},
				{1, 3, 1, 4},
				{5, 4, 2, 3},
			},
			r0:    0,
			c0:    1,
			color: 10,
		}, want: [][]int{
			{10, 10, 10, 5},
			{10, 2, 10, 10},
			{10, 3, 10, 4},
			{5, 4, 2, 3},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := colorBorder(tt.args.grid, tt.args.r0, tt.args.c0, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("colorBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
