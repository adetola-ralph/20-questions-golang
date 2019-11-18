package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please enter a sentence")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')

	sentence = strings.Replace(sentence, "\n", "", -1)

	// reverse the whole sentence
	reversearray := make([]string, len(sentence))
	sentencelength := len(sentence) - 1
	for i := sentencelength; i >= 0; i-- {
		reversearray[sentencelength-i] = string(sentence[i])
	}
	fmt.Println(strings.Join(reversearray, ""))

	// reverse the words in the sentence
	// reversearray := []string{}
	// sentencearray := strings.Split(sentence, " ")
	// for i := 0; i < len(sentencearray); i++ {
	// 	word := sentencearray[i]
	// 	reversewordarray := make([]string, len(word))

	// 	for j := len(word) - 1; j >= 0; j-- {
	// 		reversewordarray[len(word)-1-j] = string(word[j])
	// 	}

	// 	reverseword := strings.Join(reversewordarray, "")
	// 	reversearray = append(reversearray, reverseword)
	// }
	// fmt.Println(strings.Join(reversearray, " "))
}
