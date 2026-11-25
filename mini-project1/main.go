package main

import (
	"fmt"
	"math/rand/v2"
)


func main() {
	secret := rand.IntN(100) + 1
	var guess int
	var diff int

	fmt.Println("=== Number Guessing Game ===")
	fmt.Print("Guess a number between 1 and 100: ")
	fmt.Scanln(&guess)

	if guess < 1 || guess > secret {
		fmt.Println("Invalid guess! Pick a number between 1 and 100.")
	} else {
		if guess == secret {
			fmt.Println("Spot on! You guessed the secret number!")
		} else if guess < secret {
			fmt.Printf("Too low! The secret number was %d.\n", secret)
		} else {
			fmt.Printf("Too high! The secret number was %d.\n", secret)
		}

		if guess != secret {
			if secret > guess {
				diff = secret - guess
			} else {
				diff = guess - secret
			}

			if diff < 0 {
				diff = -diff
			}
			fmt.Printf("You were off by %d!", diff)
		}
	}
}
