package main

import "fmt"

func main() {
	numbers := []int{42, 43, 44, 45, 46, 47}

	for i, n := range numbers {
		fmt.Printf("%d: %d\n", i, n)
	}
}
