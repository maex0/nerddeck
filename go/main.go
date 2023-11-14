package main

import (
	"bufio"
	"fmt"
	"nerddeck/flashcards"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to NerdDeck!")
	fmt.Println("If you want to know more about NerdDeck, visit https://github.com/maex0/nerddeck")

	var cards []flashcards.BasicFlashCard

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Add Flash Card")
		fmt.Println("2. View Flash Cards")
		fmt.Println("3. Start Learning")
		fmt.Println("4. Exit")

		option := getUserInput("Select an option: ")

		switch option {
		case "1":
			question := getUserInput("Enter the question: ")
			answer := getUserInput("Enter the answer: ")

			newCard := flashcards.BasicFlashCard{Front: question, Back: answer}
			cards = append(cards, newCard)
			
			fmt.Println("Flash card added successfully!")

		case "2":
			fmt.Println("\nFlash Cards:")
			for i, card := range cards {
				fmt.Printf("%d. Q: %s\n   A: %s\n", i+1, card.Front, card.Back)
			}

		case "3":
			if len(cards) == 0 {
				fmt.Println("No flash cards available. Add some cards first.")
				continue
			}

			fmt.Println("Starting Learning Mode. You got this :)")
			for _, card := range cards {
				fmt.Printf("Q: %s\n", card.Front)
				getUserInput("Press Enter to reveal the answer...")
				fmt.Printf("A: %s\n\n", card.Back)
			}

		case "4":
			fmt.Println("Exiting NerdDeck. Goodbye!")
			os.Exit(0)

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	adsfasdlf
	return strings.TrimSpace(scanner.Text())
}
