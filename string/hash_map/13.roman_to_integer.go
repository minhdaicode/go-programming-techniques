package stringalgo

func RomanToInt(s string) int {
	sum := 0
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	for i, v := range s {
		sum += m[string(v)]
		if i != 0 {
			if m[string(s[i-1])] < m[string(v)] {
				sum -= 2 * m[string(s[i-1])]
			}
		}
	}
	return sum
}
