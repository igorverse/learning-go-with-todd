package main

import "fmt"

type Person struct {
	first string
	age   int
}

func (p Person) speak() {
	fmt.Printf("%v: %v\n", p.first, p.age)
}

func main() {
	p := Person{"Igor", 25}

	p.speak()
}
