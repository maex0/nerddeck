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

endSession :: Bool -> Session -> Session
endSession success session = session { sessionDeck = reviewDeck (sessionDeck session) success, sessionSuccess = success }

getCurrentCard :: Session -> Maybe Flashcard
getCurrentCard session =
    case sessionDeck session of
        [] -> Nothing
        (currentCard:_) -> Just currentCard
