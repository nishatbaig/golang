package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0

	for _, i := range nums {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println()

	for k, v := range nums {
		fmt.Println(k, v)
	}
	fmt.Println()

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println(k, v)
	}
	fmt.Println()

	for k := range kvs {
		fmt.Println(k)
	}
	fmt.Println()

	for s, c := range "go" {
		fmt.Println(s, c)
	}
	fmt.Println()
}
