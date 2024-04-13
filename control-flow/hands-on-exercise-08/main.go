package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		rand := rand.Intn(10)

		if rand == 4 {
			fmt.Println("You're lucky")
			break
		}

		fmt.Println("You're not lucky")

	}
}
