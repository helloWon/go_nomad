package main

import "fmt"

func superAdd(numbers ...int) int {
	total := 0
	for idx, number := range numbers {
		fmt.Println(idx, number)
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
