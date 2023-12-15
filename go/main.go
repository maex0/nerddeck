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

	printWelcomeMessage()

	var cards []flashcards.FlashCard

	// Load flashcards from a file, if available
	loadFlashCards(&cards)

	for {
		printMainMenu()

		option := getUserInput("Select an option: ")

		switch option {
		case "0":
			printInstructions()
		case "1":
			addFlashCard(&cards)
		case "2":
			viewFlashCards(&cards)
		case "3":
			startLearning(&cards)
		case "4":
			saveAndExit(cards)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func printMainMenu() {
	fmt.Println("\n\n================================")
	fmt.Println("Main Menu:")
	fmt.Println("0. Instructions")
	fmt.Println("1. Add Flash Card")
	fmt.Println("2. View Flash Cards")
	fmt.Println("3. Start Learning")
	fmt.Println("4. Exit")
	fmt.Println("================================\n\n")
}

func printWelcomeMessage() {
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
}

func printInstructions() {
	fmt.Println("\nInstructions:")
	fmt.Println("1. Add Flash Card: Enter a question and answer to create a new flash card.")
	fmt.Println("2. View Flash Cards: Display all existing flash cards.")
	fmt.Println("3. Start Learning: Review flash cards that are due for learning today.")
	fmt.Println("   - Press Enter to reveal the answer.")
	fmt.Println("   - Rate your memory from 1 to 4:")
	fmt.Println("     - 1: I don't remember at all. :(")
	fmt.Println("     - 2: I remember a little.     :|")
	fmt.Println("     - 3: I remember well.         :)")
	fmt.Println("     - 4: I remember perfectly.    :D")
	fmt.Println("   - The SM2 spaced repetition algorithm will adjust the card's review interval.")
	fmt.Println("4. Exit: Save flash cards and exit the application.")
	fmt.Println("================================\n\n")
}

func addFlashCard(cards *[]flashcards.FlashCard) {
	question := getUserInput("Enter the question: ")
	answer := getUserInput("Enter the answer: ")

	newCard := flashcards.NewFlashCard(question, answer)
	*cards = append(*cards, newCard)

	// Presave cards
	saveFlashCards(*cards)

	fmt.Println("Flash card added successfully!")
}

func viewFlashCards(cards *[]flashcards.FlashCard) {
	if len(*cards) == 0 {
		fmt.Println("No flash cards available. Add some cards first.")
		return
	}
	for i, card := range *cards {
		fmt.Printf("%d. Q: %s\n   A: %s\n", i+1, card.Question, card.Answer)
	}

}

func startLearning(cards *[]flashcards.FlashCard) {
	if len(*cards) == 0 {
		fmt.Println("No flash cards available. Add some cards first.")
		return
	}

	// Check for due flashcards based on the current date
	dueFlashcards := getDueFlashcards(*cards)

	if len(dueFlashcards) == 0 {
		fmt.Println("No flashcards are due for review today.")
	} else {
		fmt.Println("Due Flash Cards:")
		fmt.Println("Starting Learning Mode. You got this :)")
		for _, dueCard := range dueFlashcards {
			card := findCardByID(*cards, dueCard.ID)
			fmt.Printf("Q: %s\n", card.Question)
			getUserInput("Press Enter to reveal the answer...")
			fmt.Printf("A: %s\n\n", card.Answer)

			// Apply SM2 algorithm
			grade := getUserInput("How well did you remember this card 1-4\n")
			card.ApplySM2Algorithm(grade)
		}

		saveFlashCards(*cards)
	}
}

func saveAndExit(cards []flashcards.FlashCard) {
	saveFlashCards(cards)
	fmt.Println("Exiting NerdDeck. Goodbye!")
	os.Exit(0)
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func loadFlashCards(cards *[]flashcards.FlashCard) {
	file, err := os.ReadFile("flashcards.json")
	if err == nil {
		err = json.Unmarshal(file, &cards)
		if err != nil {
			fmt.Println("Error loading flashcards:", err)
		}
	}
}

func saveFlashCards(cards []flashcards.FlashCard) {
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
func getDueFlashcards(cards []flashcards.FlashCard) []flashcards.FlashCard {
	var dueFlashcards []flashcards.FlashCard

	currentDate := time.Now()

	for _, card := range cards {
		if currentDate.After(card.NextReview) || currentDate.Equal(card.NextReview) {
			dueFlashcards = append(dueFlashcards, card)
		}
	}

	return dueFlashcards
}

// Function to find a card by its ID
func findCardByID(cards []flashcards.FlashCard, id string) *flashcards.FlashCard {
	for i, card := range cards {
		if card.ID == id {
			return &cards[i]
		}
	}
	return nil
}
