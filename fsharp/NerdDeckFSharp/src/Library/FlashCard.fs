namespace Library
open System

type FlashCard = {
    ID: string
    Question: string
    Answer: string
    Repetitions: int
    EFactor: float
    NextReview: DateTime
}