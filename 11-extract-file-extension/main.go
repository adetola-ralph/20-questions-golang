package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("oreofe")
	files := []string{"/Users/mykeels/Documents/file.txt",
		"/Users/mykeels/Documents/file.csv",
		"/Users/mykeels/Music/a-whole-new-world.mp3",
		"/Users/mykeels/Movies/a-day-to-remember.mp4"}

	for _, file := range files {
		fmt.Printf("%s\n", strings.Split(file, ".")[1])
	}
}
