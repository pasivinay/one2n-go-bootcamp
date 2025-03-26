package main

type Player struct {
	Score  int     `json:"score"`
	Name   string  `json:"name"`
	Target int     `json:"target"`
	Wins   float64 `json:"wins"`
}

type Game struct {
	P1       *Player `json:"p1"`
	P2       *Player `json:"p2"`
	Diceroll [][]int `json:"diceroll"`
	Matches  int     `json:"matches"`
}
