package prime

//GetPrimeFactors returns all prime factors of this given number
func GetPrimeFactors(number int) []int {
	var factors []int

	for ii := 2; ii <= number; ii++ {
		if isDividableBy(number, ii) {
			factors = append(factors, ii)
			number = number / ii
			moreFactors := GetPrimeFactors(number)
			factors = append(factors, moreFactors...)
			return factors
		}

	}

	return factors
}
