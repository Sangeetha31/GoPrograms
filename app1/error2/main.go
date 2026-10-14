package main

import (
	"errors"
	"fmt"
)

type InsuffiecientFundsError struct {
	Balance float32
	Amount  float32
}

func (e *InsuffiecientFundsError) Error() string {
	return fmt.Sprintf(
		"Insufficient Funds in Account: Balance is: %f, withdraw amount: %f",
		e.Balance,
		e.Amount,
	)
}

func withdraw(balance, amount float32) (float32, error) {
	if amount <= 0 {
		return balance, errors.New("Withdraw amount should be positive")
	}
	if amount > balance {
		return balance, &InsuffiecientFundsError{
			Balance: balance,
			Amount:  amount,
		}
	}
	return balance - amount, nil
}
func main() {
	balance := 5000

	availbal, err := withdraw(float32(balance), 6500)

	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Transaction Complete\nAvailable Balance", availbal)
}
