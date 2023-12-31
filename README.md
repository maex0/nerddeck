# NerdDeck

![Build Status](https://github.com/maex0/nerddeck/actions/workflows/dotnet.yml/badge.svg)
![Build Status](https://github.com/maex0/nerddeck/actions/workflows/go.yml/badge.svg)
![Build Status](https://github.com/maex0/nerddeck/actions/workflows/latex.yml/badge.svg)

Welcome to NerdDeck, a project initiated during the first semester of my Master's program at [TH Rosenheim](https://www.th-rosenheim.de) in the course "Concepts of Programming Languages." In this project, I will be comparing two programming languages, [Go](https://go.dev) and [F#](https://dotnet.microsoft.com/languages/fsharp), within the context of functional programming.

## Project Overview

This repository is dedicated to exploring the paradigms of functional programming in the context of two distinct programming languages: Go and F#.

<img src="ScreenshotNerdDeck.png" alt="ScreenshotNerdDeck" width="500"/>

## Table of Contents

- [Introduction](#nerddeck)
- [Project Overview](#project-overview)
- [Goals](#goals)
- [Getting Started](#getting-started)
- [Database](#database)
  - [Database Schema](#database-schema)
- [Model](#Model)
- [Contributing](#contributing)
- [License](#license)

## Goals

- Compare and contrast the key features of Go and F# as functional programming languages.
- Provide practical examples and use cases to demonstrate the strengths and weaknesses of each language.
- Implement the spaced repetition algorithm SM-2. Information about this algorithm can be found [here](https://www.supermemo.com/en/blog/application-of-a-computer-to-improve-the-results-obtained-in-working-with-the-supermemo-method).

## Getting Started

## Database

**Disclaimer:** For the purpose of this project, a single json file is used as a simple and lightweight database to store flashcards. While this approach is suitable for educational and illustrative purposes, it may not be suitable for production usage due to limitations in scalability and concurrent access.

In a production environment, a more robust database solution should be considered, such as a relational database (e.g., PostgreSQL, MySQL) or a NoSQL database (e.g., MongoDB). The choice of the database will depend on the specific requirements of the application.

### Database Schema

The flashcard data is stored in a JSON file with the following structure:

```json
[
  {
    "ID": "37268335dd6931045bdcdf926",
    "Question": "What algebraic data types does F# use?",
    "Answer": "Record types and discriminated unions",
    "Repetitions": 1,
    "EasinessFactor": 1.3,
    "NextReview": "2023-12-18T18:17:22.438077+01:00"
  }
]
```

## Model

ID is the primary key, generated from question and answer.

![Nerddeck model](NerddeckModel.png)

## Contributing

Contributions are welcome! If you have insights, suggestions, or additional examples to contribute, feel free to contact me.

## License

This project is open-source and available under the [MIT License](LICENSE). See the [License](LICENSE) file for more details.
