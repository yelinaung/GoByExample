package main

import "fmt"
import "math"

// const declares a constant value
const s string = "constant"

func main() {
	fmt.Println(s)

	// A const state can appear anywhere
	const n = 5000000

	// arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// explict cast
	fmt.Println(int64(d))

	// math package
	fmt.Println(math.Sin(n))

}
