module Main where

import Model.Flashcard
import Model.Deck
import Model.Session
import SM2

main :: IO ()
main = do
    let sampleCard1 = Flashcard 1 "Question" "Answer" 1 Ease1 0 (Deck 1 "Math")
    
    let initialDeck = addToDeck [] sampleCard1
    let session = startNewSession initialDeck

    putStrLn "Initial Deck:"
    mapM_ (putStrLn . show) (sessionDeck session)

    putStrLn "\nEnd Session (Success):"
    let updatedDeck = endSession session { sessionSuccess = True }
    mapM_ (putStrLn . show) updatedDeck
