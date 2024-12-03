package array_test

import (
	array "go-programming-techniques/array/two_pointers"
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	var tests = []struct {
		name    string
		input1  []int
		length1 int
		input2  []int
		length2 int
		want    []int
	}{
		// input1 = [1,2,3,0,0,0]
		// length1 = 3
		// input2 = [2,5,6]
		// length2 = 3
		// want = [1,2,2,3,5,6]
		{
			name:    "test case 1",
			input1:  []int{1, 2, 3, 0, 0, 0},
			length1: 3,
			input2:  []int{2, 5, 6},
			length2: 3,
			want:    []int{1, 2, 2, 3, 5, 6},
		},
		// input1 = [1]
		// length1 = 1
		// input2 = []
		// length2 = 0
		// want = [1]
		{
			name:    "test case 2",
			input1:  []int{1},
			length1: 1,
			input2:  []int{},
			length2: 0,
			want:    []int{1},
		},
		// input1 = [0]
		// length1 = 0
		// input2 = [1]
		// length2 = 1
		// want = [1]
		{
			name:    "test case 3",
			input1:  []int{0},
			length1: 0,
			input2:  []int{1},
			length2: 1,
			want:    []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			array.Merge(tt.input1, tt.length1, tt.input2, tt.length2)
			if !reflect.DeepEqual(tt.input1, tt.want) {
				t.Errorf("got %d ,want %d", tt.input1, tt.want)
			}
		})
	}
}
