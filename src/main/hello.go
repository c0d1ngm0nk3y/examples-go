package main

import (
	"fmt"
	"prime"
)

func main() {
	x := Sum(4, 38)
	fmt.Printf("hello world %d\n", x)
	fmt.Printf("113 is prime: %v\n", prime.IsPrime(113))
}

func Sum(x int, y int) int {
	return x + y
}
