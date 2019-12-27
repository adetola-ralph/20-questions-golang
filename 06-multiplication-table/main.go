package main

import "fmt"

func main() {
	for i := 1; i < 13; i++ {
		for j := 2; j < 13; j++ {
			fmt.Printf("%d ", i*j)
		}

		fmt.Printf("\n")
	}
}
