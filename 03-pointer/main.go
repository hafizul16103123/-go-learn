package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("value of pointer", ptr)

	myNumber := 100
	ptr2 := &myNumber
	fmt.Println("value of pointer", ptr2)
	fmt.Println("value of pointer", *ptr2)
}
