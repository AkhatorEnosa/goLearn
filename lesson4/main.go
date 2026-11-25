// Making Decisions (if / else / switch)

package main

import "fmt"


func main() {
	var billAmount float64
	var tipPercent float64
	var people int

	fmt.Print("Enter Bill Amount: $")
	fmt.Scanln(&billAmount)

	fmt.Print("Enter Tip Percentage: %")
	fmt.Scanln(&tipPercent)

	fmt.Print("Enter no. of People: ")
	fmt.Scanln(&people)

	if people <= 0 {
		fmt.Println("Number of people must be at least 1.")
 	} else if billAmount < 0 {
		fmt.Println("Bill amount cannot be negative.")
	} else {
		switch {
			case tipPercent == 15.00:
				fmt.Println("Standard tip.")
			case tipPercent >= 20.00:
				fmt.Println("Generous tipper!")
			default:
				fmt.Println("Custom tip applied.")
		}

		tipAmount := (billAmount * (tipPercent / 100))
		totalBill := (billAmount + tipAmount)
		perPerson := (totalBill / float64(people))

		fmt.Printf("Your total bill is $%.2f, including a $%.2f tip.\n Each person (%d) to pay: $%.2f", totalBill, tipAmount, people, perPerson)
	
	}

}
