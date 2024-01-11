package flashcards

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"time"
)

// Constants for the SM2 algorithm
const (
	defaultRepetition = 0
	defaultEFactor    = 2.5
	firstRepetition   = 1
	secondRepetition  = 2
	GradeOne          = 1
	gradeTwo          = 2
	gradeThree        = 3
	gradeFour         = 4
)

// FlashCard represents a flashcard with a question, answer, and SM2 algorithm parameters
type FlashCard struct {
	ID          string
	Question    string
	Answer      string
	Repetitions int
	EFactor     float64
	NextReview  time.Time
}

func NewFlashCard(question, answer string) FlashCard {
	return FlashCard{
		ID:          GenerateID(question, answer),
		Question:    question,
		Answer:      answer,
		Repetitions: defaultRepetition,           // Start with 0 for no assumed successful recall
		EFactor:     defaultEFactor,              // Default Easiness Factor
		NextReview:  time.Now().AddDate(0, 0, 0), // Default Next Review Date (1 day in the future)
	}
}

func (card *FlashCard) ApplySM2Algorithm(grade string) {
	numericGrade := convertGrade(grade)

	// 1. Update repetitions and easiness factor
	if card.Repetitions == defaultRepetition || numericGrade >= gradeThree {
		card.Repetitions++
		card.EFactor = card.EFactor + 0.1 - (5.0-float64(numericGrade))*(0.08+(5.0-float64(numericGrade))*0.02)
	} else {
		card.Repetitions = firstRepetition
		card.EFactor = 1.3
	}

	// 2. Calculate the next review interval
	if card.Repetitions == firstRepetition {
		// next review is set to tomorrow if new or grade is 1 or 2
		card.NextReview = time.Now().AddDate(0, 0, 1)
	} else if card.Repetitions == secondRepetition {
		card.NextReview = time.Now().AddDate(0, 0, 6)
	} else {
		card.NextReview = time.Now().AddDate(0, 0, int(math.Round(float64(card.Repetitions)*card.EFactor)))
	}
}

// convertGrade converts the given grade to an integer and validates it
func convertGrade(grade string) int {
	numericGrade, err := strconv.Atoi(grade)
	if err != nil || numericGrade < GradeOne || numericGrade > gradeFour {
		fmt.Printf("Invalid grade. Using default grade %d.", GradeOne)
		return GradeOne // Default to the lowest grade
	}

	return numericGrade
}

// GenerateID generates a unique ID for a flashcard based on its question and answer
func GenerateID(question, answer string) string {
	// Concatenate question and answer
	data := []byte(question + answer)

	// Calculate SHA-256 hash
	hash := sha256.New()
	hash.Write(data)
	hashInBytes := hash.Sum(nil)

	// Convert hash to a hexadecimal string
	return hex.EncodeToString(hashInBytes)
}

// Function to get flashcards that are due for review today
func GetDueFlashcards(cards []FlashCard) []FlashCard {
	var dueFlashcards []FlashCard

	currentDate := time.Now()

	for _, card := range cards {
		if currentDate.After(card.NextReview) || currentDate.Equal(card.NextReview) {
			dueFlashcards = append(dueFlashcards, card)
		}
	}

	return dueFlashcards
}

func FindCardByID(cards []FlashCard, id string) *FlashCard {
	for i, card := range cards {
		if card.ID == id {
			return &cards[i]
		}
	}
	return nil
}
