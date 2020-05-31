package cpt14_network_flow

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewEdmondsKarp(t *testing.T) {
	type args struct {
		src string
		s   int
		t   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{
			src: "network1.txt", s: 0, t: 3,
		}, want: 5},
		{name: "test2", args: args{
			src: "network2.txt", s: 0, t: 5,
		}, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEdmondsKarp(tt.args.src, tt.args.s, tt.args.t); !reflect.DeepEqual(got.Result(), tt.want) {
				t.Errorf("NewEdmondsKarp() = %v, want %v", got.Result(), tt.want)
			} else {
				fmt.Println(got)
			}
		})
	}
}
