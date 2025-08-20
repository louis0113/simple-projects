package main

import (
	"fmt"
	"math/rand"
)

func welcomeMessage() {

	fmt.Printf("Welcome to the Number Guessing Game!\nI'm thinking about a number between 1 and 100\n\nTry guess that number!\n\n")
}

func OptionMessage(levelName string) {
	fmt.Printf("\nGreat! You have selected the %s difficulty level.\nLet's start the game!\n\n", levelName)
}

func main() {
	num := rand.Intn(99) + 1
	var message string
	var status bool
	welcomeMessage()

	var level, guesses, guess int

	for {

		fmt.Printf("Please select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)\n\n")

		for {

			fmt.Print("Enter your choice: ")

			fmt.Scanln(&level)

			switch level {
			case 1:
				guesses = 10
				OptionMessage("Easy")
			case 2:
				guesses = 5
				OptionMessage("Medium")
			case 3:
				guesses = 3
				OptionMessage("Hard")
			default:
				continue
			}
			break
		}

		fmt.Printf("You have %d chances to guess the correct number", guesses)

		for x := 1; x <= guesses; x++ {

			fmt.Print("\n\nEnter your guess: ")

			fmt.Scanln(&guess)

			if x == guesses {
				fmt.Printf("\n\nTry again, you don't win the game!\n")
			} else {
				if num > guess {
					fmt.Printf("Incorrect! The number is greater than %d.", guess)
				} else if num < guess {
					fmt.Printf("Incorrect! The number is less than %d.", guess)
				} else {
					fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", x)
					break
				}

			}

		}

		fmt.Println("\n\nAre you  want continue the game? (N/Y)): ")

		fmt.Scan(&message)

		switch message {
		case `N`:
			status = true
		case `Y`:
			status = false
		default:
			fmt.Printf("Don't recognize the output %q \nTry again\n\n", message)
		}

		if status {
			break
		}

	}
}
