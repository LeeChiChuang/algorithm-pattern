package binarytree

import (
	"reflect"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	type args struct {
		a *TreeNode
	}

	arg1 := CreateTree([]int{1, 2, 3}, &TreeNode{}, 0, 3)
	arg2 := CreateTree([]int{-1, -1, -1}, &TreeNode{}, 0, 3)
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{a: arg1}, 6},
		{"t2", args{a: arg2}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxPathSum(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
