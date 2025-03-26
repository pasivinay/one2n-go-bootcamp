# **Go Bootcamp**

# Game of Pig Simulator

## Overview
This project is a **Go-based simulator** for the classic **Game of Pig**, where two players take turns rolling a die, following predefined holding strategies. The program reads game configurations from a **JSON file** and determines the winner based on dice rolls and player strategies.

## Features
- **Simulates multiple games** between two players.
- **Customizable player strategies** (holding at different thresholds).
- **Reads configurations** from a JSON file (`games.json`).
- **Implements test cases** using the `testify` package.

## Project Structure
```
📂 Project Root
├── main.go        # Entry point for the Game of Pig simulation
├── structs.go     # Defines Player and Game structs
├── main_test.go   # Unit tests for game functions
├── games.json     # JSON file with predefined game data
├── go.mod         # Go module dependencies
└── go.sum         # Go module checksums
```

## Installation & Setup
### **Prerequisites**
- Go **1.18+** installed

### **Clone the Repository**
```sh
git clone https://github.com/pasivinay/one2n-go-bootcamp.git
cd game-of-pig
```

### **Run the Simulation**
```sh
go run main.go
```

## How It Works
### **Game Logic**
1. Players **take turns rolling a die**.
2. If a player rolls a **1**, they lose their accumulated turn score, and the turn passes.
3. A player **holds when they reach their threshold (target score per turn)**.
4. The first player to **reach 100 points wins**.

### **Example Output**
```
John: Holding at 10 vs Jane: Holding at 15: wins: 4/10 (40.0%), losses: 6/10 (60.0%)
```
- **John holds at 10**, while **Jane holds at 15**.
- **John wins 4 out of 10 games**, while **Jane wins 6 out of 10 games**.

## Configuration (`games.json`)
Modify `games.json` to change player names, hold thresholds, and dice rolls.

Example:
```json
{
  "p1": {
    "name": "John",
    "target": 10
  },
  "p2": {
    "name": "Jane",
    "target": 15
  },
  "matches": 10,
  "Diceroll": [[]]
}
```

## Running Tests
Execute test cases using:
```sh
go test ./...
```
Tests are written using the `testify` package (`main_test.go`).

**Happy Coding! 🚀**