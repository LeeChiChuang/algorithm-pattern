package backtrack

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	type args struct {
		a []int
	}

	arg1 := []int{1, 2, 3}
	arg2 := []int{1}
	arg3 := []int{5, 4, 6, 2}
	want1 := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	want2 := [][]int{
		{1},
	}
	want3 := [][]int{
		{5, 4, 6, 2},
		{5, 4, 2, 6},
		{5, 6, 4, 2}, {5, 6, 2, 4}, {5, 2, 4, 6}, {5, 2, 6, 4},
		{4, 5, 6, 2}, {4, 5, 2, 6}, {4, 6, 5, 2}, {4, 6, 2, 5}, {4, 2, 5, 6},
		{4, 2, 6, 5}, {6, 5, 4, 2}, {6, 5, 2, 4}, {6, 4, 5, 2}, {6, 4, 2, 5},
		{6, 2, 5, 4}, {6, 2, 4, 5}, {2, 5, 4, 6}, {2, 5, 6, 4}, {2, 4, 5, 6},
		{2, 4, 6, 5}, {2, 6, 5, 4}, {2, 6, 4, 5},
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"t1", args{a: arg1}, want1},
		{"t2", args{a: arg2}, want2},
		{"t3", args{a: arg3}, want3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permute(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
