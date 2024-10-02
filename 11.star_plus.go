package main

import "fmt"

func printStarPlus() {
	size := 7
	mid := size / 2
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == mid || j == mid {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
