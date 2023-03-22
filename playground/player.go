package playground

import (
	"math/rand"
	"time"
)

type Player struct {
	name     string
	token    string
	position int
	moves    int
}

// move the player to a new position
func (p *Player) Move(pos Position) {
	p.position += int(pos)
	p.moves++
}

// RollDice rolls the dice and returns a random number between 1 and 6
func (p *Player) RollDice() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}
