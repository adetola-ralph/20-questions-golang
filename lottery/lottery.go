package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	// rand.Seed(time.Now().UnixNano())
}

func main() {
	// var a string
	i := 0
	// numbers := []string{}
	numbers := make([]string, 3)
	// numbers := [3]string{}
	rand.Seed(time.Now().UnixNano())

	// fmt.Println("Please press any key to continue")
	// fmt.Scanln(&a)
	// fmt.Println(len(numbers), "length")
	// fmt.Println(cap(numbers), "capacity")

	for i < 3 {
		// numbers = append(numbers, strconv.Itoa(rand.Intn(10)))
		numbers[i] = strconv.Itoa(rand.Intn(10))
		i++
	}

	result := strings.Join(numbers, " ")
	fmt.Println(result)

	// if strings.Contains(result, "7") {
	// if strings.Count(result, "7") == 2 {
	if strings.Count(result, "7") == 3 {
		fmt.Println("Congratulations!")
	} else {
		fmt.Println("Try again! Better luck next time.")
	}
}
