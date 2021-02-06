package main

import "fmt"

func main() {
	rect := rect{
		width:  10,
		height: 5,
	}

	fmt.Println(rect.area())
	fmt.Println(rect.perimeter())

	rp := &rect
	fmt.Println(rp.area())
	fmt.Println(rp.perimeter())
}

type rect struct {
	width  int
	height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perimeter() int {
	return 2*r.width + 2*r.height
}
