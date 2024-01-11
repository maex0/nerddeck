package flashcards

import (
	"testing"
	"time"
)

func TestNewFlashCard(t *testing.T) {
	// Arrange
	question := "What is the capital of France?"
	answer := "Paris"

	// Act
	card := NewFlashCard(question, answer)

	// Assert
	if card.Question != question {
		t.Errorf("Expected question %s, got %s", question, card.Question)
	}
}

func TestApplySM2Algorithm(t *testing.T) {
	// Arrange
	card := FlashCard{
		Repetitions: 0,
		EFactor:     2.5,
		NextReview:  time.Now(),
	}

	// Act
	card.ApplySM2Algorithm("3")

	// Assert
	if card.Repetitions != 1 {
		t.Errorf("Expected repetitions 1, got %d", card.Repetitions)
	}
}

func TestConvertGradeSuccess(t *testing.T) {
	// Arrange
	grade := "2"

	// Act
	result := convertGrade(grade)

	// Assert
	if result != 2 {
		t.Errorf("Expected result 2, got %d", result)
	}
}

func TestConvertGradeWrongInput(t *testing.T) {
	// Arrange
	grade := "input"

	// Act
	result := convertGrade(grade)

	// Assert
	if result != 1 {
		t.Errorf("Expected result 1, got %d", result)
	}
}

func TestGenerateID(t *testing.T) {
	// Arrange
	question := "What is the capital of France?"
	answer := "Paris"

	// Act
	result1 := GenerateID(question, answer)
	result2 := GenerateID(question, answer)

	// Assert
	if result1 != result2 {
		t.Errorf("Expected same id, instead got %s and %s", result1, result2)
	}
}

func TestGetDueFlashcards(t *testing.T) {
	cards := []FlashCard{
		{ID: "1", NextReview: time.Now().Add(-24 * time.Hour)},
		{ID: "2", NextReview: time.Now().Add(24 * time.Hour)},
	}

	dueCards := GetDueFlashcards(cards)

	if len(dueCards) != 1 {
		t.Errorf("Expected 1 card, got %d", len(dueCards))
	}

	if dueCards[0].ID != "1" {
		t.Errorf("Expected card with ID 1, got %s", dueCards[0].ID)
	}
}

func TestFindCardByID(t *testing.T) {
	cards := []FlashCard{
		{ID: "1", NextReview: time.Now().Add(-24 * time.Hour)},
		{ID: "2", NextReview: time.Now().Add(24 * time.Hour)},
	}

	card := FindCardByID(cards, "1")

	if card == nil {
		t.Errorf("Expected to find card, got nil")
	}

	if card.ID != "1" {
		t.Errorf("Expected card with ID 1, got %s", card.ID)
	}

	card = FindCardByID(cards, "3")

	if card != nil {
		t.Errorf("Expected to not find card, got %v", card)
	}
}

func TestApplySM2AlgorithmElseParts(t *testing.T) {
	card := NewFlashCard("Question", "Answer")

	// Test the case where the grade is less than 3 and the repetitions are not 0
	card.Repetitions = 1
	card.ApplySM2Algorithm("2")
	if card.Repetitions != firstRepetition || card.EFactor != 1.3 {
		t.Errorf("Expected repetitions to be %d and EFactor to be %f, got %d and %f", firstRepetition, 1.3, card.Repetitions, card.EFactor)
	}

	// Test the case where the repetitions are more than 2
	card.Repetitions = 3
	card.ApplySM2Algorithm("4")
	if card.Repetitions != 4 || card.NextReview.Before(time.Now().AddDate(0, 0, int(card.Repetitions))) {
		t.Errorf("Expected repetitions to be 4 and NextReview to be in the future, got %d and %s", card.Repetitions, card.NextReview)
	}
}
