package main

import "fmt"

func main() {
	sum := add(1, 2)
	fmt.Println(sum)

	plus := plusPlus(1, 2, 3)
	fmt.Println(plus)
}

func add(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}
