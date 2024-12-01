package array_test

import (
	"go-programming-techniques/hashtable/array"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		name   string
		input  []int
		target int
		want   []int
	}{
		// input = [2,7,11,15]
		// target = 9
		// want = [0,1]
		{
			name:   "test case 1",
			input:  []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		// input = [3,2,4]
		// target = 6
		// want = [1,2]
		{
			name:   "test case 2",
			input:  []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		// input = [3,3]
		// target = 6
		// want = [0,1]
		{
			name:   "test case 3",
			input:  []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := array.TwoSum(tt.input, tt.target)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d ,want %d", ans, tt.want)
			}
		})
	}
}
