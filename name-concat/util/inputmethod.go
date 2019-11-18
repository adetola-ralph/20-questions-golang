package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// function used to enapsulate getting the user's information
func GetUserInputFmt() (string, string, int) {
	var firstName, lastName string
	var age int

	fmt.Println("Please enter your firstname")
	fmt.Scanln(&firstName)

	fmt.Println("Please enter your lastname")
	fmt.Scanln(&lastName)

	fmt.Println("Please enter your age")
	fmt.Scanln(&age)

	return firstName, lastName, age
}

func GetUserInputBufio() (string, string, int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your firstname")
	firstname, _ := reader.ReadString('\n')
	firstname = strings.Replace(firstname, "\n", "", -1)

	fmt.Println("Please enter your lastname")
	lastname, _ := reader.ReadString('\n')
	lastname = strings.Replace(lastname, "\n", "", -1)

	fmt.Println("Please enter your age")
	ageinstring, _ := reader.ReadString('\n')
	ageinstring = strings.Replace(ageinstring, "\n", "", -1)
	age, err := strconv.Atoi(ageinstring)

	if err != nil {
		panic(err)
	}

	return firstname, lastname, age
}
