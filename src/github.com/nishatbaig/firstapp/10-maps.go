package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["k1"] = 7
	map1["k2"] = 13
	map1["k3"] = 20
	fmt.Println("map", map1)

	v1 := map1["mk1"]
	fmt.Println(v1)

	fmt.Println("Length: ", len(map1))

	delete(map1, "k1")
	fmt.Println("Post Delete: ", map1)

	_, err := map1["k1"]
	fmt.Println("err:", err)

	n := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println(n)
}
