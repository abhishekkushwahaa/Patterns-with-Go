package main

import "fmt"

func printAlphabetSquare() {
	alphabet := 'A'
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%c ", alphabet)
			alphabet++
			if alphabet > 'E' {
				alphabet = 'A'
			}
		}
		fmt.Println()
	}
}
