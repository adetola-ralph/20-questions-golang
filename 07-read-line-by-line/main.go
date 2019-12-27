package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Open("./data/07-line-by-line.txt")
	fileToWrite, fileToWriteErr := os.Create("./data/07-line-by-line-write.txt")

	defer file.Close()
	defer fileToWrite.Close()

	if err != nil || fileToWriteErr != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(fileToWrite)

	for scanner.Scan() {
		text := scanner.Text()
		// fmt.Println(text)
		byteWritten, writeErr := writer.WriteString(text + "\n")

		if writeErr != nil {
			fmt.Printf("Error writing %s into the file %s \n", text, fileToWrite.Name())
			panic(writeErr)
		}

		fmt.Printf("wrote %s (%d bytes) into file %s \n", text, byteWritten, fileToWrite.Name())
		writer.Flush()
		time.Sleep(1 * time.Second)
	}

	fileReadErr := scanner.Err()

	if fileReadErr != nil {
		fmt.Println("There was an error reading from yout file")
		panic(fileReadErr)
	}
}
