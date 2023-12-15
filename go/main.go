package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"nerddeck/flashcards"
	"os"
	"strings"
	"sync"
	"time"
)

var mu sync.Mutex

func main() {
	nerdDeckASCII :=
		`
    _   __              ______            __  
   / | / /__  _________/ / __ \___  _____/ /__
  /  |/ / _ \/ ___/ __  / / / / _ \/ ___/ //_/
 / /|  /  __/ /  / /_/ / /_/ /  __/ /__/ ,<   
/_/ |_/\___/_/   \__,_/_____/\___/\___/_/|_|  
`
	fmt.Println("\nWelcome to\n", nerdDeckASCII)
	fmt.Println("Developed by Maximilian Gobbel")
	fmt.Println("If you want to know more about NerdDeck, visit https://github.com/maex0/nerddeck")

	var cards []flashcards.BasicFlashCard

	// Load flashcards from a file, if available
	loadFlashCards(&cards)

	for {
		fmt.Println("\n\n================================")
		fmt.Println("Main Menu:\n")
		fmt.Println("0. Instructions")
		fmt.Println("1. Add Flash Card")
		fmt.Println("2. View Flash Cards")
		fmt.Println("3. Start Learning")
		fmt.Println("4. Exit")

		fmt.Println("================================\n\n")

		option := getUserInput("Select an option: ")
		fmt.Println("================================")

		switch option {
		case "0":
			fmt.Println("Add instructions here")
		case "1":
			question := getUserInput("Enter the question: ")
			answer := getUserInput("Enter the answer: ")

			newCard := flashcards.NewBasicFlashCard(question, answer)
			cards = append(cards, newCard)

			// Presave cards
			saveFlashCards(cards)

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

			// Check for due flashcards based on the current date
			dueFlashcards := getDueFlashcards(cards)

			if len(dueFlashcards) == 0 {
				fmt.Println("No flashcards are due for review today.")
			} else {
				fmt.Println("Due Flash Cards:")
				fmt.Println("Starting Learning Mode. You got this :)")
				for _, dueCard := range dueFlashcards {
					card := findCardByID(cards, dueCard.ID)
					fmt.Printf("Q: %s\n", card.Question)
					getUserInput("Press Enter to reveal the answer...")
					fmt.Printf("A: %s\n\n", card.Answer)

					// Apply SM2 algorithm
					grade := getUserInput("How well did you remember this card 1-4\n")
					card.ApplySM2Algorithm(grade)
				}

				saveFlashCards(cards)
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
		}
	}
}

func saveFlashCards(cards []flashcards.BasicFlashCard) {
	mu.Lock()
	defer mu.Unlock()

	file, err := json.MarshalIndent(cards, "", "  ")
	if err != nil {
		fmt.Println("Error saving flashcards:", err)
		return
	}

	err = os.WriteFile("flashcards.json", file, 0644)
	if err != nil {
		fmt.Println("Error saving flashcards:", err)
	}
}

// Function to get flashcards that are due for review today
func getDueFlashcards(cards []flashcards.BasicFlashCard) []flashcards.BasicFlashCard {
	var dueFlashcards []flashcards.BasicFlashCard

	currentDate := time.Now()

	for _, card := range cards {
		if currentDate.After(card.NextReview) || currentDate.Equal(card.NextReview) {
			dueFlashcards = append(dueFlashcards, card)
		}
	}

	return dueFlashcards
}

// Function to find a card by its ID
func findCardByID(cards []flashcards.BasicFlashCard, id string) *flashcards.BasicFlashCard {
	for i, card := range cards {
		if card.ID == id {
			return &cards[i]
		}
	}
	return nil
}
