package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for key, value := range m {
		fmt.Printf("%s : %d\n", key, value)
	}
}
