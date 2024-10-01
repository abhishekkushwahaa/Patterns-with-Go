package main

import "fmt"

func printSolidRectangle() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
