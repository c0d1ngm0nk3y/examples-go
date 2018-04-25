package main

import "fmt"

func main() {
	x := Sum(4, 38)
	fmt.Printf("hello world %d\n", x)
}

func Sum(x int, y int) int {
	return x + y
}
