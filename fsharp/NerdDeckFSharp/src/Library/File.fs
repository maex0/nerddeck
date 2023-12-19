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
    | ex -> Error ex

let saveFlashCards (cards: FlashCardDeck) =
    try
        let options = JsonSerializerOptions(WriteIndented = true)
        let file = JsonSerializer.Serialize(cards, options)
        File.WriteAllText(flashcardsFile, file)
        Ok ()
    with
    | ex -> Error ex