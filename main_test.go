package main

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestSame(t *testing.T) {
	type args struct {
		t1 *tree.Tree
		t2 *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Same(tree.New(1), tree.New(2)) should be false",
			args: args{
				t1: tree.New(1), t2: tree.New(2),
			},
			want: false,
		},
		{
			name: "Same(tree.New(1), tree.New(1)) should be true",
			args: args{
				t1: tree.New(1), t2: tree.New(1),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Same(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("Same() = %v, want %v", got, tt.want)
			}
		})
	}
}
