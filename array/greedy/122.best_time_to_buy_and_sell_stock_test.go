package array_test

import (
	array "go-programming-techniques/array/greedy"
	"reflect"
	"testing"
)

func TestMaxProfitII(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		want  int
	}{
		// input = [7,1,5,3,6,4]
		// want = 7
		{
			name:  "test case 1",
			input: []int{7, 1, 5, 3, 6, 4},
			want:  7,
		},
		// input = [1,2,3,4,5]
		// want = 4
		{
			name:  "test case 2",
			input: []int{1, 2, 3, 4, 5},
			want:  4,
		},
		// input = [7,6,4,3,1]
		// want = 0
		{
			name:  "test case 3",
			input: []int{7, 6, 4, 3, 1},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			profit := array.MaxProfitII(tt.input)
			if !reflect.DeepEqual(profit, tt.want) {
				t.Errorf("got %d ,want %d", profit, tt.want)
			}
		})
	}
}
