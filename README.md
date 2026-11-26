# goLearn

A beginner-friendly Go learning repository with hands-on lessons covering Go fundamentals, from basic syntax to variables, types, and formatted output.

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)
[![GitHub last commit](https://img.shields.io/github/last-commit/AkhatorEnosa/goLearn?style=for-the-badge)](https://github.com/AkhatorEnosa/goLearn/commits/main)

---

## Table of Contents

- [About](#about)
- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Lessons Overview](#lessons-overview)
- [How to Run](#how-to-run)
- [Contributing](#contributing)
- [License](#license)
- [Author](#author)

---

## About

**goLearn** is a personal Go learning journey. Each lesson is a standalone Go program designed to teach a specific concept through practical examples. Whether you're just starting with Go or brushing up on the basics, this repository provides clear, runnable code snippets to help you learn by doing.

---

## Prerequisites

Before running the lessons, make sure you have [Go](https://golang.org/doc/install) installed on your machine.

Verify your installation:

```bash
go version
```

---

## Project Structure

```
goLearn/
├── README.md
├── lesson1/
│   └── main.go          # Lesson 1: Hello World & basic printing
├── lesson2/
│   └── main.go          # Lesson 2: Variables, types, and formatted output
├── lesson3/
│   └── main.go          # Lesson 3: Tip calculator with user input
├── lesson4/
│   └── main.go          # Lesson 4: Conditionals and loops
├── lesson5/
│   └── main.go          # Lesson 5: For loops (number guessing game)
├── lesson6/
│   └── main.go          # Lesson 6: Slices (shopping list)
├── lesson7/
│   └── main.go          # Lesson 7: Maps (key-value pairs)
├── lesson8/
│   └── main.go          # Lesson 8: Package declaration (stub)
├── mini-project1/
│   └── main.go          # Mini Project 1: Practical application of learned concepts
└── mini-project2/
    └── main.go          # Mini Project 2: Interactive CLI To-Do List Manager
```

---

## Lessons Overview

| Lesson | File | Topics Covered |
|--------|------|----------------|
| **Lesson 1** | `lesson1/main.go` | Package declaration, importing `fmt`, `fmt.Println` for basic output |
| **Lesson 2** | `lesson2/main.go` | Variable declaration, types (`string`, `int`, `float64`, `bool`), `fmt.Printf` with format verbs |
| **Lesson 3** | `lesson3/main.go` | User input with `fmt.Scanln`, arithmetic operations, tip calculator |
| **Lesson 4** | `lesson4/main.go` | Conditionals (`if/else`, `switch`), loops (`for`, `range`) |
| **Lesson 5** | `lesson5/main.go` | `for` loops, user input, number guessing game |
| **Lesson 6** | `lesson6/main.go` | Slices, `append`, `range`, dynamic lists |
| **Lesson 7** | `lesson7/main.go` | Maps, key-value pairs, counting items |
| **Mini Project 1** | `mini-project1/main.go` | Practical application combining multiple concepts from previous lessons |
| **Mini Project 2** | `mini-project2/main.go` | Interactive CLI To-Do List Manager |

---

## How to Run

Navigate to any lesson directory and run the program:

```bash
# Lesson 1
cd lesson1
go run main.go

# Lesson 2
cd lesson2
go run main.go

# Lesson 3
cd lesson3
go run main.go

# Lesson 4
cd lesson4
go run main.go

# Lesson 5
cd lesson5
go run main.go

# Lesson 6
cd lesson6
go run main.go

# Lesson 7
cd lesson7
go run main.go

# Lesson 8
cd lesson8
go run main.go

# Mini Project 1
cd mini-project1
go run main.go

# Mini Project 2
cd mini-project2
go run main.go
```

### Expected Output

**Lesson 1:**
```
Hello, Osa!
I love to just be.
```

**Lesson 2:**
```
Hello, osa!
 Your current level is 2 and your account balance is $100.50.
 Online status: false
```

**Lesson 3:** (interactive - enter values when prompted)
```
Enter Bill Amount: $100
Enter Tip Percentage: $15
Enter no. of People: 4
Your total bill is $115.00, including a $15.00 tip.
 Each person to pay: $28.75
```

---

## Contributing

Contributions are welcome! If you'd like to add a new lesson or improve an existing one:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/lesson-name`)
3. Add your lesson in a new `lessonN/` directory with a `main.go` file
4. Commit your changes (`git commit -m "Add lesson N: description"`)
5. Push to the branch (`git push origin feature/lesson-name`)
6. Open a Pull Request

Please follow the existing structure and include a brief comment at the top of your `main.go` explaining what the lesson teaches.

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Author

**AkhatorEnosa**  
[GitHub](https://github.com/AkhatorEnosa) • [Repository](https://github.com/AkhatorEnosa/goLearn)

---

Happy learning! 🚀