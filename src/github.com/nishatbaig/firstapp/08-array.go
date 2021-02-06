package main

import "fmt"

func main() {
	var arr1 [5]int
	fmt.Println(arr1)

	arr1[4] = 100
	fmt.Println("Set: ", arr1)
	fmt.Println("Get: ", arr1[4])
	fmt.Println("Lenght: ", len(arr1))
	fmt.Println()

	arr2 := []int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	fmt.Println()

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
