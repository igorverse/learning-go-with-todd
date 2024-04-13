package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := [5]int{}

	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(9)
	}

	for _, n := range array {
		fmt.Println(n)
	}
}
