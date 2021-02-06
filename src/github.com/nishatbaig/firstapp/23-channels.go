package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "abc"
	}()

	msg := <-ch
	fmt.Println(msg)
}
