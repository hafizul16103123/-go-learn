package main

import "fmt"

type Car interface{
	getCar() string
}
// cat type -1
type SudanCar struct{ // App code
	name string
}
func (sc *SudanCar) getCar() string{
	return "Honda City"
}

func NewSudanCar() *SudanCar{
	return &SudanCar{}
}

//car type -2
type HatchBackCar struct{ // App code
	name string
}
func (hbc *HatchBackCar) getCar() string{
	return "Polo"
}
func NewHatchBackCar() *HatchBackCar{
	return &HatchBackCar{}
}

func main() { // Client
	getCarFactory(1)
	getCarFactory(2)
}
//Factory 
func getCarFactory(cartype int) { // Factory Class/Object
	var car Car
	if cartype == 1 {
		car = NewHatchBackCar()
	} else {
		car = NewSudanCar()
	}
	carInfo := car.getCar()
	fmt.Println(carInfo)
}