package main

import "fmt"

func main() {

	mySlice := []int{0,1, 2, 3, 4}
	printArr(mySlice[1:3])
	printArr(mySlice[1:])
	printArr(mySlice[:])

}
func printArr(arr []int) {

	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Println("Break")
}
