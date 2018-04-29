package prime

//GetPrimeFactors returns all prime factors of this given number
func GetPrimeFactors(number int) []int {
	var factors []int

	if number < 2 {
		return factors
	}

	factors = append(factors, number)
	return factors
}
