package main

import "fmt"

type person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := person{
		"Igor",
		"Silva",
		[]string{"Vanilla", "Strawberry"},
	}

	p2 := person{
		"Maria Eduarda",
		"Pereira",
		[]string{"Pistache", "Cocoa"},
	}

	fmt.Printf("%v %v - Sabores de sorvete favoritos: \n", p1.firstName, p1.lastName)
	for _, v := range p1.favoriteIceCreamFlavors {
		fmt.Printf("%v\t", v)
	}

	fmt.Printf("\n%v %v - Sabores de sorvete favoritos: \n", p2.firstName, p2.lastName)
	for _, v := range p2.favoriteIceCreamFlavors {
		fmt.Printf("%v\t", v)
	}
}
