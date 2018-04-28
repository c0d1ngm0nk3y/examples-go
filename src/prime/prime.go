package prime

import "math"

//IsPrime checks if it is a Prime
func IsPrime(number int) bool {
	applicable, result := isPrimeSpecialCase(number)
	if applicable {
		return result
	}

	maxDivider := calculateMaxPossibleDivider(number)
	for ii := 3; ii <= maxDivider; ii = ii + 2 {
		dividerFound := isDividableBy(number, ii)
		if dividerFound {
			return false
		}
	}
	return true
}

func isPrimeSpecialCase(number int) (bool, bool) {
	if number < 2 {
		return true, false
	} else if number == 2 {
		return true, true
	} else if isDividableBy(number, 2) {
		return true, false
	}

	return false, false
}

func calculateMaxPossibleDivider(number int) int {
	return int(math.Floor(math.Sqrt(float64(number))))
}

func isDividableBy(number int, divider int) bool {
	return (number % divider) == 0
}
