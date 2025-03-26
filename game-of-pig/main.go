package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

const result = "%v: Holding at %v vs %v: Holding at %v: wins: %v/10 (%.1f%%), losses: %v/10 (%.1f%%)\n"

func ParseJSON(filePath string) *Game {
	var game Game
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	err = json.Unmarshal(file, &game)
	if err != nil {
		log.Fatalf("Error Unmarshalling json file: %v", err)
	}
	return &game
}

func (p *Game) GameofPig() string {

	i := 0
	for i < p.Matches {
		p.Diceroll = append(p.Diceroll, []int{})
		winner := p.Winner(i)
		winner.Wins += 1
		p.P1.Score, p.P2.Score = 0, 0
		i++
	}
	return fmt.Sprintf(result, p.P1.Name, p.P1.Target, p.P2.Name, p.P2.Target, p.P1.Wins, p.P1.Wins*10, p.P2.Wins, p.P2.Wins*10)
}

func (p *Game) Winner(match int) *Player {
	var winner *Player
	curr_player := p.P1
	score := 0
	if len(p.Diceroll[match]) == 0 {
		p.rolldice(match)
	}
	i := 0
	for i < len(p.Diceroll[match]) {
		if i == len(p.Diceroll[match])-1 {
			p.rolldice(match)
		}
		if curr_player.Score+p.Diceroll[match][i]+score >= 100 {
			curr_player.Score += p.Diceroll[match][i] + score
			winner = curr_player
			break
		} else if p.Diceroll[match][i] == 1 {
			if curr_player == p.P1 {
				curr_player = p.P2
			} else {
				curr_player = p.P1
			}
			score = 0
		} else if curr_player.Target <= score+p.Diceroll[match][i] {
			curr_player.Score += score + p.Diceroll[match][i]
			if curr_player == p.P1 {
				curr_player = p.P2
			} else {
				curr_player = p.P1
			}
			score = 0
		} else {
			score += p.Diceroll[match][i]
		}
		i += 1
	}
	return winner
}

func (p *Game) rolldice(match int) {
	roll := rand.Intn(6) + 1
	p.Diceroll[match] = append(p.Diceroll[match], roll)
}

func main() {
	p := ParseJSON("games.json")
	fmt.Print(p.GameofPig())
}
