package main

import "fmt"

// mobile interface to accept any mobile type
type IMobile interface{
	appleCharger()
}

// apple charger
type Apple struct{}
func (amc Apple) appleCharger()  {
	fmt.Println("APPLE mobile charging")
}

// android adaptee
type Android struct{}
func (amc Android) androidCharger()  {
	fmt.Println("ANDROID mobile charging")
}
//android Adapter
type AndroidAdapter struct{
	android *Android
}
func (ad AndroidAdapter) appleCharger()  {
	ad.android.androidCharger()
}

// laptop adaptee

type Laptop struct{}
func (lch Laptop) laptopCharger()  {
	fmt.Println("Laptop mobile charging")
}
// laptop Adapter
type LaptopAdapter struct{
	laptop *Laptop
}
func (la LaptopAdapter) appleCharger(){
	la.laptop.laptopCharger()
}




// Client
type Client struct{}
func (c Client) charge(mob IMobile)  {
	mob.appleCharger()
}



func main(){
	client:=&Client{}

	apple:=&Apple{}

	android:=&Android{}
	androidAdapter:= &AndroidAdapter{android}

	laptop:=&Laptop{}
	laptopAdapter:=&LaptopAdapter{laptop}
	
	client.charge(apple)
	client.charge(androidAdapter)
	client.charge(laptopAdapter)
}