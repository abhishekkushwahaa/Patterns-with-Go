package main

import "fmt"

func printOddNumberTriangle() {
	for i := 1; i <= 10; i += 2 {
		for j := 0; j < (i+1)/2; j++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}
