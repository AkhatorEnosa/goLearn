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
└── lesson3/
    └── (placeholder)    # Lesson 3: Coming soon
```

---

## Lessons Overview

| Lesson | File | Topics Covered |
|--------|------|----------------|
| **Lesson 1** | `lesson1/main.go` | Package declaration, importing `fmt`, `fmt.Println` for basic output |
| **Lesson 2** | `lesson2/main.go` | Variable declaration, types (`string`, `int`, `float64`, `bool`), `fmt.Printf` with format verbs |

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