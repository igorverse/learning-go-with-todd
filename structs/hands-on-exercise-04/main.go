package main

import "fmt"

func main() {
	anonymous := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "Lucas",
		friends:   map[string]int{"Igor": 25, "Daniel": 15},
		favDrinks: []string{"Guaraná", "Água"},
	}

	fmt.Println("Amigos do", anonymous.first)
	for k, v := range anonymous.friends {
		fmt.Printf("%v : %v\n", k, v)
	}
}
