module Model.Session
(Session(..)
,startNewSession
,endSession) where

import Model.Flashcard

data Session = Session
    { sessionDeck :: FlashcardDeck
    , sessionSuccess :: Bool
    } deriving (Show)

startNewSession :: FlashcardDeck -> Session
startNewSession deck = Session { sessionDeck = deck, sessionSuccess = False }

endSession :: Session -> FlashcardDeck
endSession session = reviewDeck (sessionDeck session) (sessionSuccess session)

