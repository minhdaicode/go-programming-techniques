package array_test

import (
	array "go-programming-techniques/array/binary_search"
	"reflect"
	"testing"
)

func TestCheckIfNAndItsDoubleExist(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		want  bool
	}{
		// input = [10,2,5,3]
		// want = true
		{
			name:  "test case 1",
			input: []int{10, 2, 5, 3},
			want:  true,
		},
		// input = [3,1,7,11]
		// want = false
		{
			name:  "test case 2",
			input: []int{3, 1, 7, 11},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(array.CheckIfExist(tt.input), tt.want) {
				t.Errorf("got %v ,want %v", array.CheckIfExist(tt.input), tt.want)
			}
		})
	}
}
