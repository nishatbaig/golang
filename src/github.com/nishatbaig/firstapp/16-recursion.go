package main

import "fmt"

func main() {
	fmt.Println(fact(7))
}

func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)
}
