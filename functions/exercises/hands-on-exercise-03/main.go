package main

import "fmt"

// LIFO
func main() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)

	fmt.Println("Contagem: ")
}
