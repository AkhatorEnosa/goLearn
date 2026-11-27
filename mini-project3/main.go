// Personal Expense Tracker (with File Saving)

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Expense struct {
	Category string
	Amount   float64
}

func (e Expense) Line() string {
	return fmt.Sprintf("%s,%.2f\n", e.Category, e.Amount)
}

func addExpense(expenses []Expense, category string, amount float64) ([]Expense, error) {
	if amount <= 0 {
		return nil, errors.New("Amount should be more than 0")
	}
	expenses = append(expenses, Expense{Category: category, Amount: amount})
	return expenses, nil
}

func summarizeByCategory(expenses []Expense) (map[string]float64, float64) {
	totals := map[string]float64{}
	var grandTotal float64
	for _, exp := range expenses {
		totals[exp.Category] += exp.Amount
		grandTotal += exp.Amount
	}
	return totals, grandTotal
}

func saveExpenses(filename string, expenses []Expense) error {
	var sb strings.Builder
	for _, exp := range expenses {
		sb.WriteString(exp.Line())
	}
	return os.WriteFile(filename, []byte(sb.String()), 0644)
}

func main() {
	expenses := []Expense{}
	var choice int

	for {
		fmt.Println("\n=== Expense Tracker ===")
		fmt.Println("1. Add expense")
		fmt.Println("2. View summary by category")
		fmt.Println("3. Save to expenses.txt")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)

		var category string
		var amount float64

		switch choice {
			case 1:
				// Add expense
				fmt.Print("Add a category: ")
				fmt.Scanln(&category)
				fmt.Print("Add amount: ")
				fmt.Scanln(&amount)

				result, err := addExpense(expenses, category, amount)

				if err == nil {
					fmt.Println("Expense Added")
					expenses = result
				} else {
					fmt.Println("Error:", err)
				}
			case 2:
				// View summary by category

				if len(expenses) <= 0 {
					fmt.Println("No expenses recorded yet.")
					return
				} else {
					totals, grandTotals := summarizeByCategory(expenses)
					for key, value := range totals {
						fmt.Printf("%s | %.2f\n", key, value)
					}

					fmt.Printf("Grand Total: %.2f\n", grandTotals)
				}

			case 3:
				// Save to expenses.txt using saveExpenses("expenses.txt", expenses)
				err := saveExpenses("expenses.txt", expenses)	
				if err != nil {
					fmt.Println("Error saving expenses:", err)
				} else {
					fmt.Println("Expenses saved to expenses.txt")
				}
			case 4:
				fmt.Println("Goodbye!")
				return
			default:
				fmt.Println("Invalid option. Choose 1-4.")
		}
	}
}