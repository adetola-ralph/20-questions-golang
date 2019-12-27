package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file1, file1Err := os.Open("./data/08-file-1.txt")

	if file1Err != nil {
		fmt.Println("Error opening file 2")
		panic(file1Err)
	}

	file2, file2Err := os.Open("./data/08-file-2.txt")

	if file2Err != nil {
		fmt.Println("Error opening file 2")
		panic(file2Err)
	}

	defer file1.Close()
	defer file2.Close()

	file1Scanner := bufio.NewScanner(file1)
	file2Scanner := bufio.NewScanner(file2)

	file1Scanner.Split(bufio.ScanLines)
	file2Scanner.Split(bufio.ScanLines)

	fileToWrite, fileToWriteErr := os.Create("./data/08-file-write.txt")

	if fileToWriteErr != nil {
		fmt.Println("Error creating file to write")
		panic(fileToWriteErr)
	}

	writer := bufio.NewWriter(fileToWrite)

	for file1Scanner.Scan() && file2Scanner.Scan() {
		newText := fmt.Sprintf("%s %s \n", file1Scanner.Text(), file2Scanner.Text())
		byteWritten, writeErr := writer.WriteString(newText)

		if writeErr != nil {
			fmt.Printf("Error writing %s into the file %s \n", newText, fileToWrite.Name())
			panic(writeErr)
		}

		fmt.Printf("wrote %s (%d bytes) into file %s \n", newText, byteWritten, fileToWrite.Name())
		writer.Flush()
	}

	file1ScannerErr := file1Scanner.Err()
	file2ScannerErr := file1Scanner.Err()

	if file1ScannerErr != nil {
		fmt.Println("There was an error reading from file 1")
		panic(file1ScannerErr)
	}

	if file2ScannerErr != nil {
		fmt.Println("There was an error reading from file 2")
		panic(file2ScannerErr)
	}
}
