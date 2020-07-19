package backtrack

import (
	"reflect"
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {
	type args struct {
		a []int
		s int
	}

	arg1 := []int{1, 1, 1, 1, 1}
	arg2 := 3
	want1 := 5
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{a: arg1, s: arg2}, want1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTargetSumWays(tt.args.a, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
