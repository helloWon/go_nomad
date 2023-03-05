package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	foods := []string{"bread", "chicken"}
	// nico := person{"nico", 12, foods}
	nico := person{name: "nico", age: 12, favFood: foods}
	fmt.Println(nico.name)
}
