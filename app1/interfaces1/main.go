package main

import (
	"fmt"
)

type Sender interface {
	Send(message string)
}

type Logger interface {
	Log(message string)
}

type NotificationService interface {
	Sender
	Logger
}

type EmailService struct{}

func (email EmailService) Send(message string) {
	fmt.Println("Email Sent Successfully:", message)

}

func (email EmailService) Log(message string) {
	fmt.Println("Logging Completed", message)
}

func notify(service NotificationService) {
	service.Log("Sending E-Mail Notification")
	service.Send("Your Order Placed Successfully!")
}
func main() {
	email := EmailService{}
	notify(email)
}
