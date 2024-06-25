package main

import (
	"fmt"
)

// func multiply(a int, b int) int {
// 	return a * b
// }

// func multiply(a, b int) int {
// 	return a * b
// }

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// fmt.Println(multiply(2, 2))

	// totalenght, upperName := lenAndUpper("niko")
	// totalenght, _ := lenAndUpper("niko")
	// fmt.Println(totalenght)

	repeatMe("nico", "lynn", "dal", "marl", "flynn")
}
