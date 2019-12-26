package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	userPrompt :=
		`1. Celsius to Fahrenheit
2. Fahrenheit to Celsius
3. Celsius to Kelvin
4. Kelvin to Celsius
5. Fahrenheit to Kelvin
6. Kelvin to Fahrenheit`

	fmt.Println(userPrompt)

	reader := bufio.NewReader(os.Stdin)
	userValue, _ := reader.ReadString('\n')
	userValue = strings.Replace(userValue, "\n", "", -1)

	switch userValue {
	case "1":
		fmt.Println("Enter a celsius value:")
		celsiusValue := getUserValue(reader)
		fmt.Println(celsiusToFahrenheit(float32(celsiusValue)))
	case "2":
		fmt.Println("Enter a fahrenheit value:")
		fahrenheitValue := getUserValue(reader)
		fmt.Println(fahrenheitToCelsius(float32(fahrenheitValue)))
	case "3":
		fmt.Println("Enter a celsius value:")
		celsiusValue := getUserValue(reader)
		fmt.Println(celsiusToKelvin(celsiusValue))
	case "4":
		fmt.Println("Enter a kelvin value:")
		kelvinValue := getUserValue(reader)
		fmt.Println(kelvinToCelsius(kelvinValue))
	case "5":
		fmt.Println("Enter a fahrenheit value:")
		fahrenheitValue := getUserValue(reader)
		fmt.Println(fahrenheitToKelvin(fahrenheitValue))
	case "6":
		fmt.Println("Enter a kelvin value:")
		kelvinValue := getUserValue(reader)
		fmt.Println(kelvinToFahrenheit(kelvinValue))
	default:
		fmt.Println("Please select a valid value")
	}
}

func getUserValue(reader *bufio.Reader) float32 {
	value, _ := reader.ReadString('\n')
	value = strings.Replace(value, "\n", "", -1)
	floatValue, _ := strconv.ParseFloat(value, 32)
	return float32(floatValue)
}

func celsiusToFahrenheit(value float32) float32 {
	return (value * (9.0 / 5.0)) + 32
}

func fahrenheitToCelsius(value float32) float32 {
	return (value - 32) * (5.0 / 9.0)
}

func celsiusToKelvin(value float32) float32 {
	return value + 273.15
}

func kelvinToCelsius(value float32) float32 {
	return value - 273.15
}

func kelvinToFahrenheit(value float32) float32 {
	return celsiusToFahrenheit(value - 273.15)
}

func fahrenheitToKelvin(value float32) float32 {
	return fahrenheitToCelsius(value-273.15) + 273.15
}
