package main

import "fmt"
/*
Code is open for extension but closed for modification.
*/

// interface can hold all type of payment
 type IPaymentProcessor interface{
	payment()
 }
// payment methods
// Bkash
 type BkashPaymentProcessor struct{}

 func (bpp *BkashPaymentProcessor) payment(){
 	fmt.Println("Bkash payment")
 }
//Nagod
 type NagodPaymentProcessor struct{}

 func (npp *NagodPaymentProcessor) payment(){
 	fmt.Println("Nogod payment")
 }

 //Bank
 type BankPaymentProcessor struct{}
 func (bpp *BankPaymentProcessor) payment(){
	fmt.Println("Bank payment")
 }

 func main(){
	bkashPaymentProcessor:=&BkashPaymentProcessor{}
	nagodPaymentProcessor:=&NagodPaymentProcessor{}
	bankPaymentProcessor:=&BankPaymentProcessor{}

	bkashPaymentProcessor.payment()

	nagodPaymentProcessor.payment()

	bankPaymentProcessor.payment()
 }