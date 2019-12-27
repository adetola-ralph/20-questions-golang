package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your salart: ")
	salary, _ := reader.ReadString('\n')
	salary = strings.Replace(salary, "\n", "", -1)
	float64Salary, _ := strconv.ParseFloat(salary, 32)
	floatSalary := float32(float64Salary)

	if floatSalary < 0 {
		panic(errors.New("Input must be a positive number"))
	}

	switch {
	case floatSalary < 50000:
		fmt.Printf("Tax %.2f\n", (5.0/100.0)*floatSalary)
	case floatSalary < 200000:
		fmt.Printf("Tax: %.2f\n", (10.0/100.0)*floatSalary)
	default:
		fmt.Printf("Tax: %.2f\n", (15.0/100.0)*floatSalary)
	}
}
