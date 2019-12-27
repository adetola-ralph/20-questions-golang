package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(source)

	red := randGenerator.Intn(256)
	green := randGenerator.Intn(256)
	blue := randGenerator.Intn(256)

	fmt.Printf("%d, %d, %d\n", red, green, blue)
	fmt.Printf("#%s%s%s\n", padValue(toBase16(red)), padValue(toBase16(green)), padValue(toBase16(blue)))
}

func toBase16(value int) string {
	remainders := []string{}
	var divisor = value

	remainder := divisor % 16
	divisor = divisor / 16

	remainders = append(remainders, base16Values(remainder))

	for divisor != 0 {
		remainder = divisor % 16
		divisor = divisor / 16
		remainders = append(remainders, base16Values(remainder))
	}

	for i := 0; i < len(remainders)/2; i++ {
		j := len(remainders) - i - 1
		remainders[i], remainders[j] = remainders[j], remainders[i]
	}

	return strings.Join(remainders, "")
}

func base16Values(number int) string {
	value := strconv.Itoa(number)
	switch value {
	case "10":
		return "A"
	case "11":
		return "B"
	case "12":
		return "C"
	case "13":
		return "D"
	case "14":
		return "E"
	case "15":
		return "F"
	default:
		return value
	}
}

func padValue(value string) string {
	if len(value) < 2 {
		value = "0" + value
	}

	return value
}
