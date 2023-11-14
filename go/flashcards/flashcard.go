package flashcards

type Flashcard interface {
    ShowQuestion() string
    ShowAnswer() string
}