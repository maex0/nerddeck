module Tests

open NUnit.Framework
open FlashCard
open System

[<SetUp>]
let Setup () = ()

[<Test>]
let ``FlashCard creation test`` () =
    let id = "1"
    let question = "What is the capitol of france?"
    let answer = "Paris"
    let repetitions = 0
    let easinessFactor = 2.5
    let nextReview = DateTime.Now

    let card =
        { ID = id
          Question = question
          Answer = answer
          Repetitions = repetitions
          EasinessFactor = easinessFactor
          NextReview = nextReview }

    Assert.AreEqual(id, card.ID)
    Assert.AreEqual(question, card.Question)
    Assert.AreEqual(answer, card.Answer)
    Assert.AreEqual(repetitions, card.Repetitions)
    Assert.AreEqual(easinessFactor, card.EasinessFactor)
    Assert.AreEqual(nextReview, card.NextReview)

[<Test>]
let ``FlashCard NextReview date test`` () =
    let id = "1"
    let question = "What is the capitol of france?"
    let answer = "Paris"
    let repetitions = 0
    let easinessFactor = 2.5
    let nextReview = DateTime.Now.AddDays(easinessFactor * float repetitions)

    let card =
        { ID = id
          Question = question
          Answer = answer
          Repetitions = repetitions
          EasinessFactor = easinessFactor
          NextReview = nextReview }

    let expectedNextReview = DateTime.Now.AddDays(easinessFactor * float repetitions)
    Assert.AreEqual(expectedNextReview.Date, card.NextReview.Date)


[<Test>]
let ``Test addFlashCard function`` () =
    let cards = []
    let question = "What is F#?"
    let answer = "A functional-first programming language."
    let newCards = addFlashCard question answer cards
    Assert.AreEqual(1, newCards.Length)
    Assert.AreEqual(question, newCards.Head.Question)
    Assert.AreEqual(answer, newCards.Head.Answer)

[<Test>]
let ``Test getDueFlashCards function`` () =
    let card1 =
        { ID = "1"
          Question = "Q1"
          Answer = "A1"
          Repetitions = 0
          EasinessFactor = 2.5
          NextReview = DateTime.Now.AddDays(-1.0) }

    let card2 =
        { ID = "2"
          Question = "Q2"
          Answer = "A2"
          Repetitions = 0
          EasinessFactor = 2.5
          NextReview = DateTime.Now.AddDays(1.0) }

    let cards = [ card1; card2 ]
    let dueCards = getDueFlashCards cards
    Assert.AreEqual(1, dueCards.Length)
    Assert.AreEqual(card1, dueCards.Head)

[<Test>]
let ``Test findCardByID function`` () =
    let card1 =
        { ID = "1"
          Question = "Q1"
          Answer = "A1"
          Repetitions = 0
          EasinessFactor = 2.5
          NextReview = DateTime.Now.AddDays(-1.0) }

    let card2 =
        { ID = "2"
          Question = "Q2"
          Answer = "A2"
          Repetitions = 0
          EasinessFactor = 2.5
          NextReview = DateTime.Now.AddDays(1.0) }

    let cards = [ card1; card2 ]
    let foundCard = findCardByID cards "1"
    Assert.AreEqual(Some card1, foundCard)

[<Test>]
let ``Test convertGrade function`` () =
    Assert.AreEqual(1, convertGrade "1")
    Assert.AreEqual(defaultGrade, convertGrade "5")
    Assert.AreEqual(defaultGrade, convertGrade "abc")

[<Test>]
let ``Test calculateNewEFactor function`` () =
    Assert.AreEqual(2.3600000000000003, calculateNewEFactor 2.5 3)

[<Test>]
let ``Test applySM2Algorithm function`` () =
    let card =
        { ID = "1"
          Question = "Q1"
          Answer = "A1"
          Repetitions = 0
          EasinessFactor = 2.5
          NextReview = DateTime.Now.AddDays(-1.0) }

    let updatedCard = applySM2Algorithm card "3"
    Assert.AreEqual(1, updatedCard.Repetitions)
    Assert.AreEqual(2.3600000000000003, updatedCard.EasinessFactor)

[<Test>]
let ``Test applySM2Algorithm function second case`` () =
    let card =
        { ID = "1"
          Question = "Q1"
          Answer = "A1"
          Repetitions = 1
          EasinessFactor = 2.5
          NextReview = DateTime.Now.AddDays(-1.0) }

    let updatedCard = applySM2Algorithm card "4"
    Assert.AreEqual(2, updatedCard.Repetitions)
    Assert.AreEqual(2.5, updatedCard.EasinessFactor)

[<Test>]
let ``Test applySM2Algorithm function third case`` () =
    let card =
        { ID = "1"
          Question = "Q1"
          Answer = "A1"
          Repetitions = 1
          EasinessFactor = 2.5
          NextReview = DateTime.Now.AddDays(-1.0) }

    let updatedCard = applySM2Algorithm card "1"
    Assert.AreEqual(1, updatedCard.Repetitions)
    Assert.AreEqual(1.3, updatedCard.EasinessFactor)
