package flashcards

type BasicFlashCard struct {
	Front string
	Back  string
}

func (card BasicFlashCard) ShowQuestion() string {
    return card.Front
}

func (card BasicFlashCard) ShowAnswer() string {
    return card.Back
}