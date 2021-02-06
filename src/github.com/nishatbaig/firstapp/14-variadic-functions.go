package main

import "fmt"

func main() {
	variad(1, 2, 3, 4, 56)
	variad(1, 2, 3)
	arr := []int{7, 8, 9}
	variad(arr...)
}

func variad(k ...int) {
	fmt.Print(k, " ")
	total := 0
	for _, l := range k {
		total += l
	}
	fmt.Println(total)
}
