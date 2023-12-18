open System
open System.IO
open System.Text.Json

type FlashCard = {
    ID: string
    Question: string
    Answer: string
    Repetitions: int
    EFactor: float
    NextReview: DateTime
}

let flashcardsFile = "flashcards.json"

let loadFlashCards() =
    try
        let file = File.ReadAllText(flashcardsFile)
        let cards = JsonSerializer.Deserialize<FlashCard list>(file)
        Ok cards
    with
    | :? FileNotFoundException -> Ok []
    | ex -> Error ex

let printMainMenu() =
    printfn "\n\n================================"
    printfn "Main Menu:"
    printfn "0. Instructions"
    printfn "1. Add Flash Card"
    printfn "2. View Flash Cards"
    printfn "3. Start Learning"
    printfn "4. Exit"
    printfn "================================\n\n"

let printWelcomeMessage() =
    let nerdDeckASCII =
        """
    _   __              ______            __  
   / | / /__  _________/ / __ \___  _____/ /__
  /  |/ / _ \/ ___/ __  / / / / _ \/ ___/ //_/
 / /|  /  __/ /  / /_/ / /_/ /  __/ /__/ ,<   
/_/ |_/\___/_/   \__,_/_____/\___/\___/_/|_|  
"""
    printfn "\nWelcome to\n%s" nerdDeckASCII
    printfn "Developed by Maximilian Gobbel"
    printfn "If you want to know more about NerdDeck, visit https://github.com/maex0/nerddeck"
    printfn "For the best experience go full screen mode."

printWelcomeMessage()
printMainMenu()
match loadFlashCards() with
    | Ok cards -> printfn "Loaded %d flashcards." (List.length cards)
    | Error ex -> printfn "Error loading flashcards: %s" (ex.Message)