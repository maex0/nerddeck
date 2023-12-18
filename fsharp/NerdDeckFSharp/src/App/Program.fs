open System
open FlashCard
open File



let printMainMenu () =
    printfn
        "\n\n=====================================
🚀 Main Menu, please make a choice
Options:
0. Instructions
1. Add Flash Card
2. View Flash Cards
3. Start Learning
4. Exit
=====================================\n\n"

let printInstructions () =
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

let printWelcomeMessage () =
    let nerdDeckASCII =
        """
    _   __              ______            __  
   / | / /__  _________/ / __ \___  _____/ /__
  /  |/ / _ \/ ___/ __  / / / / _ \/ ___/ //_/
 / /|  /  __/ /  / /_/ / /_/ /  __/ /__/ ,<   
/_/ |_/\___/_/   \__,_/_____/\___/\___/_/|_|  
"""

    printfn $"\nWelcome to\n%s{nerdDeckASCII}"
    printfn "Developed by Maximilian Gobbel"
    printfn "If you want to know more about NerdDeck, visit https://github.com/maex0/nerddeck"
    printfn "For the best experience go full screen mode."
    printfn "This program is written in F#."


let getUserInput (prompt: string) : string =
    printf $"%s{prompt}"
    Console.ReadLine().Trim()


let createNewFlashCard (cards: FlashCardDeck) : FlashCardDeck =
    let question = getUserInput ("Question: ")
    let answer = getUserInput ("Answer: ")

    let updatedDeck = addFlashCard question answer cards

    match saveFlashCards updatedDeck with
    | Ok _ ->
        printfn "Flash card added successfully!"
        updatedDeck
    | Error err ->
        printfn $"Error saving flashcards: %s{err.Message}"
        updatedDeck

let viewFlashCards (cards: FlashCardDeck) =
    if List.isEmpty cards then
        printfn "No flash cards available. Add some cards first."
    else
        cards
        |> List.iteri (fun i card -> printfn $"%d{i + 1}. Q: %s{card.Question}\n   A: %s{card.Answer}\n")

    cards

let startLearning (cards: FlashCardDeck) : FlashCardDeck =
    if List.isEmpty cards then
        printfn "No flash cards available. Add some cards first."
        cards
    else
        // Check for due flashcards based on the current date
        let dueFlashcards = getDueFlashCards cards

        if List.isEmpty dueFlashcards then
            printfn "No flashcards are due for review today."
            cards
        else
            printfn "Due Flash Cards:"
            printfn "Starting Learning Mode. You got this :)"

            dueFlashcards
            |> List.fold
                (fun updatedCards dueCard ->
                    match findCardByID cards dueCard.ID with
                    | Some card ->
                        printfn $"Q: %s{card.Question}\n"
                        let _ = getUserInput ("Press Enter to reveal the answer...")
                        printfn $"A: %s{card.Answer}\n\n"

                        let grade = getUserInput ("How well did you remember this card 1-4\n")
                        let updatedCard = applySM2Algorithm card grade
                        // Replace the old card with the updated one
                        let updatedCards =
                            List.map (fun c -> if c.ID = updatedCard.ID then updatedCard else c) updatedCards

                        match saveFlashCards updatedCards with
                        | Ok _ -> updatedCards
                        | Error err ->
                            printfn $"Error saving flashcards: %s{err.Message}"
                            updatedCards
                    | None -> updatedCards)
                cards

let rec mainLoop (cards: FlashCardDeck) : unit =
    printMainMenu ()

    let option = getUserInput ("Select an option: ")

    match option with
    | "0" ->
        printInstructions ()
        cards |> mainLoop
    | "1" -> cards |> createNewFlashCard |> mainLoop
    | "2" -> cards |> viewFlashCards |> mainLoop
    | "3" -> cards |> startLearning |> mainLoop
    | "4" ->
        printfn "\n\n================================"
        printfn "Exiting NerdDeck. Goodbye!"
    | _ ->
        printfn "Invalid option. Please try again."
        mainLoop cards


printWelcomeMessage ()
printMainMenu ()

match loadFlashCards () with
| Ok cards -> mainLoop cards
| Error ex -> printfn $"Error loading flashcards: %s{ex.Message}"
