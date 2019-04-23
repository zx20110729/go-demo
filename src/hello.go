package main

import (
	"fmt"
)

func add(a int, b int) int {
	var sum int
	sum = a + b
	return sum
}

func main() {
	fmt.Println("Hello World!")
	var sum int
	sum = add(100, 5)
	fmt.Println("sum = ", sum)
}
