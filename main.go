package main

import (
	"fmt"
)

func main() {
	// names := [5]string{"nico", "lynn", "dal"}
	// names[3] = "alala"
	// names[4] = "alala"
	// names[5] = "alala"
	// fmt.Println(names)

	names := []string{"nico", "lynn", "dal"}
	names = append(names, "flynn")

	fmt.Println(names)
}
