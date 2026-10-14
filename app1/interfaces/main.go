package main

import "fmt"

type paymentProcessor interface {
	Pay(amount float32) error
}
type CreditCard struct{}

func (c CreditCard) Pay(amount float32) error {
	fmt.Println("Payment Done Through CreditCard for Amount:", amount)
	return nil
}

type UPI struct{}

func (u UPI) Pay(amount float32) error {
	fmt.Println("Payment Done Through UPI for Amount:", amount)
	return nil
}

func checkout(payment paymentProcessor, amount float32) error {
	return payment.Pay(amount)
}

func main() {

	ccard := CreditCard{}
	_ = checkout(ccard, 5200)

	upi := UPI{}
	_ = checkout(upi, 6500)

}
