module Model.Flashcard
    ( Flashcard(..)
    , Ease(..)
    , presentFlashcard
    , reviewFlashcard
    , addToDeck
    , FlashcardDeck
    , reviewDeck
    , initialInterval
    ) where

import Data.Time (UTCTime)

data Flashcard = Flashcard
    { flashcardId :: Int
    , question :: String
    , answer :: String
    , repetitions :: Int
    , efactor :: Ease
    , nextReview :: UTCTime
    } deriving (Show)

data Ease = Ease1 | Ease2 | Ease3 | Ease4 | Ease5 deriving (Show)

initialInterval :: Int
initialInterval = 1

type FlashcardDeck = [Flashcard]

easeLevel :: Ease -> Int
easeLevel e = case e of
    Ease1 -> 1
    Ease2 -> 2
    Ease3 -> 3
    Ease4 -> 4
    Ease5 -> 5

presentFlashcard :: Flashcard -> Flashcard
presentFlashcard card = card { repetitions = 1 }

reviewFlashcard :: Flashcard -> Bool -> Flashcard
reviewFlashcard card success =
    if success
        then card { repetitions = nextInterval (repetitions card) }
        else card { repetitions = 1 }

nextInterval :: Int -> Int
nextInterval currentInterval = currentInterval * 2

addToDeck :: FlashcardDeck -> Flashcard -> FlashcardDeck
addToDeck deck card = card : deck

reviewDeck :: FlashcardDeck -> Bool -> FlashcardDeck
reviewDeck deck success = map (\card -> reviewFlashcard card success) deck

convertGrade :: String -> Int
convertGrade grade = case grade of
    "1" -> 1
    "2" -> 2
    "3" -> 3
    "4" -> 4
    _ -> 1
