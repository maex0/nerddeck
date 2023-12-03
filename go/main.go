package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"nerddeck/flashcards"
	"os"
	"strings"
)

func main() {
	nerdDeckASCII :=
	`
    _   __              ______            __  
   / | / /__  _________/ / __ \___  _____/ /__
  /  |/ / _ \/ ___/ __  / / / / _ \/ ___/ //_/
 / /|  /  __/ /  / /_/ / /_/ /  __/ /__/ ,<   
/_/ |_/\___/_/   \__,_/_____/\___/\___/_/|_|  
`
	fmt.Println("\nWelcome to\n",nerdDeckASCII)
	fmt.Println("Developed by Maximilian Gobbel")
	fmt.Println("If you want to know more about NerdDeck, visit https://github.com/maex0/nerddeck")

	var cards []flashcards.BasicFlashCard

	// Load flashcards from a file, if available
	loadFlashCards(&cards)

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

			newCard := flashcards.BasicFlashCard{Question: question, Answer: answer}
			cards = append(cards, newCard)
			
			fmt.Println("Flash card added successfully!")

		case "2":
			fmt.Println("\nFlash Cards:")
			for i, card := range cards {
				fmt.Printf("%d. Q: %s\n   A: %s\n", i+1, card.Question, card.Answer)
			}

		case "3":
			if len(cards) == 0 {
				fmt.Println("No flash cards available. Add some cards first.")
				continue
			}

			fmt.Println("Starting Learning Mode. You got this :)")
			for _, card := range cards {
				fmt.Printf("Q: %s\n", card.Question)
				getUserInput("Press Enter to reveal the answer...")
				fmt.Printf("A: %s\n\n", card.Answer)
			}

		case "4":
			// Save flashcards to a file before exiting
			saveFlashCards(cards)
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
	return strings.TrimSpace(scanner.Text())
}


func loadFlashCards(cards *[]flashcards.BasicFlashCard) {
	file, err := os.ReadFile("flashcards.json")
	if err == nil {
		err = json.Unmarshal(file, &cards)
		if err != nil {
			fmt.Println("Error loading flashcards:", err)
		} else {
			fmt.Println("Flashcards loaded successfully.")
		}
	}
}

func saveFlashCards(cards []flashcards.BasicFlashCard) {
	file, err := json.MarshalIndent(cards, "", "  ")
	if err != nil {
		fmt.Println("Error saving flashcards:", err)
		return
	}

	err = os.WriteFile("flashcards.json", file, 0644)
	if err != nil {
		fmt.Println("Error saving flashcards:", err)
	} else {
		fmt.Println("Flashcards saved successfully.")
	}
}