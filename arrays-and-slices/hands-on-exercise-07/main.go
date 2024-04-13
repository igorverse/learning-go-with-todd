package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	jm := []string{"Miss", "Moneypenny", "I'm 008"}

	xp := [][]string{jb, jm}

	fmt.Printf("%T\n", xp)

	for i, v := range xp {
		fmt.Println(i, v)
	}
}
