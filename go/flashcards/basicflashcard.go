package flashcards

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"time"
)

type BasicFlashCard struct {
	ID          string
	Question    string
	Answer      string
	Repetitions int
	EFactor     float64
	NextReview  time.Time
}

func NewBasicFlashCard(question, answer string) BasicFlashCard {
	id := generateID(question, answer)
	return BasicFlashCard{
		ID:          id,
		Question:    question,
		Answer:      answer,
		Repetitions: 0,                           // Start with 0 for no assumed successful recall
		EFactor:     2.5,                         // Default Easiness Factor
		NextReview:  time.Now().AddDate(0, 0, 0), // Default Next Review Date (1 day in the future)
	}
}

func (card BasicFlashCard) ShowQuestion() string {
	return card.Question
}

func (card BasicFlashCard) ShowAnswer() string {
	return card.Answer
}

func (card *BasicFlashCard) ApplySM2Algorithm(grade string) {
	numericGrade := convertGrade(grade)

	// 1. Update repetitions and easiness factor
	if card.Repetitions == 0 || numericGrade >= 3 {
		card.Repetitions++
		card.EFactor = card.EFactor + 0.1 - (5.0-float64(numericGrade))*(0.08+(5.0-float64(numericGrade))*0.02)
	} else {
		card.Repetitions = 1
		card.EFactor = 1.3
	}

	// 2. Calculate the next review interval
	if card.Repetitions == 1 {
		// next review is set to tomorrow if new or grade is 1 or 2
		card.NextReview = time.Now().AddDate(0, 0, 1)
	} else if card.Repetitions == 2 {
		card.NextReview = time.Now().AddDate(0, 0, 6)
	} else {
		card.NextReview = time.Now().AddDate(0, 0, int(math.Round(float64(card.Repetitions)*card.EFactor)))
	}
}

func convertGrade(grade string) int {
	numericGrade, err := strconv.Atoi(grade)
	if err != nil || numericGrade < 1 || numericGrade > 4 {
		fmt.Println("Invalid grade. Using default value.")
		return 1 // Default to the lowest grade
	}

	return numericGrade
}

func generateID(question, answer string) string {
	// Concatenate question and answer
	data := []byte(question + answer)

	// Calculate SHA-256 hash
	hash := sha256.New()
	hash.Write(data)
	hashInBytes := hash.Sum(nil)

	// Convert hash to a hexadecimal string
	return hex.EncodeToString(hashInBytes)
}
