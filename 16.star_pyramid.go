package main

import "fmt"

func printStarPyramid() {
	for i := 1; i < 5; i++ {
		// Print spaces
		for j := 1; j < 5-i; j++ {
			fmt.Print(" ")
		}
		// Print stars
		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
