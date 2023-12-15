package storage

import (
	"encoding/json"
	"fmt"
	"nerddeck/flashcards"
	"os"
	"sync"
)

var mu sync.Mutex

func LoadFlashCards(cards *[]flashcards.FlashCard) {
	file, err := os.ReadFile("flashcards.json")
	if err == nil {
		err = json.Unmarshal(file, &cards)
		if err != nil {
			fmt.Println("Error loading flashcards:", err)
		}
	}
}

func SaveFlashCards(cards []flashcards.FlashCard) {
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
