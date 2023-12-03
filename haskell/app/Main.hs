module Main where

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

mainMenu :: IO (Int)
mainMenu = do
  putStrLn "\nOptions:"
  putStrLn "1. Add Flash Card"
  putStrLn "2. View Flash Cards"
  putStrLn "3. Start Learning"
  putStrLn "4. Exit"
  option <- getLine
  putStrLn $ "Your choice: " ++ option
  -- here map option to values
  -- chartoint
  return 1

endlessMenu :: IO()
endlessMenu = do
    mainMenu
    endlessMenu

    


main :: IO ()
main = do
    putStrLn $ "\nWelcome to\n" ++ nerdDeckASCII
    putStrLn "Developed by Maximilian Gobbel"
    putStrLn "If you want to know more about NerdDeck, visit https://github.com/maex0/nerddeck"

    endlessMenu
  
    let sampleCard1 = Flashcard 1 "Question" "Answer" 1 Ease1 0 (Deck 1 "Math")
    
    let initialDeck = addToDeck [] sampleCard1
    let session = startNewSession initialDeck

    putStrLn "Initial Deck:"
    mapM_ (putStrLn . show) (sessionDeck session)

    putStrLn "\nEnd Session (Success):"
    let updatedDeck = endSession session { sessionSuccess = True }
    mapM_ (putStrLn . show) updatedDeck
