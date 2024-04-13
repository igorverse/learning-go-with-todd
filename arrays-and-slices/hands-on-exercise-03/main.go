package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slice2 := slice[:5]
	slice3 := slice[5:]
	slice4 := slice[2:7]
	slice5 := slice[1:6]

	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
}
