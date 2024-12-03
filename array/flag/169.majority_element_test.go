package array_test

import (
	"fmt"
	array "go-programming-techniques/array/flag"
	"reflect"
	"testing"
)

func TestMaJorityElement(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		want  int
	}{
		// input = [3,2,3]
		// want = 3
		{
			name:  "test case 1",
			input: []int{3, 2, 3},
			want:  3,
		},
		// input = [2,2,1,1,1,2,2]
		// want = 2
		{
			name:  "test case 2",
			input: []int{2, 2, 1, 1, 1, 2, 2},
			want:  2,
		},
		// input = [3,3,4]
		// want = 3
		{
			name:  "test case 3",
			input: []int{3, 3, 4},
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := array.MajorityElement(tt.input)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("got %d ,want %d", result, tt.want)
			}
			fmt.Println(result)
		})
	}
}
