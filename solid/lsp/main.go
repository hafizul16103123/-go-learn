package main

import "fmt"

/*
Liskov Substitution Principle (LSP)
Subtypes should be substitutable for their base types.
*/
// Bird interface
type Bird interface {
	Fly()
}

// Sparrow satisfies Bird
type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow is flying!")
}

// Crow satisfies Bird
type Crow struct{}

func (c Crow) Fly() {
	fmt.Println("Crow is flying!")
}

func main() {
	// Example
	var birds []Bird = []Bird{
		Sparrow{},
		Crow{},
	}

	for _, bird := range birds {
		bird.Fly() // Works for all Bird types
	}
}
