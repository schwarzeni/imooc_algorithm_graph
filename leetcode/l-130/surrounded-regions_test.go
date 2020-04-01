package l_130

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		name string
		args [][]byte
		want [][]byte
	}{
		{name: "test1", args: [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}, want: [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'X'},
		}},
		{name: "test2", args: [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'O', 'X'},
		}, want: [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'O', 'X'},
		}},
		{name: "test3", args: [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
			{'X', 'O', 'O', 'O'},
			{'X', 'X', 'X', 'X'},
		}, want: [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'X'},
			{'X', 'O', 'O', 'O'},
			{'X', 'X', 'X', 'X'},
		}},
		{name: "test4", args: [][]byte{
			{'X', 'O', 'X', 'X'},
		}, want: [][]byte{
			{'X', 'O', 'X', 'X'},
		}},
		{name: "test5", args: [][]byte{
			{'X'},
			{'X'},
			{'X'},
			{'O'},
			{'X'},
			{'X'},
		}, want: [][]byte{
			{'X'},
			{'X'},
			{'X'},
			{'O'},
			{'X'},
			{'X'},
		}},
		{name: "test_1", args: [][]byte{
			{'O', 'O', 'O', 'O', 'X', 'X'},
			{'O', 'O', 'O', 'O', 'O', 'O'},
			{'O', 'X', 'O', 'X', 'O', 'O'},
			{'O', 'X', 'O', 'O', 'X', 'O'},
			{'O', 'X', 'O', 'X', 'O', 'O'},
			{'O', 'X', 'O', 'O', 'O', 'O'},
		}, want: [][]byte{
			{'O', 'O', 'O', 'O', 'X', 'X'},
			{'O', 'O', 'O', 'O', 'O', 'O'},
			{'O', 'X', 'O', 'X', 'O', 'O'},
			{'O', 'X', 'O', 'O', 'X', 'O'},
			{'O', 'X', 'O', 'X', 'O', 'O'},
			{'O', 'X', 'O', 'O', 'O', 'O'},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args)
			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("solve() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
