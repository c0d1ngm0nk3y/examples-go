package prime

import "math"

//IsPrime checks if it is a Prime
func IsPrime(number int) bool {
	if number < 2 {
		return false
	} else if number == 2 {
		return true
	} else if isDividableBy(number, 2) {
		return false
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

func calculateMaxPossibleDivider(number int) int {
	return int(math.Floor(math.Sqrt(float64(number))))
}

func isDividableBy(number int, divider int) bool {
	return (number % divider) == 0
}
