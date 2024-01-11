package storage

import (
	"encoding/json"
	"fmt"
	"nerddeck/flashcards"
	"os"
	"sync"
)

// Constants for the file name, file mode, and JSON indentation
const (
	flashcardsFile = "flashcards.json"
	filemode       = 0644
	indent         = "  "
)

// Mutex for synchronizing access to the file
var mu sync.Mutex

// LoadFlashCards loads the flashcards from the JSON file
func LoadFlashCards() ([]flashcards.FlashCard, error) {
	var cards []flashcards.FlashCard

	// Lock the mutex to prevent concurrent reads/writes
	mu.Lock()
	defer mu.Unlock()

	// Read the file
	file, err := os.ReadFile(flashcardsFile)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into the flashcards slice
	err = json.Unmarshal(file, &cards)
	if err != nil {
		return nil, fmt.Errorf("Error loading flashcards: %w", err)
	}

	return cards, nil
}

// SaveFlashCards saves the flashcards to the JSON file
func SaveFlashCards(cards []flashcards.FlashCard) error {
	// Lock the mutex to prevent concurrent reads/writes
	mu.Lock()
	defer mu.Unlock()

	// Marshal the flashcards slice into JSON data
	file, err := json.MarshalIndent(cards, "", indent)
	if err != nil {
		return fmt.Errorf("Error saving flashcards: %w", err)
	}

	// Write the JSON data to the file
	err = os.WriteFile(flashcardsFile, file, filemode)
	if err != nil {
		return fmt.Errorf("Error saving flashcards: %w", err)
	}

	return nil
}
