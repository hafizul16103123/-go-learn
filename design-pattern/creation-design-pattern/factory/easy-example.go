package main

import "fmt"

type SudanCar struct{ // App code
	name string
}

func main(){  // Client
	car:=getName() //client does not know anything about SudanCar,only know factory.

	fmt.Println(car)
}

func getName() *SudanCar{  //Factory
return &SudanCar{name:"Honda City"}
}