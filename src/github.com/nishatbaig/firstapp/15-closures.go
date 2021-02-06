package main

import "fmt"

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println()

	newInt := intSeq()
	fmt.Println(newInt())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
