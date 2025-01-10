package main

import "fmt"

// Notifier interface
type Notifier interface {
	Send(message string)
}

// EmailNotifier
type EmailNotifier struct{}

func (en EmailNotifier) Send(message string) {
	fmt.Printf("Email: %s\n", message)
}

// SMSNotifier
type SMSNotifier struct{}

func (sn SMSNotifier) Send(message string) {
	fmt.Printf("SMS: %s\n", message)
}

// NotificationService depends on Notifier abstraction
type NotificationService struct {
	Notifier Notifier
}

func (ns NotificationService) Notify(message string) {
	ns.Notifier.Send(message)
}

func main() {
	// Example
	emailService := NotificationService{Notifier: EmailNotifier{}}
	smsService := NotificationService{Notifier: SMSNotifier{}}

	emailService.Notify("Welcome, Alice!")
	smsService.Notify("Your OTP is 123456.")
}
