package main

import "fmt"

func main() {
	// empty slice
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("len", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5] // this slice up to (but excluding) s[2]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	//slice in single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
}
