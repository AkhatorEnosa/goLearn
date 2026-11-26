// Lesson 9: Structs & Pointers

package main

import "fmt"

type BankAccount struct{
	Owner string
	Balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		a.Balance += amount
	}
}


func (a *BankAccount) Withdraw(amount float64) bool {
	if amount > 0 && amount <= a.Balance {
		a.Balance -= amount
		return true
	} else {
		fmt.Printf("Insufficient funds. Current balance: $%.2f\n", a.Balance)
		return false
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

	account.Withdraw(130)

	fmt.Println(account.Info())
}