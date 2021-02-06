package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set: ", s)
	fmt.Println("Get: ", s[1])
	fmt.Println("Len: ", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Append: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy: ", c)

	l := s[2:5]
	fmt.Println("Slice1: ", l)

	l = s[2:]
	fmt.Println("Slice2: ", l)

	l = s[:5]
	fmt.Println("Slice3: ", l)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
