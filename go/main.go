package main

import (
	"bufio"
	"fmt"
	"nerddeck/flashcards"
	"nerddeck/storage"
	"os"
	"strings"
)

func main() {
	printWelcomeMessage()
	cards, err := storage.LoadFlashCards()
	if err != nil {
		fmt.Printf("Error loading flashcards: %v\n", err)
		os.Exit(1)
	}

	for {
		printMainMenu()

		option := getUserInput("Select an option: ")

		switch option {
		case "0":
			printInstructions()
		case "1":
			cards, err = addFlashCard(cards)
			if err != nil {
				fmt.Printf("Error adding flashcard: %v\n", err)
			}
		case "2":
			viewFlashCards(cards)
		case "3":
			cards, err = startLearning(cards)
			if err != nil {
				fmt.Printf("Error starting learning: %v\n", err)
			}
		case "4":
			fmt.Println("\n\n================================")
			fmt.Println("Exiting NerdDeck. Goodbye!")
			os.Exit(0)
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
	fmt.Println("For the best experience go full screen mode.")
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

func addFlashCard(cards []flashcards.FlashCard) ([]flashcards.FlashCard, error) {
	question := getUserInput("Enter the question: ")
	answer := getUserInput("Enter the answer: ")

	// Verify unique ids (question + answer must be unique)
	id := flashcards.GenerateID(question, answer)
	existingCard := flashcards.FindCardByID(cards, id)
	if existingCard != nil {
		return cards, fmt.Errorf("A card with the same question and answer already exists.")
	}

	newCard := flashcards.NewFlashCard(question, answer)
	cards = append(cards, newCard)

	err := storage.SaveFlashCards(cards)
	if err != nil {
		return cards, fmt.Errorf("Error saving flashcards: %v", err)
	}

	fmt.Println("Flash card added successfully!")
	return cards, nil
}

func viewFlashCards(cards []flashcards.FlashCard) {
	if len(cards) == 0 {
		fmt.Println("No flash cards available. Add some cards first.")
		return
	}
	for i, card := range cards {
		fmt.Printf("%d. Q: %s\n   A: %s\n", i+1, card.Question, card.Answer)
	}
}

func startLearning(cards []flashcards.FlashCard) ([]flashcards.FlashCard, error) {
	if len(cards) == 0 {
		return cards, fmt.Errorf("No flash cards available. Add some cards first.")
	}

	// Check for due flashcards based on the current date
	dueFlashcards := flashcards.GetDueFlashcards(cards)

	if len(dueFlashcards) == 0 {
		fmt.Println("No flashcards are due for review today.")
		return cards, fmt.Errorf("No flashcards are due for review today.")
	}
	fmt.Println("Due Flash Cards:")
	fmt.Println("Starting Learning Mode. You got this :)")
	for _, dueCard := range dueFlashcards {
		card := flashcards.FindCardByID(cards, dueCard.ID)
		fmt.Printf("Q: %s\n", card.Question)
		getUserInput("Press Enter to reveal the answer...")
		fmt.Printf("A: %s\n\n", card.Answer)

		grade := getUserInput("How well did you remember this card 1-4\n")
		card.ApplySM2Algorithm(grade)
	}

	err := storage.SaveFlashCards(cards)
	if err != nil {
		return cards, fmt.Errorf("Error saving flashcards: %v", err)
	}

	return cards, nil
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
