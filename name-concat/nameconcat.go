package main

import (
	"fmt"
	"time"

	"github.com/adetola-ralph/20-questions-golang/name-concat/util"
)

func main() {
	var firstName string
	var lastName string
	var age int
	today := time.Now()

	firstName, lastName, age = util.GetUserInputBufio()

	// fmt.Printf("Welcome, %s %s (<%d>)\n", firstName, lastName, age)
	fmt.Printf("Welcome, %s %s (<%d>)\n", firstName, lastName, today.Year()-age)
}
