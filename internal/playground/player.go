package playground

import (
	"crypto/rand"
	"math/big"
)

type Player struct {
	name     string
	token    string
	position int
	moves    int
}

const diceNumberOfSides = 6

// move the player to a new position
func (p *Player) Move(pos Position) {
	p.position += int(pos)
	p.moves++
}

// RollDice rolls the dice and returns a random number between 1 and 6
func (p *Player) RollDice() int {
	n, err := rand.Int(rand.Reader, big.NewInt(diceNumberOfSides))
	if err != nil {
		panic(err)
	}

	return int(n.Int64()) + 1
}
