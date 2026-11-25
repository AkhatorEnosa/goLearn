// Repeating Actions (for loops)

package main

import (
	"fmt"
	"math/rand/v2"
)


func main() {
	secret := rand.IntN(100) + 1
	won := false
	maxAttempts := 10
	var guess int

	for attempt := 1; attempt <= maxAttempts; {
		remainingAttempts := maxAttempts - attempt
		fmt.Println("=== Number Guessing Game ===")
		fmt.Print("Guess a number between 1 and 100: ")
		fmt.Scanln(&guess)

		if guess < 1 || guess > 100 {
			fmt.Println("Invalid guess! Pick a number between 1 and 100.")
			continue
		} else {
			if guess == secret {
				fmt.Printf("Spot on! You guessed the secret number! It took you %d attempts.\n", attempt)
				won = true
				break
			} else if guess < secret {
				fmt.Printf("Too low! Total attempts remaining is %d.\n", remainingAttempts)
			} else {
				fmt.Printf("Too high! Total attempts remaining is %d.\n", remainingAttempts)
			}
			attempt++
		}
	}

	if !won {
		fmt.Printf("Out of attempts! The secret number was %d.\n", secret)
	} 

}
