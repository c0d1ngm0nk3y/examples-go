package main

import (
	"fmt"
	"prime"
)

func main() {
	x := Sum(4, 38)
	fmt.Printf("hello world %d\n", x)
	fmt.Printf("113 is prime: %v\n", prime.IsPrime(113))
	factors := prime.GetPrimeFactors(60)
	fmt.Printf("Factors of 60: %v (%d)", factors, len(factors))
}

//Sum is guessing what?
func Sum(x int, y int) int {
	return x + y
}
