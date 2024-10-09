package main

func printPascalTriangle() {
	for i := 0; i < 5; i++ {
		count := 1
		for j := 0; j <= i; j++ {
			print(count, " ")
			count = count * (i - j) / (j + 1)
		}
		println("")
	}
}
