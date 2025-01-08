package main

import "fmt"

func main() {
	var arr [5]int
	var arr2 = [5]int{1, 2, 3, 4, 5}
	arr3 := [5]int{1, 2, 3, 4, 5}

	printArr(arr)
	printArr(arr2)
	printArr(arr3)
}

func printArr(arr [5]int) {
	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Println("Break")
}
