package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)

	fmt.Printf("O valor de x é %v\t", x)

	switch {
	case x <= 100:
		fmt.Println("Entre 0 e 100")
	case x >= 101 && x <= 200:
		fmt.Println("Entre 101 e 200")
	case x >= 201 && x <= 250:
		fmt.Println("Entre 201 e 250")
	default:
		fmt.Println("Comportamento não esperado!")
	}
}
