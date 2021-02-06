package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 10; num < 0 {
		fmt.Println(num, "is negative number")
	} else if num < 9 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}