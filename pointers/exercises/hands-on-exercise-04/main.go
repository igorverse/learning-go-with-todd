package main

import "fmt"

type person struct {
	first string
}

func changeName(p person, s string) person {
	p.first = s

	return p
}

func changeNamePtr(p *person, s string) {
	p.first = s
}

func main() {
	p := person{first: "Foo"}
	fmt.Println(p)
	p = changeName(p, "Bar")
	fmt.Println(p)

	changeNamePtr(&p, "Baz")
	fmt.Println(p)
}
