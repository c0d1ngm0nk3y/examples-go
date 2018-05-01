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

		apply := true
		for apply {
			apply = false
			x, canBeSubstracted := xCanBeSubstracted(number, k)

			if canBeSubstracted {
				number = number + x
				result += values[x]
			}

			for k <= number {
				value := values[k]
				number = number - k
				result = result + value
				apply = true
			}
		}

	}

	return result
}

func xCanBeSubstracted(number int, k int) (int, bool) {
	if isBetween(number, k-2, k) && (k <= 10) && (k >= 5) {
		return 1, true
	} else if isBetween(number, k-11, k) && (k <= 100) && (k >= 50) {
		return 10, true
	} else if isBetween(number, k-101, k) && (k <= 1000) && (k >= 500) {
		return 100, true
	}

	return 0, false
}

func isBetween(number int, lower int, upper int) bool {
	return (lower < number) && (number < upper)
}
