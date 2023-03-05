package main

import "fmt"

func main() {
	// pointers
	a := 2
	b := &a
	a = 10
	fmt.Println(a, b)

	// arrays
	names := [6]string{"nico", "lynn", "dal"}
	names[3] = "blahblah"
	names[4] = "blahblah"
	names[5] = "blahblah"
	fmt.Println(names)

	// slices
	names2 := []string{"nico", "lynn", "dal"}
	names2 = append(names2, "blah")
	fmt.Println(names2)
}
