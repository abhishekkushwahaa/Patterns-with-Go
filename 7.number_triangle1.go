package main

import "fmt"

func printNumberTriangle1() {
	for i := 0; i <= 5; i++ {
		count := 1
		for j := 0; j < 5-i; j++ {
			fmt.Print("  ")
		}
		for k := 0; k < i; k++ {
			fmt.Print(count, " ")
			count++
		}
		fmt.Println()
	}
}
