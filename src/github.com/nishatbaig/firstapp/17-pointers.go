package main

import "fmt"

func main() {
	zval := 10
	fmt.Println(zval)
	zeroval(zval)
	fmt.Println(zval)

	pval := 15
	fmt.Println(pval)
	zeroPointer(&pval)
	fmt.Println(pval)
	fmt.Println(&pval)
}

func zeroval(num int) {
	num = 0
}

func zeroPointer(num *int) {
	*num = 0
}
