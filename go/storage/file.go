package storage

import (
	"encoding/json"
	"fmt"
	"nerddeck/flashcards"
	"os"
	"sync"
)

const (
	flashcardsFile = "flashcards.json"
	filemode       = 0644
	indent         = "  "
)

var mu sync.Mutex

func LoadFlashCards() ([]flashcards.FlashCard, error) {
	var cards []flashcards.FlashCard
	mu.Lock()
	defer mu.Unlock()

	file, err := os.ReadFile(flashcardsFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &cards)
	if err != nil {
		return nil, fmt.Errorf("Error loading flashcards: %w", err)
	}

	return cards, nil
}

func SaveFlashCards(cards []flashcards.FlashCard) error {
	mu.Lock()
	defer mu.Unlock()

	file, err := json.MarshalIndent(cards, "", indent)
	if err != nil {
		return fmt.Errorf("Error saving flashcards: %w", err)
	}

	err = os.WriteFile(flashcardsFile, file, filemode)
	if err != nil {
		return fmt.Errorf("Error saving flashcards: %w", err)
	}

	return nil
}
