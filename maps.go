package main

import "fmt"

func main() {

	// make ([key-type]val-type)
	m := make(map[string]string)

	m["k1"] = "aok"
	m["k2"] = "bok"

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("maps:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
