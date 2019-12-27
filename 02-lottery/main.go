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
	numberstring := make([]string, 3)
	numbers := make([]int, 3)

	// numbers := [3]string{}
	rand.Seed(time.Now().UnixNano())

	// fmt.Println("Please press any key to continue")
	// fmt.Scanln(&a)
	// fmt.Println(len(numbers), "length")
	// fmt.Println(cap(numbers), "capacity")

	for i < 3 {
		// numbers = append(numbers, strconv.Itoa(rand.Intn(10)))
		y := rand.Intn(10)
		numbers[i] = y
		numberstring[i] = strconv.Itoa(y)
		i++
	}

	result := strings.Join(numberstring, " ")
	fmt.Println(result)

	// if strings.Contains(result, "7") {
	// if strings.Count(result, "7") == 2 {
	// if strings.Count(result, "7") == 3 {
	// if allEven(numbers) {
	if allOdd(numbers) {
		fmt.Println("Congratulations!")
	} else {
		fmt.Println("Try again! Better luck next time.")
	}
}

func allEven(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			return false
		}
	}

	return true
}

func allOdd(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			return false
		}
	}

	return true
}
