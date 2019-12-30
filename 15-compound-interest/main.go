package main

import (
	"fmt"
)

func main() {
	var principal, rate float32
	var timeInYears, formula int

	// not necessary but decided to try it out
	interestMap := make(map[int]func(float32, float32, int) float32)
	interestMap[1] = simpleInterest
	interestMap[2] = compoundInterest

	fmt.Println(`
1. For simple interest
2. For compound interest
	`)

	fmt.Scanln(&formula)

	fmt.Println("Enter the principal:")
	fmt.Scanln(&principal)

	fmt.Println("Enter the rate")
	fmt.Scanln(&rate)

	fmt.Println("Enter the time in years")
	fmt.Scanln(&timeInYears)

	result := interestMap[formula](principal, rate, timeInYears)
	fmt.Println(result)
	fmt.Printf("%.2f\n", result)
}

func simpleInterest(principal, rate float32, timeInYears int) float32 {
	interest := (principal * rate * float32(timeInYears)) / 100
	return interest + principal
}

func compoundInterest(principal, rate float32, timeInYears int) float32 {
	for ; timeInYears > 0; timeInYears-- {
		interest := principal * (rate / 100)
		principal += interest
	}

	return principal
}
