package main

import (
	"nerddeck/flashcards"
	"testing"
	"time"
)

func TestViewFlashCards(t *testing.T) {
	cards := []flashcards.FlashCard{
		{Question: "Question1", Answer: "Answer1"},
		{Question: "Question2", Answer: "Answer2"},
	}

	// Test viewing cards
	viewFlashCards(cards) // This should print the cards to the console
}

func TestStartLearning(t *testing.T) {
	cards := []flashcards.FlashCard{
		{Question: "Question1", Answer: "Answer1", NextReview: time.Now().Add(-24 * time.Hour)},
		{Question: "Question2", Answer: "Answer2", NextReview: time.Now().Add(24 * time.Hour)},
	}

	// Test starting learning
	cards, err := startLearning(cards) // This should print the due cards to the console and update their review intervals
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestStartLearningAllCardsDue(t *testing.T) {
	cards := []flashcards.FlashCard{
		{Question: "Question1", Answer: "Answer1", NextReview: time.Now().Add(-24 * time.Hour)},
		{Question: "Question2", Answer: "Answer2", NextReview: time.Now().Add(-24 * time.Hour)},
	}

	// Test starting learning with all cards due
	cards, err := startLearning(cards)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(cards) != 2 {
		t.Errorf("Expected 2 cards, got %d", len(cards))
	}
}

func TestStartLearningNoCardsDue(t *testing.T) {
	cards := []flashcards.FlashCard{
		{Question: "Question1", Answer: "Answer1", NextReview: time.Now().Add(24 * time.Hour)},
		{Question: "Question2", Answer: "Answer2", NextReview: time.Now().Add(24 * time.Hour)},
	}

	// Test starting learning with no cards due
	cards, err := startLearning(cards)
	if err == nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
