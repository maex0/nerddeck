data Flashcard = Flashcard
    { flashcardId :: Int
    , question :: String
    , answer :: String
    , interval :: Int
    , ease :: Ease
    , reviewDate :: Int -- todo: This should be a real date later
    , deck :: Deck
    } deriving (Show)

data Deck = Deck
    { deckId :: Int
    , name :: String
    } deriving (Show)

data Ease = Ease1 | Ease2 | Ease3 | Ease4 | Ease5 deriving (Show)

easeLevel :: Ease -> Int
easeLevel e = case e of
  Ease1 -> 1
  Ease2 -> 2
  Ease3 -> 3
  Ease4 -> 4
  Ease5 -> 5

initialInterval = 1

presentFlashcard :: Flashcard -> Flashcard
presentFlashcard card = card { interval = 1 }

type FlashcardDeck = [Flashcard]

addToDeck :: FlashcardDeck -> Flashcard -> FlashcardDeck
addToDeck deck card = card : deck


reviewFlashcard :: Flashcard -> Bool -> Flashcard
reviewFlashcard card success =
    if success
        then card { interval = nextInterval (interval card) }
        else card { interval = 1 }

nextInterval :: Int -> Int
nextInterval currentInterval = currentInterval * 2


reviewDeck :: FlashcardDeck -> Bool -> FlashcardDeck
reviewDeck deck success = map (\card -> reviewFlashcard card success) deck


data Session = Session
    { sessionDeck :: FlashcardDeck
    , sessionSuccess :: Bool
    } deriving (Show)

startNewSession :: FlashcardDeck -> Session
startNewSession deck = Session { sessionDeck = deck, sessionSuccess = False }

endSession :: Session -> FlashcardDeck
endSession session = reviewDeck (sessionDeck session) (sessionSuccess session)



main :: IO ()
main = do
    let sampleCard1 = Flashcard 1 "Question" "Answer" 1 Ease1 0 (Deck 1 "Math")
    let sampleCard2 = Flashcard 2 "Question" "Answer" 1 Ease1 0 (Deck 1 "Math")
    
    let initialDeck = addToDeck [] sampleCard1
    let session = startNewSession initialDeck

    putStrLn "Initial Deck:"
    mapM_ (putStrLn . show) (sessionDeck session)

    putStrLn "\nEnd Session (Success):"
    let updatedDeck = endSession session { sessionSuccess = True }
    mapM_ (putStrLn . show) updatedDeck