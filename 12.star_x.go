package main

import "fmt"

func printStarX() {
	size := 9
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j == i || j == size-i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
