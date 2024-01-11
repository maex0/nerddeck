module App.Tests

open NUnit.Framework
open Program
open System.IO
open System

[<TestFixture>]
type TestClass() =

    [<Test>]
    member _.TestPrintMainMenu() =
        let writer = new StringWriter()
        Console.SetOut(writer)
        printMainMenu ()
        let output = writer.ToString()
        Assert.IsTrue(output.Contains("Main Menu, please make a choice"))

    [<Test>]
    member _.TestPrintInstructions() =
        let writer = new StringWriter()
        Console.SetOut(writer)
        printInstructions ()
        let output = writer.ToString()
        Assert.IsTrue(output.Contains("Instructions:"))

    [<Test>]
    member _.TestPrintWelcomeMessage() =
        let writer = new StringWriter()
        Console.SetOut(writer)
        printWelcomeMessage ()
        let output = writer.ToString()
        Assert.IsTrue(output.Contains("Welcome to"))
