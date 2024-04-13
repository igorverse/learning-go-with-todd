package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, n := range slice {
		fmt.Printf("Value: %v, type: %T\n", n, n)
	}
}
