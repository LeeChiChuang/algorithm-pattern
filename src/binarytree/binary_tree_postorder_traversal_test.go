package binarytree

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	type args struct {
		a *TreeNode
	}
	arg1 := CreateTree([]int{1, 2, 3, 4}, &TreeNode{}, 0, 4)
	arg2 := CreateTree([]int{1, 2, 3, 4, 5, 6}, &TreeNode{}, 0, 6)
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{a:arg1}, []int{4, 2, 3, 1} },
		{"test2", args{a:arg2}, []int{4, 5, 2, 6, 3, 1} },
	}

	for _, tt := range tests  {
		t.Run(tt.name, func(t *testing.T) {
			if got := PostorderTraversal(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostorderTraversal2(t *testing.T) {
	type args struct {
		a *TreeNode
	}
	arg1 := CreateTree([]int{1, 2, 3, 4}, &TreeNode{}, 0, 4)
	arg2 := CreateTree([]int{1, 2, 3, 4, 5, 6}, &TreeNode{}, 0, 6)
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{a:arg1}, []int{4, 2, 3, 1} },
		{"test2", args{a:arg2}, []int{4, 5, 2, 6, 3, 1} },
	}

	for _, tt := range tests  {
		t.Run(tt.name, func(t *testing.T) {
			if got := PostorderTraversal2(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostorderTraversal2() = %v, want %v", got, tt.want)
			}
		})
	}
}