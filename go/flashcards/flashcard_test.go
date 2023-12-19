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
