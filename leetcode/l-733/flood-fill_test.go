package l_733

import (
	"reflect"
	"testing"
)

func Test_floodFill(t *testing.T) {
	type args struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test1", args: args{
			image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			sr:       1,
			sc:       1,
			newColor: 2,
		}, want: [][]int{
			{2, 2, 2},
			{2, 2, 0},
			{2, 0, 1},
		}},
		{name: "test2", args: args{
			image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
				{1, 0, 1},
			},
			sr:       2,
			sc:       2,
			newColor: 2,
		}, want: [][]int{
			{1, 1, 1},
			{1, 1, 0},
			{1, 0, 2},
			{1, 0, 2},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.newColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
