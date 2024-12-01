package array_test

import (
	array "go-programming-techniques/array/two_pointers"
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		step  int
		want  []int
	}{
		// input = [1,2,3,4,5,6,7]
		// step = 3
		// want = 2
		{
			name:  "test case 1",
			input: []int{1, 2, 3, 4, 5, 6, 7},
			step:  3,
			want:  []int{5, 6, 7, 1, 2, 3, 4},
		},
		// input = [0,1,2,2,3,0,4,2]
		// val = 3
		// want = 2
		{
			name:  "test case 2",
			input: []int{-1, -100, 3, 99},
			step:  2,
			want:  []int{3, 99, -1, -100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			array.Rotate(tt.input, tt.step)

			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("got %d ,want %d", tt.input, tt.want)
			}
		})
	}
}
