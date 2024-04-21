package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I'm walking")
}

type yougin interface {
	walk()
	run()
}

func youngRun(y yougin) {
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()
	// youngRun(d1)

	d2 := &dog{"Padget"}
	d2.walk()
	d2.run()
	youngRun(d2)
}

// Quando um valor é do tipo pointeiro, é possível
// usar pointer receivers ou value receivers
// para implementar uma interface
