module File

open System.IO
open System.Text.Json
open FlashCard

let flashcardsFile = "flashcards.json"

let loadFlashCards() =
    try
        let file = File.ReadAllText(flashcardsFile)
        let cards = JsonSerializer.Deserialize<FlashCard list>(file)
        Ok cards
    with
    | :? FileNotFoundException -> Ok []
    | ex -> Error "An error occurred while loading the flashcards."

let saveFlashCards (cards: FlashCardDeck) =
    try
        let options = JsonSerializerOptions(WriteIndented = true)
        let file = JsonSerializer.Serialize(cards, options)
        File.WriteAllText(flashcardsFile, file)
        Ok ()
    with
    | ex -> Error "An error occurred while saving the flashcards."