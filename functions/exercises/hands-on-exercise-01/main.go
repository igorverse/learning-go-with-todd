package main

import "fmt"

func foo() int {
	return 42
}

func bar() (int, string) {
	return foo(), "Quarenta e dois"
}

func main() {
	foo := foo()
	fmt.Println(foo)

	barInt, barStr := bar()
	fmt.Println(barInt)
	fmt.Println(barStr)
}
