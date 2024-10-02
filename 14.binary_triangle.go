package main

import "fmt"

func binaryTriangle() {
	for i := 0; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}
