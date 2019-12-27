package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(source)

	for true {
		red := randGenerator.Intn(255)
		green := randGenerator.Intn(255)
		blue := randGenerator.Intn(255)

		fmt.Printf("rgb(%d, %d, %d)\n", red, green, blue)
		time.Sleep(time.Second)
	}
}
