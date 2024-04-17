package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	y := &x
	fmt.Printf("%v\t%T\t%v\n", y, *y, *y)
	fmt.Println(*&x)

	*y = 4
	fmt.Println(x)
	fmt.Println(*y)
}
