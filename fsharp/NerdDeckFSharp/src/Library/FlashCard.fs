module FlashCard
open System

type FlashCard = {
    ID: string
    Question: string
    Answer: string
    Repetitions: int
    EFactor: float
    NextReview: DateTime
}

type FlashCardDeck = List<FlashCard>

let addFlashCard(question: string)(answer:string)(cards: FlashCardDeck) : FlashCardDeck =
    let newCard = {
        ID = Guid.NewGuid().ToString()
        Question = question
        Answer = answer
        Repetitions = 0
        EFactor = 2.5
        NextReview = DateTime.Now
    }
    
    newCard :: cards
