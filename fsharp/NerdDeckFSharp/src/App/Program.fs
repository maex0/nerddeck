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
    printfn "🚀 Main Menu, please make a choice"
    printfn "options"
    printfn "0. Instructions"
    printfn "1. Add Flash Card"
    printfn "2. View Flash Cards"
    printfn "3. Start Learning"
    printfn "4. Exit"
    printfn "================================\n\n"

let printInstructions() =
    printfn "\nInstructions:"
    printfn "1. Add Flash Card: Enter a question and answer to create a new flash card."
    printfn "2. View Flash Cards: Display all existing flash cards."
    printfn "3. Start Learning: Review flash cards that are due for learning today."
    printfn "   - Press Enter to reveal the answer."
    printfn "   - Rate your memory from 1 to 4:"
    printfn "     - 1: I don't remember at all. :("
    printfn "     - 2: I remember a little.     :|"
    printfn "     - 3: I remember well.         :)"
    printfn "     - 4: I remember perfectly.    :D"
    printfn "   - The SM2 spaced repetition algorithm will adjust the card's review interval."
    printfn "4. Exit: Save flash cards and exit the application."
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
    printfn "This program is written in F#."


let getUserInput(prompt: string) =
    printf "%s" prompt
    Console.ReadLine().Trim()

let rec mainLoop(cards: FlashCard list) =
    printMainMenu()

    let option = getUserInput("Select an option: ")

    match option with
    | "0" -> printInstructions(); mainLoop cards
    | "1" -> printfn "Selected option 1"; mainLoop cards
    | "2" -> printfn "Selected option 2"; mainLoop cards
    | "3" -> printfn "Selected option 3"; mainLoop cards
    | "4" -> printfn "\n\n================================"; printfn "Exiting NerdDeck. Goodbye!";
    | _ -> printfn "Invalid option. Please try again."; mainLoop cards

printWelcomeMessage()
printMainMenu()
match loadFlashCards() with
    | Ok cards -> mainLoop cards
    | Error ex -> printfn "Error loading flashcards: %s" (ex.Message)
