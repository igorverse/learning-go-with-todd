package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("\n%v: \n", i)
		for j := 4; j >= 0; j-- {
			fmt.Printf("%v\n", j)
		}
	}
}
