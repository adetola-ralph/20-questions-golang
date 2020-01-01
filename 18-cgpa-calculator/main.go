package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("oreofe")

	csvFile, err := os.Open("./data/18-student-data.csv")
	defer csvFile.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(csvFile)
	scanner.Split(bufio.ScanLines)

	for i := 0; scanner.Scan(); i++ {
		if i != 0 {
			text := scanner.Text()
			seperatedStrings := strings.Split(text, ",")
			id := seperatedStrings[0]
			name := seperatedStrings[1]
			sum := 0

			for _, v := range seperatedStrings[2:] {
				// sum += strconv.Atoi(v)
				value, _ := strconv.Atoi(v)
				sum += value
			}

			average := float32(sum) / 5.0
			cgpa := (float32(sum) / 20.0) / 5.0

			fmt.Printf("%s (%s): %.1f %.1f\n", name, id, average, cgpa)
		}
	}
}
