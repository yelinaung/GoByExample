package main

import "fmt"

func main() {
	i := 1

	// single condition loop
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("xxxxxxxxxxxx")
	// class initial / condiiton / after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("xxxxxxxxxxxx")

	// for without a condition will loop repeatedly until you 'break'
	for {
		fmt.Println("loop")
		break
	}
}
