data Flashcard = Flashcard
    { flashcardId :: Int
    , question :: String
    , answer :: String
    , interval :: Int
    , ease :: Int
    , reviewDate :: Int -- todo: This should be a real date later
    , deck :: Deck
    } deriving (Show)

data Deck = Deck
    { deckId :: Int
    , name :: String
    } deriving (Show)


main :: IO ()
main = print $ Deck 0 "Math"
 