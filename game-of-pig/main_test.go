package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJSON(t *testing.T) {
	assert := assert.New(t)
	Expected := "John"
	Game := ParseJSON("games.json")
	assert.Equal(Expected, Game.P1.Name, "First ")
}

func TestGameofPig(t *testing.T) {
	assert := assert.New(t)
	game := ParseJSON("games.json")
	Actual := game.GameofPig()
	Expected := "John: Holding at 10 vs Jane: Holding at 15: wins: 4/10 (40.0%), losses: 6/10 (60.0%)\n"
	assert.Equal(Expected, Actual, "John should be winning losing with 4 wins and 6 losses")
}

func TestWinner(t *testing.T) {
	assert := assert.New(t)
	Diceroll := [][]int{{3, 1, 5, 5, 2, 6, 3, 6, 2, 5, 3, 2, 2, 5, 6, 2, 3, 5, 4, 3, 4, 6, 4, 4, 4, 3, 6, 1, 2, 2, 6, 3, 1, 4, 1, 2, 6, 5, 1, 1, 4, 5, 5, 4, 5, 3, 3, 6, 3, 2, 4}}
	game := Game{Diceroll: Diceroll, P1: &Player{Name: "Jane", Score: 0, Target: 10, Wins: 0}, P2: &Player{Name: "John", Score: 0, Target: 15, Wins: 0}}
	Actual := game.Winner(0).Name
	Expected := "John"
	assert.Equal(Expected, Actual, "The winner should be John for this game")
}
