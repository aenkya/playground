package playground

import "fmt"

type Board struct {
	size            int
	snakePositions  map[int]Position
	ladderPositions map[int]Position
}

/*
MovePlayer moves the player to a new position on the board
*/
func (b *Board) MovePlayer(p *Player, pos Position) {
	newPos := p.position + int(pos)
	if newPos > b.size {
		fmt.Printf("\n%s with token %s cannot move beyond %d", p.name, p.token, b.size)
		return
	}
	// check if the player has landed on a snake
	if newPos, ok := b.snakePositions[newPos]; ok {
		p.Move(newPos)
		fmt.Printf("\n%s with token %s landed on a snake and moved to %d", p.name, p.token, p.position)
		return
	}

	// check if the player has landed on a ladder
	if newPos, ok := b.ladderPositions[newPos]; ok {
		p.Move(newPos)
		fmt.Printf("\n%s with token %s landed on a ladder and moved to %d", p.name, p.token, p.position)
		return
	}

	// move the player to the new position
	p.Move(Position(newPos))
	fmt.Printf("\n%s with token %s moved to %d", p.name, p.token, p.position)
}

/*
check if the player has won the game
by reaching the last position on the board
*/
func (b *Board) CheckWinner(p *Player) bool {
	return p.position == b.size
}
