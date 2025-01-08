package main

import (
	"fmt"
)

func main() {
	// constant variable
	const PI = 3.1416
	
	// var with explicit type
	var num1 int

	// var with type inferred
	var num2 = 2

	// short syntax declaration with type inferred
	num3 := 2

	// multiple variable at a time
	var a, b, c int
	var d, e, f = 1, 2, 3
	g, h := 5, 6

	// common var with explicit type
	var (
		i, j, k int
		l       string
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l,num1,num2,num3,PI)
	// initialize value in side if block
	if i := 10; i > 2 {
		fmt.Println("I is greater than 2")
	} else {
		fmt.Println("I is less than 2")
	}
	// Initial value in switch case
	switch x := 1; x {
	case 1:
		fmt.Println("x is 1")
	}
}
