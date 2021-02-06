package main

import (
	"fmt"
)

func main() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Alice", age: 30})
	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	fmt.Println()
	p1 := Person{
		name: "Sumi",
		age:  18,
	}
	fmt.Println(p1)

	var p2 Person
	p2 = Person{
		name: "Nishat",
		age:  20,
	}
	fmt.Println(p2)
	fmt.Println(p2.name)

	sp := &p2
	fmt.Println(sp.age)
	sp.age = 52
	fmt.Println(sp.age)
	fmt.Println(p2.age)
}

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	return &p
}
