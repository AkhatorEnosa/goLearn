package main

import "fmt"


func main() {
	var billAmount float64
	var tipPercent float64
	var people int

	fmt.Print("Enter Bill Amount: $")
	fmt.Scanln(&billAmount)

	fmt.Print("Enter Tip Percentage: $")
	fmt.Scanln(&tipPercent)

	fmt.Print("Enter no. of People: ")
	fmt.Scanln(&people)

	tipAmount := (billAmount * (tipPercent / 100))
	totalBill := (billAmount + tipAmount)
	perPerson := (totalBill / float64(people))

	fmt.Printf("Your total bill is $%.2f, including a $%.2f tip.\n Each person to pay: $%.2f", totalBill, tipAmount, perPerson)
}
