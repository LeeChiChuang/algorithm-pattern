package backtrack

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	type args struct {
		a []int
	}

	arg1 := []int{1, 2, 3}
	want1 := [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 3},
		{2},
		{2, 3},
		{3},
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"t1", args{a: arg1}, want1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subsets(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
