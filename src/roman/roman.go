package roman

import (
	"sort"
)

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

//FromArabic number -> string
func FromArabic(number int) string {
	values := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	result := ""

	keys := []int{}
	for k := range values {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, k := range keys {

		if k == number+1 {
			number++
			result += "I"
		} else if k == number+10 {
			number = number + 10
			result += "X"
		} else if k == number+100 {
			number = number + 100
			result += "C"
		}

		for k <= number {
			value := values[k]
			number = number - k
			result = result + value
		}

		if number == 0 {
			break
		}
	}

	return result
}
