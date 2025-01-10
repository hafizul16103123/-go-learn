package main

import "fmt"

// Target
type mobile interface{
	appleCharger()
}
// Concrete implementation
type apple struct{}
func (a apple) appleCharger(){
 fmt.Println("Charging apple mobile")
}
// Adaptee
type android struct{}
func (a *android) androidCharger(){
 fmt.Println("Charging android mobile")
}

// Adapter
type appleToAndroidAdapter struct{
	android *android
}
func (ad *appleToAndroidAdapter) appleCharger(){
	ad.android.androidCharger()
}


// Client
type client struct{}
func (c *client) chargeMobile(mobile mobile){
	mobile.appleCharger()
}

func main(){ 
	// initial requirements
	client:=&client{}
	apple:=&apple{}
	client.chargeMobile(apple)
	
	// Extended requirements
	android:=&android{}
	appleToAndroidAdapter:=&appleToAndroidAdapter{android}
	client.chargeMobile(appleToAndroidAdapter)
}
