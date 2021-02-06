package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "abc"
	ch <- "def"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
