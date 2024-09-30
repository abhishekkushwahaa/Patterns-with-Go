package main

import "fmt"

func printNumberSquare() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(j+1, " ")
		}
		fmt.Println()
	}
}
