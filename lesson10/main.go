// Lesson 10: Error Handling

package main

import (
	"errors"
	"fmt"
)

type BankAccount struct{
	Owner string
	Balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		a.Balance += amount
	}
}


func (a *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Withdrawal amount must be greater tha zero.")	
	} else if amount > a.Balance {
		return fmt.Errorf("insufficient funds: balance is $%.2f", a.Balance)
	} else {
		a.Balance -= amount
		return nil
	}

}

func (a BankAccount) Info() string {
	return fmt.Sprintf("%s - Balance: $%.2f\n", a.Owner, a.Balance)
}

func main() {
	account := BankAccount{
		Owner: "Osa",
		Balance: 0,
	}

	account.Deposit(500)

	err := account.Withdraw(130)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err2 := account.Withdraw(530)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}

	fmt.Println(account.Info())
}