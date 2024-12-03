package array_test

import (
	"fmt"
	array "go-programming-techniques/array/two_pointers"
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		val   int
		want  int
	}{
		// input = [3,2,2,3]
		// val = 3
		// want = 2
		{
			name:  "test case 1",
			input: []int{3, 2, 2, 3},
			val:   3,
			want:  2,
		},
		// input = [0,1,2,2,3,0,4,2]
		// val = 3
		// want = 2
		{
			name:  "test case 2",
			input: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:   2,
			want:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			length := array.RemoveElement(tt.input, tt.val)

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
