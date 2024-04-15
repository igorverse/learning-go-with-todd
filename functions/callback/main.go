package main

import "fmt"

func main() {
	x := doMath(40, 2, add)
	fmt.Println(x)

	y := doMath(44, 2, substract)
	fmt.Println(y)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func substract(a int, b int) int {
	return a - b
}
