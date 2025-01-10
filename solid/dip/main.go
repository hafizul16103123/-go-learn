package main

import "fmt"

/*
Dependency Inversion Principle (DIP)
High-level modules should depend on abstractions, not concrete implementations.
*/

// interface can hold all type of payment (abstraction)
 type IPaymentProcessor interface{
	payment()
 }

 
// payment methods (Concrete implementation)
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



 // payment process (High level module) depends on abstraction not concrete implementation
type PaymentProcessor struct{
	paymentProcessor IPaymentProcessor
 }
 func (bp *PaymentProcessor) process(){
	bp.paymentProcessor.payment()
}

 func main(){
	bkashPaymentProcessor:=&BkashPaymentProcessor{}
	nagodPaymentProcessor:=&NagodPaymentProcessor{}
	bankPaymentProcessor:=&BankPaymentProcessor{}

	paymentProcessor:=&PaymentProcessor{bkashPaymentProcessor}
	paymentProcessor.process()

	paymentProcessor2:=&PaymentProcessor{nagodPaymentProcessor}
	paymentProcessor2.process()

	paymentProcessor3:=&PaymentProcessor{bankPaymentProcessor}
	paymentProcessor3.process()
 }