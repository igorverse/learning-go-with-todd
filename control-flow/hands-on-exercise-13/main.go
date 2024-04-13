package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	age := m["James"]
	fmt.Println(age)

	if value, ok := m["Q"]; !ok {
		fmt.Println("There is no Q! The value found was: ", value)
	}
}
