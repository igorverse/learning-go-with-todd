package main

import "fmt"

func addOne(v int) int {
	v = v + 1
	return v
}

func main() {
	i := 1
	fmt.Println(addOne(i))

	fmt.Println(i)
}
