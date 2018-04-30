package roman

//ToArabic string ->int
func ToArabic(romanString string) int {
	literals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := 0
	lastValue := literals["M"]

	for ii := 0; ii < len(romanString); ii++ {
		nextChar := romanString[ii : ii+1]
		value := literals[nextChar]

		if lastValue < value {
			sum = sum - 2*lastValue
		}

		sum += value
		lastValue = value
	}

	return sum
}
