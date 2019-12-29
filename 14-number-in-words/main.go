package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ones      = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	tenTwenty = [10]string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens      = [10]string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eigthy", "ninety"}
	suffixes  = [4]string{"hundred", "thousand", "million", "billion"}
)

func main() {
	fmt.Println("Enter a number")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	value = strings.Replace(value, "\n", "", -1)
	valueInt, _ := strconv.Atoi(value)

	if valueInt < 0 {
		panic(errors.New("Please enter a positive number"))
	}

	numberInWords := ""

	// 1,234,567
	// 1,000,000 - one million
	// 0,200,000 - two hundred thousand
	// 0,030,000 - thirty thousand
	// 0,004,000 - four thousand
	// 0,000,500 - five hundred
	// 0,000,060 - sixty
	// 0,000,007 - seven

	// one million two hundred and thrity four thousand, five hundred and sixty seven
	// break each section into hundreds and process them
	// TODO: I thhink this needs o be refactored some more, will get to it later

	if valueInt >= 1000000000 {
		numberInWords = fmt.Sprintf("%s billion", processHundred(valueInt/1000000000))
		valueInt = valueInt % 1000000000
	}

	if valueInt >= 1000000 {
		numberInWords = fmt.Sprintf("%s %s million", numberInWords, processHundred(valueInt/1000000))
		valueInt = valueInt % 1000000
	}

	if valueInt >= 1000 {
		numberInWords = fmt.Sprintf("%s %s thousand", numberInWords, processHundred(valueInt/1000))
		valueInt = valueInt % 1000
	}

	numberInWords = fmt.Sprintf("%s %s", numberInWords, processHundred(valueInt))

	fmt.Printf("your number in words: %s\n", strings.TrimSpace((numberInWords)))
}

func processHundred(value int) string {
	returnstring := ""
	first := value / 100

	if first > 0 {
		returnstring = fmt.Sprintf("%s hundred", ones[first])
	}

	second := value % 100 / 10
	third := value % 100 % 10

	if second > 0 || third > 0 {
		if returnstring != "" {
			returnstring = fmt.Sprintf("%s and", returnstring)
		}

		if second >= 1 && second < 2 {
			return strings.TrimSpace(fmt.Sprintf("%s %s", returnstring, tenTwenty[third]))
		}

		if second > 0 {
			returnstring = fmt.Sprintf("%s %s", returnstring, tens[second])
		}

		if third > 0 {
			returnstring = fmt.Sprintf("%s %s", returnstring, ones[third])
		}
	}

	return strings.TrimSpace(returnstring)
}
