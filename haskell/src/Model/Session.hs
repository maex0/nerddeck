module Model.Session
(Session(..)
,startNewSession
,endSession
,getCurrentCard) where

import Model.Flashcard

data Session = Session
    { sessionDeck :: FlashcardDeck
    , sessionSuccess :: Bool
    } deriving (Show)

startNewSession :: FlashcardDeck -> Session
startNewSession deck = Session { sessionDeck = deck, sessionSuccess = False }

endSession :: Session -> FlashcardDeck
endSession session = reviewDeck (sessionDeck session) (sessionSuccess session)

getCurrentCard :: Session -> Maybe Flashcard
getCurrentCard session =
    case sessionDeck session of
        [] -> Nothing
        (currentCard:_) -> Just currentCard
