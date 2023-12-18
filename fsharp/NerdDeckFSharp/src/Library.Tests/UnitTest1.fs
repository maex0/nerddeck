module Library.Tests

open NUnit.Framework
open FlashCard
open System

[<SetUp>]
let Setup () =
    ()

[<Test>]
let ``FlashCard creation test`` () =
    let id = "1"
    let question = "What is the capitol of france?"
    let answer = "Paris"
    let repetitions = 0
    let eFactor = 2.5
    let nextReview = DateTime.Now

    let card = { ID = id; Question = question; Answer = answer; Repetitions = repetitions; EFactor = eFactor; NextReview = nextReview }

    Assert.AreEqual(id, card.ID)
    Assert.AreEqual(question, card.Question)
    Assert.AreEqual(answer, card.Answer)
    Assert.AreEqual(repetitions, card.Repetitions)
    Assert.AreEqual(eFactor, card.EFactor)
    Assert.AreEqual(nextReview, card.NextReview)

[<Test>]
let ``FlashCard NextReview date test`` () =
    let id = "1"
    let question = "What is the capitol of france?"
    let answer = "Paris"
    let repetitions = 0
    let eFactor = 2.5
    let nextReview = DateTime.Now.AddDays(eFactor * float repetitions)

    let card = { ID = id; Question = question; Answer = answer; Repetitions = repetitions; EFactor = eFactor; NextReview = nextReview }

    let expectedNextReview = DateTime.Now.AddDays(eFactor * float repetitions)
    Assert.AreEqual(expectedNextReview.Date, card.NextReview.Date)