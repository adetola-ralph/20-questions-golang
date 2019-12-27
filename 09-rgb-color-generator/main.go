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
		red := randGenerator.Intn(256)
		green := randGenerator.Intn(256)
		blue := randGenerator.Intn(256)

		fmt.Printf("rgb(%d, %d, %d)\n", red, green, blue)
		time.Sleep(time.Second)
	}
}
