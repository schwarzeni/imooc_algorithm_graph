package l_529

import (
	"reflect"
	"testing"
)

func Test_updateBoard(t *testing.T) {
	type args struct {
		board [][]byte
		click []int
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{name: "test1", args: args{
			board: [][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'}},
			click: []int{3, 0},
		}, want: [][]byte{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'M', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'}}},
		{name: "test2", args: args{
			board: [][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'}},
			click: []int{1, 2},
		}, want: [][]byte{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'X', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'}}},
		{name: "test3", args: args{
			board: [][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'}},
			click: []int{1, 1},
		}, want: [][]byte{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'M', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'}},
		},
		{name: "test4", args: args{
			board: [][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'}},
			click: []int{3, 0},
		}, want: [][]byte{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'M', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'}},
		},
		{name: "test5", args: args{
			board: [][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'}},
			click: []int{100, 100},
		}, want: [][]byte{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'M', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'}},
		},
		{name: "test6", args: args{
			board: [][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'X', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'}},
			click: []int{1, 2},
		}, want: [][]byte{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'X', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'}}},
		{name: "test_1", args: args{
			board: [][]byte{
				{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'M'},
				{'E', 'E', 'M', 'E', 'E', 'E', 'E', 'E'},
				{'M', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'M', 'M', 'E', 'E', 'E', 'E'}},
			click: []int{0, 0},
		}, want: [][]byte{
			{'B', 'B', 'B', 'B', 'B', 'B', '1', 'E'},
			{'B', '1', '1', '1', 'B', 'B', '1', 'M'},
			{'1', '2', 'M', '1', 'B', 'B', '1', '1'},
			{'M', '2', '1', '1', 'B', 'B', 'B', 'B'},
			{'1', '1', 'B', 'B', 'B', 'B', 'B', 'B'},
			{'B', 'B', 'B', 'B', 'B', 'B', 'B', 'B'},
			{'B', '1', '2', '2', '1', 'B', 'B', 'B'},
			{'B', '1', 'M', 'M', '1', 'B', 'B', 'B'}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateBoard(tt.args.board, tt.args.click); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want \n%s, but got \n%s", printByteArr(tt.want), printByteArr(got))
			}
		})
	}
}

func printByteArr(arr [][]byte) (result string) {
	for _, i := range arr {
		for _, j := range i {
			result += string(j) + string(' ')
		}
		result += "\n"
	}
	return
}
