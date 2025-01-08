package main

import (
	"fmt"
)

func main() {
	// // Map using make function
	// var myMap = make(map[string]int)
	// myMap["Hafiz"] = 29
	// myMap["Nusayeb"] = 3
	// for _, v := range myMap {
	// 	fmt.Println(v)
	// }

	// // Empty Map using map literal
	// var myMap2 = map[string]int{}
	// myMap2["Hafiz"] = 29
	// myMap2["Nusayeb"] = 3
	// for _, v := range myMap2 {
	// 	fmt.Println(v)
	// }

	// // Initialize Map using map literal having default value
	// var myMap3 = map[string]int{
	// 	"Hafiz":   29,
	// 	"Nusayeb": 3,
	// }

	// for key, v := range myMap3 {
	// 	fmt.Println(key, ">>>>>>>>>", v)
	// }

	// // storing to a nil map causes a panic
	// var ages map[string]int
	// fmt.Println(ages == nil)
	// // "true"
	// fmt.Println(len(ages) == 0) // "true"
	// ages["Hafiz"] = 2

	ages := map[string]int{}
	age, ok := ages["bob"]
	fmt.Println(age, !ok)
	if !ok {
		fmt.Println("value not exist")
	}

	// isEqual := equal(map[string]int{"A": 0}, map[string]int{"A": 0})
	// fmt.Println(isEqual)

}

// func equal(x, y map[string]int) bool {
// 	if len(x) != len(y) {
// 		return false
// 	}
// 	for k, xv := range x {
// 		fmt.Println(xv)
// 		if yv, ok := y[k]; !ok || yv != xv {
// 			return false
// 		}
// 	}
// 	return true
// }
