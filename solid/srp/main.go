package main

import (
	"fmt"
)

/**
Single Responsibility Principle (SRP)
Each struct or function should have one responsibility.

*/

type User struct{
	name string
	email string
	age int
}

type EmailSender struct{}

func (es EmailSender) sendEmail(){
	fmt.Println("Email sent ....")
}


func main(){
	user:=User{"Hafiz","hafiz@gmail.com",25}
	emailSender:=&EmailSender{}
	emailSender.sendEmail()
	fmt.Println(user)
}
