package main

import (
	"fmt"
)

func square(n int) int {
	return n * n
}

func printSquare(f func(int) int, r int) string {
	return fmt.Sprintf("O quadrado de %d é %d", r, f(r))
}

func main() {
	fmt.Println(printSquare(square, 2))
}
