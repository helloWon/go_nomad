package main

import (
	"fmt"
	"strings"
)

// func multiply(a int, b int) int {
func multiply(a, b int) int {
	return a * b
}

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main()  {
	// 곱셈
	fmt.Println(multiply(2,2))

	// 값 여러 개 return
	totalLenght, upperName := lenAndUpper("sw")
	fmt.Println(totalLenght, upperName)

	// args 여러 개
	repeatMe("nico", "sw", "lynn")
}