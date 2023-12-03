module Model.Deck
    ( Deck(..)
    ) where

data Deck = Deck
    { deckId :: Int
    , name :: String
    } deriving (Show)
