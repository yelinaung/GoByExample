package main

import "fmt"

func main() {

	// create an array that will hold exactly 5 ints
	var a [5]int
	fmt.Println("emp:", a)

	// array[index] = value
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// get len of array
	fmt.Println("len", len(a))

	// one line delcaration
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl", b)

	// creating two d array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
