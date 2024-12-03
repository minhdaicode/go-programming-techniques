package stringalgo_test

import (
	stringalgo "go-programming-techniques/string/hash_map"
	"reflect"
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "test case 1",
			input: "III",
			want:  3,
		},
		{
			name:  "test case 2",
			input: "LVIII",
			want:  58,
		},
		{
			name:  "test case 3",
			input: "MCMXCIV",
			want:  1994,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stringalgo.RomanToInt(tt.input)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("got %d ,want %d", result, tt.want)
			}
		})
	}
}
