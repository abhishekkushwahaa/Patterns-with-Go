package main

import "fmt"

func printStarTriangleReverse() {
	for i := 0; i <= 5; i++ {
		fmt.Println()
		for j := 0; j < 5-i; j++ {
			fmt.Print("* ")
		}
	}
}
