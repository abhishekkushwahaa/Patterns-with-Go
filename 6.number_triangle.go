package main

import "fmt"

func printNumberTriangle() {
	for i := 1; i <= 5; i++ {
		count := 1
		for j := 0; j < i; j++ {
			fmt.Print(count, " ")
			count++
		}
		fmt.Println("")
	}
}
