package main

import "fmt"

func main() {
	a, b := vals(1, 2)
	fmt.Println(a, b)

	_, c := vals(2, 3)
	fmt.Println(c)
}

func vals(a, b int) (int, int) {
	return a, b
}
