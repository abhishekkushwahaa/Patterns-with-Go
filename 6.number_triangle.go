package main

import "fmt"

func printNumberTriangle() {
	//count := 1
	for i := 0; i <= 6; i++ {
		count := 1
		for j := 1; j < i; j++ {
			fmt.Print(count, " ")
			count++
		}
		fmt.Println("")
	}
}
