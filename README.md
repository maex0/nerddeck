# NerdDeck

![Build Status](https://github.com/maex0/nerddeck/actions/workflows/build.yml/badge.svg)

Welcome to NerdDeck, a project initiated during the first semester of my Master's program at [TH Rosenheim](https://www.th-rosenheim.de) in the course "Concepts of Programming Languages." In this project, I will be comparing two programming languages, [Go](https://go.dev) and [F#](https://dotnet.microsoft.com/languages/fsharp), within the context of functional programming.

## Project Overview

This repository is dedicated to exploring the paradigms of functional programming in the context of two distinct programming languages: Go and F#.

## Table of Contents

- [Introduction](#nerddeck)
- [Project Overview](#project-overview)
- [Goals](#goals)
- [Getting Started](#getting-started)
- [Database](#database)
  - [Database Schema](#database-schema)
- [Model](#Model)
- [Functional Programming in Go](#functional-programming-in-go)
- [Functional Programming in F#](#functional-programming-in-fsharp)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Goals

- Compare and contrast the key features of Go and F# as functional programming languages.
- Provide practical examples and use cases to demonstrate the strengths and weaknesses of each language.
- Implement the spaced repetition algorithm SM-2. Information about this algorithm can be found [here](https://www.supermemo.com/en/blog/application-of-a-computer-to-improve-the-results-obtained-in-working-with-the-supermemo-method)

## Getting Started

## Database

**Disclaimer:** For the purpose of this project, a single json file is used as a simple and lightweight database to store flashcards. While this approach is suitable for educational and illustrative purposes, it may not be suitable for production usage due to limitations in scalability and concurrent access.

In a production environment, a more robust database solution should be considered, such as a relational database (e.g., PostgreSQL, MySQL) or a NoSQL database (e.g., MongoDB). The choice of the database will depend on the specific requirements of the application.

### Database Schema

The flashcard data is stored in a JSON file with the following structure:

```json
[
  {
    "question": "What is functional programming?",
    "answer": "Functional programming is a programming paradigm that treats computation as the evaluation of mathematical functions and avoids changing-state and mutable data."
  }
]
```

## Model

Current Model (this probably will change after time)

![Nerddeck model](NerddeckModel.png)

## Contributing

Contributions are welcome! If you have insights, suggestions, or additional examples to contribute, feel free to contact me.

## License

This project is open-source and available under the [MIT License](LICENSE). See the [License](LICENSE) file for more details.
