module FlashCard
open System

let defaultRepetition = 0
let gradeThree = 3
let firstRepetition = 1
let secondRepetition = 2
let defaultGrade = 1

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

let getDueFlashCards(cards: FlashCardDeck) : FlashCardDeck =
    cards |> List.filter(fun card -> DateTime.Now >= card.NextReview)
    
let findCardByID(cards: FlashCardDeck) (id: string) : FlashCard option =
    cards |> List.tryFind (fun card -> card.ID = id)
    
let convertGrade(grade: string) : int =
    match Int32.TryParse grade with
    | true, numericGrade when numericGrade >= 1 && numericGrade <= 4 -> numericGrade
    | _ -> 
        printfn $"Invalid grade. Using default grade {defaultGrade}."
        defaultGrade
    
let calculateNewEFactor (oldEFactor: float) (grade: int) : float =
    oldEFactor + 0.1 - (5.0 - float grade) * (0.08 + (5.0 - float grade) * 0.02)
    
let applySM2Algorithm card grade =
    let numericGrade = convertGrade grade

    // 1. Update repetitions and easiness factor
    let repetitions, eFactor =
        if card.Repetitions = defaultRepetition || numericGrade >= gradeThree then
            card.Repetitions + 1, calculateNewEFactor card.EFactor numericGrade
        else
            firstRepetition, 1.3

    // 2. Calculate the next review interval
    let nextReview = match repetitions with
                        | r when r = firstRepetition -> DateTime.Now.AddDays(1.0)
                        | r when r = secondRepetition -> DateTime.Now.AddDays(6.0)
                        | _ -> DateTime.Now.AddDays(float repetitions * eFactor)

    { card with Repetitions = repetitions; EFactor = eFactor; NextReview = nextReview }
    