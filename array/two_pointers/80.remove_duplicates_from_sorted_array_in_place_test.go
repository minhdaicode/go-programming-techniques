package array_test

import (
	"fmt"
	array "go-programming-techniques/array/two_pointers"
	"reflect"
	"testing"
)

func TestRemoveDuplicatesFromSortedArrayInPlace(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		want  int
	}{
		// input = [1,1,1,2,2,3]
		// want = 5
		// nums = [1,1,2,2,3]
		{
			name:  "test case 1",
			input: []int{1, 1, 1, 2, 2, 3},
			want:  5,
		},
		// input = [0,0,1,1,1,1,2,3,3]
		// want = 7
		// nums = [0,0,1,1,2,3,3]
		{
			name:  "test case 2",
			input: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want:  7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			length := array.RemoveDuplicatesInPlace(tt.input)

			if !reflect.DeepEqual(length, tt.want) {
				t.Errorf("got %d ,want %d", length, tt.want)
			}

			fmt.Print("[ ")
			for i := 0; i < length; i++ {
				fmt.Printf("%d ", tt.input[i])
			}
			fmt.Print("]")
		})
	}
}
