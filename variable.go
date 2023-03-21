package main

import (
	"fmt"
)

func main() {
	// var greetme = "hello,"
	// var school string = "33"
	school := 33 * 35
	fmt.Print(school)
	fmt.Println()
	// fmt.Print(greetme, name)

	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}

}
