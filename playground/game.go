package playground

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Game interface {
	Start()
}

type Board struct {
	size            int
	snakePositions  map[int]Position
	ladderPositions map[int]Position
}

func (b *Board) Print() {
	// print the board layout
	for i := 1; i <= b.size; i++ {
		fmt.Printf("%d ", i)
		// check if the position has a snake
		if _, ok := b.snakePositions[i]; ok {
			fmt.Printf("S ")
		}
		// check if the position has a ladder
		if _, ok := b.ladderPositions[i]; ok {
			fmt.Printf("L ")
		}
		if i%10 == 0 {
			fmt.Println()
		}
	}
}

func (b *Board) MovePlayer(p *Player, pos Position) {
	// move the player to a new position
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

func (b *Board) CheckWinner(p *Player) bool {
	// check if the player has won the game
	return false
}

type Position int

type Snake struct {
	start int
	end   int
}

type Ladder struct {
	start int
	end   int
}

type Player struct {
	name     string
	token    string
	position int
	moves    int
}

func (p *Player) Move(pos Position) {
	// move the player to a new position
	p.position += int(pos)
	p.moves++
}

func (p *Player) RollDice() int {
	// generate a random number between 1 and 6
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

type SnakesAndLaddersGame struct {
	board   *Board
	snakes  []*Snake
	ladders []*Ladder
	players []*Player
}

func (s *SnakesAndLaddersGame) Start() {
	var wg sync.WaitGroup
	wg.Add(1)
	// Intro message
	fmt.Println("Welcome to Snakes & Ladders!")
	// Print the board layout
	s.board.Print()

	var numOfPlayers int
	// Ask for number of players
	go func() {
		defer wg.Done()
		fmt.Print("Enter number of players: ")
		_, err := fmt.Scanln(&numOfPlayers)
		if err != nil {
			fmt.Println("Invalid number of players.")
		}
	}()

	wg.Wait()
	// Ask for name of each player
	for i := 0; i < numOfPlayers; i++ {
		name := ""
		fmt.Printf("\nEnter name of player %d: ", i+1)
		fmt.Scan(&name)
		// Ask for token of each player
		token := ""
		fmt.Printf("\nEnter token of player %d: ", i+1)
		fmt.Scan(&token)
		s.players = append(s.players, &Player{name: name, token: token})
	}

	// Ask each player to roll the dice
	for {
		for _, player := range s.players {
			fmt.Printf("\n%s's turn. Press enter to roll the dice.", player.name)
			fmt.Scanln()
			// Roll the dice
			dice := player.RollDice()
			fmt.Printf("\n%s rolled a %d", player.name, dice)
			// Move the player on the board
			s.board.MovePlayer(player, Position(dice))
			// Check if the player has won the game
			if s.board.CheckWinner(player) {
				fmt.Printf("\n%s won the game!", player.name)
				return
			}
		}
	}
}

func NewGame() Game {
	// initialize the board
	board := &Board{size: 100}

	// initialize the snakes
	snakes := []*Snake{
		{start: 99, end: 54},
		{start: 70, end: 55},
		{start: 52, end: 42},
		{start: 25, end: 2},
	}
	board.snakePositions = make(map[int]Position, len(snakes))
	// add the snakes to the board
	for _, snake := range snakes {
		board.snakePositions[snake.start] = Position(snake.end)
	}

	// initialize the ladders
	ladders := []*Ladder{
		{start: 6, end: 25},
		{start: 11, end: 40},
		{start: 60, end: 85},
		{start: 46, end: 90},
	}
	board.ladderPositions = make(map[int]Position, len(ladders))
	// add the ladders to the board
	for _, ladder := range ladders {
		board.ladderPositions[ladder.start] = Position(ladder.end)
	}

	// initialize the game
	game := &SnakesAndLaddersGame{
		board:   board,
		snakes:  snakes,
		ladders: ladders,
	}
	return game
}
