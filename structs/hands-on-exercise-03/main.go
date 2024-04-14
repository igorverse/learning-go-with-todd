package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	door  int
	color string
}

func main() {
	v1 := vehicle{engine: engine{electric: false}, make: "hyunday", model: "i30", door: 4, color: "preto"}
	v2 := vehicle{engine: engine{electric: true}, make: "tesla", model: "S", door: 4, color: "branco"}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Printf("%v: %v\n", v1.model, v1.electric)
	fmt.Printf("%v: %v\n", v2.model, v2.electric)
}
