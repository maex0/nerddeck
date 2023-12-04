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

import Model.Deck

data Flashcard = Flashcard
    { flashcardId :: Int
    , question :: String
    , answer :: String
    , interval :: Int
    , ease :: Ease
    , reviewDate :: Int -- todo: This should be a real date later
    , deck :: Deck
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
presentFlashcard card = card { interval = 1 }

reviewFlashcard :: Flashcard -> Bool -> Flashcard
reviewFlashcard card success =
    if success
        then card { interval = nextInterval (interval card) }
        else card { interval = 1 }

nextInterval :: Int -> Int
nextInterval currentInterval = currentInterval * 2

addToDeck :: FlashcardDeck -> Flashcard -> FlashcardDeck
addToDeck deck card = card : deck

reviewDeck :: FlashcardDeck -> Bool -> FlashcardDeck
reviewDeck deck success = map (\card -> reviewFlashcard card success) deck

