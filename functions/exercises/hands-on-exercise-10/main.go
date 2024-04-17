package main

import (
	"fmt"
	"math"
)

func powinator(n float64) func() float64 {
	c := 0.0
	return func() float64 {
		c++
		return math.Pow(n, c)
	}
}

func main() {
	x := powinator(2)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}
