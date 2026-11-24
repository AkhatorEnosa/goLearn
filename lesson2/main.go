package main

import "fmt"

func main() {
	username := "osa"
	level := 1
	accountBalance := 100.50
	isOnline := false

	level = level + 1

	fmt.Printf("Hello, %s!\n Your current level is %d and your account balance is $%.2f.\n Online status: %t\n", username, level, accountBalance, isOnline)
}
