package main

import "fmt"

func main() {
	persons := map[string][]string{
		"bond_james":       {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}

	persons["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	for name, slice := range persons {
		fmt.Printf("%v\n", name)
		for i, value := range slice {
			fmt.Printf("%d: %v\t", i, value)
		}
		fmt.Println()
	}
}
