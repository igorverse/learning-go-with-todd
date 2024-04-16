package main

import "fmt"

func foo(numbers ...int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

func bar(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

func main() {
	x := []int{1, 1, 2, 3, 5, 8}
	s := foo(x...)
	s2 := bar(x)

	fmt.Println(s, s2)
}
