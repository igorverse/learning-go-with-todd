package main

import "fmt"

func addOne(v *int) int {
	*v += 1

	return *v
}

func main() {
	a := 1
	fmt.Println("a: ", a)
	fmt.Println("addOne(&a): ", addOne(&a))
	fmt.Println("a: ", a)
}
