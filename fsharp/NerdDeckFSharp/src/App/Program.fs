open System
open System.IO
open System.Text.Json
open Library


let flashcardsFile = "flashcards.json"

let loadFlashCards() =
    try
        let file = File.ReadAllText(flashcardsFile)
        let cards = JsonSerializer.Deserialize<FlashCard list>(file)
        Ok cards
    with
    | :? FileNotFoundException -> Ok []
    | ex -> Error ex

let saveFlashCards (cards: FlashCard list) =
    try
        let options = JsonSerializerOptions(WriteIndented = true)
        let file = JsonSerializer.Serialize(cards, options)
        File.WriteAllText(flashcardsFile, file)
        Ok ()
    with
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
    printf $"%s{prompt}"
    Console.ReadLine().Trim()

let addFlashCard(cards: FlashCard list) =
    let question = getUserInput("Question: ")
    let answer = getUserInput("Answer: ")

    let newCard = {
        ID = Guid.NewGuid().ToString()
        Question = question
        Answer = answer
        Repetitions = 0
        EFactor = 2.5
        NextReview = DateTime.Now
    }

    let updatedCards = newCard :: cards
    match saveFlashCards updatedCards with
        | Ok _ -> printfn "Flash card added successfully!"; cards
        | Error err -> printfn "Error saving flashcards: %s" err.Message; cards

let viewFlashCards(cards: FlashCard list) =
    if List.isEmpty cards then
        printfn "No flash cards available. Add some cards first."
    else
        cards
        |> List.iteri (fun i card -> printfn "%d. Q: %s\n   A: %s\n" (i+1) card.Question card.Answer)
    cards

let startLearning(cards: FlashCard list) =
    printfn "Not implemented yet."
    cards

let rec mainLoop(cards: FlashCard list) =
    printMainMenu()

    let option = getUserInput("Select an option: ")

    match option with
    | "0" -> printInstructions(); cards |> mainLoop
    | "1" -> cards |> addFlashCard |> mainLoop
    | "2" -> cards |> viewFlashCards |> mainLoop
    | "3" -> cards |> startLearning |> mainLoop
    | "4" -> printfn "\n\n================================"; printfn "Exiting NerdDeck. Goodbye!";
    | _ -> printfn "Invalid option. Please try again."; mainLoop cards


printWelcomeMessage()
printMainMenu()
match loadFlashCards() with
    | Ok cards -> mainLoop cards
    | Error ex -> printfn "Error loading flashcards: %s" (ex.Message)
