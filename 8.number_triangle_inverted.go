package main

import "fmt"

func printNumberTriangleInverted() {
	for i := 0; i <= 5; i++ {
		count := 1
		for j := 1; j <= 5-i; j++ {
			fmt.Print(count, " ")
			count++
		}
		fmt.Println()
	}
}
