package main

import (
	"bufio"
	"fmt"
	rand "math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {

	var chances uint8
	var scanner = bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
	
	Welcome to the Number Guessing Game!
	I'm thinking of a number between 1 and 100.

	Please select the difficulty level:
	1. Easy (10 chances)
	2. Medium (5 chances)
	3. Hard (3 chances)

	`)
		var lvl string
		for {

			fmt.Printf("Enter your choice (1,2, or 3): ")
			scanner.Scan()
			choice := scanner.Text()
			switch choice {
			case "1":
				chances = 10
				lvl = "Easy"

			case "2":
				chances = 5
				lvl = "Medium"

			case "3":
				chances = 3
				lvl = "Hard"

			default:
				fmt.Printf("Enter a valid choice.\n")
				continue
			}
			break
		}

		fmt.Printf("Great! You've selected the %s difficulty level.\n", lvl)
		startGame(chances, scanner)

		fmt.Printf("Do you want to play again? (y/n): ")
		scanner.Scan()
		again := strings.ToLower(scanner.Text())
		if again != "y" {
			fmt.Printf("Thank you for playing! Goodbye!\n")
			break
		}
	}

}

func startGame(chances uint8, scanner *bufio.Scanner) {

	var randNum = uint8(rand.IntN(100) + 1)

	for i := chances; i > 0; i-- {
		fmt.Printf("Enter your guess: ")
		scanner.Scan()
		input := scanner.Text()

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Invalid input. Try again.\n")
			i++
			continue
		}

		attempts := chances - i + 1

		g := uint8(guess)
		if g == randNum {
			fmt.Printf("Congratulations! You guessed the correct number in %d %s.\n", attempts, Pluralize(attempts, "attempt", "attempts"))
			return
		} else if g < randNum {
			fmt.Printf("Incorrect! The number is greater than %d. \n", g)

		} else {
			fmt.Printf("Incorrect! The number is less than %d. \n", g)

		}
		if i > 1 {
			fmt.Printf("You have %d %s left.\n", i-1, Pluralize(i-1, "chance", "chances"))
			continue
		}

	}
	fmt.Printf("Sorry, you've run out of chances. The correct number was %d. Better luck next time!\n", randNum)
}

func Pluralize(count uint8, singular, plural string) string {
	if count == 1 {
		return singular
	}
	return plural
}
