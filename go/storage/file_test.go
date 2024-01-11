package storage

import (
	"encoding/json"
	"nerddeck/flashcards"
	"os"
	"testing"
)

func TestLoadFlashCards(t *testing.T) {
	// Prepare a test file
	cards := []flashcards.FlashCard{
		{Question: "Question1", Answer: "Answer1"},
		{Question: "Question2", Answer: "Answer2"},
	}
	data, _ := json.Marshal(cards)
	os.WriteFile(flashcardsFile, data, filemode)

	// Test the function
	loadedCards, err := LoadFlashCards()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(loadedCards) != len(cards) {
		t.Errorf("Expected %d cards, got %d", len(cards), len(loadedCards))
	}

	// Clean up
	os.Remove(flashcardsFile)
}

func TestSaveFlashCards(t *testing.T) {
	// Prepare some cards
	cards := []flashcards.FlashCard{
		{Question: "Question1", Answer: "Answer1"},
		{Question: "Question2", Answer: "Answer2"},
	}

	// Test the function
	err := SaveFlashCards(cards)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check the file
	data, _ := os.ReadFile(flashcardsFile)
	var loadedCards []flashcards.FlashCard
	json.Unmarshal(data, &loadedCards)

	if len(loadedCards) != len(cards) {
		t.Errorf("Expected %d cards, got %d", len(cards), len(loadedCards))
	}

	// Clean up
	os.Remove(flashcardsFile)
}
