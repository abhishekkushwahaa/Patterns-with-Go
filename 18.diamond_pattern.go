package main

import "fmt"

func printDiamondPattern() {
	n := 5

	// Upper half of the pattern (including the middle row)
	for i := 1; i <= n; i++ {
		// Print leading spaces
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}

		// Print the repeated number i
		for j := 1; j <= (2*i - 1); j++ {
			fmt.Print(i)
		}

		// Move to the next line
		fmt.Println()
	}

	// Lower half of the pattern (below the middle row)
	for i := n - 1; i >= 1; i-- {
		// Print leading spaces
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}

		// Print the repeated number i
		for j := 1; j <= (2*i - 1); j++ {
			fmt.Print(i)
		}

		// Move to the next line
		fmt.Println()
	}
}
