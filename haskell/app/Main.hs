module Main where

import System.Exit (exitSuccess)
import Data.Char (ord)
import Model.Flashcard
import Model.Deck
import Model.Session
import SM2

import Data.Char (ord)

nerdDeckASCII :: String
nerdDeckASCII =
    "     _   __              ______            __  \n" ++
    "    / | / /__  _________/ / __ \\___  _____/ /__\n" ++
    "   /  |/ / _ \\/ ___/ __  / / / / _ \\/ ___/ //_/\n" ++
    "  / /|  /  __/ /  / /_/ / /_/ /  __/ /__/ ,<   \n" ++
    " /_/ |_/\\___/_/   \\__,_/_____/\\___/\\___/_/|_|  \n"   

charToInt :: Char -> Int
charToInt char = ord char - ord '0'

mainMenu :: FlashcardDeck -> IO ()
mainMenu deck = do
    putStrLn "🚀 Main Menu, please make a choice"
    putStrLn "Options:"
    putStrLn "1. Add Flash Card"
    putStrLn "2. View Flash Cards"
    putStrLn "3. Start Learning"
    putStrLn "4. Exit"
    option <- getLine

    let choice = charToInt (head option)
    handleMenuOption choice deck

handleMenuOption :: Int -> FlashcardDeck -> IO ()
handleMenuOption 1 deck = do
    putStrLn "You chose option 1 (Add Flash Card)"
    putStrLn "Enter the question: "
    question <- getLine
    putStrLn "Enter the answer: "
    answer <- getLine

    let newCard = Flashcard
            { flashcardId = length deck + 1
            , question = question
            , answer = answer
            , interval = initialInterval
            , ease = Ease1
            , reviewDate = 0
            , deck = Deck 1 "Math"
            }
    let updatedDeck = addToDeck deck newCard
    putStrLn "Flash card added successfully!" 
    mainMenu updatedDeck

handleMenuOption 2 deck = do
    putStrLn "You chose option 2 (View Flash Cards)"
    putStrLn "Flash Cards:"
    mapM_ (putStrLn . show) deck
    mainMenu deck

handleMenuOption 3 deck = do
    putStrLn "You chose option 3 (Start Learning)"

    -- Start a new session
    let newSession = startNewSession deck

    -- At first check if there are any flashcards in the deck
    let maybeCurrentCard = getCurrentCard newSession
    case maybeCurrentCard of
        Nothing -> do
            putStrLn "There are no Flashcards in the deck. Add Flashcards first before learning."
        Just currentCard -> do
            putStrLn $ "\n\n---Question---\n" ++ question currentCard
            putStrLn "---Press Enter to reveal the answer..."
            _ <- getLine  -- Wait for Enter
            putStrLn $ "---Answer---\n" ++ answer currentCard ++ "\n\n"

            -- putStrLn "Did you recall the answer correctly? (y/n): "
            -- successOption <- getLine
        
            -- let success = successOption == "y" || successOption == "Y"

            -- let updatedSession = endSession success newSession
            -- putStrLn "Review recorded!"

            -- -- Recursively call handleMenuOption 3 with the updated session and deck
            -- handleMenuOption 3 (sessionDeck updatedSession)
   
    mainMenu deck

handleMenuOption 4 deck = do
    putStrLn "Exiting program."
    exitSuccess

handleMenuOption _ deck = do
    putStrLn "Invalid option. Please choose a valid option (1-4)."
    mainMenu deck 


main :: IO ()
main = do
    putStrLn "\n------------------------------------------------------------------------------------\n"
    putStrLn $ "Welcome to\n" ++ nerdDeckASCII
    putStrLn "Developed by Maximilian Gobbel"
    putStrLn "If you want to know more about NerdDeck, visit https://github.com/maex0/nerddeck"
    putStrLn "\n------------------------------------------------------------------------------------\n"

    let initialDeck = []  -- Start with an only one empty deck for the moment
    mainMenu initialDeck
