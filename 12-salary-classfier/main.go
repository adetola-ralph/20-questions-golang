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
	intSalary, _ := strconv.Atoi(salary)

	if intSalary < 0 {
		panic(errors.New("Input must be a positive number"))
	}

	switch {
	case intSalary < 50000:
		fmt.Println("Basic Earner")
	case intSalary < 200000:
		fmt.Println("Mid Earner")
	default:
		fmt.Println("High Earner")
	}
}
