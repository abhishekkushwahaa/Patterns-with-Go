package main

import "fmt"

func floydTriangle() {
	count := 1
	for i := 0; i <= 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(count, " ")
			count++
		}
		fmt.Println()
	}
}
