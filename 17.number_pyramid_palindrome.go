package main

import "fmt"

func printNumberPalindromePyramid() {
	n := 5 // Number of rows

	for i := 1; i <= n; i++ {
		// Print leading spaces
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}

		// Print increasing numbers
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}

		// Print decreasing numbers
		for j := i - 1; j >= 1; j-- {
			fmt.Print(j)
		}

		// Move to the next line
		fmt.Println()
	}
}
