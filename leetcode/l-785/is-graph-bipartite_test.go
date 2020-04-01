package l_785

import "testing"

func Test_isBipartite(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{graph: [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}}, want: true},
		{name: "test2", args: args{graph: [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}}, want: false},
		{name: "testx1", args: args{graph: [][]int{{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8}}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBipartite(tt.args.graph); got != tt.want {
				t.Errorf("isBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}
