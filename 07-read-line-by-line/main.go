package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./data/07-line-by-line.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fileReadErr := scanner.Err()

	if fileReadErr != nil {
		fmt.Println("There was an error reading from yout file")
		panic(fileReadErr)
	}
}
