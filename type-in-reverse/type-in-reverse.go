package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tjarratt/babble"
)

// not necessary but doing for the sake of doing it
type myString string

// also this
func (value myString) reverseWord() string {
	reversearray := make([]string, len(value))
	valuelength := len(value) - 1
	for i := valuelength; i >= 0; i-- {
		reversearray[valuelength-i] = string(value[i])
	}

	return strings.Join(reversearray, "")
}

func main() {
	babbler := babble.NewBabbler()

	// and this
	wordToReverse := myString(babbler.Babble())

	fmt.Printf("Type %s in reverse: ", wordToReverse)
	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')

	word = strings.Replace(word, "\n", "", -1)

	// yes here also
	if word != wordToReverse.reverseWord() {
		fmt.Println("❌")
	} else {
		fmt.Println("✅")
	}
}
