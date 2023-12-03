package flashcards

import (
	"time"
)

type BasicFlashCard struct {
	ID int
	Question string
	Answer  string
	Interval int
	Ease int
	ReviewDate time.Time
	Deck Deck
}

func (card BasicFlashCard) ShowQuestion() string {
    return card.Question
}

func (card BasicFlashCard) ShowAnswer() string {
    return card.Answer
}