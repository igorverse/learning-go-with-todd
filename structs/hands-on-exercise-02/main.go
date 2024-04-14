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

	lastName := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range lastName {
		fmt.Println(v)
		for _, v2 := range v.favoriteIceCreamFlavors {
			fmt.Println(v2)
		}
	}
}
