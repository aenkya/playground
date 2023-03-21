package playground

import "fmt"

type Game interface {
	Start()
}

type Board struct {
	size int
}

func (b *Board) Print() {
	// print the board in CLI
}

func (b *Board) MovePlayer(p *Player, pos Position) {
	// move the player to a new position
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
}

func (p *Player) RollDice() int {
	// roll a dice and return the number
	return 0
}

type SnakesAndLaddersGame struct {
	board   *Board
	snakes  []*Snake
	ladders []*Ladder
	players []*Player
}

func (s *SnakesAndLaddersGame) Start() {
	// Print the layout of the board with snakes and ladders on it
	// Intro message
	fmt.Println("Welcome to Snakes & Ladders!")
	// Ask for number of players
	// Ask for name of each player
	// Ask for token of each player
	// Initialize the players
	// Start the game
	s.board.Print()
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
	// initialize the ladders
	ladders := []*Ladder{
		{start: 6, end: 25},
		{start: 11, end: 40},
		{start: 60, end: 85},
		{start: 46, end: 90},
	}
	// initialize the game
	game := &SnakesAndLaddersGame{
		board:   board,
		snakes:  snakes,
		ladders: ladders,
	}
	return game
}
