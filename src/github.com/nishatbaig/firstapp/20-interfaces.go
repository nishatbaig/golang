package main

import (
	"fmt"
	"math"
)

func main() {
	r := rect{width: 3, height: 4}
	measure(r)

	c := circle{radius: 5}
	measure(c)
}

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())

}
